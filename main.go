package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gocolly/colly"
	"github.com/immersivesky/s3-upload-emojis/html"
	"github.com/immersivesky/s3-upload-emojis/internal/adapters/repository"
	"github.com/immersivesky/s3-upload-emojis/location"
	"github.com/immersivesky/s3-upload-emojis/logger"
	rw "github.com/immersivesky/s3-upload-emojis/redis"
	"github.com/immersivesky/s3-upload-emojis/s3"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"sync"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	client := s3.NewS3("emojis")
	fmt.Printf("Bucket: %s | Count: %s, count: %s\n", client.BucketName, client.GetSize(), client.GetObjectCount())

	db, err := repository.NewDB(os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		panic(err)
	}

	db.PushEmojiPack()

	cli := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	redisWriter := rw.NewRedisWriter(cli, "error_logs")

	logger := logger.NewZapLogger(redisWriter)

	forEarchLocations(client, db, location.GetLocations(), logger)
}

func forEarchLocations(s3 *s3.S3, db *repository.DB, locations []location.Location, logger *zap.Logger) {
	wg := &sync.WaitGroup{}

	for _, loc := range locations {
		wg.Add(1)
		go scrapingHTML(s3, db, loc, logger, wg)
	}

	wg.Wait()
}

func scrapingHTML(s3 *s3.S3, db *repository.DB, location location.Location, logger *zap.Logger, wg *sync.WaitGroup) {
	defer wg.Done()

	collector := colly.NewCollector()
	collector.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	collector.OnHTML("div.col-md-2.col-sm-3.col-xs-6.platform-grid-child.p-0", html.On(s3, db, location, logger))

	if err := collector.Visit(location.URL); err != nil {
		panic(err)
	}
}

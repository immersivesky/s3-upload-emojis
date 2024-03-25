package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"io"
	"math"
	"strconv"
	"strings"
)

type S3 struct {
	Client     *s3.Client
	BucketName string
}

func NewS3(bucketName string) *S3 {
	sess, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	return &S3{
		Client:     s3.NewFromConfig(sess),
		BucketName: bucketName,
	}
}

func (service S3) Upload(reader io.Reader, objectName string) string {
	uploader := manager.NewUploader(service.Client)
	out, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: &service.BucketName,
		Key:    &objectName,
		Body:   reader,
		ACL:    types.ObjectCannedACLPublicReadWrite,
	})
	if err != nil {
		panic(err)
	}

	return out.Location
}

func (service S3) GetSize() string {
	entries, err := service.Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &service.BucketName,
	})
	if err != nil {
		panic(err)
	}

	var size int64

	for _, object := range entries.Contents {
		size += *object.Size
	}

	return fmt.Sprintf("%fMiB", float64(size)/(1<<20))
}

func (service S3) GetObjectCount() string {
	entries, err := service.Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &service.BucketName,
	})
	if err != nil {
		panic(err)
	}

	var count int

	for _, object := range entries.Contents {
		_ = object
		count++
	}

	return Comma(int64(count), " ")
}

func Comma(num int64, symbol string) string {
	var sign string

	if num == math.MinInt64 {
		return "-9,223,372,036,854,775,808"
	}

	if num < 0 {
		sign = "-"
		num = 0 - num
	}

	parts := []string{"", "", "", "", "", "", ""}
	j := len(parts) - 1

	for num > 999 {
		parts[j] = strconv.FormatInt(num%1000, 10)
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		num = num / 1000
		j--
	}

	parts[j] = strconv.Itoa(int(num))
	return sign + strings.Join(parts[j:], symbol)
}

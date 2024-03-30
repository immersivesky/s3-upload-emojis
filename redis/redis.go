package redis

import "github.com/go-redis/redis"

type RedisWriter struct {
	cli     *redis.Client
	listKey string
}

func NewRedisWriter(cli *redis.Client, listKey string) *RedisWriter {
	return &RedisWriter{
		cli:     cli,
		listKey: listKey,
	}
}

func (w *RedisWriter) Write(b []byte) (int, error) {
	n, err := w.cli.RPush(w.listKey, b).Result()
	return int(n), err
}

package logger

import (
	"github.com/immersivesky/s3-upload-emojis/redis"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewZapLogger(redisWriter *redis.RedisWriter) *zap.Logger {
	lowPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.DebugLevel
	})
	jsonEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	stdCore := zapcore.NewCore(jsonEncoder, zapcore.Lock(os.Stdout), lowPriority)

	syncer := zapcore.AddSync(redisWriter)
	redisCode := zapcore.NewCore(jsonEncoder, syncer, lowPriority)

	core := zapcore.NewTee(stdCore, redisCode)

	return zap.New(core).WithOptions(zap.AddCaller())
}

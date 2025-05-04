package services

import (
	"fmt"
	"os"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/redis/go-redis/v9"
)

func ShortenUrl(originalUrl string) (string, error) {
	existingId, err := Redis.Get(Ctx, "url:"+originalUrl).Result()
	if err == nil {
		return fmt.Sprintf("%s/short?url=%s", os.Getenv("BASE_URL"), existingId), nil
	}

	var id string
	for {
		id, _ = gonanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", 4)
		_, err := Redis.Get(Ctx, "short:"+id).Result()
		if err == redis.Nil {
			break
		}
	}

	expiration := time.Hour * 24 * 30
	err1 := Redis.Set(Ctx, "short:"+id, originalUrl, expiration).Err()
	err2 := Redis.Set(Ctx, "url:"+originalUrl, id, expiration).Err()
	if err1 != nil || err2 != nil {
		return "", fmt.Errorf("erro ao salvar no redis")
	}

	return fmt.Sprintf("%s/short?url=%s", os.Getenv("BASE_URL"), id), nil
}

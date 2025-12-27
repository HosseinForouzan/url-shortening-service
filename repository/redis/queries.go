package redis

import (
	"context"
	"fmt"
	"time"
)

func (rdb *Redis) RedisSet(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error) {
	res , err := rdb.rdb.Set(ctx, key, value, expiration).Result()
	if err != nil {
		return "", fmt.Errorf("failed to set %w", err)
	}
	return res, nil
}

func (rdb *Redis) RedisGet(ctx context.Context, key string) (string, error) {
	res, err := rdb.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get %w", err)
	}

	return res, nil
}
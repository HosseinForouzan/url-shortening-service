package redis

import (
	"context"
	"fmt"
	"time"
)

func (rdb *Redis) CachSet(key string, value interface{}, expiration time.Duration) (string, error) {
	res , err := rdb.rdb.Set(context.Background(), key, value, expiration).Result()
	if err != nil {
		return "", fmt.Errorf("failed to set %w", err)
	}
	return res, nil
}

func (rdb *Redis) CachGet(key string) (string, error) {
	res, err := rdb.rdb.Get(context.Background(), key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get %w", err)
	}

	return res, nil
}
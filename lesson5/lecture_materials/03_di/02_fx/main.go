package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

type Cache interface {
	Store(key any, val any)
	Load(key any) (val any, ok bool)
}

type UserService struct {
	cache Cache
}

func NewUserService(lc fx.Lifecycle, cache Cache) *UserService {
	return &UserService{
		cache: cache,
	}
}

func main() {
	fx.New(
		fx.Provide(NewRedisCache),
		fx.Provide(NewRedisOptions),
		fx.Provide(NewUserService),

		fx.Invoke(func(us *UserService) {
			fmt.Println("app started:", us)
		}),
	).Run()
}

func NewRedisOptions(lc fx.Lifecycle) *redis.Options {
	return &redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	}
}

type RedisCache struct {
	rdb *redis.Client
}

func NewRedisCache(lc fx.Lifecycle, options *redis.Options) Cache {
	rdb := redis.NewClient(options)

	return &RedisCache{
		rdb: rdb,
	}
}

func (rc *RedisCache) Store(key any, val any) {
	err := rc.rdb.Set(context.Background(), fmt.Sprint(key), val, 0).Err()
	if err != nil {
		panic(err)
	}
}

func (rc *RedisCache) Load(key any) (val any, ok bool) {
	val, err := rc.rdb.Get(context.Background(), fmt.Sprint(key)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, false
		}
		panic(err)
	}
	return val, true
}

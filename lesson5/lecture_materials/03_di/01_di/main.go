package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Store interface {
	Store(key any, val any)
	Load(key any) (val any, ok bool)
}

type UserService struct {
	store Store
}

func NewUserService(store Store) *UserService {
	return &UserService{
		store: store,
	}
}

func main() {
	// store := &sync.Map{}
	cache := NewRedisCache("", "", 0)
	_ = NewUserService(cache)
}

type RedisStore struct {
	rdb *redis.Client
}

func NewRedisCache(addr, password string, db int) *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisStore{
		rdb: rdb,
	}
}

func (rc *RedisStore) Store(key any, val any) {
	err := rc.rdb.Set(context.Background(), fmt.Sprint(key), val, 0).Err()
	if err != nil {
		panic(err)
	}
}

func (rc *RedisStore) Load(key any) (val any, ok bool) {
	val, err := rc.rdb.Get(context.Background(), fmt.Sprint(key)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, false
		}
		panic(err)
	}
	return val, true
}

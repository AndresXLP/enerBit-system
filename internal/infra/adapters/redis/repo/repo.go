package repo

import (
	"enerBit-system/internal/domain/ports/redis/repo"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/gommon/log"
)

type redisRepo struct {
	logs redis.Conn
}

func NewRedisRepository(logs redis.Conn) repo.Repository {
	return &redisRepo{logs}
}

func (r redisRepo) SendStreamLog(message string) {
	if _, err := r.logs.Do("XADD", "mystream", "*", "message", message); err != nil {
		log.Error(err)
	}
}

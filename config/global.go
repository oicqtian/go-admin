package config

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	REDIS  *redis.Client
	CONFIG Midware
)

const (
	SIGNING_KEY = "OdeCXCGULv9lJ89F"
)

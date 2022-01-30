package role_cache

import (
	"fmt"
	gin_redis "github.com/fellowme/gin_common_library/redis"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"go_project/app/role/role_const"
)

func getRoleMenuTireRedisKeyCacheFormat(key string) string {
	return fmt.Sprintf(role_const.RoleMenuTireRedisKeyFormatString, key, role_const.RedisKeyVersion)
}

func SetRoleMenuMapRedisKeyCache(key string, value string, name ...string) {
	redisName := ""
	if len(name) != 0 {
		redisName = name[0]
	}
	err := gin_redis.SetKeyValue(redisName, getRoleMenuTireRedisKeyCacheFormat(key), value)
	if err != nil {
		zap.L().Error("SetPhoneRedisKeyCache error", zap.Any("error", err), zap.Any("key", key),
			zap.String("key", key), zap.String("value", value))
	}
}

func GetRoleMenuMapRedisKeyCache(key string, name ...string) ([]byte, error) {
	redisName := ""
	if len(name) != 0 {
		redisName = name[0]
	}
	value, err := redis.Bytes(gin_redis.GetKeyByte(redisName, getRoleMenuTireRedisKeyCacheFormat(key)))
	if err != nil {
		zap.L().Error("GetRoleMenuTireRedisKeyCache error", zap.Any("error", err), zap.Any("key", key),
			zap.String("key", key), zap.Any("value", value))
	}
	return value, err
}

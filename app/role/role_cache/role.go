package role_cache

import (
	"encoding/json"
	"errors"
	"fmt"
	gin_redis "github.com/fellowme/gin_common_library/redis"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"go_project/app/role/role_const"
	"go_project/app/role/role_param"
	"strconv"
)

func getRoleMenuRedisKeyCacheFormat(key string) string {
	return fmt.Sprintf(role_const.RoleMenuTireRedisKeyFormatString, key, role_const.RedisKeyVersion)
}

func SetRoleMenuMapRedisKeyCache(roleMap map[int][]role_param.MenuResponseParam, name ...string) error {
	if len(roleMap) == 0 {
		return errors.New("roleMap not empty")
	}
	redisName := "default"
	if len(name) != 0 {
		redisName = name[0]
	}
	redisList := make([]string, 0)
	for key, value := range roleMap {
		jsonByte, _ := json.Marshal(value)
		newKey := getRoleMenuRedisKeyCacheFormat(strconv.Itoa(key))
		redisList = append(redisList, newKey)
		redisList = append(redisList, string(jsonByte))
	}
	err := gin_redis.MSetKey(nil, redisName, redisList)
	if err != nil {
		zap.L().Error("SetRoleMenuMapRedisKeyCache error", zap.Any("error", err),
			zap.Any("roleMap", roleMap))
	}
	return err
}

func GetRoleMenuMapRedisKeyCache(key string, name ...string) ([]byte, error) {
	redisName := "default"
	if len(name) != 0 {
		redisName = name[0]
	}
	value, err := redis.Bytes(gin_redis.GetKey(nil, redisName, getRoleMenuRedisKeyCacheFormat(key)))
	if err != nil {
		zap.L().Error("GetRoleMenuTireRedisKeyCache error", zap.Any("error", err), zap.Any("key", key),
			zap.String("key", key), zap.Any("value", value))
	}
	return value, err
}

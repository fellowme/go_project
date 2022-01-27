package account_cache

import (
	"encoding/json"
	"fmt"
	gin_redis "github.com/fellowme/gin_common_library/redis"
	gin_util "github.com/fellowme/gin_common_library/util"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"go_project/app/account/account_const"
	"go_project/app/account/account_param"
)

func getTodayLoginUserKeyCacheFormat(key string, flag bool) string {
	if flag {
		fmt.Sprintf(account_const.UserLoginFormatString, key, account_const.RedisKeyVersion)
	}
	return fmt.Sprintf(account_const.UserLoginOutFormatString, key, account_const.RedisKeyVersion)
}

func getPhoneCodeKeyCacheFormat(key string) string {
	return fmt.Sprintf(account_const.PhoneFormatString, key, account_const.RedisKeyVersion)
}

func getSessionKeyCacheFormat(key int32) string {
	return fmt.Sprintf(account_const.SessionFormatString, key, account_const.RedisKeyVersion)
}

func GetPhoneRedisKeyCache(key string, name ...string) string {
	redisName := ""
	if len(name) != 0 {
		redisName = name[0]
	}
	code, err := redis.String(gin_redis.GetKeyByte(redisName, getPhoneCodeKeyCacheFormat(key)))
	if err != nil && err != redis.ErrNil {
		zap.L().Error("GetRedisKeyCache error", zap.Any("error", err), zap.Any("key", key))
		return ""
	}
	return code
}

func DeletePhoneRedisKeyCache(key string, name ...string) error {
	redisName := ""
	if len(name) != 0 {
		redisName = name[0]
	}
	err := gin_redis.DeleteKey(redisName, getPhoneCodeKeyCacheFormat(key))
	if err != nil && err != redis.ErrNil {
		zap.L().Error("DeletePhoneRedisKeyCache error", zap.Any("error", err), zap.Any("key", key))
	}
	return err
}

func SetPhoneRedisKeyCache(key string, value string, name ...string) {
	redisName := ""
	if len(name) != 0 {
		redisName = name[0]
	}
	err := gin_redis.SetKeyValue(redisName, getPhoneCodeKeyCacheFormat(key), value,
		account_const.VerificationCodeExpireKeyTimeSecond)
	if err != nil {
		zap.L().Error("SetPhoneRedisKeyCache error", zap.Any("error", err), zap.Any("key", getPhoneCodeKeyCacheFormat(key)),
			zap.String("key", key), zap.String("value", value),
			zap.Int("expire_time", account_const.VerificationCodeExpireKeyTimeSecond))
	}
}

func GetUserBitRedisKeyCache(key int, flag bool, name ...string) (int, error) {
	redisName := ""
	if len(name) != 0 {
		redisName = name[0]
	}
	data, err := gin_redis.GetBitmapKey(redisName, getTodayLoginUserKeyCacheFormat(gin_util.NowTimeToString(), flag), key)
	if err != nil {
		zap.L().Error("GetUserBitRedisKeyCache error", zap.Any("error", err), zap.Any("key", key))
		return data, err
	}
	return data, nil
}

func SetUserBitRedisKeyCache(key, value int, flag bool, name ...string) error {
	redisName := ""
	if len(name) != 0 {
		redisName = name[0]
	}
	err := gin_redis.SetBitmapKey(redisName, getTodayLoginUserKeyCacheFormat(gin_util.NowTimeToString(), flag), key, value)
	if err != nil {
		zap.L().Error("SetUserBitRedisKeyCache error", zap.Any("error", err), zap.Any("key", key), zap.Any("value", value))

	}
	return err
}

func GetUserRedisKeyCache(key int32, name ...string) (account_param.SessionUserParam, error) {
	redisName := ""
	if len(name) != 0 {
		redisName = name[0]
	}
	var data account_param.SessionUserParam
	dataByte, err := redis.Bytes(gin_redis.GetKeyByte(redisName, getSessionKeyCacheFormat(key)))
	if err != nil {
		zap.L().Error("GetUserRedisKeyCache error", zap.Any("error", err), zap.Any("key", key))
		return data, err
	}
	jsonError := json.Unmarshal(dataByte, &data)
	if jsonError != nil {
		zap.L().Error("GetUserRedisKeyCache json error", zap.Any("error", err), zap.Any("key", key), zap.Any("jsonData", string(dataByte)))
	}
	return data, jsonError
}

func SetUserRedisKeyCache(key int32, value interface{}, name ...string) error {
	redisName := ""
	if len(name) != 0 {
		redisName = name[0]
	}
	valueByte, _ := json.Marshal(value)
	err := gin_redis.SetKeyValue(redisName, getSessionKeyCacheFormat(key), string(valueByte),
		account_const.SessionExpireKeyTimeSecond)
	if err != nil {
		zap.L().Error("SetUserRedisKeyCache error", zap.Any("error", err), zap.Any("key", key),
			zap.Any("string(valueByte)", string(valueByte)), zap.Any("value", value),
			zap.Int("expire_time", account_const.VerificationCodeExpireKeyTimeSecond))
	}
	return err
}

func DeleteUserRedisKeyCache(key int32, name ...string) error {
	redisName := ""
	if len(name) != 0 {
		redisName = name[0]
	}
	err := gin_redis.DeleteKey(redisName, getSessionKeyCacheFormat(key))
	if err != nil && err != redis.ErrNil {
		zap.L().Error("DeleteUserRedisKeyCache error", zap.Any("error", err), zap.Any("key", key))
	}
	return err
}

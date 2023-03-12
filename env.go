package pkgo

import (
	"fmt"
	"os"
	"strconv"
)

func GetEnv(key string, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func GetEnvInt(key string, def int64) int64 {
	if v := os.Getenv(key); v != "" {
		iv, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			return iv
		} else {
			fmt.Printf("GetEnvInt: parse env[%s] to int64 failed: %s", key, err.Error())
		}
		return def
	}
	return def
}

package auth

import (
	"github.com/alicebob/miniredis"
	"log"
)

func mockedIsExistRedis() {
	IsExistRedis = func(key string) bool {
		s, err := miniredis.Run()
		if err != nil {
			panic(err)
		}
		defer s.Close()

		redisKey := "6c3a65d23c5f26fc529f6c5ce01a6b31"

		redErr := s.Set(redisKey, "")
		if redErr != nil {
			log.Println(redErr)
		}

		if s.Exists(key) {
			return true
		}
		return false
	}
}

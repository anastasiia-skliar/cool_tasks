package auth

func mockedIsExistRedis(key string){
	IsExistRedis = func(key string) bool {
		redisKey := "6c3a65d23c5f26fc529f6c5ce01a6b31"
		if key == redisKey {
			return true
		}
		return false
	}
}

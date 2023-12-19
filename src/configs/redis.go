package configs

import "github.com/redis/go-redis/v9"

func ConnectRedis() *redis.Client {
	_, redisURI := EnvData()
	opts, err := redis.ParseURL(redisURI)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opts)

	return client
}

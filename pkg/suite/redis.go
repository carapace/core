package test

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func Redis() (db *redis.Client, cleanup func()) {
	Logger.Debug("creating redis instance")

	decrement := resourceLimit.allow()

	resource, err := pool.Run("redis", "latest", []string{})
	if err != nil {
		Logger.Panic("Could not start redis resource:", zap.Error(err))
	}

	Logger.Debug("created redis resource")

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	err = pool.Retry(func() error {
		db = redis.NewClient(&redis.Options{
			Addr:     "localhost:" + resource.GetPort("6379/tcp"),
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		pong, pongErr := db.Ping().Result()
		if pongErr != nil {
			Logger.Debug("retrying redis ping", zap.Error(err))
			return pongErr
		}

		if pong != "PONG" {
			Logger.Panic("Could not connect to redis docker image")
		}
		return nil
	})
	Logger.Debug("connected to redis resource")

	if err != nil {
		Logger.Panic("Could not connect to redis docker image", zap.Error(err))
	}
	return db, func() {
		err := db.Close()
		if err != nil {
			Logger.Warn("unable to close redis", zap.Error(err))
		}
		err = resource.Close()
		if err != nil {
			Logger.Warn("unable to close redis resource", zap.Error(err))
		}

		decrement()
	}
}

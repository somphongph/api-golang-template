package server

import (
	"api-places/configs"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/jinzhu/copier"
	"github.com/tripgator/lib-golang-packages/xcache"
	"github.com/tripgator/lib-golang-packages/xdb"
	"github.com/tripgator/lib-golang-packages/xlogger"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	port  int
	mongo *mongo.Client
	redis *redis.Client
}

func NewServer() *http.Server {
	conf := configs.GetConfig()
	sect := configs.GetSecret()

	xlogger.Init(
		conf.Server.LogLevel,
		conf.App.Name,
		conf.Server.PrettyPrint,
	)

	var mongoConfig xdb.Mongo
	copier.Copy(&mongoConfig, sect.Mongo)
	mongo, err := xdb.NewMongoClient(context.Background(), &mongoConfig)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %s", err)
	}

	var redisConfig xcache.Redis
	copier.Copy(&redisConfig, sect.Redis)
	redis, err := xcache.NewRedis(&redisConfig)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %s", err)
	}

	newServer := &Server{
		port:  conf.App.Port,
		mongo: mongo,
		redis: redis.Client(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      newServer.RegisterRoutes(),
		IdleTimeout:  conf.Server.IdleTimeout,
		ReadTimeout:  conf.Server.ReadTimeout,
		WriteTimeout: conf.Server.WriteTimeout,
	}

	xlogger.Infof("Server is running on port: %d", conf.App.Port)

	return server
}

package main

import (
	"context"
	"fmt"
	"github.com/Islam-Miko/go-mongodb/config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
)

var (
	server *gin.Engine
	ctx context.Context
	mongoclient *mongo.Client
	redisclient *redis.Client
)

func init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load ENVS", err)
	}

	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI(config.DBUri)
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("MongoDB successfully connected...")

	redisclient = redis.NewClient(&redis.Options{
		Addr: config.RedisUri,
	})

	if _, err := redisclient.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	err = redisclient.Set(ctx, "ping-test", "Welcome to Redis", 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis client connected...")

	server = gin.Default()
}

func main(){
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)

	}
	defer mongoclient.Disconnect(ctx)

	value, err := redisclient.Get(ctx, "ping-test").Result()
	if err == redis.Nil{
		fmt.Println("Key: ping-test does not exists")
	} else if err != nil {
		panic(err)
	}
	router := server.Group("/api/v1")

	router.GET("/healthcheker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": value})
	})

	log.Fatal(server.Run(":" + config.Port))
}
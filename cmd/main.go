package main

import (
	"github.com/IsThatASkyline/wb_l0"
	"github.com/IsThatASkyline/wb_l0/controllers"
	"github.com/IsThatASkyline/wb_l0/models"
	"github.com/IsThatASkyline/wb_l0/pkg/handler"
	"github.com/IsThatASkyline/wb_l0/pkg/repository"
	"github.com/IsThatASkyline/wb_l0/pkg/repository/cache"
	"github.com/IsThatASkyline/wb_l0/pkg/repository/postgres"
	"github.com/IsThatASkyline/wb_l0/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	orderCache := cache.NewCache() //cache

	db, err := postgres.NewPostgresDB()
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	var orders []models.Order

	orders, err = postgres.GetOrders(db)
	if err != nil {
		log.Fatalf("error by getting orders from postgres :%s", err)
	}
	// set orders into cache
	for _, order := range orders {
		orderCache.Set(order.Order_uid, order)
	}

	// nats

	sc, err := stan.Connect("test-cluster", "test", stan.NatsURL("nats://nats_streaming:4222"))
	if err != nil {
		log.Fatalf("failed to connect to nats-streaming : %s", err)
	}

	// Simple Async Subscriber
	preTime, _ := time.ParseDuration("1m")
	sub, _ := sc.Subscribe("orders", controllers.MsgHandler(orderCache, db), stan.StartAtTimeDelta(preTime))

	repos := repository.NewRepository(orderCache)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(wb_l0.Server)
	if err := srv.Run(os.Getenv("SERVER_PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

	// Unsubscribe sudo docker-compose -f nats&nats-streaming.yaml up -d
	sub.Unsubscribe()

	// Close connection
	sc.Close()
}

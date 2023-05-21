package main

import (
	"datainjector/internal/helper"
	"datainjector/internal/mongodb"
	"datainjector/internal/postgresdb"
	"github.com/go-co-op/gocron"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syreclabs.com/go/faker"
	"syscall"
	"time"
)

var orderRepository *mongodb.OrderRepository
var productEventRepository *mongodb.ProductEventRepository
var userRepository *postgresdb.UserRepository

func main() {
    log.Println("Starting data injector...")
	mongoDb := InitMongo()
	defer mongoDb.Close()

	postgres := InitPostgres()
	defer postgres.Close()

	initProducts()
	initUsers()
	initOrders()

	s := gocron.NewScheduler(time.UTC)
	//TODO set to min sync
	_, err := s.Every(10).Second().Do(insertOrder)
	if err != nil {
		log.Panic("Cannot schedule order job", err)
	}

	_, err = s.Every(1).Minute().Do(updateProducts)
	if err != nil {
		log.Panic("Cannot schedule product job", err)
	}
	s.StartAsync()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
}

func InitMongo() *mongodb.MongoDB {
	mongoConfig := mongodb.MongoConfig{
		Uri: GetEnvOrDefault(
			"MONGO_URL",
			"mongodb://admin:admin123@localhost:10020",
		),
	}
	mongoDB := mongoConfig.Connect()
	orderRepository = mongodb.NewOrderRepository(mongoDB)
	productEventRepository = mongodb.NewProductStreamRepository(mongoDB)
	mongoDB.Test()
	return mongoDB
}

func GetEnvOrDefault(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func InitPostgres() *postgresdb.PostgresDB {
	postgresConfig := postgresdb.PostgresConfig{
		DSN: GetEnvOrDefault(
			"POSTGRES_DSN",
			"host=localhost port=10010 user=admin password=admin123 dbname=postgres",
		),
	}
	postgresDB := postgresConfig.Connect()
	userRepository = postgresdb.NewUserRepository(postgresDB)
	userRepository.CreateTable()
	return postgresDB
}

func initOrders() {
	count, err := orderRepository.Count()
	if err != nil {
		log.Panic(err)
	}
	//poor man's db init check
	if count == 0 {
		for i := 0; i < 100; i++ {
			order := &mongodb.Order{
				ProductId: rand.Intn(30) + 1,
				UserId:    rand.Intn(10) + 1,
				TimeStamp: helper.RandomTime(),
			}
			err := orderRepository.Insert(order)
			if err != nil {
				log.Panic("Cannot write order into db", err)
			}
		}
		log.Println("Orders has been initialized")
	}
}

func initProducts() {
	count, err := productEventRepository.Count()
	if err != nil {
		log.Panic(err)
	}
	//poor man's db init check
	if count == 0 {
		for i := 1; i <= 30; i++ {
			product := &mongodb.Product{
				ProductId:  i,
				Name:       faker.Commerce().ProductName(),
				Price:      rand.Intn(500) + 200,
				ModifiedAt: time.Now().UTC(),
			}
			err := productEventRepository.Insert(product)
			if err != nil {
				log.Panic("Cannot write product into db", err)
			}
		}
		log.Println("Products has been initialized")
	}
}

func initUsers() {
	count := userRepository.Count()
	if count == 0 {
		for i := 1; i <= 10; i++ {
			user := postgresdb.User{
				ID:    i,
				Name:  faker.Name().Name(),
				Email: faker.Internet().Email(),
			}
			err := userRepository.Insert(user)
			if err != nil {
				log.Panic("Cannot initialize users. ", err)
			}
		}
		log.Println("Users have been initialized")
	}

}

func insertOrder() {
	count := rand.Intn(20) + 10
	for i := 0; i < count; i++ {
		order := &mongodb.Order{
			ProductId: rand.Intn(30) + 1,
			UserId:    rand.Intn(10) + 1,
			TimeStamp: helper.RandomTime(),
		}
		err := orderRepository.Insert(order)
		if err != nil {
			log.Println("Cannot insert new order")
		}
	}
	log.Printf("%d orders were generated\n", count)
}

func updateProducts() {
	for i := 0; i < rand.Intn(10)+3; i++ {
		index := rand.Intn(30) + 1
		product := &mongodb.Product{
			ProductId:  index,
			Price:      rand.Intn(500) + 200,
			ModifiedAt: time.Now().UTC(),
		}
		err := productEventRepository.Insert(product)
		if err != nil {
			log.Panic("Cannot write product into db", err)
		}
	}

	log.Println("Products were updated")

}

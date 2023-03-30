package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Polidoro-root/go-expert-classes/18_clean_architecture/configs"
	"github.com/Polidoro-root/go-expert-classes/18_clean_architecture/internal/event/handler"
	"github.com/Polidoro-root/go-expert-classes/18_clean_architecture/internal/infra/web/webserver"
	"github.com/Polidoro-root/go-expert-classes/18_clean_architecture/pkg/events"
	"github.com/streadway/amqp"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel("amqp://guest:guest@localhost:5672")

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	// createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)

	webserver.AddHandler("/order", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			webOrderHandler.List(w, r)
		} else if r.Method == http.MethodPost {
			webOrderHandler.Create(w, r)
		} else {
			http.Error(w, "/order accept only GET and POST methods", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Starting web server on port ", configs.WebServerPort)

	webserver.Start()

}

func getRabbitMQChannel(url string) *amqp.Channel {
	conn, err := amqp.Dial(url)

	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	return ch
}

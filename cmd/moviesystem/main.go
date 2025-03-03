package main

import (
	"database/sql"
	"fmt"

	"monte_clone_go/configs"
	"monte_clone_go/internal/event/handler"
	"monte_clone_go/internal/infra/web/webserver"
	"monte_clone_go/pkg/events"

	"github.com/streadway/amqp"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Carrega as configurações do ambiente
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	// Conecta ao banco de dados MySQL
	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Obtém o canal do RabbitMQ
	rabbitMQChannel := getRabbitMQChannel()
	defer rabbitMQChannel.Close()

	// Configuração do EventDispatcher e handlers de eventos
	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("MovieCreated", &handler.MovieCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})
	eventDispatcher.Register("MovieUpdated", &handler.MovieUpdatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})
	eventDispatcher.Register("MovieDeleted", &handler.MovieDeletedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	// Configuração do WebServer e handlers HTTP
	webserver := webserver.NewWebServer(configs.WebServerPort)
	webMovieHandler := NewWebMovieHandler(db, eventDispatcher)

	// Adicionando rotas
	webserver.AddHandler("POST /movie", webMovieHandler.Create)
	webserver.AddHandler("GET /movie", webMovieHandler.Get)
	webserver.AddHandler("PUT /movie", webMovieHandler.Update)
	webserver.AddHandler("DELETE /movie", webMovieHandler.Delete)

	// Inicia o servidor HTTP
	fmt.Println("Starting web server on port", configs.WebServerPort)
	webserver.Start() // A função Start não retorna valor, então não há necessidade de verificar o retorno
}

// Função para obter o canal do RabbitMQ
func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}

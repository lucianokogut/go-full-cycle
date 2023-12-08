package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"br.com.lucianokogut/go-full-cycle-esquenta/internal/infra/akafka"
	"br.com.lucianokogut/go-full-cycle-esquenta/internal/infra/repository"
	"br.com.lucianokogut/go-full-cycle-esquenta/internal/infra/web"
	"br.com.lucianokogut/go-full-cycle-esquenta/internal/usecase"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewProductRepositoryMysql(db)
	createProductUsecase := usecase.NewCreateProductUseCase(repository)
	listProductsUsecase := usecase.NewListProductsUseCase(repository)

	productHandlers := web.NewProductHandlers(createProductUsecase, listProductsUsecase)

	r := chi.NewRouter()
	r.Post("/products", productHandlers.CreateProductHandler)
	r.Get("/products", productHandlers.ListProductsHandler)

	go http.ListenAndServe(":8000", r)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			fmt.Println("Erro na aplicação, logando mensagem...")
		}
		_, err = createProductUsecase.Execute(dto)
	}
}

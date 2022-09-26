package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"pay_gate/services"
	"pay_gate/services/actions"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("Server initializing...")
	initializeServer()

	router := gin.Default()
	router.GET("/echo", getServerStatus)

	router.POST("/payments", services.GetAction(actions.CreatePaymentActionType))

	router.Run("localhost:8080")
}

func getServerStatus(c *gin.Context) {

	mysql_status := "running"

	db, err := sql.Open("mysql", "zxqiu:1115@tcp(localhost:3306)/payment_service")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		mysql_status = "down"
	} else {
		var version string
		err = db.QueryRow("SELECT VERSION()").Scan(&version)

		if err != nil {
			log.Fatal(err)
			mysql_status = "down"
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"service_status": "running", "db_status": &mysql_status})
}

func initializeServer() {
	initializeKafka()
}

func initializeKafka() {
	hostname, err := os.Hostname()

	if err != nil {
		log.Fatalln("cannot got current hostname")
	}

	_, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         hostname,
		"acks":              "all"})

	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

}

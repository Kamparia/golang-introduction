package main

import (
	"02_useragent/database"
	"02_useragent/handler"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	// database connection
	db := database.DbConnect()

	// routers
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/useragent", handler.UserAgentHandler(db))
	http.HandleFunc("/report", handler.ReportHandler(db))

	// start the web server (http://localhost:8080/)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

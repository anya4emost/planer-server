package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/anya4emost/planer-server/internal/config"
	"github.com/anya4emost/planer-server/internal/server"
)

func main() {
	// app := fiber.New()
	// db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=planer_db sslmode=disable")

	// if err != nil {
	// 	pqErr := err.(*pq.Error)
	// 	log.Fatalf("Error connecting to database: %v\n", pqErr.Code)
	// }

	// defer db.Close()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	var result int
	// 	err = db.Get(&result, "SELECT 1")
	// 	if err != nil {
	// 		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	// 	}

	// 	return c.Status(200).JSON(fiber.Map{
	// 		"hello":    "world",
	// 		"database": "available",
	// 	})
	// })

	// log.Fatal(app.Listen("localhost:4000"))

	cfg := config.Load()
	s := server.NewServer(cfg)
	go func() {
		if err := s.Start(); err != nil {
			log.Fatalf("Server failed to start: %v\n", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down server...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Stop(); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server shutdown complete")
}

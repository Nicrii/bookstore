package main

import (
	"github.com/Nicrii/bookstore/internal/db"
	repo "github.com/Nicrii/bookstore/internal/repository"
	"github.com/Nicrii/bookstore/internal/server"
	"log"
	"net"
	"os"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)
	logger.Println("Service started")

	//initialise database
	database, err := db.InitGormMysql()
	if err != nil {
		log.Fatalf("Failed to initialise db: %v", err)
	}

	//create a new repository
	repository := repo.NewRepository(database)

	// create new rpc server
	srv := server.NewRPCServer(repository)

	// start rpc server
	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

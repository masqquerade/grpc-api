package main

import (
	"context"
	"log"
	"msqrd/pkg/user"
	"net"
	"time"

	api "msqrd/pkg/api"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := grpc.NewServer()
	srv := &user.GRPCServer{}

	srv.Store = srv.Store.New()

	err = srv.Store.Open(ctx)
	if err != nil {
		log.Fatal(err)
	}

	api.RegisterUserServiceServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

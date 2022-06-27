package main

import (
	"context"
	"flag"
	"log"
	"time"

	api "msqrd/pkg/api"

	"google.golang.org/grpc"
)

var emailf string
var usernamef string
var passwordf string
var methodf string
var idf int64

func init() {
	flag.StringVar(&emailf, "email", "", "your email for create a user")
	flag.StringVar(&usernamef, "username", "newuser", "your usename for create a user")
	flag.StringVar(&passwordf, "password", "root", "your password for create a user")
	flag.StringVar(&methodf, "method", "create", "method you want to call")
	flag.Int64Var(&idf, "id", 0, "id to delete user")
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewUserServiceClient(conn)

	switch methodf {
	case "create":
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		res, err := c.Create(ctx, &api.CreateRequest{
			Email:    emailf,
			Username: usernamef,
			Password: passwordf,
		})

		if err != nil {
			log.Fatal(err)
		}

		log.Println(res.GetId())

	case "delete":
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		res, err := c.Delete(ctx, &api.DeleteRequest{
			Id: int32(idf),
		})

		if err != nil {
			log.Fatal(err)
		}

		log.Println(res.GetRes())
	}
}

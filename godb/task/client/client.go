package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"github.com/atul161/godb/task/prototype"
)

func main() {
	fmt.Println("Helllo i am  a client")
	//grpinsecure is for ssl
	conn, err := grpc.Dial("localhost:50056", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	//we have to close connection
	defer conn.Close()
	//new client request
	c := prototype.NewUserProfilesClient(conn)

	req := &prototype.CreateUserProfileRequest{
		UserProfile: &prototype.UserProfile{
			Id:        "1",
			FirstName: "Sahaj",
			LastName:  "Khandelwal",
			Email:     "sahaj@gmail.com",
		},
	}

	res, err := c.CreateUserProfile(context.Background(), req)
	//checking the error
	if err != nil {
		log.Fatalf("Error while doing calci operation %v", err)
	}

	log.Printf("response from calci:%v", res.Email)
}

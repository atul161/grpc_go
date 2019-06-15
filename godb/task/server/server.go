package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/atul161/godb/task/prototype"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)

type server struct {
}

//create user profile
func (*server) CreateUserProfile(ctx context.Context, req *prototype.CreateUserProfileRequest) (*prototype.UserProfile, error) {
	fmt.Printf("Function Invoked %v", req)
	//now we need to implement this function
	//get the information
	id := req.GetUserProfile().GetId()
	firstname := req.GetUserProfile().GetFirstName()
	lastname := req.GetUserProfile().GetLastName()
	email := req.GetUserProfile().GetEmail()
	//now we can perform like insertion in database
	//todo
	//sending data to the databse
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//sql open takes two parameter the first one is driver name and second one tells hoe to
	//connect with database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//db.Ping() we force our code to actually open up a connection
	//to the database which will validate whether or not our connection
	//string was 100% correct.
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	sqlStatement := `
INSERT INTO userprofile (id, firstname, lastname, email)
VALUES ($1, $2, $3, $4)
RETURNING id`
	iid := 0
	err = db.QueryRow(sqlStatement, 31, "sallu", "kallu", "kallu@gmail.com").Scan(&iid)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	res := &prototype.UserProfile{
		Id:        id,
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
	}
	return res, nil

}

//get User profile
func main() {

	/*psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//sql open takes two parameter the first one is driver name and second one tells hoe to
	//connect with database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//db.Ping() we force our code to actually open up a connection
	//to the database which will validate whether or not our connection
	//string was 100% correct.
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	sqlStatement := `
	INSERT INTO userprofile (id, firstname, lastname, email)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, 30, "kishan", "anand", "atulanand1996dbg@gmail.com").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")*/
	//creating server
	fmt.Println("Welcome to the server")
	//listen tcp and port no
	lis, err := net.Listen("tcp", "0.0.0.0:50056")
	if err != nil {
		log.Fatalf("Failed to listen1 %v", err)
	}
	//creating grpc server
	s := grpc.NewServer()
	prototype.RegisterUserProfilesServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

}

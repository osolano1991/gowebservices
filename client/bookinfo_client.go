package main

import (
	"context"
	pb "github.com/osolano1991/gowebservices/booksapp"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

func main() {
	address := os.Getenv("ADDRESS")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBookInfoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddBook(ctx, &pb.Book{
		Id:        "1",
		Title:     "Operating System Concepts",
		Edition:   "9th",
		Copyright: "2012",
		Language:  "ENGLISH",
		Pages:     "976",
		Author:    "Abraham Silberschatz",
		Publisher: "John Wiley & Sons"})
	if err != nil {
		log.Fatalf("Could not add book: %v", err)
	}

	log.Printf("Book ID: %s added successfully", r.Value)
	book, err := c.GetBook(ctx, &pb.BookID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get book: %v", err)
	}
	log.Printf("Book: ", book.String())
}

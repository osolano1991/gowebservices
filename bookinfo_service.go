package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gofrs/uuid"
	pb "github.com/osolano1991/gowebservices/booksapp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	bookMap map[string]*pb.Book
}

func (s *server) AddBook(ctx context.Context, in *pb.Book) (*pb.BookID, error) {
	_, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"Error while generating Book ID", err)
	}
	//in.Id = out.String()
	if s.bookMap == nil {
		s.bookMap = make(map[string]*pb.Book)
	}
	s.bookMap[in.Id] = in
	return &pb.BookID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *server) GetBook(ctx context.Context, in *pb.BookID) (*pb.Book, error) {
	value, exists := s.bookMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Book does not exist.", in.Value)
}

func (s *server) DeleteBook(ctx context.Context, in *pb.BookID) (*pb.BookID, error) {
	_, exists := s.bookMap[in.Value]
	if exists {
		delete(s.bookMap, in.Value)
		return &pb.BookID{Value: in.Value}, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Book does not exist.", in.Value)
}

func (s *server) UpdateBook(ctx context.Context, in *pb.Book) (*pb.BookID, error) {
	if s.bookMap == nil {
		s.bookMap = make(map[string]*pb.Book)
	}
	s.bookMap[in.Id] = in
	return &pb.BookID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *server) ReadCSV(ctx context.Context, in *pb.File) (*pb.BookID, error) {

	f, _ := os.Open(in.Value)
	bookline, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	if s.bookMap == nil {
		s.bookMap = make(map[string]*pb.Book)
	}

	// Se recorren todos los registros
	for _, book := range bookline {
		fmt.Printf("\n\n======================================================================================")
		fmt.Printf("\n\n1) Creando libro con ID -> ", book[0])
		newBook, err := s.AddBook(ctx, &pb.Book{
			Id:        book[0],
			Title:     book[1],
			Edition:   book[2],
			Copyright: book[3],
			Language:  book[4],
			Pages:     book[5],
			Author:    book[6],
			Publisher: book[7]})

		if err != nil {
			log.Fatalf("\n\nError al agregar el libro: %v", err)
		}
		fmt.Printf("\n\tLibro agregado correctamente: ", newBook.String())

		bookGet, err := s.GetBook(ctx, &pb.BookID{Value: book[0]})
		if err != nil {
			fmt.Printf("\n\nEl libro consultado no existe: %v", err)
		} else {
			fmt.Printf("\n\n2) Consultando libro agregado \n\t-> ", bookGet.String())
		}
	}

	return &pb.BookID{Value: "1"}, status.New(codes.OK, "").Err()
}

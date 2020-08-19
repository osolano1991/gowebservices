package main

import (
	"context"
	pb "github.com/osolano1991/gowebservices/booksapp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	bookMap map[string]*pb.Book
}

func (s *server) AddBook(ctx context.Context, in *pb.Book) (*pb.BookID, error) {
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

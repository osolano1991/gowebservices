package main

import (
    "context"
    pb "github.com/osolano1991/gowebservices/booksapp"
    "github.com/gofrs/uuid"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type server struct {
    bookMap map[string]*pb.Book
}

func (s *server) AddBook(ctx context.Context, in *pb.Book) (*pb.BookID, error) {
    //out, err := uuid.NewV4()
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
   // out, err := uuid.NewV4()
   // if err != nil {
   //     return nil, status.Errorf(codes.Internal,
   //         "Error while generating Book ID", err)
   // }
   // in.Id = out.String()
    if s.bookMap == nil {
        s.bookMap = make(map[string]*pb.Book)
    }
    s.bookMap[in.Id] = in
    return &pb.BookID{Value: in.Id}, status.New(codes.OK, "").Err()
}
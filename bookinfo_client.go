package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/osolano1991/gowebservices/booksapp"
	"google.golang.org/grpc"
)

func main() {
	address := os.Getenv("ADDRESS")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBookInfoClient(conn)

	// Agregar libro
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
		log.Fatalf("\n\n\nNo se pudo agregar el libro: %v", err)
	}

	// Obtener libro
	log.Printf("\n\n\nLibro creado ID: %s", r.Value)
	book, err := c.GetBook(ctx, &pb.BookID{Value: r.Value})
	if err != nil {
		log.Fatalf("\n\n\nEl libro consultado no existe: %v", err)
	}
	log.Printf("\n\n\nLibro consultado: ", book.String())

	// Eliminar libro
	bookDel, err := c.DeleteBook(ctx, &pb.BookID{Value: r.Value})
	if err != nil {
		log.Fatalf("\n\n\nNo se pudo eliminar el libro: %v", err)
	}
	log.Printf("\n\n\nLibro eliminado: ", bookDel.String())

	// Obtener libro nuevamente para comprobar si se elimino
	bookGet, err := c.GetBook(ctx, &pb.BookID{Value: r.Value})
	if err != nil {
		//log.Fatalf("\n\nEl libro consultado no existe: %v", err)
		log.Printf("\n\n\nEl libro consultado no existe: %v", err)
	} else {
		log.Printf("\n\n\nLibro consultado: ", bookGet.String())
	}

	// Actualizar libro
	upd, err := c.UpdateBook(ctx, &pb.Book{
		Id:        "1",
		Title:     "UPDATED Book",
		Edition:   "10th",
		Copyright: "2012",
		Language:  "ENGLISH",
		Pages:     "976",
		Author:    "Abraham Silberschatz",
		Publisher: "John Wiley & Sons"})
	if err != nil {
		log.Fatalf("\n\n\nNo se pudo actualizar el libro: %v", err)
	} else {
		// Obtener libro nuevamente para comprobar si se actualizo
		bookGetUpdated, err := c.GetBook(ctx, &pb.BookID{Value: upd.Value})
		if err != nil {
			log.Fatalf("\n\n\nEl libro consultado no existe: %v", err)
		}
		log.Printf("\n\n\nLibro actualizado correctamente: ", bookGetUpdated.String())
	}

	// Leer CSV
	c.ReadCSV(ctx, &pb.File{Value: "books.csv"})
	log.Printf("\n\n\nProcesando y creando libros desde archivo CSV")

}

package main

import (
	"context"
	"fmt"
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
		log.Fatalf("did not connect -> %v", err)
	}
	defer conn.Close()
	c := pb.NewBookInfoClient(conn)

	// Agregar libro
	fmt.Println("\n\n1) Creando libro...")
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
		log.Fatalf("\n\tError al agregar el libro: -> ", err)
	}

	// Obtener libro
	book, err := c.GetBook(ctx, &pb.BookID{Value: r.Value})
	if err != nil {
		log.Fatalf("\n\tEl libro consultado no existe -> ", err)
	}
	fmt.Println("\n\n2) Obteniendo libro creado...")
	fmt.Println("\n\t", book.String())

	// Actualizar libro
	fmt.Println("\n\n3) Actualizando libro...")
	upd, err := c.UpdateBook(ctx, &pb.Book{
		Id:        "1",
		Title:     "Libro Actualizado",
		Edition:   "3th",
		Copyright: "2020",
		Language:  "SPANISH",
		Pages:     "1500",
		Author:    "Solano Mora",
		Publisher: "Juan Perez"})
	if err != nil {
		log.Fatalf("\n\tError al actualizar el libro: %v", err)
	} else {
		// Obtener libro nuevamente para comprobar si se actualizo
		bookGetUpdated, err := c.GetBook(ctx, &pb.BookID{Value: upd.Value})
		if err != nil {
			log.Fatalf("\n\tEl libro consultado no existe: %v", err)
		}
		fmt.Println("\n\tActualizado -> ", bookGetUpdated.String())
	}

	// Eliminar libro
	bookDel, err := c.DeleteBook(ctx, &pb.BookID{Value: r.Value})
	if err != nil {
		log.Fatalf("\n\n\nError al eliminar el libro -> ", err)
	}
	fmt.Println("\n\n4) Eliminando libro...")
	fmt.Println("\n\tLibro eliminado -> ", bookDel.String())

	// Obtener libro nuevamente para comprobar si se elimino
	fmt.Println("\n\n5) Obteniendo libro eliminado...")
	bookGet, err := c.GetBook(ctx, &pb.BookID{Value: r.Value})
	if err != nil {
		fmt.Println("\n\tEl libro consultado no existe -> ", err)
	} else {
		fmt.Println("\n\tLibro consultado: ", bookGet.String())
	}

	// Leer CSV
	c.ReadCSV(ctx, &pb.File{Value: "books.csv"})
	fmt.Println("\n\nLeyendo libros desde archivo books.csv -> Servidor")

}

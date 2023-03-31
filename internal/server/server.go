package server

import (
	"context"
	"fmt"
	"github.com/Nicrii/bookstore/api"
	"github.com/Nicrii/bookstore/internal/models"
	"github.com/Nicrii/bookstore/internal/repository"
	"google.golang.org/grpc"
)

type grpcServer struct {
	BookstoreRepo repository.BookstoreRepo
	api.UnimplementedBookstoreServer
}

func NewRPCServer(repository repository.BookstoreRepo) *grpc.Server {
	srv := grpcServer{
		BookstoreRepo: repository,
	}

	gsrv := grpc.NewServer()
	api.RegisterBookstoreServer(gsrv, &srv)

	return gsrv
}

func (s *grpcServer) FindAuthorsByBook(ctx context.Context, req *api.Book) (*api.FindAuthorsResponse, error) {
	book := &models.Book{
		ID:   req.Id,
		Name: req.GetName(),
	}

	authors, err := s.BookstoreRepo.FindAuthorsByBook(book)
	if err != nil {
		return nil, fmt.Errorf("can't find authors: %w", err)
	}

	res := &api.FindAuthorsResponse{}
	data := make([]*api.Author, len(authors))

	for i, author := range authors {
		a := &api.Author{
			Id:        author.ID,
			Firstname: author.Firstname,
			Lastname:  author.Lastname,
		}

		data[i] = a
	}

	res.Authors = data

	return res, nil
}

func (s *grpcServer) FindBooksByAuthor(ctx context.Context, req *api.Author) (*api.FindBooksResponse, error) {
	author := &models.Author{
		ID:        req.Id,
		Firstname: req.GetFirstname(),
		Lastname:  req.GetLastname(),
	}

	books, err := s.BookstoreRepo.FindBooksByAuthor(author)
	if err != nil {
		return nil, fmt.Errorf("can't find books: %w", err)
	}

	res := &api.FindBooksResponse{}
	data := make([]*api.Book, len(books))

	for i, book := range books {
		b := &api.Book{
			Id:   book.ID,
			Name: book.Name,
		}

		data[i] = b
	}

	res.Books = data

	return res, nil
}

package server

import (
	"context"
	"errors"
	"github.com/Nicrii/bookstore/api"
	"github.com/Nicrii/bookstore/internal/models"
	"github.com/Nicrii/bookstore/internal/repository"
	"reflect"
	"testing"
)

type RepositorySuccessMock struct {
}

func (r RepositorySuccessMock) FindBooksByAuthor(author *models.Author) ([]*models.Book, error) {
	return []*models.Book{
		{
			ID:   1,
			Name: "случайная книга",
		},
	}, nil
}

func (r RepositorySuccessMock) FindAuthorsByBook(book *models.Book) ([]*models.Author, error) {
	return []*models.Author{
		{
			ID:        123,
			Firstname: "Имя",
			Lastname:  "Случайного автора"},
	}, nil
}

type RepositoryErrorMock struct {
}

func (r RepositoryErrorMock) FindBooksByAuthor(author *models.Author) ([]*models.Book, error) {
	return nil, errors.New("repository error")
}

func (r RepositoryErrorMock) FindAuthorsByBook(book *models.Book) ([]*models.Author, error) {
	return nil, errors.New("repository error")
}

func Test_grpcServer_FindAuthorsByBook(t *testing.T) {
	type fields struct {
		BookstoreRepo                repository.BookstoreRepo
		UnimplementedBookstoreServer api.UnimplementedBookstoreServer
	}

	type args struct {
		ctx context.Context
		req *api.Book
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.FindAuthorsResponse
		wantErr bool
	}{
		{name: "successful case",
			fields: fields{
				BookstoreRepo:                RepositorySuccessMock{},
				UnimplementedBookstoreServer: api.UnimplementedBookstoreServer{},
			},

			args: args{
				ctx: context.TODO(),
				req: &api.Book{
					Id:   1,
					Name: "случайная книга",
				},
			},

			want: &api.FindAuthorsResponse{
				Authors: []*api.Author{
					{
						Id:        123,
						Firstname: "Имя",
						Lastname:  "Случайного автора"},
				},
			},
			wantErr: false,
		},

		{name: "repository error",
			fields: fields{
				BookstoreRepo:                RepositoryErrorMock{},
				UnimplementedBookstoreServer: api.UnimplementedBookstoreServer{},
			},

			args: args{
				ctx: context.TODO(),
				req: &api.Book{
					Id:   1,
					Name: "случайная книга",
				},
			},

			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &grpcServer{
				BookstoreRepo:                tt.fields.BookstoreRepo,
				UnimplementedBookstoreServer: tt.fields.UnimplementedBookstoreServer,
			}
			got, err := s.FindAuthorsByBook(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAuthorsByBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAuthorsByBook() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grpcServer_FindBooksByAuthor(t *testing.T) {
	type fields struct {
		BookstoreRepo                repository.BookstoreRepo
		UnimplementedBookstoreServer api.UnimplementedBookstoreServer
	}
	type args struct {
		ctx context.Context
		req *api.Author
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.FindBooksResponse
		wantErr bool
	}{
		{name: "successful case",
			fields: fields{
				BookstoreRepo:                RepositorySuccessMock{},
				UnimplementedBookstoreServer: api.UnimplementedBookstoreServer{},
			},

			args: args{
				ctx: context.TODO(),
				req: &api.Author{

					Id:        123,
					Firstname: "Имя",
					Lastname:  "Случайного автора",
				},
			},

			want: &api.FindBooksResponse{
				Books: []*api.Book{
					{Id: 1,
						Name: "случайная книга",
					},
				},
			},
			wantErr: false,
		},

		{name: "repository error",
			fields: fields{
				BookstoreRepo:                RepositoryErrorMock{},
				UnimplementedBookstoreServer: api.UnimplementedBookstoreServer{},
			},

			args: args{
				ctx: context.TODO(),
				req: &api.Author{

					Id:        123,
					Firstname: "Имя",
					Lastname:  "Случайного автора",
				},
			},

			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &grpcServer{
				BookstoreRepo:                tt.fields.BookstoreRepo,
				UnimplementedBookstoreServer: tt.fields.UnimplementedBookstoreServer,
			}
			got, err := s.FindBooksByAuthor(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindBooksByAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindBooksByAuthor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

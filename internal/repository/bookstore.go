package repository

import (
	"github.com/Nicrii/bookstore/internal/models"
	"gorm.io/gorm"
)

type BookstoreRepo interface {
	FindBooksByAuthor(author *models.Author) ([]*models.Book, error)
	FindAuthorsByBook(book *models.Book) ([]*models.Author, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) FindBooksByAuthor(author *models.Author) ([]*models.Book, error) {
	var books []*models.Book

	err := r.db.Table("Books").Joins("JOIN AuthorBook ON Books.ID = AuthorBook.BookID").
		Where("AuthorBook.AuthorID = ?", author.ID).Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r Repository) FindAuthorsByBook(book *models.Book) ([]*models.Author, error) {
	var authors []*models.Author

	err := r.db.Table("Authors").Joins("JOIN AuthorBook ON Authors.ID = AuthorBook.AuthorID").
		Where("AuthorBook.BookID = ?", book.ID).Find(&authors).Error
	if err != nil {
		return nil, err
	}

	return authors, nil
}

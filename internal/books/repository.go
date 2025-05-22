package books

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetBook(id uint) (*Book, error)
	CreateBook(book *Book) error
	UpdateBook(book *Book) error
	DeleteBook(id uint) error
	ListBooks() ([]Book, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db}
}

func (r *repo) GetBook(id uint) (*Book, error) {
	var book Book
	if err := r.db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *repo) CreateBook(book *Book) error {
	println("book created# :====>", book.Title, book.Author, book.Year, book.Publisher)
	return r.db.Create(book).Error
}

func (r *repo) UpdateBook(book *Book) error {
	return r.db.Save(book).Error
}

func (r *repo) DeleteBook(id uint) error {
	return r.db.Delete(&Book{}, id).Error
}

func (r *repo) ListBooks() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	return books, err
}

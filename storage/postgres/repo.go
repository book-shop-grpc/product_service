package postgres

import (
	"context"

	"github.com/book-shop-grpc/product_service/genproto/product_service"
)

type AuthorRepoI interface {
	CreateAuthor(ctx context.Context, author *product_service.Author) (*product_service.Author, error)
	GetAuthor(ctx context.Context, authorID string) (*product_service.Author, error)
	UpdateAuthor(ctx context.Context, author *product_service.Author) (*product_service.Author, error)
	DeleteAuthor(ctx context.Context, authorID string) error
	ListAuthors(ctx context.Context) (*product_service.AuthorListResponse, error)
}

type BookRepoI interface {
	CreateBook(ctx context.Context, book *product_service.Book) (*product_service.Book, error)
	GetBook(ctx context.Context, bookID string) (*product_service.Book, error)
	UpdateBook(ctx context.Context, book *product_service.Book) (*product_service.Book, error)
	DeleteBook(ctx context.Context, bookID string) error
	ListBooks(ctx context.Context) (*product_service.BookListResponse, error)
}

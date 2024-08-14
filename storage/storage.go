package storage

import (
	"github.com/book-shop-grpc/product_service/storage/postgres"
	"github.com/jackc/pgx/v5"
)

type StorageI interface {
	GetBookRepo()postgres.BookRepoI
	GetAuthorRepo()postgres.AuthorRepoI
}

type storage struct{
	bookRepo postgres.BookRepoI
	authorRepo postgres.AuthorRepoI
}

func NewStorage(db *pgx.Conn)StorageI{
	return &storage{
		bookRepo: postgres.NewBookRepo(db),
		authorRepo: postgres.NewAuthorRepo(db),
	}
}

func(s *storage)GetBookRepo()postgres.BookRepoI{
	return s.bookRepo
}
func(s *storage)GetAuthorRepo()postgres.AuthorRepoI{
	return s.authorRepo
}
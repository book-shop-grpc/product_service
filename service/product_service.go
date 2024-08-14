package service

import (
	"context"

	"github.com/book-shop-grpc/product_service/genproto/product_service"
	"github.com/book-shop-grpc/product_service/storage"
)

type ProductService struct {
	storage storage.StorageI
	product_service.UnimplementedProductServiceServer
}

func NewProductService(storage storage.StorageI) *ProductService {
	return &ProductService{storage: storage}
}

// Author CRUD operations
func (s *ProductService) CreateAuthor(ctx context.Context, req *product_service.CreateAuthorRequest) (*product_service.Author, error) {
	return nil, nil
}

func (s *ProductService) GetAuthor(ctx context.Context, req *product_service.GetAuthorRequest) (*product_service.Author, error) {
	return nil, nil
}

func (s *ProductService) UpdateAuthor(ctx context.Context, req *product_service.UpdateAuthorRequest) (*product_service.Author, error) {
	return nil, nil
}

func (s *ProductService) DeleteAuthor(ctx context.Context, req *product_service.DeleteAuthorRequest) (*product_service.Empty, error) {
	return nil, nil
}

func (s *ProductService) ListAuthors(ctx context.Context, req *product_service.Empty) (*product_service.AuthorListResponse, error) {
	return nil, nil
}


// Book CRUD operations
func (s *ProductService) CreateBook(ctx context.Context, req *product_service.CreateBookRequest) (*product_service.Book, error) {
	return nil, nil
}

func (s *ProductService) GetBook(ctx context.Context, req *product_service.GetBookRequest) (*product_service.Book, error) {
	return nil, nil
}

func (s *ProductService) UpdateBook(ctx context.Context, req *product_service.UpdateBookRequest) (*product_service.Book, error) {
	return nil, nil
}

func (s *ProductService) DeleteBook(ctx context.Context, req *product_service.DeleteBookRequest) (*product_service.Empty, error) {
	return nil, nil
}

func (s *ProductService) ListBooks(ctx context.Context, req *product_service.Empty) (*product_service.BookListResponse, error) {
	return nil, nil
}

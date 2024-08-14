package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/book-shop-grpc/product_service/genproto/product_service" // Adjust the import path
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type BookRepo struct {
	db *pgx.Conn // or your database connection type
}

func NewBookRepo(db *pgx.Conn) *BookRepo {
	return &BookRepo{db: db}
}

func (r *BookRepo) CreateBook(ctx context.Context, book *product_service.Book) (*product_service.Book, error) {
	log.Println("req here")
	book.BookId = uuid.New().String()
	query := `
		INSERT INTO 
			books(
				book_id, 
				book_name, 
				price, 
				created_at, 
				quantity, 
				is_sail, 
				author_id
			) 
	        VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(
		ctx, query,
		book.BookId,
		book.BookName,
		book.Price,
		book.CreatedAt,
		book.Quantity,
		book.IsSail,
		book.AuthorId,
	)
	if err != nil {
		log.Printf("Error creating book: %v", err)
		return nil, err
	}
	return book, nil
}

func (r *BookRepo) ListBooks(ctx context.Context) (*product_service.BookListResponse, error) {
	query := `
	SELECT 
		book_id, 
		book_name, 
		price, 
		created_at, 
		quantity, 
		is_sail, 
		author_id 
	FROM 
		books`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error listing books: %v", err)
		return nil, err
	}
	defer rows.Close()

	var books []*product_service.Book
	for rows.Next() {
		book := &product_service.Book{}
		err := rows.Scan(
			&book.BookId,
			&book.BookName,
			&book.Price,
			&book.CreatedAt,
			&book.Quantity,
			&book.IsSail,
			&book.AuthorId,
		)
		if err != nil {
			log.Printf("Error scanning book row: %v", err)
			return nil, err
		}
		books = append(books, book)
	}

	return &product_service.BookListResponse{Books: books}, nil
}

func (r *BookRepo) GetBook(ctx context.Context, bookID string) (*product_service.Book, error) {
	query := `
		SELECT 
			book_id, 
			book_name, 
			price, 
			created_at, 
			quantity, 
			is_sail, 
			author_id 
		FROM 
			books 
		WHERE 
			book_id = $1`
	row := r.db.QueryRow(ctx, query, bookID)

	book := &product_service.Book{}
	err := row.Scan(
		&book.BookId,
		&book.BookName,
		&book.Price,
		&book.CreatedAt,
		&book.Quantity,
		&book.IsSail,
		&book.AuthorId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Book not found: %v", bookID)
			return nil, nil
		}
		log.Printf("Error getting book: %v", err)
		return nil, err
	}
	return book, nil
}

func (r *BookRepo) UpdateBook(ctx context.Context, book *product_service.Book) (*product_service.Book, error) {
	query := `
	UPDATE 
		books 
	SET 
		book_name = $1, 
		price = $2, 
		created_at = $3, 
		quantity = $4, 
		is_sail = $5, 
		author_id = $6 
	WHERE book_id = $7`
	_, err := r.db.Exec(
		ctx, query,
		book.BookName,
		book.Price,
		book.CreatedAt,
		book.Quantity,
		book.IsSail,
		book.AuthorId,
		book.BookId,
	)
	if err != nil {
		log.Printf("Error updating book: %v", err)
		return nil, err
	}
	return book, nil
}

func (r *BookRepo) DeleteBook(ctx context.Context, bookID string) error {
	query := `
	DELETE FROM 
		books 
	WHERE 
		book_id = $1`
	_, err := r.db.Exec(ctx, query, bookID)
	if err != nil {
		log.Printf("Error deleting book: %v", err)
		return err
	}
	return nil
}

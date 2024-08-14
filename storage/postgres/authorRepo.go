package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/book-shop-grpc/product_service/genproto/product_service"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// AuthorRepo is the implementation of AuthorRepoI interface
type AuthorRepo struct {
	db *pgx.Conn
}

func NewAuthorRepo(db *pgx.Conn) *AuthorRepo {
	return &AuthorRepo{db: db}
}

func (r *AuthorRepo) CreateAuthor(ctx context.Context, author *product_service.Author) (*product_service.Author, error) {
	author.AuthorId = uuid.New().String()
	query := `INSERT INTO
				 authors(
				 	author_id, 
					full_name, 
					date_of_birth, 
					date_of_death
				) 
	          VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(
		ctx, query,
		author.AuthorId,
		author.FullName,
		author.DateOfBirth,
		author.DateOfDeath,
	)
	if err != nil {
		log.Printf("Error creating author: %v", err)
		return nil, err
	}
	return author, nil
}

func (r *AuthorRepo) ListAuthors(ctx context.Context) (*product_service.AuthorListResponse, error) {
	query := `
		SELECT 
			author_id, 
			full_name, 
			date_of_birth, 
			date_of_death 
		FROM 
			authors
		`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error listing authors: %v", err)
		return nil, err
	}
	defer rows.Close()

	var authors []*product_service.Author
	for rows.Next() {
		author := &product_service.Author{}
		err := rows.Scan(&author.AuthorId, &author.FullName, &author.DateOfBirth, &author.DateOfDeath)
		if err != nil {
			log.Printf("Error scanning author row: %v", err)
			return nil, err
		}
		authors = append(authors, author)
	}

	return &product_service.AuthorListResponse{Authors: authors}, nil
}

func (r *AuthorRepo) GetAuthor(ctx context.Context, authorID string) (*product_service.Author, error) {
	query := `  SELECT 
					author_id, 
					full_name, 
					date_of_birth, 
					date_of_death 
				FROM 
					authors 
				WHERE 
					author_id = $1`
	row := r.db.QueryRow(ctx, query, authorID)

	author := &product_service.Author{}
	err := row.Scan(
		&author.AuthorId,
		&author.FullName,
		&author.DateOfBirth,
		&author.DateOfDeath,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Author not found: %v", authorID)
			return nil, nil
		}
		log.Printf("Error getting author: %v", err)
		return nil, err
	}
	return author, nil
}

func (r *AuthorRepo) UpdateAuthor(ctx context.Context, author *product_service.Author) (*product_service.Author, error) {
	query := `	UPDATE 
					authors 
				SET 
					full_name = $1, 
					date_of_birth = $2, 
					date_of_death = $3 
				WHERE 
					author_id = $4`
	_, err := r.db.Exec(ctx, query, author.FullName, author.DateOfBirth, author.DateOfDeath, author.AuthorId)
	if err != nil {
		log.Printf("Error updating author: %v", err)
		return nil, err
	}
	return author, nil
}

func (r *AuthorRepo) DeleteAuthor(ctx context.Context, authorID string) error {
	query := `
		DELETE FROM 
			authors 
		WHERE 
			author_id = $1`
	_, err := r.db.Exec(ctx, query, authorID)
	if err != nil {
		log.Printf("Error deleting author: %v", err)
		return err
	}
	return nil
}

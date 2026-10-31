package repositories

import "database/sql"

type Repository struct {
	db *sql.DB
}
func NewRepository(database *sql.DB) *Repository {
	return &Repository{db: database}
}
func (db *Repository) FindAll() {}

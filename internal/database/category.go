package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name, description string) (Category, error) {
	ID := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO category (id, name, description) VALUES ($1, $2, $3)", ID, name, description)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: ID, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (c *Category) FindByCourseID(courseID string) (*Category, error) {
	row := c.db.QueryRow(
		`SELECT A.id, A.name, A.description FROM category A
		 INNER JOIN course B 
		 ON A.id = B.category_id
		 WHERE B.id = $1
		`, courseID,
	)
	var category Category
	err := row.Scan(
		&category.ID,
		&category.Name,
		&category.Description,
	)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

package repository

import (
	"database/sql"

	"github.com/budimansol/pairproject/internal/model"
)

type MenuRepository struct {
	db *sql.DB
}

func NewMenuRepository(db *sql.DB) *MenuRepository {
	return &MenuRepository{db: db}
}

func (r *MenuRepository) Create(menu model.Menu) error {
	query := `
		INSERT INTO menus (name, category, price, stock)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.Exec(query, menu.Name, menu.Category, menu.Price, menu.Stock)
	return err
}

func (r *MenuRepository) GetAll() ([]model.Menu, error) {
	rows, err := r.db.Query(`SELECT id, name, category, price, stock, created_at FROM menus ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menus []model.Menu
	for rows.Next() {
		var m model.Menu
		if err := rows.Scan(&m.ID, &m.Name, &m.Category, &m.Price, &m.Stock, &m.CreatedAt); err != nil {
			return nil, err
		}
		menus = append(menus, m)
	}
	return menus, nil
}

func (r *MenuRepository) Update(menu model.Menu) error {
	query := `
		UPDATE menus
		SET name = $1, category = $2, price = $3, stock = $4
		WHERE id = $5
	`
	_, err := r.db.Exec(query, menu.Name, menu.Category, menu.Price, menu.Stock, menu.ID)
	return err
}

func (r *MenuRepository) Delete(id int) error {
	query := `DELETE FROM menus WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
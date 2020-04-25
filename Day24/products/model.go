package main

import (
	"database/sql"
	"errors"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *Product) getProduct(db *sql.DB) error {
    stmt, err := db.Prepare("SELECT name, price FROM products WHERE id=$1")
	if err != nil {
		return err
	}
    defer stmt.Close()

	err = stmt.QueryRow(p.ID).Scan(&p.Name, &p.Price)
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) updateProduct(db *sql.DB) error {
    stmt, err := db.Prepare("UPDATE products SET name=$1, price=$2 WHERE id=$3")
	if err != nil {
		return err
    }

	_, err = stmt.Exec(p.Name, p.Price, p.ID)
	if err != nil {
		return err
    }
    return nil
}

func (p *Product) deleteProduct(db *sql.DB) error {
	return errors.New("NotImplemented")
}

func (p *Product) createProduct(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO products(name, price) VALUES($1, $2)")
	if err != nil {
		return err
    }

	res, err := stmt.Exec(p.Name, p.Price)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
    }

	p.ID = int(id)
	return nil
}

func getProducts(db *sql.DB, start, count int) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products LIMIT $1 OFFSET $2", count, start)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

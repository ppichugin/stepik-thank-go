package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// начало решения

// SQLMap представляет карту, которая хранится в SQL-базе данных
type SQLMap struct {
	key  string
	val  any
	conn *sql.DB
	stmt map[string]*sql.Stmt
}

// NewSQLMap создает новую SQL-карту в указанной базе
func NewSQLMap(db *sql.DB) (*SQLMap, error) {
	stmt, err := db.Prepare(`create table if not exists map(key text primary key, val blob)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return nil, err
	}
	setStmt, _ := db.Prepare(`insert into map(key, val) values (?, ?) on conflict (key) do update set val = excluded.val`)
	delStmt, _ := db.Prepare(`delete from map where key = ?`)
	getStmt, _ := db.Prepare(`select val from map where key = ?`)
	stmts := map[string]*sql.Stmt{
		"set":    setStmt,
		"delete": delStmt,
		"get":    getStmt,
	}

	return &SQLMap{conn: db, stmt: stmts}, nil
}

// Get возвращает значение для указанного ключа.
// Если такого ключа нет - возвращает ошибку sql.ErrNoRows.
func (m *SQLMap) Get(key string) (any, error) {
	stmt := m.stmt["get"]
	row := stmt.QueryRow(key)
	err := row.Scan(&m.val)
	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	} else if err != nil {
		return nil, err
	}
	return m.val, nil
}

// Set устанавливает значение для указанного ключа.
// Если такой ключ уже есть - затирает старое значение (это не считается ошибкой).
func (m *SQLMap) Set(key string, val any) error {
	stmt := m.stmt["set"]
	_, err := stmt.Exec(key, val)
	if err != nil {
		return err
	}
	return nil
}

// SetItems устанавливает значения указанных ключей.
func (m *SQLMap) SetItems(items map[string]any) error {
	stmt := m.stmt["set"]
	tx, err := m.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	txStmt := tx.Stmt(stmt)
	for k, v := range items {
		_, err = txStmt.Exec(k, v)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

// Delete удаляет запись карты с указанным ключом.
// Если такого ключа нет - ничего не делает (это не считается ошибкой).
func (m *SQLMap) Delete(key string) error {
	stmt := m.stmt["delete"]
	_, err := stmt.Exec(key)
	if err != nil {
		return err
	}
	return nil
}

// Close освобождает ресурсы, занятые картой в базе.
func (m *SQLMap) Close() error {
	if m.stmt == nil {
		return nil
	}
	for _, stmt := range m.stmt {
		err := stmt.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// конец решения

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	m, err := NewSQLMap(db)
	if err != nil {
		panic(err)
	}
	defer m.Close()

	m.Set("name", "Alice")

	items := map[string]any{
		"name": "Bob",
		"age":  42,
	}
	m.SetItems(items)

	name, err := m.Get("name")
	fmt.Printf("name = %v, err = %v\n", name, err)
	// name = Bob, err = <nil>

	age, err := m.Get("age")
	fmt.Printf("age = %v, err = %v\n", age, err)
	// age = 42, err = <nil>
}

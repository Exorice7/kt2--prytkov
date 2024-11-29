package database

import (
	"AnimalsBD/models"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB(dbUser, dbPassword, dbName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", dbUser, dbPassword, dbName))
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к БД: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка при проверке соединения с БД: %w", err)
	}
	return db, nil
}

func CreateTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS animals (
			id INT AUTO_INCREMENT PRIMARY KEY,
			type VARCHAR(255),
			sound VARCHAR(255),
			move VARCHAR(255),
			age INT
		)
	`)
	if err != nil {
		return fmt.Errorf("ошибка создания таблицы: %w", err)
	}
	return nil
}

func InsertAnimal(db *sql.DB, animalType string, animal models.Animal) error {
	_, err := db.Exec(`
		INSERT INTO animals (type, sound, move, age) VALUES (?, ?, ?, ?)
	`, animalType, animal.Sound(), animal.Move(), animal.Age())
	if err != nil {
		return fmt.Errorf("ошибка вставки данных в БД: %w", err)
	}
	return nil
}

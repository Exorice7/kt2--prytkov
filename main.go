package main

import (
	"AnimalsBD/database"
	"AnimalsBD/models"
	"database/sql"
	"fmt"
)

func inputAnimal() (string, models.Animal, error) {
	var animalType string
	var age int

	fmt.Print("Введите тип животного с большой буквы (Зебра, Тигр, Панда): ")
	fmt.Scan(&animalType)
	fmt.Print("Введите возраст животного: ")
	fmt.Scan(&age)

	var animal models.Animal
	switch animalType {
	case "Зебра":
		animal = &models.Zebra{AgeValue: age}
	case "Тигр":
		animal = &models.Tiger{AgeValue: age}
	case "Панда":
		animal = &models.Panda{AgeValue: age}
	default:
		return "", nil, fmt.Errorf("неизвестный тип животного: %s", animalType)
	}

	return animalType, animal, nil
}

func insertAnimals(db *sql.DB, count int) error {
	for i := 0; i < count; i++ {
		fmt.Printf("Введите данные для животного %d:\n", i+1)
		animalType, animal, err := inputAnimal()
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		insertErr := database.InsertAnimal(db, animalType, animal)
		if insertErr != nil {
			if err := errorhandler.HandleInsertError(db, animalType, animal, insertErr); err != nil {
				fmt.Println("Ошибка при обработке данных:", err)
				return err
			}
		}
	}
	return nil
}

func displayAnimals(db *sql.DB) error {
	fmt.Println("\nИнформация о животных в БД:")
	rows, err := db.Query("SELECT * FROM animals")
	if err != nil {
		return fmt.Errorf("ошибка при запросе: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var animalType, sound, move string
		var age int
		if err := rows.Scan(&id, &animalType, &sound, &move, &age); err != nil {
			return fmt.Errorf("ошибка при чтении строки: %w", err)
		}
		fmt.Printf("ID: %d, Тип: %s, Звук: %s, Движение: %s, Возраст: %d\n", id, animalType, sound, move, age)
	}
	return nil
}

func main() {
	dbUser, dbPassword, dbName := "your_db_user", "your_db_password", "your_db_name"

	db, err := database.ConnectToDB(dbUser, dbPassword, dbName)
	if err != nil {
		fmt.Println("Ошибка при подключении к БД:", err)
		return
	}
	defer db.Close()

	if err := database.CreateTable(db); err != nil {
		fmt.Println("Ошибка при создании таблицы:", err)
		return
	}

	var count int
	fmt.Print("Сколько животных вы хотите добавить? ")
	fmt.Scan(&count)

	if err := insertAnimals(db, count); err != nil {
		fmt.Println("Ошибка при добавлении животных:", err)
		return
	}

	if err := displayAnimals(db); err != nil {
		fmt.Println("Ошибка при отображении данных:", err)
	}
}

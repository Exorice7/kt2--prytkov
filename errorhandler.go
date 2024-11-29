package errorhandler

import (
	"AnimalsBD/database"
	"AnimalsBD/models"
	"database/sql"
	"fmt"
)

func HandleInsertError(db *sql.DB, animalType string, animal models.Animal, err error) error {
	if err != nil {
		fmt.Printf("ошибка вставки данных для типа %s: %v\n", animalType, err)

		fmt.Println("Повторная попытка записи данных...")
		retryErr := database.InsertAnimal(db, animalType, animal)
		if retryErr != nil {
			return fmt.Errorf("не удалось записать данные для типа %s после повторной попытки: %w", animalType, retryErr)
		}
	}
	return nil
}

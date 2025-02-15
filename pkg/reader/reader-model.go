package reader

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/property-lang/go-property-lang/contracts"
	"os"
)

func ReadModel(path string) (contracts.ModelContract, error) {
	var model contracts.ModelContract

	// Проверяем, существует ли файл
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return model, errors.New("файл не найден")
	}

	// Читаем файл
	file, err := os.ReadFile(path)
	if err != nil {
		return model, fmt.Errorf("ошибка при чтении файла: %v", err)
	}

	// Парсим JSON в структуру
	if err := json.Unmarshal(file, &model); err != nil {
		return model, fmt.Errorf("ошибка при парсинге JSON: %v", err)
	}

	return model, nil
}

package reader

import (
	"errors"
	"fmt"
	"github.com/property-lang/go-property-lang/contracts"
	"os"
	"path/filepath"
)

func ReadDirectory(path string) ([]contracts.ModelContract, error) {
	var models []contracts.ModelContract

	// Проверяем, существует ли директория
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("директория не найдена")
	}

	// Читаем содержимое директории
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении директории: %v", err)
	}

	// Проходим по всем файлам в директории
	for _, file := range files {
		// Проверяем, что это JSON файл
		if filepath.Ext(file.Name()) == ".json" {
			// Получаем полный путь к файлу
			fullPath := filepath.Join(path, file.Name())

			// Читаем модель из файла
			model, err := ReadModel(fullPath)
			if err != nil {
				return nil, fmt.Errorf("ошибка при чтении файла %s: %v", file.Name(), err)
			}

			// Добавляем модель в массив
			models = append(models, model)
		}
	}

	return models, nil
}

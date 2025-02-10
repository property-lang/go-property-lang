package validation

import (
	"errors"
	"fmt"
	"go-property-lang/contracts"
)

func PropValidate(p contracts.PropertyContract, value interface{}) error {
	switch p.Type {

	case "string":
		return ValidateString(p, value)

	case "int":
		return ValidateInt(p, value)

	case "float":
		return ValidateFloat(p, value)

	case "select":
		return ValidateSelect(p, value)

	default:
		return errors.New("неподдерживаемый тип: " + p.Type)
	}
}

// Валидация для строк
func ValidateString(p contracts.PropertyContract, value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("значение должно быть строкой")
	}

	// Проверяем, что строка не пустая, если поле обязательное
	if p.Required && str == "" {
		return errors.New("обязательное поле не может быть пустым")
	}

	// Проверка минимальной длины строки
	if len(str) < p.Min {
		return fmt.Errorf("длина строки должна быть не менее %d символов", p.Min)
	}

	// Проверка максимальной длины строки
	if p.Max > 0 && len(str) > p.Max {
		return fmt.Errorf("длина строки должна быть не более %d символов", p.Max)
	}
	return nil
}

func ValidateInt(p contracts.PropertyContract, value interface{}) error {
	num, ok := value.(int)
	if !ok {
		return errors.New("значение должно быть целым числом")
	}

	// Проверка минимального значения
	if num < p.Min {
		return fmt.Errorf("число должно быть не меньше %d", p.Min)
	}

	// Проверка максимального значения
	if p.Max > 0 && num > p.Max {
		return fmt.Errorf("число должно быть не больше %d", p.Max)
	}
	return nil
}

func ValidateFloat(p contracts.PropertyContract, value interface{}) error {
	floatValue, ok := value.(float64) // Используем float64 из-за природы `interface{}`
	if !ok {
		return errors.New("значение должно быть числом с плавающей точкой")
	}

	// Проверка минимального значения
	if floatValue < float64(p.Min) {
		return fmt.Errorf("число должно быть не меньше %d", p.Min)
	}

	// Проверка максимального значения
	if p.Max > 0 && floatValue > float64(p.Max) {
		return fmt.Errorf("число должно быть не больше %d", p.Max)
	}
	return nil
}

func ValidateSelect(p contracts.PropertyContract, value interface{}) error {
	selectedOption, ok := value.(string)
	if !ok {
		return errors.New("значение должно быть строкой")
	}

	// Проверяем, что поле является обязательным
	if p.Required && selectedOption == "" {
		return errors.New("обязательное поле не может быть пустым")
	}

	// Проверка, что выбранное значение из списка допустимых
	for _, option := range p.Options {
		if option.Value == selectedOption {
			// Найдено допустимое значение
			return nil
		}
	}
	return errors.New("значение должно быть одним из доступных вариантов")
}

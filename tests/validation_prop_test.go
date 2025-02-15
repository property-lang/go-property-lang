package tests

import (
	"github.com/property-lang/go-property-lang/contracts"
	"github.com/property-lang/go-property-lang/pkg/validation"
	"testing"
)

func TestValidateString(t *testing.T) {
	// Валидное значение
	prop := contracts.PropertyContract{Type: "string", Min: 3, Max: 10, Required: true}
	err := validation.PropValidate(prop, "hello")
	if err != nil {
		t.Errorf("ожидалось, что ошибка не вернётся, но получили: %s", err)
	}

	// Проверка на слишком короткую строку
	err = validation.PropValidate(prop, 2)
	if err == nil || err.Error() != "значение должно быть строкой" {
		t.Errorf("ожидалось, что вернётся ошибка: %s", err)
	}

	// Проверка на слишком короткую строку
	err = validation.PropValidate(prop, "hi")
	if err == nil || err.Error() != "длина строки должна быть не менее 3 символов" {
		t.Errorf("ожидалось, что вернётся ошибка: длина строки должна быть не менее 3 символов, но ошибка: %s", err)
	}

	// Проверка на слишком длинную строку
	err = validation.PropValidate(prop, "verylongstringvalue")
	if err == nil || err.Error() != "длина строки должна быть не более 10 символов" {
		t.Errorf("ожидалось, что вернётся ошибка: длина строки должна быть не более 10 символов, но ошибка: %s", err)
	}
}

func TestValidateInt(t *testing.T) {
	prop := contracts.PropertyContract{Type: "int", Min: 10, Max: 50}

	// Валидное значение
	err := validation.PropValidate(prop, 25)
	if err != nil {
		t.Errorf("ожидалось, что ошибка не вернётся, но получили: %s", err)
	}

	// Слишком маленькое значение
	err = validation.PropValidate(prop, 5)
	if err == nil || err.Error() != "число должно быть не меньше 10" {
		t.Errorf("ожидалось, что вернётся ошибка: число должно быть не меньше 10, но ошибка: %s", err)
	}

	// Слишком большое значение
	err = validation.PropValidate(prop, 75)
	if err == nil || err.Error() != "число должно быть не больше 50" {
		t.Errorf("ожидалось, что вернётся ошибка: число должно быть не больше 50, но ошибка: %s", err)
	}
}

func TestValidateFloat(t *testing.T) {
	prop := contracts.PropertyContract{Type: "float", Min: 0, Max: 100}

	// Валидное значение
	err := validation.PropValidate(prop, 50.5)
	if err != nil {
		t.Errorf("ожидалось, что ошибка не вернётся, но получили: %s", err)
	}

	// Слишком маленькое значение
	err = validation.PropValidate(prop, -1.0)
	if err == nil || err.Error() != "число должно быть не меньше 0" {
		t.Errorf("ожидалось, что вернётся ошибка: число должно быть не меньше 0, но ошибка: %s", err)
	}

	// Слишком большое значение
	err = validation.PropValidate(prop, 150.0)
	if err == nil || err.Error() != "число должно быть не больше 100" {
		t.Errorf("ожидалось, что вернётся ошибка: число должно быть не больше 100, но ошибка: %s", err)
	}
}

func TestValidateSelect(t *testing.T) {
	prop := contracts.PropertyContract{
		Type: "select",
		Options: []contracts.OptionPropertyContract{
			{Label: "Option 1", Value: "opt1"},
			{Label: "Option 2", Value: "opt2"},
		},
		Required: true,
	}

	// Валидное значение
	err := validation.PropValidate(prop, "opt1")
	if err != nil {
		t.Errorf("ожидалось, что ошибка не вернётся, но получили: %s", err)
	}

	// Некорректное значение
	err = validation.PropValidate(prop, "invalid")
	if err == nil || err.Error() != "значение должно быть одним из доступных вариантов" {
		t.Errorf("ожидалось, что вернётся ошибка: значение должно быть одним из доступных вариантов, но ошибка: %s", err)
	}

	// Пустое значение, если поле обязательное
	err = validation.PropValidate(prop, "")
	if err == nil || err.Error() != "обязательное поле не может быть пустым" {
		t.Errorf("ожидалось, что вернётся ошибка: обязательное поле не может быть пустым, но ошибка: %s", err)
	}
}

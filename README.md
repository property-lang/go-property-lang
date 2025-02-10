 

# go-property-lang

### [Концепция](https://github.com/property-lang/conception-property-lang)   


# Подключение

    go get https://github.com/property-lang/go-property-lang


1) Создаем папку для моделей, например proplangs
2) Добавляем модель например  userExample.proplang.json  
3) Подгружаем модель


    model, err := reader.ReadModel("./proplangs/userExample.proplang.json")
    if err != nil {
    fmt.Println(err)
    return
    }
      
4) Валидируем данные по свойству


	err = validation.ModelPropKeyValidate(model, "name", "t")
	if err != nil {
		fmt.Println(err)
	}


5) Валидируем данные по тегу модели


	err = validation.ModelValidateByTag(model, "name", map[string]interface{}{
		"login":  "asfsaf",
		"age": 6346436,
	})
	if err != nil {
		fmt.Println(err)
	}


 

 

 

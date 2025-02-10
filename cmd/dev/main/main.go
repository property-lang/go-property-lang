package main

import (
	"fmt"
	"go-property-lang/pkg/reader"
)

func main() {

	model, err := reader.ReadModel("./proplangs/userExample.proplang.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(model)
	fmt.Println(model.Model)
	fmt.Println(model)
	fmt.Println(model.Properties[0].GetFullName())

}

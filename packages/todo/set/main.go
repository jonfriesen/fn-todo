package main

import (
	"fmt"

	"github.com/jonfriesen/todo"
)

func Main(args map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	value, ok := args["item"].(string)
	if !ok {
		res["body"] = "no item found"
		return res
	}

	updatedList, err := todo.Set(value)
	if err != nil {
		fmt.Println(">>>", err.Error())
		res["body"] = err.Error()
		return res
	}

	res["body"] = updatedList
	return res
}

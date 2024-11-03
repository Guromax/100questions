package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Преобразование структуры в JSON
	person := Person{Name: "John Doe", Age: 30}
	jsonData, _ := json.Marshal(person)
	fmt.Println(string(jsonData))

	// Чтение JSON в структуруvar decodedPerson Person
	json.Unmarshal(jsonData, &decodedPerson)
	fmt.Println(decodedPerson.Name, decodedPerson.Age)
}

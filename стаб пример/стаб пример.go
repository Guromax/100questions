package main

import "fmt"

type DatabaseStub struct{}

func (db *DatabaseStub) GetUserName(id int) string {
	return "Alice"
}

type Database interface {
	GetUserName(id int) string
}

func PrintUserName(db Database, id int) {
	name := db.GetUserName(id)
	fmt.Println(name)
}

func main() {
	dbStub := &DatabaseStub{}
	PrintUserName(dbStub, 1)
}

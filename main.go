package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"simple-graphql-go/graph/handler"
	"simple-graphql-go/graph/resolvers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	/**
	Get (read) todo by id
	http://localhost:8080/todo?query={todo(id:1){id,text,done}}
	http://localhost:8080/todo?query={todo(id:1){id,text,done,user{id,name}}}

	Get (read) todo list
	http://localhost:8080/todo?query={list{id,text,done}}
	http://localhost:8080/todo?query={list{id,user_id,text,done,user{id,name}}}

	Create todo
	http://localhost:8080/todo?query=mutation+_{create(text:"Todo Test",user_id:1,done:false){id,text,done}}

	Update todo by id
	http://localhost:8080/todo?query=mutation+_{update(id:2,text:"Todo Test Update",user_id:1,done:false){id,text,done}}

	Delete todo by id
	http://localhost:8080/todo?query=mutation+_{delete(id:2){id,text,done}}
	*/
	connectDB()
	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		result := handler.TodoHandler(r.URL.Query().Get("query"))
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func connectDB() {
	dsn := "root:password@tcp(localhost:3306)/simple_graphql_go?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Exec("SET SESSION group_concat_max_len = 100000000;")

	resolvers.DB = db
}

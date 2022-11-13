package resolvers

import (
	"simple-graphql-go/graph/models"

	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetTodos(params graphql.ResolveParams) (interface{}, error) {
	var todos []models.Todo
	if err := DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func GetTodoByID(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(int)
	if ok {
		var todo models.Todo
		if err := DB.First(&todo, id).Error; err != nil {
			return nil, err
		}
		return todo, nil
	}
	return nil, nil
}

func CreateTodo(params graphql.ResolveParams) (interface{}, error) {
	todo := models.Todo{
		UserID: params.Args["user_id"].(int),
		Text:   params.Args["text"].(string),
		Done:   params.Args["done"].(bool),
	}

	if err := DB.Create(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func UpdateTodoByID(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(int)
	if ok {
		var todo models.Todo
		if err := DB.First(&todo, id).Error; err != nil {
			return nil, err
		}

		todo.UserID = params.Args["user_id"].(int)
		todo.Text = params.Args["text"].(string)
		todo.Done = params.Args["done"].(bool)

		if err := DB.Save(&todo).Error; err != nil {
			return nil, err
		}
		return todo, nil
	}
	return nil, nil
}

func DeleteTodoByID(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(int)
	if ok {
		var todo models.Todo
		if err := DB.First(&todo, id).Error; err != nil {
			return nil, err
		}

		if err := DB.Delete(&todo).Error; err != nil {
			return nil, err
		}
		return todo, nil
	}
	return nil, nil
}

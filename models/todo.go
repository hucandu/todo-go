package models
import (
	forms "github.com/hucandu/todo-go/forms"
)

//todo model
type TodoModel struct {
	ID          string    `db:"id, primarykey" json:"id"`
	Data        string    `db:"data" json:"data"`
	CreatedAt   string    `db:"created_at" json:"created_at"`
	Status      string    `db:"status" json:"status"`
	CreatedBy   string    `db:"created_by" json:"created_by"`
}


func (model *TodoModel) createTodo (data forms.CreateTodo){
}

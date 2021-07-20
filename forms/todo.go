package forms

type CreateTodo struct {
	Data  string `form:"data" binding:"required"`
}
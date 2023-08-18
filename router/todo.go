package router

import (
	"net/http"
	"strconv"
	"todo-app/database"
	"todo-app/dto"
	"todo-app/middleware"
	"todo-app/model"

	"github.com/gin-gonic/gin"
)

func model2Dto(m *model.Todo) dto.TodoReturnDto {
	return dto.TodoReturnDto{ID: m.ID, Name: m.Name, IsComplete: m.IsComplete}
}

// @security ApiKeyAuth
// @summary get todos
// @tags todo
// @produce json
// @success 200 {array} dto.TodoReturnDto
// @router /api/todo [get]
func getTodos(c *gin.Context) {
	d := database.Connect()
	todos := make([]model.Todo, 0)
	d.Find(&todos)
	data := make([]dto.TodoReturnDto, 0, len(todos))
	for _, v := range todos {
		data = append(data, model2Dto(&v))
	}
	c.JSON(http.StatusOK, data)
}

// @security ApiKeyAuth
// @summary get a specific todo by id
// @tags todo
// @param id path integer true "id"
// @produce json
// @success 200 {object} dto.TodoReturnDto
// @router /api/todo/{id} [get]
func getTodoById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	d := database.Connect()
	todo := model.Todo{}
	if res := d.Where(&model.Todo{ID: id}).Take(&todo); res.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, model2Dto(&todo))
}

// @security ApiKeyAuth
// @summary create a new todo
// @tags todo
// @accept json
// @param payload body dto.TodoPostDto true "payload"
// @produce json
// @success 201 {object} dto.TodoReturnDto
// @router /api/todo [post]
func createPost(c *gin.Context) {
	todoPostDto := dto.TodoPostDto{}
	if err := c.ShouldBind(&todoPostDto); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	d := database.Connect()
	todo := model.Todo{Name: todoPostDto.Name, IsComplete: todoPostDto.IsComplete}
	d.Save(&todo)
	c.JSON(http.StatusCreated, model2Dto(&todo))
}

// @security ApiKeyAuth
// @summary update a specific todo by id, if not exist, create it
// @tags todo
// @accept json
// @param id path integer true "id"
// @param payload body dto.TodoPutDto true "payload"
// @produce json
// @success 200 {object} dto.TodoReturnDto
// @router /api/todo/{id} [put]
func putTodoById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	todoPutDto := dto.TodoPutDto{}
	if err := c.ShouldBind(&todoPutDto); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	d := database.Connect()
	todo := model.Todo{ID: id, Name: todoPutDto.Name, IsComplete: todoPutDto.IsComplete}
	d.Save(&todo)
	c.JSON(http.StatusOK, model2Dto(&todo))
}

// @security ApiKeyAuth
// @summary delete a specific todo by id
// @tags todo
// @param id path integer true "id"
// @success 204
// @router /api/todo/{id} [delete]
func deleteTodoById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	d := database.Connect()
	todo := model.Todo{ID: id}
	d.Delete(&todo)
	c.Status(http.StatusNoContent)
}

func LoadTodoRouter(e *gin.Engine) {
	router := e.Group("/api/todo").Use(middleware.AuthRequire)
	{
		router.GET("", getTodos)
		router.GET("/:id", getTodoById)
		router.POST("", createPost)
		router.PUT("/:id", putTodoById)
		router.DELETE("/:id", deleteTodoById)
	}
}

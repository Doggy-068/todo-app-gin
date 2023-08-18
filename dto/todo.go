package dto

type TodoReturnDto struct {
	ID         uint64 `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	IsComplete bool   `json:"is_complete" binding:"required"`
}

type TodoPostDto struct {
	Name       string `json:"name" binding:"required"`
	IsComplete bool   `json:"is_complete" binding:"required"`
}

type TodoPutDto struct {
	Name       string `json:"name" binding:"required"`
	IsComplete bool   `json:"is_complete" binding:"required"`
}

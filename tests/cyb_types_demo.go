package tests

type DemoResponse struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

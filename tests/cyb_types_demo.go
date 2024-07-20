package tests

type DemoRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type DemoResponse struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

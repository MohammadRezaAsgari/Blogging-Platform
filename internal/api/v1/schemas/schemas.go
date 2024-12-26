package schemas

type LoginRegisterInput struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

type CreateArticleRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type UpdateArticleRequest struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

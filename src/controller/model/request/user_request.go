package request

type UserRequest struct {
	Name     string `json:"name" binding:"required,min=4,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%^&*."`
	Age      int8   `json:"age" binding:"required,min=18,max=120"`
}

type UserUpdateRequest struct {
	Name string `json:"name" binding:"omitempty,min=4,max=50"`
	Age  int8   `json:"age" binding:"omitempty,min=18,max=120"`
}

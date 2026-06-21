package entity

type UserEntity struct {
	ID	     string `psql:"_id,omitempty"`
	Email    string `psql:"email"`
	Password string `psql:"password"`
	Name     string `psql:"name"`
	Age      int8   `psql:"age"`
}
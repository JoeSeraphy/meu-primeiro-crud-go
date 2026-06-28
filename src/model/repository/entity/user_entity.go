package entity

type UserEntity struct {
	ID       string `db:"id" json:"id,omitempty"`
	Email    string `db:"email" json:"email,omitempty"`
	Password string `db:"password" json:"password,omitempty"`
	Name     string `db:"name" json:"name,omitempty"`
	Age      int8   `db:"age" json:"age"`
}

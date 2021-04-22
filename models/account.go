package models

type Account struct {
	Id        int    `db:"id" json:"id"`
	Username  string `db:"username" json:"username"`
	Passowrd  string `db:"password" json:"password"`
	Name      string `db:"name" json:"name"`
	Create_at string `db:"create_at" json:"-"`
	Update_at string `db:"update_at" json:"-"`
}

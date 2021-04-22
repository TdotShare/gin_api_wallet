package models

type History struct {
	Id        int    `db:"id" json:"id"`
	Bank_id   int    `db:"bank_id" json:"bank_id"`
	Price     string `db:"price" json:"price"`
	Status    int    `db:"status" json:"status"`
	Create_at string `db:"create_at" json:"-"`
	Update_at string `db:"update_at" json:"-"`
}

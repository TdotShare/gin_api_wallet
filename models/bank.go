package models

type Bank struct {
	Id         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Account_id int    `db:"account_id" json:"account_id"`
	Create_at  string `db:"create_at" json:"-"`
	Update_at  string `db:"update_at" json:"-"`
}

type BankAccount struct {
	Id         int       `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	Account_id int       `db:"account_id" json:"account_id"`
	Create_at  string    `db:"create_at" json:"-"`
	Update_at  string    `db:"update_at" json:"-"`
	Account    Account   `db:"account"`
	History    []History `db:"history"`
}

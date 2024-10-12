package models

type ProductDto struct {
	Id    uint32 `db:"id"`
	Name  string `db:"name"`
	Price uint32 `db:"price"`
}

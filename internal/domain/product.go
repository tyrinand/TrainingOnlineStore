package domain

type Product struct {
	Id    uint32
	Name  string
	Price uint32
}

type PageInfo struct {
	Page     uint32
	PageSize uint32
}

package repositories

type product struct {
	ID      int
	Name    string
	Quatity int
}

type ProductRepository interface {
	GetProducts() ([]product, error)
}

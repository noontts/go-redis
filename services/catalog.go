package services

type Product struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Quatity int    `json:"Quantity"`
}

type CatalogService interface {
	GetProducts() ([]Product, error)
}

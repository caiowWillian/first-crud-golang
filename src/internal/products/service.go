package products

type Service interface {
	CreateProduct(product Product) bool
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) CreateProduct(product Product) bool {
	return true
}

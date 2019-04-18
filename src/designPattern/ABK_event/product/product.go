package product

type Product struct {
	name string
	canChanged bool
}

func NewProduct(productManager *ProductManager, name string) *Product {
	res := &Product{}
	if productManager.isPermittedCreate {
		res.name = name
		res.canChanged = true
	}
	return res
}

func (s *Product)Name() string {
    return s.name
}

func (s *Product)SetName(name string) {
    s.name = name
}

func (s *Product)Clone() Product {
    return *s
}

package product

type ProductManager struct {
	isPermittedCreate bool
}

func NewProductManager() *ProductManager {
	return &ProductManager{isPermittedCreate: false}
}

func (s *ProductManager)IsPermittedCreate() bool {
    return s.isPermittedCreate
}

func (s *ProductManager)CreateProduct(name string) *Product {
    s.isPermittedCreate = true
    p := NewProduct(s, name)
    NewProductEvent(*p, NEW_PRODUCT)
	return p
}

func (s *ProductManager)AbandonProduct(p *Product) {
    NewProductEvent(*p, DEL_PRODUCT)
    p = nil
}

func (s *ProductManager)EditProduct(p *Product, name string) {
    NewProductEvent(*p, EDIT_PRODUCT)
    p.SetName(name)
}

func (s *ProductManager)Clone(p *Product) Product {
    return p.Clone()
}
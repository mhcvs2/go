package product

type ProductManager struct {
	isPermittedCreate bool
}

func (s *ProductManager)IsPermittedCreate() bool {
    return s.isPermittedCreate
}
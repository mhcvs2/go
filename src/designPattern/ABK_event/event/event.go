package event

type ProductEventType int

const (
	NEW_PRODUCT ProductEventType = iota
	DEL_PRODUCT
	EDIT_PRODUCT
	CLONE_PRODUCT
)

type ProductEvent struct {
	source *Product
}
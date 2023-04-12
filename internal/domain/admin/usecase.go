package admin

type UseCase interface {
	Create() (*Model, error)
	FindById(id string) (*Model, error)
	FindAll() []*Model
	Update() (*Model, error)
	Delete(id string) error
}

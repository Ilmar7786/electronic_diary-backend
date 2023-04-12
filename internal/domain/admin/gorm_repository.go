package admin

type Repository interface {
	Create(user *Model) error
	GetByID(id uint) (*Model, error)
	GetByEmail(email string) (*Model, error)
	Update(user *Model) error
	Delete(user *Model) error
}

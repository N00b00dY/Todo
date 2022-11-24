package data

type DBRepository interface {
	GetAll() ([]*Todo, error)
	GetOne(id int) (*Todo, error)
	Update(todo Todo) error
	DeleteByID(id int) error
	Insert(todo Todo) (int, error)
}

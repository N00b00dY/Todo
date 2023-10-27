package data

import "db-service/dbs"

type DBRepository interface {
	GetAll() ([]*dbs.Todo, error)
	GetOne(id int) (*Todo, error)
	Update(todo Todo) error
	DeleteByID(id int) error
	Insert(todo Todo) (int, error)
}

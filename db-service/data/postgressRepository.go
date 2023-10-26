package data

import (
	"context"
	"database/sql"
	"db-service/dbs"
	"fmt"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

type PostgresRepository struct {
	Conn *sql.DB
	Todo Todo
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		Conn: db,
	}
}

// New is the function used to create an instance of the data package. It returns the type
// Model, which embeds all the types we want to be available to our application.
// func New(dbPool *sql.DB) Models {
// 	db = dbPool

// 	return Models{
// 		Todo: Todo{},
// 	}
// }

// Models is the type for this package. Note that any model that is included as a member
// in this type is available to us throughout the application, anywhere that the
// app variable is used, provided that the model is also added in the New function.
// type Models struct {
// 	Todo Todo
// }

type Models struct {
	Todo Todo
}

// Todo is the structure which holds one todo from the database.
type Todo struct {
	ID        int       `json:"id"`
	Todo      string    `json:"todo"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetAll returns a slice of all todo, sorted by last name
func (u *PostgresRepository) GetAll() ([]*dbs.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, todo, todo_active, created_at, updated_at
	from todos order by id`

	rows, err := u.Conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*Todo
	var dbsTodos []*dbs.Todo

	for rows.Next() {
		var todo Todo
		var dbsTodo dbs.Todo
		err := rows.Scan(
			&todo.ID,
			&todo.Todo,
			&todo.Active,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
		dbsTodo.ID = int32(todo.ID)
		dbsTodo.Todo = todo.Todo
		dbsTodo.Active = int32(todo.Active)

		todos = append(todos, &todo)
		dbsTodos = append(dbsTodos, &dbsTodo)
	}

	return dbsTodos, nil
}

// GetOne returns one todo by id
func (u *PostgresRepository) GetOne(id int) (*Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, todo, todo_active, created_at, updated_at from todos where id = $1`

	var todo Todo
	row := u.Conn.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&todo.ID,
		&todo.Todo,
		&todo.Active,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

// Update updates one todo in the database, using the information
// stored in the receiver u
func (u *PostgresRepository) Update(todo Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update todos set
		todo = $1,
		todo_active = $2,
		updated_at = $3
		where id = $4
	`

	_, err := u.Conn.ExecContext(ctx, stmt,
		todo.Todo,
		todo.Active,
		time.Now(),
		todo.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

// DeleteByID deletes one todo from the database, by ID
func (u *PostgresRepository) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	fmt.Println(id)
	stmt := `delete from todos where id = $1`

	_, err := u.Conn.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// Insert inserts a new todo into the database, and returns the ID of the newly inserted row
func (u *PostgresRepository) Insert(todo Todo) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID int
	stmt := `insert into todos (todo, todo_active, created_at, updated_at)
		values ($1, $2, $3, $4) returning id`

	err := u.Conn.QueryRowContext(ctx, stmt,
		todo.Todo,
		todo.Active,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

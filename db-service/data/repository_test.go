package data

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"log"
	"os"
	"testing"
	"time"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "password"
	dbname   = "todo"
	port     = "5435"
	dsn      = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=5"
)

var (
	resource *dockertest.Resource
	pool     *dockertest.Pool
	testDB   *sql.DB
	testRepo DBRepository
)

func TestMain(m *testing.M) {
	// connect to docker; fail if docker not running
	p, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("could not connect to docker; is it running? %s", err)
	}

	pool = p

	// set up our docker options, specify image and so forth

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14.5",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + dbname,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0,", HostPort: port},
			},
		},
	}

	// ge a resource (docker image)
	resource, err = pool.RunWithOptions(&opts)
	if err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("could not start resource: %s", err)
	}

	// start the image and wait until its ready
	if err := pool.Retry(func() error {
		var err error
		testDB, err = sql.Open("pgx", fmt.Sprintf(dsn, host, port, user, password, dbname))
		if err != nil {
			log.Println("Error: ", err)
			return err
		}
		return testDB.Ping()
	}); err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("could not connect to databse: %s", err)
	}

	// populate teh database with empty tables
	err = createTables()
	if err != nil {
		log.Fatalf("could not create tables: %s", err)
	}

	testRepo = &PostgresRepository{
		Conn: testDB,
		Todo: Todo{},
	}

	// run the tests
	code := m.Run()

	// clean up
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("could not purge resource: %s", err)
	}

	os.Exit(code)
}

func createTables() error {
	tableSQL, err := os.ReadFile("./testdata/todo.sql")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = testDB.Exec(string(tableSQL))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func Test_pingDb(t *testing.T) {
	err := testDB.Ping()
	if err != nil {
		t.Errorf("could not ping database: %s", err)
	}
}

func Test_PostgresRepository_Insert(t *testing.T) {
	todo := Todo{
		ID:        1,
		Todo:      "First test todo",
		Active:    int(0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	id, err := testRepo.Insert(todo)
	if err != nil {
		t.Errorf("Error inserting todo: %v", err)
	}
	if id != 1 {
		t.Errorf("Expected id to be 1, got %d", id)
	}
}

func Test_PostgresRepository_GetAll(t *testing.T) {
	oneTodo, err := testRepo.GetAll()
	if err != nil {
		t.Errorf("Error GetAll todo: %v", err)
	}

	if len(oneTodo) != 1 {
		t.Errorf("Expected 1 todo , got %d", len(oneTodo))
	}
}

func Test_PostgresRepository_GetOne(t *testing.T) {
	oneTodo, err := testRepo.GetOne(1)
	if err != nil {
		t.Errorf("Error GetOne todo: %v", err)
	}

	if oneTodo.ID != 1 {
		t.Errorf("Expected todo.ID to be 1, got %d", oneTodo.ID)
	}
}

func Test_PostgresRepository_Update(t *testing.T) {
	// a todo to test the update is needed
	todo := Todo{
		ID:        1,
		Todo:      "Update todo",
		Active:    2,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// calling the  tested function
	err := testRepo.Update(todo)
	if err != nil {
		t.Errorf("Error GetAll todo: %v", err)
	}

	NewTodo, err := testRepo.GetOne(1)
	if NewTodo.Active != 2 {
		t.Errorf("Expected todo.active to be 2, got %d", NewTodo.Active)
	}
}

func Test_PostgresRepository_DeleteByID(t *testing.T) {
	todo := Todo{
		ID:        1,
		Todo:      "First test todo",
		Active:    int(0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	id, err := testRepo.Insert(todo)
	if err != nil {
		t.Errorf("Error inserting todo: %v", err)
	}

	err = testRepo.DeleteByID(id)
	if err != nil {
		t.Errorf("Error inserting todo: %v", err)
	}
}

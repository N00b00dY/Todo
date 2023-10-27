package main

import (
	"database/sql"
	"db-service/data"
	"fmt"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"log"
	"os"
	"testing"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "password"
	dbname   = "todo"
	port     = "5436"
	dsn      = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=5"
)

var (
	resource     *dockertest.Resource
	pool         *dockertest.Pool
	testDB       *sql.DB
	testRepo     data.PostgresRepository
	testDBServer DBServer
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

	// populate the database with empty tables
	err = createTables()
	if err != nil {
		log.Fatalf("could not create tables: %s", err)
	}

	testRepo = data.PostgresRepository{
		Conn: testDB,
		Todo: data.Todo{},
	}

	testDBServer = DBServer{
		PostgresRepository: testRepo,
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
	tableSQL, err := os.ReadFile("../../data/testdata/todo.sql")
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

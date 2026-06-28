package postgre

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var (
	POSTGRES_URL = "POSTGRES_URL"
	//POSTGRES_USER_DB = "POSTGRES_USER_DB"
)

func NewPostgreConnection(
	ctx context.Context,
) (*pgx.Conn, error) {
	postgreURL := os.Getenv(POSTGRES_URL)
	//userDB := os.Getenv(POSTGRES_USER_DB)
	//postDB_URL := postgreURL + " " + userDB
	conn, err := pgx.Connect(ctx, postgreURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	defer conn.Close(ctx)

	var name string
	var email string
	var age int8
	err = conn.QueryRow(ctx, "SELECT name, email, age FROM user WHERE id=$1", 1).Scan(&name, &email, &age)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}
	fmt.Println(name, email, age)

	return conn, nil
}

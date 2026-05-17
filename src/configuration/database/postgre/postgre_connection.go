package postgre

import (
	"fmt"
	"context"
	"os"
	"github.com/jackc/pgx/v5"
)

var (
	POSTGRES_URL = "POSTGRES_URL"
	POSTGRES_USER_DB = "POSTGRES_USER_DB"
)

func NewPostgreConnection(
	ctx context.Context,
) (*pgx.Conn, error) {
		postgreURL := os.Getenv(POSTGRES_URL)
		userDB := os.Getenv(POSTGRES_USER_DB)
	conn, err := pgx.Connect(ctx, postgreURL, userDB)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	defer conn.Close(ctx)

	var name string
	var weight int64
	err = conn.QueryRow(ctx, "SELECT name, weight FROM pets WHERE id=$1", 1).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}
	fmt.Println(name, weight)

	return conn, nil
	
}

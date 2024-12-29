package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {
	host := `localhost`
	port := `5432`
	user := `postgres`
	pass := ``
	db := `trial`

	conn, err := pgx.Connect(context.Background(), "postgresql://"+user+":"+pass+"@"+host+":"+port+"/"+db+"?sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}
	return conn
}

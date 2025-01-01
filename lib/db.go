package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {
	host := `103.93.58.89`
	port := `54326`
	user := `postgres`
	pass := `123`
	db := `trial`

	conn, err := pgx.Connect(context.Background(), "postgresql://"+user+":"+pass+"@"+host+":"+port+"/"+db+"?sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}
	return conn
}

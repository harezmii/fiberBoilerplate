package ent

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
)

func Conn() (*Client, error) {
	client, err := Open("mysql", "<user>:<pass>@tcp(<host>:<port>)/<database>?parseTime=True")
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// Run the auto migration tool.
	if errCreate := client.Schema.Create(context.Background()); err != nil {
		return nil, errCreate
	}
	return client, nil
}

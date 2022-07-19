package repository

import "fiberBoilerplate/pkg/ent"

type Repo struct {

	// Database connection
	Conn ent.Client

	// Model
	Model struct{}
}

type IRepo interface {
	Create() (struct{}, error)
	Index() (struct{}, error)
	Show() (struct{}, error)
	Update() (struct{}, error)
	Delete() (struct{}, error)
}

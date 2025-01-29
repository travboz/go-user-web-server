package main

type Storage interface {
	Insert(u User) User
	Get(id int) (User, error)
	GetAll() []User
	Delete(id int) error
}

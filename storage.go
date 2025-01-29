package main

type Storage interface {
	Insert(u User) User
	Get(id int) (User, error)
	Delete(id int) error
}

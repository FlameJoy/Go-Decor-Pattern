package main

import "fmt"

type DB interface {
	Store(string) error
}

type Store struct{}

func (s *Store) Store(val string) error {
	fmt.Println("storing into DB:", val)
	return nil
}

// 3rd party lib function
type ExecuteFn func(string)

func myExecuteFn(db DB) ExecuteFn {
	return func(s string) {
		fmt.Println(s)
		db.Store(s)
	}
	// access to Db

}

func Execute(fn ExecuteFn) {
	fn("Foo foo bar")
}

package main

import "fmt"

func main() {
	s := newServer(withTLS, withMaxConn(99))
	fmt.Printf("Opts: %+v\n", s)
	store := &Store{}
	Execute(myExecuteFn(store))
}

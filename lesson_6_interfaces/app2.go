package main

import "log"

func main() {
	usr := newUser("FIO", "address", "phone")

	log.Println(usr)

}

type User interface {
	ChangeFio(newFIO string)
	ChangeAddress(newAddress string)
}

type user struct {
	FIO, Address, Phone string
}

func (usr *user) ChangeFio(newFIO string) {
	usr.FIO = newFIO
}

func (usr *user) ChangeAddress(newAddress string) {
	usr.Address = newAddress
}

func newUser(fio, address, phone string) User {
	usr := user{
		FIO:     fio,
		Address: address,
		Phone:   phone,
	}
	return &usr
}

var _ User = &user{}

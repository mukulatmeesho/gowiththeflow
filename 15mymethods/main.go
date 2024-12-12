package main

import "fmt"

type User struct {
	ID     int8
	Name   string
	Email  string
	Age    int16
	Active bool
}

func (u User) GetStatus() {
	fmt.Println("Is user active :", u.Active)
}

func (u User) newEmail() {
	u.Email = "test@dkksdkd.com"
	fmt.Println("New Email is :", u.Email)

}
func main() {
	fmt.Print("methods in Go")

	mukul := User{
		1,
		"Mukul",
		"mukul.saini@meesho.com",
		23,
		true}
	fmt.Println(mukul)
	fmt.Printf("Mukul's details are : %+v\n", mukul)
	mukul.GetStatus()
	mukul.newEmail()
	fmt.Printf("Mukul's details are : %+v\n", mukul)

}

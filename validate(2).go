package main

import "regexp"
import "fmt"

func main(){
	Email := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	Phone := regexp.MustCompile(`^[+]+62[0-9]{7,14}$`)
	Username := regexp.MustCompile(`^[a-z]{5,9}$`)
	Password := regexp.MustCompile(`^[A-Z]+[#]+[a-z0-9~!@#$%^&*-_+=\|;:'",./?]{6,200}$`)

	email := "a@a.a"
	fmt.Println(Email.MatchString(email))

	phone:= "+6221287263624"
	fmt.Println(Phone.MatchString(phone))

	username:= "admin"
	fmt.Println(Username.MatchString(username))

	password:= "A#sssssss2!#"
	fmt.Println(Password.MatchString(password))
}
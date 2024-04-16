package user_package

import (
	"fmt"
	"regexp"
)

type Users struct {
	Name  string
	Age   int
	Email string
}

func (u Users) ValidateUser() error {
	fmt.Print("Name: ")
	fmt.Scanln(&u.Name)
	fmt.Print("Age: ")
	fmt.Scanln(&u.Age)
	fmt.Print("Email: ")
	fmt.Scanln(&u.Email)

	if len(u.Name) >= 6 {
		return fmt.Errorf("length cannot be less than a 6 characters")
	} else if u.Name == "" {
		return fmt.Errorf("couldn't be empty")
	}

	if u.Age < 0 {
		return fmt.Errorf("must be greater than 0")
	} else if u.Age > 120 {
		return fmt.Errorf("couldn't be more than 120")
	}else if u.Age==0{
		return fmt.Errorf("age may not be equal to zero")
	}

	if u.Email == "" {
		return fmt.Errorf("couldn't be empty")
	} else {
		email := regexp.MustCompile(`[^@]+@[^@]+\.[^@]+`)
		if !email.MatchString(u.Email) {
			return fmt.Errorf("invalid email format")
		}
	}
	return nil
}

func VAlid() {
	var user Users
	resault:=user.ValidateUser()
	if resault!=nil{
		fmt.Println("Error: ",resault )
	}else{
		fmt.Println("No errors")
	}
}
package utils

import "testing"

func TestHashPassword(t *testing.T){
	password:="mypassword"

	hashed,err:=HashPassword(password)

	if err != nil{
		t.Errorf("failed to hash password: %v", err)
	}

	if hashed == password {
		t.Errorf("password was not hashed")
	}
}

func TestCheckPasswordHash(t *testing.T) {

	password := "mypassword"

	hashed, _ := HashPassword(password)

	valid := CheckPassword(password, hashed)

	if !valid {
		t.Errorf("password hash check failed")
	}
}

func TestCheckPasswordHashWrongPassword(t *testing.T) {

	password := "mypassword"

	hashed, _ := HashPassword(password)

	valid := CheckPassword("wrongpassword", hashed)

	if valid {
		t.Errorf("wrong password passed hash check")
	}
}
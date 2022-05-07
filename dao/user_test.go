package dao

import (
	"fmt"
	"testing"
)

func TesGetUserByPhone(t *testing.T) {
	user, _ := GetUserByPhone("15827059700")
	fmt.Println("user=", *user)
}

func TestSum(t *testing.T) {
	fmt.Println("123")
	fmt.Println("123")
	fmt.Println("123")
	fmt.Println("123")
}

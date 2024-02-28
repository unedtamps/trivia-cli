package seed

import "math/rand"

var char string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var len_char = len(char)

func random_name() string {
	name := ""
	for i := 0; i < 20; i++ {
		name += string(char[rand.Intn(len_char)])
	}
	return name
}
func random_email() string {
	email := ""
	for i := 0; i < 25; i++ {
		email += string(char[rand.Intn(len_char)])
	}
	email += "@gmail.com"
	return email
}

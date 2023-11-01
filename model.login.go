package main

type User struct {
	Username string `json:"usuario"`
	Password string `json:"senha"`
}

type Users []User

package main

type UserWithCheck struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age"  binding:"required,AgeCheck"`
	Sex  string `json:"sex"  binding:"required"`
}
type UserWithoutCheck struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

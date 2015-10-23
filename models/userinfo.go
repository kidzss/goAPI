package models

type User struct {
	Id   int
	Name string
	Age  int
	Sex  int
	Tel  string

	// Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

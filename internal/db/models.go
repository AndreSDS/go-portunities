// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

type Openning struct {
	ID       int64
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

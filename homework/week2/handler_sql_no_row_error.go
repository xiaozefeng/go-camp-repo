package week2

import (
	"database/sql"
	"encoding/json"
	"errors"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

type PersonRepository struct {
}

func (pr *PersonRepository) QueryByName(name string) ([]byte, error) {
	return nil, sql.ErrNoRows
}

var pr = &PersonRepository{}

func ListPersontByName(name string) ([]*Person, error) {
	res, err := pr.QueryByName(name)
	if err != nil && IsNoRowsErr(err) {
		return []*Person{}, nil
	}
	var persons []*Person
	// fake unmarshal
	err = json.Unmarshal(res, &persons)
	if err != nil {
		return nil, err
	}
	return persons, nil
}

func IsNoRowsErr(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

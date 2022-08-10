package db

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func RunSchema(path string) error {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	query := string(c)
	stmt, err := DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	return err
}

func CreateTables() error {
	folder := filepath.Join("db", "schemas")
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return err
	}
	for _, f := range files {
		path := filepath.Join(folder, f.Name())
		err := RunSchema(path)
		fmt.Println(path, " ", err)
		if err != nil {
			return err
		}
	}
	return nil
}

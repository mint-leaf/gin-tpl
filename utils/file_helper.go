package utils

import (
	"io/ioutil"
	"os"
)

// LoadDataFromFile load data from file
func LoadDataFromFile(filename string, marshaller func([]byte, interface{}) error, meta interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return marshaller(data, meta)
}

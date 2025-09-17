package files

import (
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {

	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return data, nil

}

func WriteFile(name string, content []byte) {

	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		file.Close()
		fmt.Println(err.Error())
	}
}

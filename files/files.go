package files

import (
	"fmt"
	"os"
)

func ReadFile() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(data))
}

func WriteFile(content string, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err.Error())
		defer file.Close()
	}
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Запись успешна")
}

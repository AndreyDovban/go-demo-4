package files

import (
	"fmt"
	"os"
)

func ReadFile() {
	fmt.Println("read file")
}

func WriteFile(name string, content string) {
	fmt.Println("write file")

	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	_, err = file.Write([]byte(content))
	if err != nil {
		file.Close()
		fmt.Println(err.Error())
	}
}

package files

import (
	"fmt"
	"os"
)

func ReadFile(url string) {
	data, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println("err")
		return
	}
	fmt.Println(data)

}

func WriteFile(content, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}

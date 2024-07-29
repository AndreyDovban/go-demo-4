package files

import (
	"fmt"
	"os"
)

func ReadFile(url string) {
	fmt.Println(url)

}

func WriteFile(content, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString(content)
	if err != nil {
		file.Close()
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
	file.Close()
}

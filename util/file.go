package util

import (
	"fmt"
	"os"
)

func WriteFile(filepath, content string) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(content)

	if err != nil {
	   panic(err)
	}
	fmt.Println("Done writing into a file")
}

func DeleteFile(filepath string) {
	err := os.Remove(filepath)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("File deleted successfully")
    }
}
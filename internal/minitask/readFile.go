package minitask

import (
	"fmt"
	"io"
	"os"
)

func ReadFile() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("continue...")
		}
	}()

	file, err := os.Open(".\\internal\\model\\biodata.go")
	// file, err := os.Open(".\\internal\\model\\bioaata.go")
	// file, err := os.Open(".\\internal")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	res, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", res)
}

package save

import (
	"fmt"
	"os"
)

const (
	path = "./price.txt"
)

func WritePrice(p string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("error opening %s: %s", path, err)
		return err
	}
	_, err = fmt.Fprintln(f, p)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return err
	}
	return nil
}

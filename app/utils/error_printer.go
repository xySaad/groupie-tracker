package utils

import (
	"fmt"
	"os"
)

func PrintError(err ...any) (int, error) {
	return fmt.Fprintln(os.Stderr, err...)
}

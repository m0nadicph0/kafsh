package util

import (
	"fmt"
	"os"
)

func Success(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(os.Stdout, "✅ "+format, a...)
}

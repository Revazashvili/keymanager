package keymanager

import (
	"fmt"
	"os"
)

func CheckErr(msg interface{}) {
	if msg != nil {
		fmt.Fprintln(os.Stderr, "Error:", msg)
		os.Exit(1)
	}
}

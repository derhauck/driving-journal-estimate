package main

import (
	"fmt"
)

func main() {
	err := router.Run(":8080")
	if err != nil {
		_ = fmt.Sprintf(err.Error())
	}
}

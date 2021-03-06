package main

import (
	"fmt"
	"os"

	"github.com/alexkappa/service-template-grpc/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(128)
	}
}

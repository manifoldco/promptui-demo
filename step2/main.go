package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Prompt{
		Label: "Search",
		Validate: func(input string) error {
			if len(input) < 3 {
				return errors.New("Search term must have at least 3 characters")
			}
			return nil
		},
	}

	keyword, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Search for %q\n", keyword)
}

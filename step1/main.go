package main

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Prompt{
		Label: "Search",
	}

	keyword, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Search for %q\n", keyword)
}

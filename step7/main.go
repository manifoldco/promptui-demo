package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

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

	const apiURL = "http://www.recipepuppy.com/api/"

	query := apiURL + "?q=" + keyword

	client := &http.Client{Timeout: 10 * time.Second}

	r, err := client.Get(query)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer r.Body.Close()

	type recipe struct {
		Title       string
		Href        string
		Ingredients string
	}

	type apiResponse struct {
		Results []recipe
	}

	var res apiResponse

	err = json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	funcMap := promptui.FuncMap
	funcMap["truncate"] = func(size int, input string) string {
		if len(input) <= size {
			return input
		}
		return input[:size-3] + "..."
	}

	templates := promptui.SelectTemplates{
		Active:   `🍕 {{ .Title | cyan | bold }}`,
		Inactive: `   {{ .Title | cyan }}`,
		Selected: `{{ "✔" | green | bold }} {{ "Recipe" | bold }}: {{ .Title | cyan }}`,
	}

	list := promptui.Select{
		Label:     "Recipe",
		Items:     res.Results,
		Templates: &templates,
	}

	idx, _, err := list.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(res.Results[idx].Href)
}

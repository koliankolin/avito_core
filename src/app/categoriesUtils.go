package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type Categories struct {
	Категории map[string][]string
}

func getAllCategories() (*Categories, error) {

	resp, err := http.Get("https://www.avito.ru/rossiya")

	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return nil, err
	}


	categories := new(Categories)
	categories.Категории = make(map[string][]string)

	var mainCategory string
	doc.Find("#category").Children().Each(func(i int, s *goquery.Selection) {
		class, _ := s.Attr("class")
		if class == "opt-group" {
			mainCategory = s.Text()

			categories.Категории[mainCategory] = make([]string, 0)
		} else {
			if mainCategory == "" {
				return
			}
			categories.Категории[mainCategory] = append(categories.Категории[mainCategory], s.Text())
		}
	})

	return categories, nil
}

func getAllCategoriesJson() ([]byte, error) {
	categories, err := getAllCategories()

	categoriesJson, err := json.Marshal(categories)

	if err != nil {
		return []byte(`{"message":"error has happened"}`), err
	}
	return categoriesJson, nil
}

func printHumanReadableCategories() error {
	categories, err := getAllCategories()

	if err != nil {
		return err
	}

	for category, subcats := range categories.Категории {
		fmt.Println(category)

		if len(subcats) > 0 {
			for _, subcat := range subcats[:len(subcats) - 2] {
				fmt.Printf("\t├── %s\n", subcat)
			}
			fmt.Printf("\t└── %s\n", subcats[len(subcats) - 1])
		}
	}
	return nil
}



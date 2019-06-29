package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/essentialkaos/translit"
	"net/http"
	"strconv"
	"strings"
)

type Statistics struct {
	Location string
	Total int
	Categories map[string]int
}

func convertToInt(str string) (int, error) {
	return strconv.Atoi(strings.ReplaceAll(str, " ", ""))
}

func convertToTranslit(location string) string {
	location = strings.ReplaceAll(location, " ", "_")
	location = strings.ToLower(translit.EncodeToBGN(location))
	return strings.ReplaceAll(location, "′", "")
}

func getStatisticsFromLocation(location string) (*Statistics, error) {
	resp, err := http.Get("https://www.avito.ru/" + convertToTranslit(location))

	if err != nil || resp.Status == "404" {
		return nil, nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, nil
	}
	total, err := convertToInt(doc.Find(".breadcrumbs-link-count.js-breadcrumbs-link-count").Text())
	if err != nil {
		return nil, nil
	}

	statistics := new(Statistics)
	statistics.Categories = make(map[string]int)
	statistics.Total = total
	statistics.Location = strings.Title(location)

	doc.Find(".catalog-counts__row.clearfix").Children().Each(func(i int, s *goquery.Selection) {
		s.Children().Each(func(i int, ch *goquery.Selection) {
			arr := strings.Split(ch.Text(), "    ")
			totalCategory, err := convertToInt(arr[1])
			if err != nil {
				return
			}
			categoryName := strings.ReplaceAll(arr[0], "\n", "")
			categoryName = strings.TrimSpace(categoryName)
			statistics.Categories[categoryName] = totalCategory
		})
	})

	return statistics, nil
}

func getStatisticsOneJson(location string) ([]byte, error) {
	stats, err := getStatisticsFromLocation(location)

	if stats == nil {
		return []byte{}, fmt.Errorf("no such location: " + location)
	}

	statisticsJson, err := json.Marshal(stats)
	if err != nil {
		return []byte{}, err
	}

	return statisticsJson, nil
}

func printHumanReadableStatistic(location string) error {
	stats, err := getStatisticsFromLocation(location)

	if err != nil {
		return err
	}

	if stats == nil {
		return fmt.Errorf("no such location: " + location)
	}

	fmt.Printf("Локация: %s\n", strings.Title(stats.Location))
	fmt.Printf("Общее количество объявлений: %d\n", stats.Total)
	for category, value := range stats.Categories {
		fmt.Printf("%s: %d\n", category, value)
	}
	return nil
}
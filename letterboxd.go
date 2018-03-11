package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	var diaryFile = "diary.csv"
	// FIXME: make the username configurable.
	var letterboxdUsername = "pdurbin"
	csvFile, tsvOpenError := os.Open(diaryFile)
	if tsvOpenError != nil {
		fmt.Println("You must export your Letterboxd diary.csv file and place it in the current directory. Error: " + tsvOpenError.Error())
		os.Exit(1)
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	_, _ = reader.Read() // delete header
	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var movie Movie
	for _, each := range csvData {
		movie.Title = each[1]
		movie.URL = each[3]
		movie.Rating = each[4]
		movie.Watched = each[6]
		movie.Slug = "https://letterboxd.com/" + letterboxdUsername + "/film/" + strings.Split(movie.URL, "/")[5] + "/"
		fmt.Printf("%s\t%s\t%-35s\t%s\n", movie.Watched, movie.Rating, movie.Title, movie.URL)
	}
}

type Movie struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	Rating  string `json:"rating"`
	Slug    string `json:"slug"`
	Watched string `json:"rating"`
}

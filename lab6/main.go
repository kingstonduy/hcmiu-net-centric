package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type Manga struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Likes  string `json:"likes"`
}
type GenreMangas struct {
	Genre  string  `json:"gerne"`
	Mangas []Manga `json:"mangas"`
}
type AllGenres struct {
	Genres []GenreMangas `json:"gernes"`
}

var genres = []string{
	"action",
	"comedy",
	"drama",
	"fantasy",
	"heartwarming",
	"historical",
	"horror",
	"mystery",
	"romance",
	"sf",
	"slice-of-life",
	"sports",
	"supernatural",
	"super-hero",
	"thriller",
	"tiptoon",
}

const (
	BASE_URL = "https://www.webtoons.com/en/genres/"
)

func Handle() {
	var allGenres AllGenres
	// Loop through each genre
	for _, genre := range genres {
		fmt.Printf("Scraping genre: %s\n", genre)

		// Request the HTML page for the genre
		url := BASE_URL + genre
		res, err := http.Get(url)
		if err != nil {
			log.Fatalf("Error fetching URL for genre %s: %v", genre, err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("Status code error for genre %s: %d %s", genre, res.StatusCode, res.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatalf("Error parsing HTML for genre %s: %v", genre, err)
		}

		// Extract manga details for the current genre
		var mangas []Manga
		doc.Find("li").Each(func(index int, item *goquery.Selection) {
			// Extract details
			title := item.Find("p.subj").Text()
			author := item.Find("p.author").Text()
			likes := item.Find("p.grade_area em.grade_num").Text()

			// Only append if all fields are non-empty
			if title != "" && author != "" && likes != "" {
				mangas = append(mangas, Manga{
					Title:  title,
					Author: author,
					Likes:  likes,
				})
			}
		})

		// Append genre and its mangas to the result
		allGenres.Genres = append(allGenres.Genres, GenreMangas{
			Genre:  genre,
			Mangas: mangas,
		})
	}

	// Save all data to a single JSON file
	file, err := os.Create("result.json")
	if err != nil {
		log.Fatal("Error creating JSON file:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(allGenres); err != nil {
		log.Fatal("Error encoding JSON data:", err)
	}

	fmt.Println("Manga details have been saved to mangas.json.")
}

func main() {
	Handle()
}

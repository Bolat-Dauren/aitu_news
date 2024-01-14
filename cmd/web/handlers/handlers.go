// cmd/web/handlers/handlers.go
package handlers

import (
	"aitu_news/pkg/models/driver"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"aitu_news/pkg/models"
)

var tmpl = template.Must(template.ParseFiles(
	filepath.Join("ui", "html", "index.html"),
	filepath.Join("ui", "html", "about.html"),
	filepath.Join("ui", "html", "contacts.html"),
	filepath.Join("ui", "html", "categories.html"),
	filepath.Join("ui", "html", "ad_article.html"), // Add the new HTML file
))

var newsRepo models.NewsRepository = &models.MemoryNewsRepository{}

func HandleRequests() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, "Error rendering home template", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "about.html", nil)
		if err != nil {
			http.Error(w, "Error rendering about template", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "contacts.html", nil)
		if err != nil {
			http.Error(w, "Error rendering contacts template", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "categories.html", nil)
		if err != nil {
			http.Error(w, "Error rendering categories template", http.StatusInternalServerError)
			return
		}
	})

	// Handle category requests
	http.HandleFunc("/category/", func(w http.ResponseWriter, r *http.Request) {
		category := strings.TrimPrefix(r.URL.Path, "/category/")
		fmt.Fprintf(w, "Selected category: %s", category)
	})

	http.HandleFunc("/ad_articles", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form", http.StatusInternalServerError)
				return
			}

			title := r.FormValue("title")
			content := r.FormValue("content")

			err = driver.AddArticle(title, content)
			if err != nil {
				http.Error(w, "Error adding article to database", http.StatusInternalServerError)
				return
			}

			// Redirect to the /ad_articles page after adding the article
			http.Redirect(w, r, "/ad_articles", http.StatusSeeOther)
			return
		}

		// Retrieve all articles from the database
		articles, err := driver.GetArticles()
		if err != nil {
			http.Error(w, "Error getting articles from database", http.StatusInternalServerError)
			return
		}

		// Render the ad_article template with the list of articles for GET requests
		err = tmpl.ExecuteTemplate(w, "ad_article.html", articles)
		if err != nil {
			http.Error(w, "Error rendering ad_article template", http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

// pkg/models/driver/db.go
package driver

import (
	"aitu_news/pkg/models"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "0000"
	dbname   = "news_db"
)

func ConnectDB() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to the database.")
}

// AddArticle adds a new article to the database
func AddArticle(title, content string) error {
	_, err := db.Exec("INSERT INTO articles (title, content) VALUES ($1, $2)", title, content)
	if err != nil {
		fmt.Println("Error adding article to database:", err)
	}
	return err
}

func GetArticles() ([]models.News, error) {
	rows, err := db.Query("SELECT * FROM articles ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []models.News
	for rows.Next() {
		var article models.News
		err := rows.Scan(&article.ID, &article.Title, &article.Content)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

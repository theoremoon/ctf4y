package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id      int64
	Title   string
	Private bool
}

func render(w http.ResponseWriter, name string, values map[string]interface{}) {
	funcMap := template.FuncMap{
		"nl2br": func(text string) template.HTML {
			return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
		},
	}
	t := template.Must(template.New("view").Funcs(funcMap).ParseFiles("templates/template.html", name))
	t.ExecuteTemplate(w, "base", values)
}

func main() {
	db, err := sql.Open("mysql", os.Getenv("DBURL"))
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		limit := r.URL.Query().Get("limit")
		if limit == "" {
			limit = "20"
		}
		rows, err := db.Query("SELECT id, title, private FROM posts ORDER BY created_at DESC LIMIT " + limit)
		if err != nil {
			log.Println(err)
			render(w, "templates/error.html", map[string]interface{}{
				"Error": "INTERNAL ERROR",
			})
			return
		}
		defer rows.Close()

		posts := make([]Post, 0, 20)
		for rows.Next() {
			var post Post
			if err := rows.Scan(&post.Id, &post.Title, &post.Private); err != nil {
				log.Println(err)
				render(w, "templates/error.html", map[string]interface{}{
					"Error": "INTERNAL ERROR",
				})
				return
			}
			posts = append(posts, post)
		}
		render(w, "templates/index.html", map[string]interface{}{
			"Posts": posts,
		})
	})
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				render(w, "templates/error.html", map[string]interface{}{
					"Error": "INVALID REQUEST",
				})
				return
			}
			title := r.Form.Get("title")
			content := r.Form.Get("content")
			private_key := r.Form.Get("key")

			stmt, err := db.Prepare("INSERT INTO posts(title, content, private, private_key) VALUES(?,?,?,?)")
			if err != nil {
				log.Println(err)
				render(w, "templates/error.html", map[string]interface{}{
					"Error": "INTERNAL ERROR",
				})
				return
			}
			defer stmt.Close()

			_, err = stmt.Exec(title, content, private_key != "", private_key)
			if err != nil {
				log.Println(err)
				render(w, "templates/error.html", map[string]interface{}{
					"Error": "INTERNAL ERROR",
				})
				return
			}

			http.Redirect(w, r, "/", 301)
		} else {
			render(w, "templates/create.html", map[string]interface{}{})
		}
	})
	http.HandleFunc("/view", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		var title string
		var content string
		err := db.QueryRow("SELECT title, content FROM posts WHERE id = ? AND private_key = ?", query.Get("id"), query.Get("key")).Scan(&title, &content)
		if err != nil {
			log.Println(err)
			render(w, "templates/error.html", map[string]interface{}{
				"Error": "Post not found or it is secret",
			})
			return
		}
		render(w, "templates/view.html", map[string]interface{}{
			"Title": title, "Content": content,
		})
	})
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		limit := r.URL.Query().Get("limit")
		if limit == "" {
			limit = "20"
		}

		var count int
		err := db.QueryRow(fmt.Sprintf("SELECT count(*) FROM posts WHERE title LIKE '%%%s%%' LIMIT %s", query, limit)).Scan(&count)
		if err != nil {
			log.Println(err)
			render(w, "templates/error.html", map[string]interface{}{
				"Error": "INTERNAL ERROR",
			})
			return
		}

		render(w, "templates/count.html", map[string]interface{}{
			"Count": count,
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, logHandler(http.DefaultServeMux)))
}

func logHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.Method, r.URL, r.RemoteAddr)
		handler.ServeHTTP(w, r)
	})
}

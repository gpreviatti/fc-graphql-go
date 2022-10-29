package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gpreviatti/fc-graphql-go/database"
	"github.com/gpreviatti/fc-graphql-go/graph"
	"github.com/gpreviatti/fc-graphql-go/graph/generated"
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err!= nil {
        log.Fatalf("failed to open database: %v", err)
    }
	defer db.Close()

	categoryDb := database.NewCategory(db)
	courseDb := database.NewCourse(db)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDb,
		CourseDB: courseDb,
	}}))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

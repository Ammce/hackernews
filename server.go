package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Ammce/hackernews/adapters/graph/generated"
	graph "github.com/Ammce/hackernews/adapters/graph/resolvers"
	"github.com/Ammce/hackernews/domain/user"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const defaultPort = ":8080"

// Defining the Graphql handler
func graphqlHandler(db *sql.DB) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file

	userService := user.UserServiceImpl{}

	domain := graph.DomainGraphQL{
		UserService: userService,
	}

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db, Domain: domain}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// db := postgres.DatabaseConfig()
	connStr := "postgres://postgres:postgres@docker.for.mac.localhost:5433/news?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()
	r.POST("/query", graphqlHandler(db))
	r.GET("/", playgroundHandler())
	r.Run(defaultPort)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

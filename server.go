package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Ammce/hackernews/adapters/graph/directives"
	"github.com/Ammce/hackernews/adapters/graph/generated"
	"github.com/Ammce/hackernews/adapters/graph/middleware"
	graph "github.com/Ammce/hackernews/adapters/graph/resolvers"
	repositories "github.com/Ammce/hackernews/adapters/postgres/repository"
	"github.com/Ammce/hackernews/domain/auth"
	"github.com/Ammce/hackernews/domain/user"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const defaultPort = ":8080"

// Defining the Graphql handler
func graphqlHandler(db *sql.DB) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	userRepo := repositories.UserRepositoryImpl{DB: db}
	userService := user.UserServiceImpl{UserRepo: userRepo}
	authService := auth.AuthServiceImpl{UserRepo: userRepo}

	domain := graph.DomainGraphQL{
		UserService: userService,
		AuthService: authService,
	}

	c := generated.Config{Resolvers: &graph.Resolver{DB: db, Domain: domain, UserDataLoader: graph.UserDataLoader(db)}}
	directives.SetupDirectives(&c)

	h := handler.NewDefaultServer(generated.NewExecutableSchema(c))

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
	r.Use(middleware.SetupContext())
	r.GET("/health", func(ctx *gin.Context) {
		ctx.Status(200)
	})
	r.POST("/query", graphqlHandler(db))
	r.GET("/", playgroundHandler())
	r.Run(defaultPort)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

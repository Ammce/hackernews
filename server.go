package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	newsapiRepo "github.com/Ammce/hackernews/adapters/clients/newsapi"
	domainGraph "github.com/Ammce/hackernews/adapters/graph"
	"github.com/Ammce/hackernews/adapters/graph/directives"
	"github.com/Ammce/hackernews/adapters/graph/generated"
	"github.com/Ammce/hackernews/adapters/graph/middleware"
	graph "github.com/Ammce/hackernews/adapters/graph/resolvers"
	repositories "github.com/Ammce/hackernews/adapters/postgres/repository"
	"github.com/Ammce/hackernews/dataloaders"
	"github.com/Ammce/hackernews/domain/article"
	"github.com/Ammce/hackernews/domain/auth"
	"github.com/Ammce/hackernews/domain/comment"
	externalarticle "github.com/Ammce/hackernews/domain/externalArticle"
	"github.com/Ammce/hackernews/domain/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const defaultPort = ":8080"

// Defining the Graphql handler
func graphqlHandler(db *sql.DB) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	userRepo := repositories.NewUserRepositoryImpl(db)
	articleRepo := repositories.NewArticleRepositoryImpl(db)
	commentRepo := repositories.NewCommentRepositoryImpl(db)
	externalArticleRepo := newsapiRepo.NewNewsApi()

	userService := user.NewUserServiceImpl(userRepo)
	authService := auth.NewAuthServiceImpl(userRepo)
	articleService := article.NewArticleServiceImpl(articleRepo)
	commentService := comment.NewCommentServiceImpl(commentRepo)
	externalArticleService := externalarticle.NewExternalArticleServiceImpl(externalArticleRepo)

	domain := domainGraph.DomainGraphQL{
		UserService:            userService,
		AuthService:            authService,
		ArticleService:         articleService,
		CommentService:         commentService,
		ExternalArticleService: externalArticleService,
	}

	c := generated.Config{Resolvers: &graph.Resolver{DB: db, Domain: domain, UserDataLoader: dataloaders.UserDataLoader(db), ArticleDataLoader: dataloaders.ArticleDataLoader(db), CommentDataLoader: dataloaders.CommentDataLoader(db)}}
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error while loading the env variables: %s", err)
	}
	fmt.Println("Env variables successfully loaded.", os.Getenv("ENV_HEALTCHECK"))
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
	r.Use(middleware.CORSMiddleware())
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

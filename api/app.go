package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/daesu/stoicadon/api/graphql/gen"
	"github.com/daesu/stoicadon/api/graphql/resolvers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
)

const (
	defaultPort = "8080"
)

// Application is a simple struct to hold API configurations
type Application struct {
	router chi.Router
	host   string
	port   string
}

// ConfigureApplication grabs environment variables and sets
// up the database, routes etc for the API to run.
func ConfigureApplication() (*Application, error) {
	host, err := os.Hostname()
	if err != nil {
		logrus.Fatal("unable to get Hostname", err)
	}

	// cors
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "SiteID", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// Setup routes
	router := chi.NewRouter()
	resolver := resolvers.NewResolver()

	// protected routes
	router.Group(func(router chi.Router) {
		router.Use(c.Handler)

		gqlConfig := gen.Config{Resolvers: resolver}
		graphqlHandler := handler.New(gen.NewExecutableSchema(gqlConfig))
		graphqlHandler.AddTransport(transport.POST{})
		router.Handle("/query", graphqlHandler)
	})

	enablePlayground := os.Getenv("ENABLE_PLAYGROUND")
	if enablePlayground != "" {
		router.Group(func(router chi.Router) {
			router.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
		})
	}

	portFlag := os.Getenv("PORT")
	if portFlag == "" {
		portFlag = defaultPort
		logrus.Infof("port not specified. Using default port %s", portFlag)
	}

	application := Application{
		router,
		host,
		portFlag,
	}

	return &application, nil
}

func StartAPI(application *Application) error {
	logrus.WithFields(logrus.Fields{
		"Host": application.host,
	}).Info("Starting API")

	serveOn := fmt.Sprintf(":%s", application.port)
	err := http.ListenAndServe(serveOn, application.router)
	if err != nil {
		panic(err)
	}

	return nil
}

package web

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"pdf-service/middleware"
)

// TempDir is the tempdir
var TempDir = ""

// RunServer runs the webserver
func RunServer(port int, tempDir string) {
	TempDir = tempDir

	// Setuop routes & cors
	router := setupRouter()
	cors := setupCORS()

	s := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		Handler:      cors(router),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.LoggerMiddleware)
	router.HandleFunc("/", HandleHome).Methods(http.MethodGet)
	router.HandleFunc("/upload", HandleHTMLUpload).Methods(http.MethodPost)

	return router
}

func setupCORS() func(http.Handler) http.Handler {
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Request-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
	)

	return cors
}

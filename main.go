package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"pdf-service/middleware"
	"pdf-service/web"
)

type cliArgs struct {
	Port int
}

func main() {
	// Check Chromium path
	checkChromium()

	// Make tmp dir
	makeTmpDir()

	// Fetch flags
	args := cliArgs{}
	flag.IntVar(&args.Port, "port", 8765, "Command line arguments")
	flag.Parse()

	fmt.Println("Starting PDF Service")
	fmt.Println("Running on port: ", args.Port)

	// Setuop routes & cors
	router := setupRouter()
	cors := setupCORS()

	s := &http.Server{
		Addr:         ":" + strconv.Itoa(args.Port),
		Handler:      cors(router),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}

func makeTmpDir() {
	_ = os.Mkdir("./tmp", 0666)
}

func checkChromium() {
	_, err := exec.LookPath("chromium")
	if err != nil {
		log.Fatal(err)
	}
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.LoggerMiddleware)
	router.HandleFunc("/", web.HandleHome).Methods(http.MethodGet)
	router.HandleFunc("/upload", web.HandleHTMLUpload).Methods(http.MethodPost)

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

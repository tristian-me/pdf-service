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

// CliArgs is a simple struct for the CLI Args (flags)
type CliArgs struct {
	Port	int
	TempDir string
}

// DefaultTempDir is the default temporary directory
const DefaultTempDir = "/tmp/pdf-service"

func main() {
	// Check Chromium path
	checkChromium()

	// Fetch flags
	args := CliArgs{}
	flag.IntVar(&args.Port, "port", 8765, "Command line arguments")
	flag.StringVar(&args.TempDir, "temp-dir", DefaultTempDir, "The pemparary directory")
	flag.Parse()

	fmt.Println("Starting PDF Service")
	fmt.Println("Running on port: ", args.Port)

	err := setupDirectories(args)
	if err != nil {
		log.Fatal("Cannot create directory, maybe not writable?")
	}

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

func setupDirectories(args CliArgs) error {
	var err error

	if args.TempDir == DefaultTempDir {
		err = os.MkdirAll(DefaultTempDir, os.ModePerm)
	} else {
		err = os.Mkdir(args.TempDir, os.ModePerm)
	}

	return err
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

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Config struct {
	Port             string   `json:"port"`
	Web              string   `json:"web"`
	AllowedFileTypes []string `json:"allowedFileTypes"`
}

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	http.HandleFunc("/", handleRequest(config))

	fmt.Printf("Starting server on port %s\n", config.Port)
	http.ListenAndServe(":"+config.Port, nil)
}

func loadConfig(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func handleRequest(config Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cleanedPath := filepath.Clean(r.URL.Path)

		if cleanedPath == "/" {
			http.Redirect(w, r, "/index.md", http.StatusMovedPermanently)
			return
		}

		path := filepath.Join(config.Web, cleanedPath)

		validExtension := false
		for _, ext := range config.AllowedFileTypes {
			if filepath.Ext(path) == ext {
				validExtension = true
				break
			}
		}
		if !validExtension {
			http.Error(w, "Invalid file type", 400)
			return
		}

		file, err := os.Open(path)
		if err != nil {
			http.Error(w, "Oops! 404 document not found", 404)
			return
		}
		defer file.Close()

		// w.Header().Add("Cache-Control", "no-cache")
		w.Header().Add("Cache-Control", "public, max-age=31536000")

		if _, err := io.Copy(w, file); err != nil {
			fmt.Printf("Failed to write response: %v\n", err)
			http.Error(w, "Failed to write response", 500)
		}
	}
}

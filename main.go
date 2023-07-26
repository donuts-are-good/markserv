package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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
		path := fmt.Sprintf("%s%s", config.Web, r.URL.Path)

		if strings.Contains(path, "..") {
			http.Error(w, "Invalid path", 400)
			return
		}

		if strings.HasSuffix(path, "/") {
			path = path + "index.md"
		} else {
			validExtension := false
			for _, ext := range config.AllowedFileTypes {
				if strings.HasSuffix(path, ext) {
					validExtension = true
					break
				}
			}
			if !validExtension {
				http.Error(w, "Invalid file type", 400)
				return
			}
		}

		file, err := os.Open(path)
		if err != nil {
			http.Error(w, "File not found", 404)
			return
		}
		defer file.Close()

		io.Copy(w, file)
	}
}

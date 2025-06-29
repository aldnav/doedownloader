package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/joho/godotenv"
)

// Global variables
var reportsDirectory string
var reportFileInput string
var cookiePath string

func readEnv(key string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}
	err := godotenv.Load(envFile)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func validateArguments() {
	reportsDirectory = readEnv("REPORTS_DIRECTORY")
	if reportsDirectory == "" {
		fmt.Println("REPORTS_DIRECTORY is not set")
		os.Exit(1)
	}
	if _, err := os.Stat(reportsDirectory); os.IsNotExist(err) {
		fmt.Println("REPORTS_DIRECTORY does not exist")
		os.Exit(1)
	}
	cookiePath = readEnv("COOKIE_PATH")
}

type ReportLog []struct {
	Name        string     `json:"name"`
	URL         string     `json:"url"`
	Description string     `json:"description"`
	Attachments [][]string `json:"attachments"`
}

func readLinksFromReport() []string {
	// Identify latest report using the format (YYYY-MM-DD_retail_pump.json)
	fmt.Println("[.] Looking for latest report")
	candidateFiles, err := filepath.Glob(reportsDirectory + "/*_retail_pump.json")
	if err != nil {
		log.Fatal(err)
	}
	if len(candidateFiles) == 0 {
		fmt.Println("No reports found")
		os.Exit(1)
	}
	sort.Strings(candidateFiles)
	reportFileInput = candidateFiles[len(candidateFiles)-1]
	fmt.Println("[.] Reading from report file: " + reportFileInput)
	reportFile, err := ioutil.ReadFile(reportFileInput)
	if err != nil {
		log.Fatal(err)
	}
	// Read the report and parse it
	data := ReportLog{}
	err = json.Unmarshal(reportFile, &data)
	if err != nil {
		log.Fatal(err)
	}
	// Return array of all links
	linksFromAttachments := []string{}
	for _, attachment := range data {
		for _, link := range attachment.Attachments {
			linksFromAttachments = append(linksFromAttachments, link[1])
		}
	}
	if len(linksFromAttachments) == 0 {
		fmt.Println("No links found")
		os.Exit(1)
	}
	return linksFromAttachments
}

func downloadReports(links []string) {
	// Create directory for the downloaded reports
	directoryPath := strings.Split(reportFileInput, "_")[0]
	err := os.MkdirAll(directoryPath, 0755)
	if err != nil {
		log.Fatal(err)
	}
	// Download all links
	for _, link := range links {
		downloadFile(link, directoryPath)
	}
}

type RawCookies []struct {
	SameSite string `json:"sameSite"`
	Name     string `json:"name,omitempty"`
	Value    string `json:"value"`
	Domain   string `json:"domain"`
	Path     string `json:"path"`
	HTTPOnly bool   `json:"httpOnly"`
	Secure   bool   `json:"secure"`
}

func getCookieFromFile(name string) (http.Cookie, error) {
	// Read cookie file
	cookieFile, err := ioutil.ReadFile(cookiePath)
	if err != nil {
		log.Fatal(err)
	}
	// Read the cookie in raw JSON
	var rawCookies RawCookies
	err = json.Unmarshal(cookieFile, &rawCookies)
	if err != nil {
		log.Fatal(err)
	}
	for _, rCookie := range rawCookies {
		if rCookie.Name == name {
			return http.Cookie{
				Name:     rCookie.Name,
				Value:    rCookie.Value,
				Domain:   rCookie.Domain,
				Path:     rCookie.Path,
				HttpOnly: rCookie.HTTPOnly,
				Secure:   rCookie.Secure,
			}, nil
		}
	}
	return http.Cookie{}, errors.New("Cookie not found")
}

func downloadFile(link, directoryPath string) {
	// Build fileName from fullPath
	fileURL, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]

	// Create blank file
	file, err := os.Create(directoryPath + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	// Setting cookie first
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Got error while creating cookie jar %s", err.Error())
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
		Jar: jar,
	}

	// Use cookie if COOKIE_PATH is provided
	if cookiePath != "" {
		cookie, err := getCookieFromFile("AKSYONSHIELD")
		if err == nil {
			urlObj, _ := url.Parse(link)
			client.Jar.SetCookies(urlObj, []*http.Cookie{&cookie})
		} else {
			fmt.Printf("[!] Warning: Could not load cookie: %v\n", err)
		}
	}
	// Put content on file
	resp, err := client.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("[.] Downloaded a file %s with size %d\n", directoryPath+"/"+fileName, size)
}

func main() {
	validateArguments()
	links := readLinksFromReport()
	downloadReports(links)
	fmt.Println("[.] Done!")
}

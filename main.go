package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jexlor/colorado"
)

// Function to sanitize the URL and create a valid directory name
func sanitizeURLForDir(urlStr string) string {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "site"
	}
	return strings.ReplaceAll(strings.ReplaceAll(u.Host, ".", "_"), ":", "_")
}

func downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath, body, 0644)
}

func downloadResource(baseURL, resource, outputDir string) error {
	resourceURL, err := url.Parse(resource)
	if err != nil {
		return err
	}

	if !resourceURL.IsAbs() {
		base, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		resourceURL = base.ResolveReference(resourceURL)
	}

	resourcePath := filepath.Join(outputDir, resourceURL.Path)
	os.MkdirAll(filepath.Dir(resourcePath), os.ModePerm)

	err = downloadFile(resourceURL.String(), resourcePath)
	if err != nil {
		fmt.Printf("Failed to download %s: %v\n", resourceURL.String(), err)
	} else {
		fmt.Printf("Downloaded %s\n", resourceURL.String())
	}

	return err
}

func extractResources(html string) []string {
	re := regexp.MustCompile(`(src|href)="([^"]+)"`)
	matches := re.FindAllStringSubmatch(html, -1)
	var resources []string
	for _, match := range matches {
		resource := match[2]
		if !strings.HasPrefix(resource, "http") && !strings.HasPrefix(resource, "#") {
			resources = append(resources, resource)
		}
	}
	return resources
}

func downloadPage(baseURL, outputDir, pageURL string) error {
	pagePath := filepath.Join(outputDir, pageURL)
	os.MkdirAll(filepath.Dir(pagePath), os.ModePerm)

	err := downloadFile(baseURL+pageURL, pagePath)
	if err != nil {
		return fmt.Errorf("failed to download page %s: %v", baseURL+pageURL, err)
	}

	html, err := ioutil.ReadFile(pagePath)
	if err != nil {
		return fmt.Errorf("failed to read downloaded page %s: %v", pagePath, err)
	}

	resources := extractResources(string(html))
	for _, resource := range resources {
		err = downloadResource(baseURL, resource, outputDir)
		if err != nil {
			fmt.Printf("Failed to download resource %s: %v\n", resource, err)
		}
	}

	return nil
}

func main() {
	text := `
                )                            
          (   ( /(       (   (     )      (   
  (      ))\  )\())(    ))\  )(   /((    ))\  
  )\ )  /((_)(_))/ )\  /((_)(()\ (_))\  /((_) 
 _(_/( (_))  | |_ ((_)(_))   ((_)_)((_)(_))   
| ' \))/ -_) |  _|(_-</ -_) | '_|\ V / / -_)  
|_||_| \___|  \__|/__/\___| |_|   \_/  \___|  
                                                                          									  
		  `

	fmt.Println(colorado.Color(text, colorado.Red, ""))
	urlStr := "example.com" //enter here site you want to download
	baseURL, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Invalid URL:", err)
		return
	}
	siteDir := sanitizeURLForDir(baseURL.String())
	outputDir := filepath.Join(".", siteDir)

	// Create the output directory if it doesn't exist
	os.MkdirAll(outputDir, os.ModePerm)

	err = downloadPage(baseURL.String(), outputDir, "index.html")
	if err != nil {
		fmt.Printf("Error downloading index page: %v\n", err)
		return
	}
	fmt.Println(colorado.Color("Content downloaded successfully! Run 'serve_content.sh' to host it.", colorado.Blue, ""))
}

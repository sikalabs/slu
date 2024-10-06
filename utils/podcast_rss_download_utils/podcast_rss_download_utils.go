package podcast_rss_download_utils

// Created using ChatGPT: https://chatgpt.com/share/cc1cddbe-2351-4678-90f9-5a088be6eb70

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/unicode/norm"
)

type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Items []Item `xml:"item"`
}

type Item struct {
	Title     string    `xml:"title"`
	PubDate   string    `xml:"pubDate"`
	Enclosure Enclosure `xml:"enclosure"`
}

type Enclosure struct {
	URL string `xml:"url,attr"`
}

func sanitizeFilename(name string) string {
	// Normalize the string and remove diacritics
	t := norm.NFD.String(name)
	t = runes.Remove(runes.In(unicode.Mn)).String(t)

	// Replace spaces with underscores
	t = strings.ReplaceAll(t, " ", "_")

	// Remove special characters
	reg := regexp.MustCompile("[^a-zA-Z0-9_.]+")
	t = reg.ReplaceAllString(t, "")

	return t
}

func parsePubDate(pubDate string) time.Time {
	// Define layouts to handle both zero-padded and non-zero-padded days
	layouts := []string{
		"Mon, 02 Jan 2006 15:04:05 MST",   // zero-padded day
		"Mon, 2 Jan 2006 15:04:05 MST",    // non-zero-padded day
		"Mon, 02 Jan 2006 15:04:05 +0200", // zero-padded day with timezone
		"Mon, 2 Jan 2006 15:04:05 +0200",  // non-zero-padded day with timezone
	}

	var t time.Time
	var err error

	// Try parsing with each layout
	for _, layout := range layouts {
		t, err = time.Parse(layout, pubDate)
		if err == nil {
			return t
		}
	}

	// If parsing fails, log the error and return the zero value of time.Time
	fmt.Println("Error parsing publication date:", err)
	return time.Time{}
}

func formatPubDate(pubDate string) string {
	// Parse the PubDate using the parsePubDate function
	t := parsePubDate(pubDate)
	if t.IsZero() {
		return "unknown_date_"
	}

	// Format the date as yyyy-mm-dd
	return t.Format("2006-01-02") + "_"
}

func downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func PodcastRssDownload(rssURL, outDir string) {
	// Create the output directory if it doesn't exist
	err := os.MkdirAll(outDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	resp, err := http.Get(rssURL)
	if err != nil {
		fmt.Println("Error fetching RSS feed:", err)
		return
	}
	defer resp.Body.Close()

	var rss RSS
	err = xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		fmt.Println("Error decoding RSS feed:", err)
		return
	}

	for _, item := range rss.Channel.Items {
		datePrefix := formatPubDate(item.PubDate)
		filename := datePrefix + sanitizeFilename(item.Title) + ".mp3"
		filepath := filepath.Join(outDir, filename)

		// Check if the file already exists
		if _, err := os.Stat(filepath); err == nil {
			fmt.Println("File already exists, skipping:", filepath)
			continue
		}

		fmt.Println("Downloading:", item.Enclosure.URL)
		err = downloadFile(item.Enclosure.URL, filepath)
		if err != nil {
			fmt.Println("Error downloading file:", err)
		} else {
			fmt.Println("Saved as:", filepath)
		}
	}
}

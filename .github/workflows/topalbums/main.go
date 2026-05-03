package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// 1. TODO: Retrieve spotify data
	albums := randomAlbums(5)

	// 2: Create spotify html using data
	spotifyhtml := render(albums)

	// 3. inject the spotify html into the readme
	readmePath := "README.md"
	src, err := os.ReadFile(readmePath)
	if err != nil {
		return fmt.Errorf("read %s: %w", readmePath, err)
	}
	out := injectSpotifyHtml(string(src), spotifyhtml)
	if err := os.WriteFile(readmePath, []byte(out), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", readmePath, err)
	}
	fmt.Printf("rendered %d albums into %s\n", len(albums), readmePath)
	return nil
}
func injectSpotifyHtml(src, html string) string {
	spliceRe := regexp.MustCompile(`(?s)<!-- spotify -->.*?<!-- /spotify -->`)
	return spliceRe.ReplaceAllString(src, "<!-- spotify -->\n"+html+"\n<!-- /spotify -->")
}

// Everything below will be removed once spotify integration is done (maybe just restructure render to actual spotify data)
type album struct {
	Artist   string
	Title    string
	URL      string
	ImageURL string
}

var artists = []string{
	"Geese", "Mitski", "Cameron Winter", "Big Thief", "Phoebe Bridgers",
	"Black Country, New Road", "Wednesday", "Alex G", "Fontaines D.C.",
	"Caroline Polachek", "MJ Lenderman", "Waxahatchee", "Snail Mail",
}
var titles = []string{
	"Heavy Metal", "3D Country", "Getting Killed", "Hellmode",
	"Punisher", "Ants From Up There", "Rat Saw God", "Desire, I Want to Turn Into You",
	"Manning Fireworks", "Tigers Blood", "Lush", "Skinty Fia",
}

func randomAlbums(n int) []album {
	out := make([]album, n)
	for i := range n {
		seed := fmt.Sprintf("%d-%d", rand.Int63(), i)
		out[i] = album{
			Artist:   artists[rand.Intn(len(artists))],
			Title:    titles[rand.Intn(len(titles))],
			URL:      "https://picsum.photos/seed/" + seed,
			ImageURL: "https://picsum.photos/seed/" + seed + "/64/64",
		}
	}
	return out
}
func render(items []album) string {
	var b strings.Builder
	b.WriteString(`<p align="left">`)
	for i, a := range items {
		if i > 0 {
			b.WriteString(" ")
		}
		title := fmt.Sprintf("%s - %s", a.Artist, a.Title)
		fmt.Fprintf(&b, `<a href="%s"><img src="%s" title="%s"></a>`, a.URL, a.ImageURL, title)
	}
	b.WriteString("</p>")
	return b.String()
}

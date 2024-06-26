package main

import (
	"embed"
	"html/template"
	"io"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

//go:embed templates/*
var fs embed.FS

func newTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseFS(fs, "templates/*.html")),
	}
}

type Page struct {
	CurrentStream LofiStream
	Current       int
	Prev          int
	Next          int
}

func newPage(lofistreams []LofiStream) Page {
	current := 0;
	currentStream := lofistreams[current]
	prev, next := getNextPrevStreams(current, len(lofistreams))

	return Page{
		CurrentStream: currentStream,
		Current:       current,
		Prev:          prev,
		Next:          next,
	}
}

// Returns the indexes of the next streams, in the form (prev, next)
func getNextPrevStreams(cur int, streamsLen int) (int, int) {
	prev := (cur - 1 + streamsLen) % streamsLen // Ensuring only positive numbers are returned
	next := (cur + 1) % streamsLen

	return prev, next
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplates()

	streams, err := readStreamsFromJSON("streams.json")
	if err != nil {
		panic(err)
	}

	page := newPage(streams)

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", page)
	})

	e.GET("/stream/next/:songid", func(c echo.Context) error {
		songIdStr := c.Param("songid")

		songId, err := strconv.Atoi(songIdStr)
		if err != nil {
			return c.String(http.StatusBadRequest, "Song ID was not a number!")
		}

		if songId >= len(streams) {
			return c.String(http.StatusNotFound, "No song with that ID exists!")
		}

		page.CurrentStream = streams[songId]
		page.Current = songId
		page.Prev, page.Next = getNextPrevStreams(songId, len(streams))

		c.Render(200, "iframe", page.CurrentStream)
		c.Render(200, "player", page)
		return c.Render(200, "title", page)
	})

	e.GET("/stream/random/:songid", func (c echo.Context) error {
		songIdStr := c.Param("songid")

		songId, err := strconv.Atoi(songIdStr)
		if err != nil {
			return c.String(http.StatusBadRequest, "Song ID was not a number!")
		}

		if songId >= len(streams) {
			return c.String(http.StatusNotFound, "No song with that ID exists!")
		}

		randId := rand.IntN(len(streams))

		for randId == songId { // Ensuring the same song is not played again
			randId = rand.IntN(len(streams))
		}

		page.CurrentStream = streams[randId];
		page.Current = randId
		page.Prev, page.Next = getNextPrevStreams(randId, len(streams))

		c.Render(200, "iframe", page.CurrentStream)
		c.Render(200, "player", page)
		return c.Render(200, "title", page)
	})

	e.GET("/pause-button", func (c echo.Context) error {
		return c.Render(200, "pauseButton", page)
	})

	e.GET("/play-button", func (c echo.Context) error {
		return c.Render(200, "playButton", page)
	})


	e.Logger.Fatal(e.Start(":8080"))
}

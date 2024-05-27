package main

import (
	"html/template"
	"io"
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

func newTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
}

type Page struct {
	CurrentStream LofiStream
	Prev          int
	Next          int
}

func newPage(lofistreams []LofiStream) Page {
	currentStream := lofistreams[0]
	prev, next := getNextPrevStreams(0, len(lofistreams))

	return Page{
		CurrentStream: currentStream,
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
		page.Prev, page.Next = getNextPrevStreams(songId, len(streams))

		c.Render(200, "iframe", page.CurrentStream)
		return c.Render(200, "player", page)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

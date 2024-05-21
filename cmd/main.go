package main

import (
	"html/template"
	"io"

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
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Page struct {
	Streams       []LofiStream
	FirstStream   LofiStream
	currentStream int
}

func (p Page) getCurrentStream() LofiStream {
	return p.Streams[p.currentStream]
}

func newPage(lofistreams []LofiStream) Page {
	return Page{
		Streams:       lofistreams,
		FirstStream:   lofistreams[0],
		currentStream: 0,
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplates()

	streams := []LofiStream{
		newLofiStream("Lofi Girl", "jfKfPfyJRdk"),
		newLofiStream("Chillhop Radio", "5yx6BWlEVcY"),
		newLofiStream("Hiphop Radio", "wkhLHTmS_GI"),
	}

	page := newPage(streams)

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", page)
	})

	e.GET("/stream/prev", func(c echo.Context) error {
		page.currentStream = (page.currentStream - 1) % len(page.Streams)
		c.Render(200, "iframe", page.getCurrentStream())
		return c.Render(200, "player", page.getCurrentStream())
	})

	e.GET("/stream/next", func(c echo.Context) error {
		page.currentStream = (page.currentStream + 1) % len(page.Streams)
		c.Render(200, "iframe", page.getCurrentStream())
		return c.Render(200, "player", page.getCurrentStream())
	})

	e.Logger.Fatal(e.Start(":8080"))
}

var id int

type LofiStream struct {
	Id      int
	Name    string
	VideoId string
}

func newLofiStream(name string, videoid string) LofiStream {
	id += 1
	return LofiStream{
		Id:      id,
		Name:    name,
		VideoId: videoid,
	}
}

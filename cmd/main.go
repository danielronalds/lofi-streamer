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
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Page struct {
	Streams       []LofiStream
	CurrentStream LofiStream
}

func newPage(lofistreams []LofiStream) Page {
	return Page{
		Streams:       lofistreams,
		CurrentStream: lofistreams[0],
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplates()

	streams := []LofiStream{
		newLofiStream("Lofi Girl", "jfKfPfyJRdk"),
		newLofiStream("Chillhop Radio", "5yx6BWlEVcY"),
	}

	page := newPage(streams)

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", page)
	})

	e.GET("/song/:id", func(c echo.Context) error {
		idString := c.Param("id")

		id, err := strconv.Atoi(idString)
		if err != nil {
			return c.String(http.StatusBadRequest, "Id was not a number!")
		}

		for _, stream := range page.Streams {
			if stream.Id == id {
				return c.Render(200, "iframe", stream)
			}
		}

		return c.String(http.StatusBadRequest, "Could not find a matching Id")
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

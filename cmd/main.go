package main

import (
	"fmt"
	"html/template"
	"io"
	tt "text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
	text      *tt.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("./views/*.html")),
		text:      tt.Must(tt.ParseGlob("./views/css/*.css")),
	}
}

type Contact struct {
	Name  string
	Email string
}

func newContact(name, email string) Contact {
	return Contact{
		Name:  name,
		Email: email,
	}
}

type Contacts = []Contact

type Data struct {
	Contacts
}

func newData() Data {
	return Data{
		Contacts: []Contact{
			newContact("john", "valid@email.com"),
			newContact("mary", "mary@email.com"),
		},
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Logger.Info("gettting configs")
	settings, err := loadSettings()
	if err != nil {
		e.Logger.Fatal("error loading startup settings: ", err)
	}
	data := newData()
	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", data)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")

		data.Contacts = append(data.Contacts, newContact(name, email))
		return c.Render(200, "contact-list", data)
	})

	e.Logger.Info("stating up server")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", settings.Port)))
}

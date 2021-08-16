package main

import (
	"fmt"
	"html/template"
	"net/http"
	"roku/roku"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	// r := roku.New()
	// r.Use(roku.Logger(), roku.Recovery())
	r := roku.Default()
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "static")

	stu1 := &student{Name: "Joe", Age: 20}
	stu2 := &student{Name: "Ouhsub", Age: 22}

	r.GET("/", func(c *roku.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	r.GET("/students", func(c *roku.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", roku.H{
			"title":  "roku",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *roku.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", roku.H{
			"title": "roku",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.GET("/panic", func(c *roku.Context) {
		names := []string{"roku"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":8090")
}

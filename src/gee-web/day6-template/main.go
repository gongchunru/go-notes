package main

import (
	"fmt"
	"gee"
	"net/http"
	"time"
)

/*
(1) render array
$ curl http://localhost:9999/date
<html>
<body>
    <p>hello, gee</p>
    <p>Date: 2021-12-26</p>
</body>
</html>
*/

/*
(2) custom render function
$ curl http://localhost:9999/students
<html>
<body>
    <p>hello, gee</p>

    <p>0: Tom is 20 years old</p>

    <p>1: Jacky is 23 years old</p>

</body>
</html>
*/

/*
❯ curl http://localhost:9999/assets/css/gchr.css
p {
    color: orange;
    font-weight: 700;
    font-size: 20px;
}

*/

import (
	"html/template"
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
	r := gee.New()
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHtmlGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Tom", Age: 20}
	stu2 := &student{Name: "Jacky", Age: 23}
	r.GET("", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Date(2021, 12, 26, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run("localhost:9999")
}

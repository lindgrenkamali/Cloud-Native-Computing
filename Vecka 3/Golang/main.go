package main

import (
	"html/template"
	"io"
	"main/data"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)
var employees = []data.Employee{
	{
	Id: 1,
	Age: 22,
	City: "Halmstad",
	Name: "David",
},
 {
	Id: 2,
	Age: 3,
	City: "Halmstad",
	Name: "Happy",
},
}

type EchoTemplate struct {
	templates *template.Template
}

type PageView struct {
	Title string
	Rubrik string
	Employees []data.Employee
}

func start(c *gin.Context){
	c.HTML(http.StatusAccepted, "templates/home.tmpl", &PageView {Title: "test", Rubrik: "Like a boss"})
}

func test(c *gin.Context){
c.IndentedJSON(http.StatusOK, employees)
}

func (t *EchoTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func employeesJson(c *gin.Context){
	var employees []data.Employee
	data.DB.Find(&employees)
	c.IndentedJSON(http.StatusOK, employees)
}


func main() {
	
	router := gin.Default()
	
	router.LoadHTMLGlob("templates/*")
	
	router.Static("/static", "./static")
	router.GET("/employees", func(c *gin.Context) {
        c.JSON(http.StatusOK, employees)
    })

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", PageView{
			Title: "2021", Rubrik: "21", Employees: employees,
		})
	})

	// router.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Hello")
	// })
	
	router.Run("localhost:8080")
	
	
	// t := tabby.New()
	// t.AddHeader("Name", "Age", "City")
	// t.AddLine("David", 22, "Halmstad")

	// print(employees.IsTrash())
	// t.Print()
	//runEcho()
}

// func home (c echo.Context) error {
// 	return c.Render(http.StatusOK, "")
// }

// func runEcho() {
// 	app := echo.New()

// 	et := &EchoTemplate{
// 		templates: template.Must(template.ParseGlob("templates/")),
// 	}

// 	app.Renderer = et

// 	app.GET("/", home)
	

// 	// app.GET("/", func(c echo.Context) error {
// 	// 	return c.Render(http.StatusOK, "/templates/home.html", &PageView {Title: "test", Rubrik: "Like a boss"})
// 	// })

// 	app.Logger.Fatal(app.Start(":8080"))
// }
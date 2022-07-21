package main

import (
	"os"
	"strings"

	"github.com/hwalim/go-project/job_scraper/goroutine/scrapper"
	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	// return c.String(http.StatusOK, "Hello, World!")
	return c.File("home.html")
}

func HandleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName) // 첨부파일을 리턴하는 기ㅇ
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", HandleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}

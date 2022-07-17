package main

// 총 페이지 수를 가져온다음 그 후에는 각 페이지별로 goroutine 생성, getPage는 각 일자리 정보 별로 goroutine 생성
import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	location string
	title    string
	name     string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	title := card.Find(".jobTitle>a").Text()
	name := card.Find(".companyName").Text()
	location := card.Find(".companyLocation").Text()
	c <- extractedJob{
		title:    title,
		name:     name,
		location: location}
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)
	// 닫아줘야함 -> 메모리 새어나가는 걸 막을 수 있음
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body) //input output
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages
}

func getPage(page int, c chan<- []extractedJob) {
	var jobs []extractedJob
	c2 := make(chan extractedJob)
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body) //input output
	checkErr(err)

	searchCards := doc.Find(".resultContent")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c2)
	})
	for i := 0; i < searchCards.Length(); i++ {
		result := <-c2
		jobs = append(jobs, result)
	}
	c <- jobs
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush() // write data to a file

	headers := []string{"Title", "Name", "Location"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{job.title, job.name, job.location}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func main() {
	start := time.Now()
	c := make(chan []extractedJob)
	var jobs []extractedJob
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		go getPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		job := <-c
		jobs = append(jobs, job...)
	}
	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
	end := time.Now()
	fmt.Println("Time:", end.Sub(start))
}

package scrapper

// 총 페이지 수를 가져온다음 그 후에는 각 페이지별로 goroutine 생성, getPage는 각 일자리 정보 별로 goroutine 생성
import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	ccsv "github.com/tsak/concurrent-csv-writer"
)

type extractedJob struct {
	location string
	title    string
	name     string
}

// Cleanstring cleans string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

// Scraper Indeed by term
func Scrape(term string) {
	var baseURL string = "https://kr.indeed.com/jobs?q=" + term + "&limit=50"
	start := time.Now()
	pageChannel := make(chan []extractedJob)
	var allJobs []extractedJob
	totalPages := getPages(baseURL)

	for i := 0; i < totalPages; i++ {
		go getPage(i, baseURL, pageChannel)
	}

	for i := 0; i < totalPages; i++ {
		jobs := <-pageChannel
		allJobs = append(allJobs, jobs...)
	}
	writeJobs(allJobs)
	fmt.Println("Done, extracted", len(allJobs))
	end := time.Now()
	fmt.Println("Time:", end.Sub(start))
}

func extractJob(card *goquery.Selection, cardChannel chan<- extractedJob) {
	title := card.Find(".jobTitle>a").Text()
	name := card.Find(".companyName").Text()
	location := card.Find(".companyLocation").Text()
	cardChannel <- extractedJob{
		title:    title,
		name:     name,
		location: location}
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
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

func getPage(page int, url string, pageChannel chan<- []extractedJob) {
	var jobs []extractedJob
	cardChannel := make(chan extractedJob)
	pageURL := url + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body) //input output
	checkErr(err)

	searchCards := doc.Find(".resultContent")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, cardChannel)
	})
	for i := 0; i < searchCards.Length(); i++ {
		result := <-cardChannel
		jobs = append(jobs, result)
	}
	pageChannel <- jobs
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

	// file, err := os.Create("jobs.csv")
	// checkErr(err)

	// w := csv.NewWriter(file)
	// defer w.Flush() // write data to a file

	// headers := []string{"Title", "Name", "Location"}

	// wErr := w.Write(headers)
	// checkErr(wErr)
	//goroutine으로 만들 수 있음
	// for _, job := range jobs {
	// 	jobSlice := []string{job.title, job.name, job.location}
	// 	jwErr := w.Write(jobSlice)
	// 	checkErr(jwErr)
	// }

	// concurrent-csv-writer 사용
	csv, err := ccsv.NewCsvWriter("jobs.csv")
	if err != nil {
		log.Fatalln("Could not open `jobs.csv` for writing")
	}
	defer csv.Close()

	writeC := make(chan bool)

	for _, job := range jobs {
		go func(job extractedJob) {
			csv.Write([]string{job.title, job.name, job.location})
			writeC <- true
		}(job)
	}
	for i := 0; i < len(jobs); i++ {
		<-writeC
	}

}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const NUMBER_WORKERS = 42
const DATE_PARSE_FORM = "1/2/2006 at 3:04 PM MST"
const DATE_OFFSET = 1167606000 // to decrease the size of outputs

// starts at http://www.speedtest.net/result/109057624
var state = 109057624
var mutex = &sync.Mutex{}

type SpeedTest struct {
	Ping     string
	Download string
	Upload   string
	ISP      string
	Date     time.Time
}

func get_ping(doc *goquery.Document) string {
	return strings.Replace(doc.Find(".share-ping p").Text(), " ms", "", -1)
}
func get_download(doc *goquery.Document) string {
	return strings.Replace(doc.Find(".share-download p").Text(), "Mb/s", "", -1)
}
func get_upload(doc *goquery.Document) string {
	return strings.Replace(doc.Find(".share-upload p").Text(), "Mb/s", "", -1)
}
func get_isp(doc *goquery.Document) string {
	return doc.Find(".share-isp p").Text()
}
func get_date(doc *goquery.Document) time.Time {
	result, err := time.Parse(DATE_PARSE_FORM, doc.Find(".share-meta-date").Text())
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func get_page(id int) *SpeedTest {
	url := "http://www.speedtest.net/result/" + strconv.Itoa(id)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		//fmt.Println("Failed to fetch page")
		return nil
	}

	if strings.Index(doc.Find(".share-metrics").Text(), "Result ID Not Valid") != -1 {
		return nil
	}

	result := &SpeedTest{
		Ping:     get_ping(doc),
		Download: get_download(doc),
		Upload:   get_upload(doc),
		ISP:      get_isp(doc),
		Date:     get_date(doc),
	}
	return result
}

func worker(done chan bool, output chan *SpeedTest) {
	for {
		mutex.Lock()
		current := state
		state++
		mutex.Unlock()
		speedTest := get_page(current)
		if speedTest != nil {
			output <- speedTest
		}
	}
	done <- true
}

func output_log_file(output chan *SpeedTest) {
	logFiles := map[string]*os.File{}

	for {
		select {
		case speedTest := <-output:
			filename := speedTest.Date.Format("2006-01.output")
			F, ok := logFiles[filename]
			if ok == false {
				if f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644); err != nil {
					fmt.Println(err)
					continue
				} else {
					F = f
					defer f.Close()
				}
			}

			line := fmt.Sprintf("%d;%s;%s;%s;%s\n", speedTest.Date.Unix()-DATE_OFFSET, speedTest.Ping, speedTest.Download, speedTest.Upload, speedTest.ISP)
			if _, err := F.WriteString(line); err != nil {
				fmt.Println(err)
			}
		}
	}

}

func main() {
	if len(os.Args) > 1 {
		var err error
		state, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
	}

	done := make(chan bool)
	output := make(chan *SpeedTest)

	for i := 0; i < NUMBER_WORKERS; i++ {
		go worker(done, output)
	}
	go output_log_file(output)

	for i := 0; i < NUMBER_WORKERS; i++ {
		<-done
	}
}

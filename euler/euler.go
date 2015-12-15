package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	problemID := *flag.Int("new", 1, "new <problem ID>")
	flag.Parse()

	folder := "p" + strconv.Itoa(problemID)
	os.Mkdir(folder, 0771)

	data := strings.Replace(problemTemplate, problemIDString, strconv.Itoa(problemID), -1)
	data = strings.Replace(data, problemDescriptionString, getProblemDescription(problemID), -1)
	file := "p" + strconv.Itoa(problemID) + ".go"
	err := ioutil.WriteFile(folder+"/"+file, []byte(data), 0664)
	panicIfError(err)

	data = strings.Replace(problemTestTemplate, problemIDString, strconv.Itoa(problemID), -1)
	file = "p" + strconv.Itoa(problemID) + "_test.go"
	err = ioutil.WriteFile(folder+"/"+file, []byte(data), 0664)
	panicIfError(err)
}
func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func getProblemDescription(problemID int) string {
	doc, err := goquery.NewDocument("https://projecteuler.net/problem=" + strconv.Itoa(problemID))
	if err != nil {
		log.Fatal(err)
	}
	info := ""
	doc.Find("h2").Each(func(i int, s *goquery.Selection) {
		// class, _ := s.Attr("id")
		// fmt.Println(class, s.Text())
		info += "Title:\n" + s.Text()
	})
	info += "\n\n"
	doc.Find(".info").Each(func(i int, s *goquery.Selection) {
		// class, _ := s.Attr("class")
		// fmt.Println(class, s.Text())
		info += "Info:\n" + s.Text()
	})
	info += "\n\n"
	doc.Find(".problem_content").Each(func(i int, s *goquery.Selection) {
		// class, _ := s.Attr("class")
		// fmt.Println(class, s.Text())
		info += "Description:" + s.Text()
	})
	return info
}

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
	problemIDFlag := flag.Int("new", 1, "new <problem ID>")
	flag.Parse()
	problemID := *problemIDFlag
	println("ProblemID: " + strconv.Itoa(problemID))
	folder := "p" + strconv.Itoa(problemID)
	err := os.Mkdir(folder, 0771)
	if err != nil {
		log.Fatal("Folder " + folder + " already exists! Remove it to generate problem template.")

	}
	data := strings.Replace(problemTemplate, problemIDString, strconv.Itoa(problemID), -1)
	data = strings.Replace(data, problemDescriptionString, getProblemDescription(problemID), -1)
	file := "p" + strconv.Itoa(problemID) + ".go"
	err = ioutil.WriteFile(folder+"/"+file, []byte(data), 0664)
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
	info := "Project Euler - Problem: " + strconv.Itoa(problemID) + "\n\n"
	doc.Find("h2").Each(func(i int, s *goquery.Selection) {
		// class, _ := s.Attr("id")
		// fmt.Println(class, s.Text())
		info += "Title:\n" + s.Text()
	})
	info += "\n\n"
	doc.Find(".info").Each(func(i int, s *goquery.Selection) {
		info += "Info:\n" + s.Text()
	})
	info += "\n\n"
	doc.Find(".problem_content").Each(func(i int, s *goquery.Selection) {
		info += "Description:\n" + s.Text() + "{{end_block}}"
	})

	yuckyStrings := map[string]string{
		"\n\n\n":           "\n\n",
		"\n{{end_block}}":  "",
		"Description:\n\n": "Description:\n",
		"Title:\n":         "Title: ",
		"You are currently using a secure connectionInfo:\n": "",
		"Problem " + strconv.Itoa(problemID):                 ""}

	for k, v := range yuckyStrings {
		for i := 0; i < 5; i++ { // Make sure all get replaced.
			info = strings.Replace(info, k, v, -1)
		}
	}
	return info
}

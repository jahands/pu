package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	problemID := flag.Int("new", 1, "new <problem ID>")
	flag.Parse()

	folder := "p" + strconv.Itoa(*problemID)
	err := os.Mkdir(folder, 0771)
	panicIfError(err)

	data := strings.Replace(problemTemplate, problemIDString, strconv.Itoa(*problemID), -1)
	file := "p" + strconv.Itoa(*problemID) + ".go"
	err = ioutil.WriteFile(folder+"/"+file, []byte(data), 0664)
	panicIfError(err)

	data = strings.Replace(problemTestTemplate, problemIDString, strconv.Itoa(*problemID), -1)
	file = "p" + strconv.Itoa(*problemID) + "_test.go"
	err = ioutil.WriteFile(folder+"/"+file, []byte(data), 0664)
	panicIfError(err)
}
func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

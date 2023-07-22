package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	outputFlag string
	assetsFlag bool
)

func init() {
	flag.StringVar(&outputFlag, "o", "", "output file name")
	flag.BoolVar(&assetsFlag, "a", false, "download all images, styles, js")
	flag.Parse()
}

func main() {
	url := "https://www.vk.com"
	wget(url)
}

func wget(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Could not perform HTTP Get method: ", err.Error())
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("bad status: %s\n", resp.Status)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Если задано имя файла, то создаем его и сохраняем туда
	if outputFlag != "" {
		out, err := os.Create(outputFlag + ".html")
		if err != nil {
			fmt.Println("Could not create file: ", err.Error())
			os.Exit(1)
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			fmt.Println("Could not copy downloaded data to file: ", err.Error())
			os.Exit(1)
		}
	} else {
		// Считываем тело ответа
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Could not read response body: ", err.Error())
		}
		fmt.Println(string(b))
	}

}

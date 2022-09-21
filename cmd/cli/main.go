package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Program starting!")

	const (
		helloWorldFilePath = "./helloworld.png"
		helloWorldImageURL = "https://raw.githubusercontent.com/otiai10/gosseract/main/test/data/001-helloworld.png"
	)

	if err := downloadFile(helloWorldFilePath, helloWorldImageURL); err != nil {
		log.Fatal(err)
	}

	client := gosseract.NewClient()
	defer client.Close()

	if err := client.SetImage(helloWorldFilePath); err != nil {
		log.Fatal(err)
	}

	text, err := client.Text()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(text)
	fmt.Println("✅ Done :D")
}

func downloadFile(filepath string, url string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return err
	}

	return nil
}

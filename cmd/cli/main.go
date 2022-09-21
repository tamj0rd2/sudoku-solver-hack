package main

import (
	"embed"
	"fmt"
	"github.com/otiai10/gosseract"
	"log"
)

//go:embed testdata/*
var testDataDir embed.FS

// run through docker using `make run`
func main() {
	fmt.Println("Program starting!")

	client := newOCRClient()
	defer client.Close()

	for _, tc := range []struct {
		filePath     string
		expectedText string
	}{
		{
			filePath:     "testdata/hello world.png",
			expectedText: "Hello world!",
		},
		{
			filePath:     "testdata/1-9.png",
			expectedText: "123456789",
		},
	} {
		b, err := testDataDir.ReadFile(tc.filePath)
		if err != nil {
			log.Fatal(err)
		}

		gotText, err := client.ReadBytesAsString(b)
		if err != nil {
			log.Fatal(err)
		}

		if gotText != tc.expectedText {
			log.Fatalf("expected %q but got %s", tc.expectedText, gotText)
		}

		fmt.Println("✅ ", tc.filePath)
	}
}

type ocrClient struct {
	tessClient *gosseract.Client
}

func newOCRClient() *ocrClient {
	tessClient := gosseract.NewClient()
	return &ocrClient{tessClient: tessClient}
}

func (c *ocrClient) Close() error {
	return c.tessClient.Close()
}

func (c *ocrClient) ReadBytesAsString(bytes []byte) (string, error) {
	if err := c.tessClient.SetImageFromBytes(bytes); err != nil {
		return "", fmt.Errorf("failed to set image from bytes: %w", err)
	}

	text, err := c.tessClient.Text()
	if err != nil {
		return "", fmt.Errorf("failed to get text from image: %w", err)
	}

	return text, nil
}

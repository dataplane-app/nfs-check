package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	start := time.Now()

	c := time.Now().UTC().String()
	filename := "test.txt"

	var testfolder string
	if os.Getenv("DP_TEST_FOLDER") == "" {
		testfolder = ""
	} else {
		testfolder = os.Getenv("DP_TEST_FOLDER")
	}

	createFile := testfolder + filename

	/* Write check. */
	err := os.WriteFile(createFile, []byte(c), 0644)
	if err != nil {
		time.Sleep(time.Second * 3)
		panic("😩 File system not ready: write failure: will exit to retry")
	}

	/* Read check. */
	dat, err := os.ReadFile(createFile)
	if err != nil {
		time.Sleep(time.Second * 3)
		panic("😩 File system not ready: read failure: will exit to retry")
	}

	if string(dat) != c {
		time.Sleep(time.Second * 3)
		panic("😩 File system not ready: read write check failure: will exit to retry")
	}

	stop := time.Now()

	log.Println("File system ready", string(dat), fmt.Sprintf("%f", float32(stop.Sub(start))/float32(time.Millisecond))+"ms")

}

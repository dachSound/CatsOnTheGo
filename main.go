package main

import("fmt"
	"net/http"
	"io"
	"os"
	"time"
	"path/filepath"
	"strconv"
)

func fetchInfo(url string) (*http.Response, error){
	fetched, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching info:", err)
		return nil, err
	}
	fmt.Println(fetched.Status)
	return fetched, nil
}

func saveFile(fetched *http.Response, catNum int) error {
	dir := "cats"
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Println("Failed to create directory:", err)
		return  err
	}

	data, err := io.ReadAll(fetched.Body)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return  err
	}

	catName := fmt.Sprintf("cat%d.jpg", catNum)
	filePath := filepath.Join(dir, catName)

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		fmt.Println("Failed to save file:", err)
		return  err
	}
	fmt.Printf("Saved file successfully: %s\n", filePath)
	return nil
}

func findMaxFileNum(dir string) int {

	pattern := filepath.Join(dir, "cat*.jpg")

	matches, err := filepath.Glob(pattern)
	if err != nil  || len(matches) == 0 {
		return -1
	}
	maxNum := -1

	for _, match := range matches {

		filename := filepath.Base(match)

		numStr := filename[3 :len(filename)-4]

		num, err := strconv.Atoi(numStr)
		if err == nil && num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func main(){
	url := "https://cataas.com/cat"

	maxNum := findMaxFileNum("cats")

	fileNum := 0
	if maxNum != -1 {
		fileNum = maxNum + 1
	}

	for {
		var Delay time.Duration

		Delay = 1500 // <- Delay in milliseconds

		time.Sleep(Delay * time.Millisecond)
		fetchedInfo, err := fetchInfo(url)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		err = saveFile(fetchedInfo, fileNum)
		if err != nil {
			fmt.Println("Error saving file:", err)
			return
		}

		fetchedInfo.Body.Close()
		fileNum++
	}
}

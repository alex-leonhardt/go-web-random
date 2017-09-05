package main

import (
	"bufio"
	"math/rand"
	"net/http"
	"os"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func rootHandler(lines []string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		word := lines[rand.Intn(len(lines))]
		w.Write([]byte(word))
	}
}

func main() {

	lines, _ := readLines("/usr/share/dict/words")
	rh := rootHandler(lines) // <-- returns a http.HandlerFunc ;)

	http.HandleFunc("/", rh)
	http.ListenAndServe(":8080", nil)

}

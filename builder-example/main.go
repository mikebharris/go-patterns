package main

import "fmt"

func main() {
	doiChecker := aDoiChecker().withZipFilepath("/a/path/to/the/zip").build()
	results, _ := doiChecker.GetDoiResults()
	fmt.Println(results)
}

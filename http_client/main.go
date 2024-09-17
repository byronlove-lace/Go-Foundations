package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://warrocketwiki.com/Every_Story_Ever")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		respOutput := "Response Status: " + resp.Status
		panic(respOutput)
	}

	htmlContent := html.New
	responseContent := bufio.NewReader(resp.Body)

	fmt.Println(responseContent.Read)

}

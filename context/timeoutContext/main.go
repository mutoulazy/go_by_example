package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Result struct {
	r *http.Response
	err error
}

func main() {
	process()
}

func process() {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan Result, 1)
	req, err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil {
		fmt.Println("http request build failed, err: ", err)
		return
	}
	go func() {
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("http request failed, err: ", err)
		}
		result := Result{r: resp, err:err}
		c <- result
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		res := <-c
		fmt.Println("http request timeout, err: ", res.err)
	case res := <-c:
		defer res.r.Body.Close()
		out, err := ioutil.ReadAll(res.r.Body)
		if err != nil {
			fmt.Println("read response failed, err: ", err)
		}
		fmt.Printf("Server response: %s", out)
	}
	return
}
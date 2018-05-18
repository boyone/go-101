package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
)

func handle(w http.ResponseWriter, res *http.Request) {
	resultChan := make(chan string, 1)
	errChan := hystrix.Go("front", func() error {
		resp, err := http.Get("http://localhost:3333/account")
		if err != nil {
			log.Print("not working")
			return err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		resultChan <- string(body)
		return nil
	}, func(err error) error {
		return err
	})

	select {
	case result := <-resultChan:
		//log.Println("success:", result)
		fmt.Fprint(w, "account: "+result)
	case err := <-errChan:
		log.Println("failure:", err)
		//w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprint(w, "account: not found")
	}
}

func main() {
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "8881"), hystrixStreamHandler)
	http.HandleFunc("/account", handle)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

package c2

import (
	"context"
	"net/http"
	"fmt"
	"time"
	"io/ioutil"
	"sync"
)

var (
	wg sync.WaitGroup
)

type ResPack struct {
	r *http.Response
	err error
}

func work2(ctx context.Context) {
	client := &http.Client{}
	defer wg.Done()
	c := make(chan ResPack, 1)

	req, _ := http.NewRequest("GET", "http://localhost:9200", nil)
	req = req.WithContext(ctx)
	go func() {
		resp, err := client.Do(req)
		pack := ResPack{r: resp, err: err}
		c <- pack
	}()

	select {
	case <-ctx.Done():
		<-c
		fmt.Println("Timeout!")
	case res:= <-c:
		if res.err != nil {
			fmt.Println(res.err)
			return
		}
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Server Response: %s", out)
	}
	return
}


func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()
	wg.Add(1)
	go work2(ctx)
	wg.Wait()
	fmt.Println("Finished")
}

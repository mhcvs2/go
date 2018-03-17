package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
	"encoding/json"
	"bytes"
)

var (
	wg sync.WaitGroup
)

type Dashboard struct {
	Dashboard map[string]interface{} `json:"dashboard"`
	FolderId  int         `json:"folderId"`
	Overwrite bool        `json:"overwrite"`
}

type ResPack struct {
	r   *http.Response
	err error
}

var filepath = "/root/test/grafana/grafana/dashboards/test.json"

func work2(ctx context.Context) {
	dashboard, _ := ioutil.ReadFile(filepath)
	data := make(map[string]interface{})
	err := json.Unmarshal(dashboard, &data)
	if err != nil {
		fmt.Println(err.Error())
	}
	db := Dashboard{Dashboard:data, FolderId:0, Overwrite:true}
	rb, _ := json.Marshal(db)
	reqBody := &bytes.Buffer{}
	reqBody.Write(rb)
	fmt.Println(reqBody.String())
	client := &http.Client{}
	defer wg.Done()
	c := make(chan ResPack, 1)

	req, _ := http.NewRequest("POST", "http://localhost:3000/api/dashboards/db", reqBody)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer eyJrIjoicURWU3JndTNiRG02NDA1Zll0cFlsYTRzWVBnWXNudHkiLCJuIjoibWhjIiwiaWQiOjF9")
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
	case res := <-c:
		if res.err != nil {
			fmt.Println(res.err)
			return
		}
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Response status: %d\n", res.r.StatusCode)
		fmt.Printf("Server Response: %s\n", out)
	}
	return
}

//Response status: 412
//Server Response: {"message":"A dashboard with the same name already exists","status":"name-exists"}
//Response status: 400
//Server Response: {"message":"Dashboard title cannot be empty"}
//Response status: 200
//Server Response: {"slug":"docker-containers","status":"success","version":2}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	wg.Add(1)
	go work2(ctx)
	wg.Wait()
	fmt.Println("Finished")
}

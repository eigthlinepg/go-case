package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func calc(x, y, z int) int {

	return (x + y) * z
}

func randHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())

	var (
		i = rand.Intn(20)
		j = rand.Intn(20)
		k = rand.Intn(20)
	)

	result := calc(i, j, k)

	_, _ = w.Write([]byte(fmt.Sprintf("%d", result)))
}

func main() {

	// 启动一个http server
	http.HandleFunc("/rand", randHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("start server failed:%v", err)
	}

}

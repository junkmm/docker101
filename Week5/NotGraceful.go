package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"sync/atomic"
)

var requestID int32

func index(w http.ResponseWriter, r *http.Request) {
	// 순차적인 요청 번호를 생성
	requestNumber := atomic.AddInt32(&requestID, 1)

	// 5초 동안 "Processing..." 로그를 찍고 5초 후에 "처리 완료" 메시지 반환
	for i := 1; i <= 5; i++ {
		// 로그를 stdout으로 출력 (컨테이너 로그로 확인 가능)
		fmt.Printf("Request #%d: Processing... #%d\n", requestNumber, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintf(w, "처리 완료")
}

func main() {
	// 서버 시작
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":80", nil))
}

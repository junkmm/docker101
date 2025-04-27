package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"syscall"
	"context"
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
	// Create a server instance with a graceful shutdown setup
	server := &http.Server{
		Addr: ":80",
	}

	// Handle HTTP requests
	http.HandleFunc("/", index)

	// Start the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe failed: %v", err)
		}
	}()

	// Create a channel to listen for interrupt signals (SIGTERM or SIGKILL)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Wait for a shutdown signal
	<-stop

	// Graceful shutdown: Allow ongoing requests to finish (5 seconds timeout)
	fmt.Println("Received shutdown signal. Waiting for 5 seconds before shutting down...")
	time.Sleep(5 * time.Second)

	// Shutdown the server gracefully with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown failed: %v", err)
	}

	fmt.Println("Server gracefully stopped")
}

package main

import (
        "fmt"
        "os"
        "strconv"
        "time"
)

const logFilePath = "/var/log/custom.log"
const defaultSleepSeconds = 5

func main() {
        file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
                fmt.Fprintf(os.Stderr, "로그 파일 열기 오류 (%s): %v\n", logFilePath, err)
                fmt.Fprintln(os.Stderr, "팁: /var/log/ 디렉토리에 파일을 쓰려면 root 권한(sudo)으로 프로그램을 실행해야 할 수 있습니다.")
                os.Exit(1)
        }
        defer file.Close()

        maxEnv := os.Getenv("MAX")
        sleepSeconds := defaultSleepSeconds
        if maxEnv != "" {
                maxVal, err := strconv.Atoi(maxEnv)
                if err == nil && maxVal > 0 {
                        sleepSeconds = maxVal
                } else {
                        fmt.Fprintf(os.Stderr, "환경 변수 MAX의 값이 올바르지 않습니다. 기본값(%d초)을 사용합니다: %v\n", defaultSleepSeconds, err)
                }
        }

        for i := 1; i <= sleepSeconds; i++ {
                message := fmt.Sprintf("Sleep #%d", i)
                fmt.Println(message)
                if _, err := fmt.Fprintln(file, message); err != nil {
                        fmt.Fprintf(os.Stderr, "로그 파일 쓰기 오류: %v\n", err)
                }
                time.Sleep(1 * time.Second)
        }

        fmt.Println("Exit")
}

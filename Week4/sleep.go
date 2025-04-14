package main

import (
        "fmt"
        "os" // 파일 시스템 작업을 위해 os 패키지를 추가합니다.
        "time"
)

// 로그 파일 경로를 상수로 정의합니다.
const logFilePath = "/var/log/custom.log"

func main() {
        // 로그 파일을 추가(Append) 모드로 열거나 생성합니다.
        // os.O_APPEND: 파일이 존재하면 끝에 내용을 추가합니다.
        // os.O_CREATE: 파일이 존재하지 않으면 생성합니다.
        // os.O_WRONLY: 쓰기 전용 모드로 엽니다.
        // 0644: 파일 권한 설정 (소유자는 읽기/쓰기, 그룹과 다른 사용자는 읽기만 가능)
        file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
                // 파일 열기 중 에러가 발생하면 에러 메시지를 표준 에러(stderr)에 출력하고 프로그램을 종료합니다.
                fmt.Fprintf(os.Stderr, "로그 파일 열기 오류 (%s): %v\n", logFilePath, err)
                // /var/log 디렉토리는 일반적으로 root 권한이 필요합니다.
                fmt.Fprintln(os.Stderr, "팁: /var/log/ 디렉토리에 파일을 쓰려면 root 권한(sudo)으로 프로그램을 실행해야 할 수 있습니다.")
                os.Exit(1) // 에러 코드 1로 종료
        }
        // defer 키워드를 사용하여 main 함수가 종료되기 직전에 파일이 반드시 닫히도록 합니다.
        // 이렇게 하면 에러가 발생하더라도 파일 핸들이 누수되지 않습니다.
        defer file.Close()

        // 1부터 5까지 반복합니다.
        for i := 1; i <= 5; i++ {
                // 콘솔에 출력하고 파일에 저장할 메시지를 생성합니다.
                message := fmt.Sprintf("Sleep #%d", i)

                // 1. 콘솔에 메시지를 출력합니다.
                fmt.Println(message)

                // 2. 파일에 동일한 메시지를 씁니다.
                // fmt.Fprintln 함수는 io.Writer 인터페이스(file 객체가 구현함)에
                // 문자열과 함께 줄바꿈 문자를 추가하여 씁니다.
                if _, err := fmt.Fprintln(file, message); err != nil {
                        // 파일 쓰기 중 에러가 발생하면 표준 에러에 메시지를 출력합니다.
                        // (여기서는 에러가 나도 프로그램을 중단하지는 않습니다.)
                        fmt.Fprintf(os.Stderr, "로그 파일 쓰기 오류: %v\n", err)
                }

                // 1초 동안 실행을 멈춥니다.
                time.Sleep(1 * time.Second)
        }

        // 반복문이 끝나면 콘솔에만 "Exit"를 출력합니다. (로그 파일에는 기록 안 함)
        fmt.Println("Exit")
}

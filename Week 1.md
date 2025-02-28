# Go 언어로 웹서버를 개발하고 컨테이너화 하기

## EC2 인스턴스 생성 후 접근
- 테스트 계정으로 EC2 인스턴스 생성 및 접근(Public)


## Go 개발환경 구성
golang 설치(Go언어 개발에 필요한 도구)
```
sudo dnf install golang -y
```

go version 확인
```
go version
```

go 개발 환경 구성
```
mkdir ~/go
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
source ~/.bashrc
```

기본 Go 코드 작성
```
mkdir -p ~/go/src/hello
cd ~/go/src/hello
vi hello.go
```

간단한 HTTP 응답 코드 구성
```
package main

import (
        "fmt"
        "log"
        "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World!!!!!!23434234")
}

func main() {
        http.HandleFunc("/", index)
        log.Fatal(http.ListenAndServe(":80", nil))
}
```

## Go 웹서버 실행
go 웹 서비스 실행
```
sudo go run hello.go
#인스턴스에 웹 요청
```

go build로 바이너리 파일 추출 후 실행
```
sudo go build hello.go
./hello.go
#인스턴스에 웹 요청
```

## 컨테이너화 진행
Docker 설치
```
sudo dnf install docker -y
sudo systemctl start docker
```

ec2-user에 docker 실행 권한 부여
```
#권장되지 않음
sudo chmod 666 /var/run/docker.sock
```

Dockerfile 생성
`vi Dockerfile`
```
FROM golang:1.20.1-alpine3.17 AS builder
WORKDIR /work
COPY . /work/
RUN go build hello.go

FROM alpine:3.14
COPY --from=builder /work/hello /work/hello
ENTRYPOINT ["/work/hello"]
```

docker build 명령으로 컨테이너 이미지 생성
```
docker build -t web .
```

컨테이너 이미지 확인
```
docker image ls
```

컨테이너 실행
```
docker run -dp 80:80 --rm web
```

docker stop 으로 기존 컨테이너 중지
```
docker stop ${container-id}
```

hello.go 코드 업데이트
```
vi hello.go
```

다시 docker build
```
docker build -t web:v2 .
```

다시 docker run
```
docker run -dp 80:80 --rm web:v2
```

도커 이미지를 Docker Hub에 업로드
```
# https://hub.docker.com/에 로그인
# 새로운 Repository 생성
docker login
docker tag web:v2 kimhj4270/${reponame}${tagname}
docker push kimhj4270/${reponame}${tagname}
```
## 새로운 EC2 인스턴스에서 컨테이너 실행
docker 설치
```
sudo dnf install docker -y
sudo systemctl start docker
sudo chmod 666 /var/run/docker.sock
```

docker run(pull)
```
docker run -dp 80:80 --rm kimhj4270/dockerstudy:0228
```

웹 접속 확인

### Builder
FROM golang:1.13-alpine as builder
RUN apk update && apk add git && apk add ca-certificates

WORKDIR /usr/src/app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o main .


### Make executable image
FROM scratch

# aws mysql에 접속하기 위한 환경 변수 설정
# 로컬 서버에 저장된 환경변수 값을 docker에서도 사용
# aws 서버 안에 환경변수를 설정하는 것이 아니라 docker 컨테이너 안에 환경변수가 있어야 aws mysql에 접속 가능
ENV TRADING_USER $TRADING_USER
ENV TRADING_PASSWORD $TRADING_PASSWORD
ENV TRADING_ENDPOINT $TRADING_ENDPOINT

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/src/app/main /main

CMD [ "/main" ]
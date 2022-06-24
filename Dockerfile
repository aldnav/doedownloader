# syntax=docker/dockerfile:1

FROM golang:1.18

RUN echo quit | openssl s_client -showcerts -servername doe.gov.ph -connect doe.gov.ph:443 2>/dev/null | awk '/BEGIN/,/END/{ if(/BEGIN/){c=""};c=c $0 "\n"}END{print c}' >ca.crt || true \
    && cp ca.crt /usr/local/share/ca-certificates \
    && update-ca-certificates \
    && rm ca.crt

VOLUME [ "/app/data" ]
WORKDIR /app
ENV REPORTS_DIRECTORY=/app/data COOKIE_PATH=/tmp/cookies.json
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /downloader

CMD [ "/downloader" ]
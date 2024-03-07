FROM golang:1.22.0 as builder

WORKDIR /build

COPY go.mod go.sum ./
COPY . ./
COPY cmd/main.go ./

RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin


FROM scratch

COPY --from=builder ./build/bin ./bin
COPY --from=builder ./config/local.yaml config.yaml

CMD ["./bin"]
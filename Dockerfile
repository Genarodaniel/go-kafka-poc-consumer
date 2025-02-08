FROM golang:latest


WORKDIR /app

COPY . .

RUN cd /app/cmd && GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server .


CMD [ "./cmd/server" ]
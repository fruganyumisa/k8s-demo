FROM golang:1.7
RUN mkdir /app 
WORKDIR /app

COPY . /app

RUN go build ./app/cmd/main.go

CMD [ "./app/cmd/main" ]


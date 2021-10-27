#FROM golang:1.12.0-alpine3.9
FROM golang:alpine 
RUN apk add git
RUN go get gitlab.com/Prosp3r/ori/pb
RUN go get google.golang.org/grpc
RUN go get google.golang.org/grpc/reflection
RUN mkdir /app
ADD ./server /app
COPY ./pb /pb
WORKDIR /app
RUN go build -o main .
EXPOSE 8080
CMD ["./app/main"]
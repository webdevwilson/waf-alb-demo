FROM golang:alpine
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 80
CMD ["./main"]
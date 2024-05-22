FROM golang:1.22-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY . .
RUN go mod download
RUN go build -o build/cmd ./cmd 
EXPOSE ${BOOKSWAP_PORT}
CMD [ "./build/cmd" ]
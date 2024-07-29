FROM golang:1.22-alpine3.20 as builder

WORKDIR /src
COPY go.sum go.mod ./
RUN go mod download

COPY . .
RUN go build -o /bin/oneauxilia-core-docs .

FROM alpine:3.20.1

RUN mkdir /app

WORKDIR /app 

COPY --from=builder /bin/oneauxilia-core-docs /app/oneauxilia-core-docs

COPY docs/oneauxilia.json docs/

RUN chmod +x /app/oneauxilia-core-docs
EXPOSE 8000

CMD ["/app/oneauxilia-core-docs"]
FROM golang:1.19 as build
WORKDIR /app
COPY . .
RUN go build -o /bin/walletview ./cmd/*

FROM golang:1.19
COPY --from=build /bin/walletview walletview
CMD ["./walletview"]

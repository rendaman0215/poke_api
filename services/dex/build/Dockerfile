#=============== Build ===============
FROM golang:1.24.4 AS build

WORKDIR /
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /main cmd/main.go

#============ Application ============
FROM gcr.io/distroless/static-debian11

COPY --from=build /main /main
CMD ["/main"]

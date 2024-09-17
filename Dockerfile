FROM golang:1.23 AS backendbuilder
WORKDIR /src
COPY backend /src
RUN go build -o /bin/svc ./cmd/httpserver/*

FROM scratch
COPY --from=backendbuilder /bin/svc /bin/svc
CMD ["/bin/svc"]

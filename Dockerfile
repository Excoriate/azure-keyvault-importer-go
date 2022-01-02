FROM golang:1.14 as build

WORKDIR /src

COPY . .

RUN go build -v -o chilecli main.go

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=build /src/chilecli /
CMD ["/chilecli"]
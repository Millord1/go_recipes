FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Install tools to crop and split pdfs (to prepare pdf before reading them)
RUN apt-get update && apt-get install -y --no-install-recommends \
    mutool \
    texlive-extra-utils \
    && rm -rf /var/lib/apt/lists/*

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-go-users

EXPOSE 8080
CMD ["go", "run", "main.go"]
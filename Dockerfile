# Specifies a parent image
FROM golang:1.21.5

# Creates an app directory to hold your app’s source code
WORKDIR /app

# Copies everything from your root directory into /app
COPY . .

# RUN go clean -modcache

# generates the files needed for HTML rendering
RUN for i in $(seq 1 10); do go install github.com/a-h/templ/cmd/templ@latest && break || sleep 5; done
RUN templ generate

# Installs Go dependencies
RUN for i in $(seq 1 10); do go mod download && break || sleep 5; done

#  fetches the relative dependencies 
RUN go get

# Builds your app with optional configuration
RUN go build -o /maxtrackr


# Tells Docker which network port your container listens on
EXPOSE 5000

# Specifies the executable command that runs when the container starts
CMD [ "./maxtrackr" ]
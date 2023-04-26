
# Set the base image to use
FROM golang:latest


# Set the working directory
WORKDIR /Go

COPY . .

RUN go get -u github.com/go-sql-driver/mysql

RUN go mod download


# Expose the port the application will listen on
EXPOSE 8080

# Run the application when the container starts
CMD ["./student-management"]

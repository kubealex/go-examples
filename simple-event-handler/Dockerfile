FROM golang as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o event-handler

FROM registry.access.redhat.com/ubi8/ubi-minimal

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/event-handler /event-handler

# Run the web service on container startup.
CMD ["/event-handler"]
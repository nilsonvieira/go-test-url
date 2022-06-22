# stage 1: build
FROM golang:1.15 as build
LABEL stage=intermediate
WORKDIR /app
COPY . .
RUN make build

# stage 2: scratch
FROM scratch as scratch
COPY --from=build /app/bin/tstnet /bin/tstnet
CMD ["tstnet"]
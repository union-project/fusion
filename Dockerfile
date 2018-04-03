FROM alpine:latest
COPY build/fusion /bin/app
WORKDIR /bin
ENTRYPOINT ["/bin/app"]
EXPOSE 8080

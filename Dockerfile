FROM scratch

WORKDIR /app
COPY ./go-automate-db .

ENTRYPOINT ["/app/go-automate-db"]
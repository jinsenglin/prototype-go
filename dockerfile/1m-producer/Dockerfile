FROM alpine:3.7
EXPOSE 8080
CMD ["./app"]

# ISSE https://github.com/GoogleCloudPlatform/google-cloud-go/issues/791
# FIX by this line
RUN apk --no-cache --update add ca-certificates

COPY ./out/1m-producer-linux-amd64 /app
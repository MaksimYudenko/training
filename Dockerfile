# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev
WORKDIR /app
# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV SRC_DIR=/go/src/github.com/MaksimYudenko/training
# Add the source code:
ADD . $SRC_DIR
# Build it:
RUN cd $SRC_DIR; go build -o training; cp training /app/
#RUN cd $SRC_DIR; cp myapp /app/
ENTRYPOINT ["./training"]

## iron/go is the alpine image with only ca-certificates added
#FROM iron/go
#WORKDIR /app
## Now just add the binary
#ADD myapp /app/
#ENTRYPOINT ["./myapp"]

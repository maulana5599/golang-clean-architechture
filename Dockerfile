from golang as builder

WORKDIR /app
ADD . .

RUN apt-get update && apt-get install -y \
    vim \
    git \
    curl 

RUN go build -o /usr/local/bin/voting

EXPOSE 8083
CMD ["/usr/local/bin/voting"]

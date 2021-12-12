FROM golang:1.17.5

RUN apt-get update && \
        apt-get -y upgrade && \
        apt-get install -y \
            gcc

RUN groupadd -r user && useradd -r -g user user

COPY --chown=root:user ./build/flag /home/user/flag
COPY --chown=root:user ./src/web.go /home/user/web.go

WORKDIR /home/user

COPY ./src/go.mod /home/user/go.mod
COPY ./src/go.sum /home/user/go.sum

RUN go mod download

RUN chmod 444 ./flag && \
    chmod 555 ./web.go

EXPOSE 30010

CMD ["go","run","web.go"]
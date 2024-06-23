FROM golang:1.23rc1-alpine

WORKDIR /app

COPY go.mod go.sum /app/src/

RUN cd /app/src && go mod download && go mod verify

COPY . /app/src

RUN cd /app/src/cmd && go build -v -o /app/bin/exe

RUN rm -rf /app/src

CMD [ "/app/bin/exe" ]
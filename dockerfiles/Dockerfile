FROM golang:1.16-alpine
ENV http_proxy ""
ENV https_proxy ""

WORKDIR /app

COPY src/server/go.mod ./
COPY src/server/go.sum ./
RUN go mod download

COPY src/server/*.go ./

RUN go build -o /server-fileupload

EXPOSE 4000

CMD [ "/server-fileupload" ]

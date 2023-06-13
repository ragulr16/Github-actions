FROM golang:alpine AS development

WORKDIR /app 

COPY . . 

RUN go build -o /app/build .

FROM alpine AS production

WORKDIR /app  

COPY --from=development /app/build . 

EXPOSE 6555

CMD [ "./build" ]

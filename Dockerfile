FROM alpine:3.15

WORKDIR /app

COPY ./skmc-parser /app/skmc-parser
RUN chmod +x ./skmc-parser

EXPOSE 8000

CMD ["./skmc-parser"]

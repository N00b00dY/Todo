FROM alpine:latest

RUN mkdir /app

COPY dbServiceApp /app

CMD [ "/app/dbServiceApp"]
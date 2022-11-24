FROM alpine:latest

RUN mkdir /app

COPY frontEndServiceApp /app

CMD [ "/app/frontEndServiceApp"]
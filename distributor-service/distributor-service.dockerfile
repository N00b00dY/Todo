FROM alpine:latest

RUN mkdir /app

COPY distributorServiceApp /app

CMD [ "/app/distributorServiceApp"]
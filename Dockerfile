FROM alpine

RUN mkdir /app
WORKDIR /app

COPY bin/app .
RUN mkdir /html
COPY html/* html/

CMD ./app

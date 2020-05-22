FROM        golang:1.14.2-alpine3.11
WORKDIR     /home
COPY        . .
ENV         GO111MODULE on
EXPOSE      1984
ENTRYPOINT  ["go", "run", "."]

FROM golang:1.19.1-bullseye as dev
WORKDIR /app
RUN echo "alias ll='ls -lahF --color=auto'" >> ~/.bashrc && . ~/.bashrc
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/k0kubun/sqldef/cmd/mysqldef@latest
CMD ["air"]

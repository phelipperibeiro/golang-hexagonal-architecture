FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN go get -u github.com/spf13/cobra@latest && \
    go install github.com/golang/mock/mockgen@v1.5.0 && \
    go install github.com/spf13/cobra-cli@latest


RUN apt-get update && apt-get install sqlite3 -y

RUN usermod -u 1000 www-data
RUN mkdir -p /var/www/.cache
RUN chown -R www-data:www-data /go
RUN chown -R www-data:www-data /var/www/.cache
USER www-data

CMD ["tail", "-f", "/dev/null"]

# acessar o container
# para entrar no sqlite, executar => sqlite3 db.sqlite
# criar table no banco => create table products(id string, name string, price float, status string);
# depois confirmar a criação com =>  .tables

# go test ./...

# mockgen -destination=application/mocks/application.go -source=application/product.go application
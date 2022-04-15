FROM golang:1.17 AS build
ADD . /app
WORKDIR /app
RUN go build main.go

#docker build -t bronim .
#docker run -p 5000:5000 bronim
#docker stop $(docker ps -a -q)
#docker rm $(docker ps -a -q)

FROM ubuntu:20.04
RUN apt-get -y update && apt-get install -y tzdata
#COPY scripts/init.sql /
#ENV PGVER 12
#RUN apt-get -y update && apt-get install -y postgresql-$PGVER

#USER postgres
#RUN /etc/init.d/postgresql start &&\
#    psql --command "ALTER USER postgres WITH SUPERUSER PASSWORD 'password';" &&\
#    /etc/init.d/postgresql stop
#EXPOSE 5432

#RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/$PGVER/main/pg_hba.conf
#RUN echo "listen_addresses='*'" >> /etc/postgresql/$PGVER/main/postgresql.conf

#VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

USER root
WORKDIR user/src/app
COPY . .
COPY --from=build /app/main .
EXPOSE 5000

#ENV PGPASSWORD password
#CMD service postgresql start && psql -h localhost -d postgres -U postgres -p 5432 -a -q -f scripts/init.sql && ./main
CMD ./main
FROM mysql/mysql-server:8.0.23

ENV MYSQL_ROOT_PASSWORD rootpassword

WORKDIR /
RUN mkdir /database
COPY . /database
ADD ./init_db.sql /docker-entrypoint-initdb.d
ADD ./load_fake_data_db.sql /docker-entrypoint-initdb.d

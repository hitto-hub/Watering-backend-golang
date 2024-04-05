FROM mysql

RUN echo "##### command install ##########"

COPY init/* /docker-entrypoint-initdb.d/

CMD ["mysqld"]

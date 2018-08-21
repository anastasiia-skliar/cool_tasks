FROM ubuntu:latest

COPY src/cool_tasks /opt/cool_tasks/
COPY config.json /opt/cool_tasks/config/

RUN chmod +x /opt/cool_tasks/cool_tasks
EXPOSE 8080

WORKDIR /opt

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
RUN chmod +x /wait

CMD cat /opt/cool_tasks/config/config.json && /wait && /opt/cool_tasks/cool_tasks -config /opt/cool_tasks/config/config.json

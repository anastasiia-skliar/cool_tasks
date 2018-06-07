FROM ubuntu:latest

COPY cool_tasks /opt/cool_tasks/bin/
COPY config.json /opt/cool_tasks/config/

RUN chmod +x /opt/cool_tasks/bin/*

EXPOSE 8080

WORKDIR /opt

ENTRYPOINT ["/opt/cool_tasks/bin/cool_tasks"]
CMD ["-config", "/opt/cool_tasks/config/config.json"]

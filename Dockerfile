FROM scratch

COPY cool_tasks /opt/cool_tasks/bin/
COPY config.json /opt/cool_tasks/config/

EXPOSE 8080

ENTRYPOINT ["/opt/cool_tasks/bin/cool_tasks"]
CMD ["-config", "/opt/cool_tasks/config/config.json"]

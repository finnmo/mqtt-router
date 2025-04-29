FROM eclipse-mosquitto:2.0

# Copy helper scripts
COPY generate_mosquitto_passwd.sh /usr/local/bin/
COPY entrypoint.sh               /usr/local/bin/
RUN chmod +x /usr/local/bin/*.sh

EXPOSE 1883
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]

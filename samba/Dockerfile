FROM ubuntu:24.04

ENV SMB_USER="some_user"
ENV SMB_PASSWD="some_password"

RUN apt-get update && apt-get install -y samba
RUN mkdir /shared

# Set up Samba
COPY samba/docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +rx /docker-entrypoint.sh

EXPOSE 139 445

# Configure user for samba
ENTRYPOINT ["/docker-entrypoint.sh"]

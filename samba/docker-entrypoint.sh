#!/bin/bash

useradd -Ms /bin/bash ${SMB_USER}
echo "${SMB_USER}:${SMB_PASSWD}" | chpasswd
echo -ne "${SMB_PASSWD}\n${SMB_PASSWD}\n" | smbpasswd -a -s ${SMB_USER}

service smbd start

tail -f /var/log/samba/log.smbd

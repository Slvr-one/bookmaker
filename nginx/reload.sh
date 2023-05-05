#!/bin/bash
set -eu

# Domain="dviross.net"

nginx=$1
reload="nginx -ts reload"
docker exec $nginx $reload
# Obtain SSL/TLS Cert
# sudo certbot --nginx -d $Domain

#Automatically Renew Let’s Encrypt Certificates
# crontab -e <<< '0 12 * * * /usr/bin/certbot renew --quiet'  #every day at noon
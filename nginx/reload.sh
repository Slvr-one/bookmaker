#!/bin/bash
set -eu

# domain="dviross.net"

nginx=$1
reload="nginx -ts reload"
docker exec $nginx $reload

# sudo certbot --nginx -d $domain # Obtain SSL/TLS Cert

#Automatically Renew Let’s Encrypt Certificates
# crontab -e <<< '0 12 * * * /usr/bin/certbot renew --quiet'  #every day at noon
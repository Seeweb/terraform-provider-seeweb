#!/bin/sh -e
#You can run this script to export the ticket to environment variables
#Ex: source get_token.sh

APIURL="https://api.seeweb.it/ecs/v2/login"


read -p "Username: " USERNAME
read -sp "Password : " PASSWORD
echo " "
echo -n "Token: "

SEEWEB_TOKEN=$(curl --silent -X POST $APIURL -F username=$USERNAME -F password=$PASSWORD |jq -r '.token')
echo $SEEWEB_TOKEN

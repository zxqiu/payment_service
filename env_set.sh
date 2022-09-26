#!/bin/bash

# set up kafka go client
# sudo apt update && sudo apt upgrade -y
# sudo apt-get -f install build-essential pkg-config
# go get gopkg.in/confluentinc/confluent-kafka-go.v1/kafka


# setup mysql on ubuntu

MYSQL_INSTALL_GROUPNAME=mysql
MYSQL_INSTALL_USERNAME=mysql
SERVICE_USERNAME=zxqiu
SERVICE_PWD=12345

sudo apt install mysql-server -y

# setup configurations
if test -f "/etc/my.cnf"; then
  echo "/etc/my.cnf already exist. It will not be updated"
else
  echo "creating new /etc/my.cnf"
  echo "[mysqld]" > my.cnf
  echo "datadir=/var/lib/mysql" >> my.cnf
  echo "socket=/tmp/mysql.sock" >> my.cnf
  echo "symbolic-links=0" >> my.cnf
  echo "character-set-server=utf8" >> my.cnf
  echo "collation-server=utf8_general_ci" >> my.cnf
  echo "[mysqld_safe]" >> my.cnf
  sudo mv my.cnf /etc/my.cnf
fi

# there might be some issues with these folders. fix them first.
sudo chown mysql:mysql /var/log/mysql
sudo chown mysql:mysql /var/lib/mysql
sudo chmod 750 /var/lib/mysql


# initialize and start service
sudo mysql --initialize-insecure --user=mysql
sudo systemctl restart mysql.service

# create user zxqiu
mysql -u root --execute="ALTER USER 'root'@'localhost' IDENTIFIED BY 'root'"
mysql -u root --password="root" --execute="CREATE USER '$SERVICE_USERNAME'@'localhost' IDENTIFIED BY '$SERVICE_PWD'"
mysql -u root --password="root" --execute="GRANT ALL PRIVILAGES ON *.* TO '$SERVICE_USERNAME'@'localhost'"

# create db
mysql -u $SERVICE_USERNAME --password="$SERVICE_PWD" --execute="CREATE DATABASE payment_service_db"

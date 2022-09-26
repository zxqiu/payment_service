#!/bin/bash

# set up kafka go client
# sudo apt update && sudo apt upgrade -y
# sudo apt-get -f install build-essential pkg-config
# go get gopkg.in/confluentinc/confluent-kafka-go.v1/kafka

MYSQL_INSTALL_GROUPNAME=mysql
MYSQL_INSTALL_USERNAME=zxqiu
SYSTEM_USERNAME=zxqiu
CODE_PATH=$(pwd)
MYSQL_PATH=/usr/local

DATABASE_USER_NAME=root
DATABASE_USER_PASSWD=root

# setup configurations
echo "[mysqld]" > my.cnf
echo "datadir="$MYSQL_PATH"/mysql/data" >> my.cnf
echo "socket=/tmp/mysql.sock" >> my.cnf
echo "symbolic-links=0" >> my.cnf
echo "character-set-server=utf8" >> my.cnf
echo "collation-server=utf8_general_ci" >> my.cnf
echo "[mysqld_safe]" >> my.cnf
sudo mv my.cnf /etc/my.cnf

echo "export PATH="$(pwd)"/bin/:$PATH" > mysql.sh
sudo mv mysql.sh /etc/profile.d/
source /etc/profile.d/mysql.sh

# start mysql server and init root user
killall mysqld
support-files/mysql.server start
bin/mysqladmin -u $DATABASE_USER_NAME password $DATABASE_USER_PASSWD
sudo cp support-files/mysql.server /etc/init.d/mysql.server

# create wallet database for project
bin/mysql -u $DATABASE_USER_NAME --password=$DATABASE_USER_PASSWD --execute="create database wallet"

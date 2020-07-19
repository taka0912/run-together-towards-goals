#!/bin/bash

# MySQLサーバーが起動するまで待機する
until mysqladmin ping -h mysql -P 3306 --silent; do
    echo 'waiting for mysqld to be connectable...'
    sleep 3
done

echo 'mysqld is connect !'

#!/bin/bash

APP_DIR=code
APP_DIR_CONF=code_conf

echo Starting MongoDB

service mongod start

echo Starting application

/var/www/$APP_DIR/bin/$APP_DIR -Dconfig.file=/var/www/$APP_DIR_CONF/production.conf -Dlogger.file=/var/www/$APP_DIR_CONF/logback.xml &

sleep 5

echo Removing log artifact
rm -r logs

echo Done!

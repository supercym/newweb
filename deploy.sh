#! /bin/sh

kill -9 $(pgrep webserver)
cd ~/newweb/
git pull https://github.com/supercym/newweb.git
cd webserver/
chmod u+x webserver
./webserver &
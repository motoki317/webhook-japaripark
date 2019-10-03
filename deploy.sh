#!/bin/bash

cd /srv/japaripark || exit
sudo git fetch
sudo git reset --hard origin/master
sudo yarn
sudo yarn build
sudo systemctl restart japaripark

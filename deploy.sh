#!/bin/bash

cd /srv/japaripark || exit
sudo git pull origin master
sudo yarn
sudo yarn build
sudo systemctl restart japaripark

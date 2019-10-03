#!/bin/bash

cd /srv/japaripark || exit
git pull origin master
yarn
yarn build
sudo systemctl restart japaripark

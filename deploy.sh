#!/bin/bash

cd /srv/japaripark || exit
yarn
yarn build
sudo systemctl restart japaripark

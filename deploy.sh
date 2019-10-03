#!/bin/bash

cd /srv || exit
yarn
yarn build
sudo systemctl restart japaripark

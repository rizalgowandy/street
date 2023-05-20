#!/usr/bin/env bash

SERVER=root@benalu.dev

( cd .. &&
GOOS=linux GOARCH=amd64 go build -o street.exe ) &&
mv ../street.exe . &&
rsync --delete -a --exclude='node_modules' --exclude 'tmp' --exclude='*.svelte' ../svelte/ svelte &&
cp ../.env.override . &&
upx street.exe &&
rsync -avz --progress . $SERVER:/home/street &&
ssh $SERVER sudo bash /home/street/reload_services.sh


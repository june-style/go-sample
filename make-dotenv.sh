#! /bin/bash

if [ ! -e ".env" ]; then
    cp .env.template .env
    echo "> Please set the execution result of the following command to the environment variable APPLICATION_HMAC_SECRET."
    echo "> openssl genrsa 4096"
fi

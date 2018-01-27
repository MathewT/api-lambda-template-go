#! /bin/bash

curl -v -XPOST -H "Content-Type: application/json" \
  -H "Authorization: Bearer someJWT" \
  -d '{ "foo":"bar"}' \
  "https://0wicntrljb.execute-api.us-east-1.amazonaws.com/dev/t2"
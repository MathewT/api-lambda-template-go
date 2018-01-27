#! /bin/bash

rm -f main deployment.zip
go build -o main
zip deployment.zip main

aws --profile=<profile> lambda update-function-code \
  --region us-east-1 \
  --function-name <function name> \
  --zip-file fileb://./deployment.zip 
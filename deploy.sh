#!/bin/bash

make && aws cloudformation package --template-file s.yaml --output-template-file example.out.yaml --s3-bucket lambda-deploy-temp; aws cloudformation deploy --template-file /naz/public/development/lambda/example.out.yaml --stack-name lambda-go-frontpage --region us-west-2 --capabilities CAPABILITY_IAM

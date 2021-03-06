#!/bin/bash
build.sh
echo "Deploying..."
aws lambda create-function --function-name FindAllAccounts \
--zip-file fileb://./deployment.zip \
--runtime go1.x --handler main \
--role arn:aws:iam::123456789012:role/FindAllAccountsRole \
--region eu-west-1 \
--profile admin
echo "Deploy completed"

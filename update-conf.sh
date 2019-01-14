#!/bin/bash
build.sh
echo "Updating configuration..."
aws lambda update-function-configuration --function-name FindAllAccounts \
--environment Variables={TABLE_NAME=rezerw-accounts} \
--region eu-west-1 \
--profile admin
echo "Configuration updated"



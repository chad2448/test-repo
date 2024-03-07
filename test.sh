#!/bin/bash

mkdir ~/.aws

cat <<EOF > ~/.aws/credentials
[default]
aws_access_key_id = $AWS_ACCESS_KEY
aws_secret_access_key = $AWS_SECRET_ACCESS_KEY
EOF

yum install awscli
aws s3 ls

#!/bin/bash

mkdir ~/.aws

cat <<EOF > ~/.aws/credentials
[default]
aws_access_key_id = $AWS_ACCESS_KEY
aws_secret_access_key = $AWS_SECRET_ACCESS_KEY
EOF

cat ~/.aws/credentials

#apt update
#apt install -y curl awscli unzip
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
tar xzf awscliv2.zip
sh ./aws/install
aws s3 ls

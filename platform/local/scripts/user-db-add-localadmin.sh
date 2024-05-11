#!/bin/bash

# MySQL connection details
DB_HOST="127.0.0.1"
DB_USER="develop"
DB_PASSWORD="example"
DB_NAME="anytore"

# SQL query to insert a user
INSERT_QUERY="INSERT INTO users (name, email, password) VALUES ('localadmin', 'localadmin@example.com', '\$2a\$10\$pTCiFiiePZd2GaHul5nzSOeDuoSVF.Z3w3KXbKQEB.r7GugsMq3my');"

# Execute the SQL query
mysql -h $DB_HOST -u $DB_USER -p$DB_PASSWORD $DB_NAME -e "$INSERT_QUERY"

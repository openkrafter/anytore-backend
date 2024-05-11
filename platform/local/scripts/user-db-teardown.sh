#!/bin/bash

# MySQL connection details
DB_HOST="127.0.0.1"
DB_USER="develop"
DB_PASSWORD="example"
DB_NAME="anytore"

# SQL query to drop the user table
DROP_TABLE_QUERY="DROP TABLE users;"

# Execute the SQL query
mysql -h $DB_HOST -u $DB_USER -p$DB_PASSWORD $DB_NAME -e "$DROP_TABLE_QUERY"

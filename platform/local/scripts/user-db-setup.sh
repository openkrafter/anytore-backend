#!/bin/bash

# MySQL connection details
# DB_HOST="127.0.0.1"
DB_HOST="user-db"
DB_USER="develop"
DB_PASSWORD="example"
DB_NAME="anytore"

# SQL query to create the user table
CREATE_TABLE_QUERY="CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);"

# Execute the SQL query
mysql -h $DB_HOST -u $DB_USER -p$DB_PASSWORD $DB_NAME -e "$CREATE_TABLE_QUERY"

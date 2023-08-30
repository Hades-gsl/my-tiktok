#!/bin/bash

# Database connection details
username="root"
password="123456"
database="tiktok"

# Loop through SQL files in the current directory
cd db/sql

for sql_file in *.sql
do
    echo "Executing $sql_file"
    mysql -u $username -p$password $database < $sql_file
    echo "Finished executing $sql_file"
    echo
done

echo "All SQL files executed."

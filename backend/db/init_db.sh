# -o errexit
set -e
# -o xtrace
set -x


mysqld &

# create database
mysql --default-character-set=utf8mb4 < env.sql
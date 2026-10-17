#!/bin/sh

# 发生错误立即退出
set -e

echo "run db migration"
source /app/app.env
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"

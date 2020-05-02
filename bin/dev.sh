#! /bin/bash

cd "$(dirname "$0")/.."

case $1 in
  "migrate") docker-compose exec migrator sql-migrate ${2} -env="developmentDocker" --config=internal/pkg/db/dbconf.yml;;
  "migrate!") docker-compose exec migrator sql-migrate ${2} -env="developmentDocker" --config=internal/pkg/db/dbconf.yml -limit=0;;
  "seed") docker-compose exec app go run cmd/seeder/main.go ;;
  "reset") ./bin/dev.sh migrate! down && ./bin/dev.sh migrate! up && ./bin/dev.sh seed;;
         *) docker-compose $@ ;;
esac
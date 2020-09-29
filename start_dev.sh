#!/usr/bin/env bash

function ctrl_c() {
  echo "# cleaning up.."
#  echo "# kill getenvoy"
#  killall getenvoy
  echo "# kill npm"
  killall npm
  echo "# kill gin"
  killall gin
  echo "# kill db"
  killall docker-compose
}

ctrl_c
trap ctrl_c SIGINT

echo "# start db :5432"
(docker-compose up postgres1) &


echo "# start npm :5000"
(cd view; npm run dev) &

source_env() {
  [ "$#" -eq 1 ] && env="$1" || env=".env"
  [ -f "$env" ] || { echo "Env file $env doesn't exist"; return 1; }
  eval $(grep -Ev '^#|^$' "$env" | sed -e 's/=\(.*\)/="\1/g' -e 's/$/"/g' -e 's/^/export /')
}

source_env database.env
source_env sf1.env

sleep 1 # wait docker compose
echo '# start gin app :9090'
(gin -i -b sf1.exe -p 9090 -a 9091)

ctrl_c

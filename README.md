
# API Demo

This is example of using PostgreSQL and Golang.

## Development mode

```
# clone source code
cd ~/go/src
git clone git@github.com:kokizzu/sf1.git

# add nodejs dependency
cd view
npm install

# gin autorecompile
go get -u -v github.com/codegangsta/gin

# add db to hosts
echo 127.0.0.1 postgres1 | sudo tee -a /etc/hosts

# start servers
./start_dev.sh

# then visit http://127.0.0.1:9090/swagger/index.html
```

## Build (Docker)

```
./build.sh
```

## Usage (Docker Compose)

```
docker-compose up

# then visit http://127.0.0.1:9091/swagger/index.html
```

## Troubleshooting

### Docker
```
# ensure already a docker user
sudo usermod -aG docker `whoami`

# or use tcp
export DOCKER_HOST=tcp://127.0.0.1:2375

# or if desperate (security hole)
sudo chmod 666 /var/run/docker.sock
```

### Database
```
docker-compose run database bash
psql --host=$POSTGRES_DB --username=$POSTGRES_USER --dbname=$POSTGRES_PASSWORD

```

## TODO
 - [x] push to docker hub
 - [ ] add envoy/caddy reverse proxy
 - [ ] add authentication
 - [ ] deploy script (with and without kube)
 - [ ] upload to rook/minio depends on the scale
 - [ ] add overseer for non kube deployment
 - [ ] cache every request (redis/aerospike/tarantool) depends on the scale
 - [ ] add sentry depends on the scale
 - [ ] add jeager depends on the scale
 - [ ] replace default log with zerolog/zap
 - [ ] implement the frontend
 
## Credits

(c) Kiswono Prayogo

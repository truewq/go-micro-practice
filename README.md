#0. generate proto buffer files
#### `cd scripts`
#### `./proto.sh`

#1. build
### 1.1 build all
#### `make`
### 1.2 build one service
#### `make cmd=auth_api`

# 2.run
### 2.1 run etcd
#### `./etcd &`
### 2.2 run micro api
#### `micro api --namespace tech.share.api --handler=http --address :8180 &`
### 2.3 run all service
#### `./scripts/start.sh`

#3. stop
### 3.1 stop micro api
#### `pkill -9 micro`
### 3.2 kill all service
#### `./scripts/stop.sh`

#4. restart
### 4.1 restart service
#### `./scripts/restart.sh`

#5. login
#### `curl -X POST "http://10.0.0.1:8180/auth/login" -H "Content-Type: application/json" -d "{ \"User\": \"tech\", \"Password\": \"share\"}"`
#6. getBill
#### `curl -X POST "http://10.0.0.1:8180/bill/getBills" -H "accept: application/json" -H "Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE2MDQ2NDI2ODEsIklkIjoiMTMyNDU2MzMyOTcwMzc0MzQ4OCIsIklzc3VlciI6InRlY2ggc2hhcmUiLCJOb3RCZWZvcmUiOjE2MDQ2MzU0ODEsIlVzZXIiOiJ0ZWNoIn0.sjNWBxhnUxKJjPbvr26tVXGjFDB3NwIb23RRmnIMoVA" -H "Content-Type: application/json" -d "{}"`
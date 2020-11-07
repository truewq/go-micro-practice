#更多技术分享视频，请关注公众号：几分钟技术分享。

#1. generate proto buffer files

`cd scripts`

`./proto.sh`

#2. build
1.1 build all services

`make`

1.2 build one service

`make cmd=auth_api`

#3. run

3.1 run etcd

`./etcd &`

3.2 run micro api

`micro api --namespace tech.share.api --handler=http --address :8180 &`

3.3 run all service

`./scripts/start.sh`

#4. stop

4.1 stop micro api

`pkill -9 micro`

4.2 kill all service

`./scripts/stop.sh`

#5. restart

5.1 restart service

`./scripts/restart.sh`

#6. login

`curl -X POST "http://10.0.0.1:8180/auth/login" -H "Content-Type: application/json" -d "{ \"User\": \"tech\", \"Password\": \"share\"}"`

#7. getBill

`curl -X POST "http://10.0.0.1:8180/bill/getBills" -H "accept: application/json" -H "Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE2MDQ2NDI2ODEsIklkIjoiMTMyNDU2MzMyOTcwMzc0MzQ4OCIsIklzc3VlciI6InRlY2ggc2hhcmUiLCJOb3RCZWZvcmUiOjE2MDQ2MzU0ODEsIlVzZXIiOiJ0ZWNoIn0.sjNWBxhnUxKJjPbvr26tVXGjFDB3NwIb23RRmnIMoVA" -H "Content-Type: application/json" -d "{}"`
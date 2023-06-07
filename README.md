# yomo-play

This program receives a streamed random number from a TCP server and splits the stream by '\n' in the source. 
The source then sends it to the zipper, the zipper sends it to the SFN, and the SFN calculates their average.


## Use yomo

```bash
# install 1.13.0 yomo cli
curl -fsSL https://get.yomo.run | sh

# run yomo zipper
yomo serve -c config.yaml

# tcp server with the source
go run main.go -broker=yomo

# run yomo sfn
cd calc && yomo build && yomo run sfn.wasm

# Mock random number to the tcp server
seq 10 | xargs -I{} -P 10 bash -c 'for (( ; ; )); do sleep 0.1; echo $RANDOM; done | telnet 192.168.31.125 8080'
```


## Use http server

```bash
# http hander as sfn
go run httpsrv/main.go

# tcp packet to http request
go run main.go -broker=http

# send mock data to the tcp server
seq 10 | xargs -I{} -P 10 bash -c 'for (( ; ; )); do sleep 0.1; echo $RANDOM; done | telnet 192.168.31.125 8080'
```
# yomo-play

This program receives a streamed random number from a TCP server and splits the stream by '\n' in the source. 
The source then sends it to the zipper, the zipper sends it to the SFN, and the SFN calculates their average.


### Use yomo 1.13.0 version.
```bash
curl -fsSL https://get.yomo.run | sh
```

### Run zipper

```bash
yomo serve -c config.yaml
```

### Run tcp server with the source

```bash
go run main.go
```

### Run sfn.

```bash
cd calc && yomo build && yomo run sfn.wasm
```

### Mock random number to the tcp server.

```bash
seq 10 | xargs -I{} -P 10 bash -c 'for (( ; ; )); do sleep 0.1; echo $RANDOM; done | telnet 192.168.31.125 8080'
```

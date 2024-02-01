# yomo playground

```bash
# start a zipper
yomo serve -c config.yaml

# run source
go run source/main.go

# build and run sfn
cd handler
yomo build -m go.mod && yomo run sfn.wasm
```
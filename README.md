# pi-go
A simple tool to download and sort pibenchmarks.com data

Inspired by this issue: https://github.com/TheRemote/PiBenchmarks/issues/26

To download all the data:

```bash
curl -k -L "https://pibenchmarks.com/api/boards/[1-37]/" -o "dump#1.json"
```

Then (for now)
```bash
mkdir data
mv dump* data/.
```

Run the script
```go
go run pi.go
```

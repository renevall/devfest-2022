# devfest-2022
Presentation about golang

# Commands


## Hello World
```
$ go run 01_hello_world/main.go
```

## Hello Webserver

```
# Tab 1
$ go run main.go
```

```
# Tab 2
$ curl http://localhost:8080
```

## Battletest Webserver

### Simple
```
# Tab 1
$ go run main.go
```


```
# Tab 2
$ echo 'GET http://localhost:8080' | vegeta attack -rate 200 -duration=30s | tee results.bin | vegeta report
```
### Pretty
```
echo 'GET http://localhost:8080' | \
    vegeta attack -rate 200 -duration 10m | vegeta encode | \
    jaggr @count=rps \
          hist\[100,200,300,400,500\]:code \
          p25,p50,p95:latency \
          sum:bytes_in \
          sum:bytes_out | \
    jplot rps+code.hist.100+code.hist.200+code.hist.300+code.hist.400+code.hist.500 \
          latency.p95+latency.p50+latency.p25 \
          bytes_in.sum+bytes_out.sum
```
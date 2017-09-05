# go-web-random
responds a random word from the words package

## pre-requisites
install the `words` package to provide `/usr/share/dict/words`

## run the service

```
go run main.go
```

or

```
go install && \
$GOPATH/bin/go-web-random
```

## run vegeta perf test

```
go get github.com/tsenart/vegeta
echo "GET http://localhost:8080" | vegeta attack -duration=60s -rate=1000 | vegeta report
```

Example output (1000rps): 
```
Requests      [total, rate]            60000, 1000.02
Duration      [total, attack, wait]    59.999532764s, 59.998999905s, 532.859µs
Latencies     [mean, 50, 95, 99, max]  538.425µs, 496.448µs, 812.888µs, 1.550336ms, 12.102291ms
Bytes In      [total, mean]            574652, 9.58
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:60000
Error Set:
```

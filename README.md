## INTRODUCTION

The purpose of this work is to get the overview of concurrency handling among Apache, NGINX and Go web servers.
This performance benchmarking is based on [C10K problem](https://en.wikipedia.org/wiki/C10k_problem).

## HOW TO
10000 requests are sent to servers at localhost concurrently and the completion time is measured for those requests.
The servers returns the same amount of data. In this scope of work, the return data is the id and the name of 249 countries with the size around 23kB/request.

For `Apache`, `NGINX` and `Go`, the requests are sent to `/static` to get the content of static file on server

For `Go` the more endpoints are created including:

-  `/static`: as mentioned above
-  `/template`: backend takes data, generates template and send to client
-  `/db`: backend gets data from PostgreSQL database, generates template and send to client


## TOOLS
- [ab - Apache HTTP server benchmarking tool](https://httpd.apache.org/docs/current/programs/ab.html): sends requests to server endpoints concurrently. Follow the below command to send requests.
  
  ```ab -n 10500 -c 10000 -g output.csv url```
  
  where `output.csv` stores the latency information for each request and `url` is the endpoint (http://localhost)
- [Gnuplot](http://www.gnuplot.info/): takes the data in `csv` files and visualizes latency in graphs. Follow the below commands to visualize data.

```
$ gnuplot
gnuplot> set title 'localhost'
gnuplot> set xlabel 'requests'
gnuplot> set ylabel 'ms'
gnuplot> set grid
gnuplot> set term png
gnuplot> set output 'result.png'
plot '/path/to/file1.csv' using 9 smooth sbezier with lines title"apache", \
/path/to/file2.csv' using 9 smooth sbezier with lines title"go", \
/path/to/file3.csv' using 9 smooth sbezier with lines title"nginx"
```



## RESULT

![static](https://github.com/truong11t2/server-benchmark/blob/main/result/static-apache-go-nginx.png)

Apache sever shows a great performance with some of requests but it takes too long for some requests. At the end, it cannot solve C10K problem.

On the otherhand, NGINX and Go have a stable performance for requests.

The difference 

### GO server

![go-result](https://github.com/truong11t2/server-benchmark/blob/main/result/static-template-db-go.png)

Only with Go server with different approaches including `static`, `template` and `database`. The server can easily handle 10000 concurrent requests

#### Static
```
Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0  545 118.6    536     939
Processing:    71  815 225.6    837    1249
Waiting:        5  251  71.8    256     362
Total:        955 1361 108.4   1373    1499

Percentage of the requests served within a certain time (ms)
  50%   1373
  66%   1424
  75%   1437
  80%   1448
  90%   1465
  95%   1476
  98%   1491
  99%   1494
 100%   1499 (longest request)
```

#### Template
```
Connection Times (ms)

              min  mean[+/-sd] median   max
Connect:        0  494 125.6    503     683
Processing:   283 1099 584.6    990    2393
Waiting:        0 1058 604.9    968    2390
Total:        699 1593 577.0   1357    3050

Percentage of the requests served within a certain time (ms)
  50%   1357
  66%   1680
  75%   2323
  80%   2392
  90%   2494
  95%   2546
  98%   2599
  99%   2816
 100%   3050 (longest request)
```

#### Database
```
Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0  525 123.5    542     705
Processing:   342 1108 577.2   1027    2414
Waiting:        0 1056 595.8    995    2411
Total:        897 1632 561.9   1462    3080

Percentage of the requests served within a certain time (ms)
  50%   1462
  66%   1763
  75%   1962
  80%   2191
  90%   2616
  95%   2700
  98%   2738
  99%   2775
 100%   3080 (longest request)
```

### Apache

```
Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0  502 502.8    178    1247
Processing:    37  528 2596.5    144   20130
Waiting:        0   80  85.7     53     512
Total:        113 1030 2588.6    419   20130

Percentage of the requests served within a certain time (ms)
  50%    419
  66%   1297
  75%   1350
  80%   1380
  90%   1420
  95%   1477
  98%   1517
  99%  20056
 100%  20130 (longest request)
 ```

### NGINX

```
Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0  461 144.5    451     868
Processing:    70  670 223.3    682    1020
Waiting:        0  153  51.4    154     870
Total:        909 1131  79.0   1134    1265

Percentage of the requests served within a certain time (ms)
  50%   1134
  66%   1170
  75%   1194
  80%   1206
  90%   1233
  95%   1247
  98%   1255
  99%   1257
 100%   1265 (longest request)
 ```

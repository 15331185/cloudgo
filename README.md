#cloudgo
--------------------------------------------------------------------------------
####代码和老师的类似，思路是利用negroni创建服务器->添加路由->运行服务器

##测试
<pre><code>
curl -v http://localhost:3000/hello/15331185
</code></pre>

<pre><code>
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 9090 (#0)
> GET /hello/15331185 HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.47.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=UTF-8
< Date: Tue, 14 Nov 2017 06:51:39 GMT
< Content-Length: 25
< 
{
  "name": "15331185"
}
* Connection #0 to host localhost left intact
</code></pre>

##压力测试
<pre><code>
ab -n 5000 -c 500 http://localhost:9090/hello/15331185
</code></pre>
<pre><code>
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 500 requests
Completed 1000 requests
Completed 1500 requests
Completed 2000 requests
Completed 2500 requests
Completed 3000 requests
Completed 3500 requests
Completed 4000 requests
Completed 4500 requests
Completed 5000 requests
Finished 5000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            9090

Document Path:          /hello/15331185
Document Length:        25 bytes

Concurrency Level:      500
Time taken for tests:   1.469 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      740000 bytes
HTML transferred:       125000 bytes
Requests per second:    3404.43 [#/sec] (mean)
Time per request:       146.867 [ms] (mean)
Time per request:       0.294 [ms] (mean, across all concurrent requests)
Transfer rate:          492.05 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   68 245.3      3    1035
Processing:     0   49  31.0     43     297
Waiting:        0   45  28.5     40     297
Total:          0  117 254.2     49    1162

Percentage of the requests served within a certain time (ms)
  50%     49
  66%     63
  75%     75
  80%     83
  90%    119
  95%   1042
  98%   1144
  99%   1152
 100%   1162 (longest request)
 </code></pre>
平均响应时间为146.867ms

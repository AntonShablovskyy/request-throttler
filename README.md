# request-throttler

request-throttler implements a throttling mechanism for HTTP GET requests. This implementation uses classic **token bucket** algorithm.

## Command-line arguments
+ **maxQPS** : Maximum queries per second. Default value is 100.
+ **requestsCount** : Total requests count. Default value is 1000.
+ **url** : Url to throttle. Default value is http://localhost:8080.

## Example run
`./request-throttler --maxQPS 1000 --requestsCount 100000 --url https://stat.mil.ru/`
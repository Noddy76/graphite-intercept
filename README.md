# graphite-intercept

Proxy to intercept Graphite traffic copy it to a file and replay the metrics on
to Graphite.

# Usage

```
./graphite-intercept -listen 0.0.0.0:2004 -target localhost:2003 -file metrics.log
```

| Option | Default        | Description                        |
| ------ | -------------- | ---------------------------------- |
| listen | localhost:0    | TCP/UDP on which to listen         |
| target | localhost:2003 | Target address and port            |
| file   | metrics.log    | File in which to write the metrics |

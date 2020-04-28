# graphite-intercept

Proxy to intercept Graphite traffic copy it to a file and replay the metrics on
to Graphite.

## Usage

```
./graphite-intercept -listen 0.0.0.0:2004 -target localhost:2003 -file metrics.log
```

| Option | Default        | Description                        |
| ------ | -------------- | ---------------------------------- |
| listen | localhost:0    | TCP/UDP on which to listen         |
| target | localhost:2003 | Target address and port            |
| file   | metrics.log    | File in which to write the metrics |

## Replaying Captured Data

Once data has been captured you can replay it into a Graphite instance. You will
need to configure the target Graphite to accept large numbers of new metrics
over a short period of time. You will also need to set the retention and
aggregation to keep values from the time at which you captured them or you can
alter the timestamps in the captured file. See `storage-schemas.conf` and
`carbon.conf` in `docker-test/configs`.

You can also launch Graphite using the `start-graphite.sh` script.

Replaying data into graphite can be done with `netcat`.

```bash
# Example replay of metrics.log into a local Graphite
<metrics.log netcat -c localhost 2003
```

# systemd

```
$> cp mapshaper-server.service.example mapshaper-server.service
```

Then adjust the settings in `mapshaper-server.service` as necessary. Now install the service:

```
$> cp mapshaper-server.service /etc/systemd/system/mapshaper-server.service
$> systemctl daemon-reload
$> systemctl enable mapshaper-server
$> systemctl start mapshaper-server
```

To test, try something like:

```
$> curl -s http://localhost:9001/api/innerpoint -d @fixtures/1745882083.geojson | jq '.features[].geometry'
{
  "type": "Point",
  "coordinates": [
    -122.38875600604932,
    37.61459515528007
  ]
}
```

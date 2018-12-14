# nginx Dynamic proxy_hide_header

This repo was created to demonstrate the problem described on StackOverflow: https://stackoverflow.com/questions/53754229/using-a-header-to-filter-proxied-response-headers/53759605#53759605

To run this add a line to your /etc/hosts:

```
127.0.0.1 example.localhost.com
```

Then spin up the app with: `docker-compose up --build`

Then hit it with curl

`Set-Cookies` filtered out
```
curl \
  http://example.localhost.com/api \
  -H 'X-No-Cookies: true' \
  -v
```

`Set-Cookies` allowed:
```
curl \
  http://example.localhost.com/api \
  -v
```

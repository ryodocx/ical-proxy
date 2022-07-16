# ical-proxy

## Usage

```sh
# edit .env
$ cat .env
ICALPROXY_REDMINE_URL="https://redmine.example.com/path/"
ICALPROXY_REDMINE_APIKEY="<redacted>"
ICALPROXY_REDMINE_QUERY="query_id=xxx"

# edit ./configs/ical.simple.rego

# run
$ docker compose up -d
$ curl http://localhost:8080/
```

## Subscribe
* macOS: https://support.apple.com/ja-jp/guide/calendar/icl1022/mac
* Outlook: https://support.microsoft.com/ja-jp/office/outlook-com-%E3%81%A7%E4%BA%88%E5%AE%9A%E8%A1%A8%E3%82%92%E3%82%A4%E3%83%B3%E3%83%9D%E3%83%BC%E3%83%88%E3%81%BE%E3%81%9F%E3%81%AF%E8%B3%BC%E8%AA%AD%E3%81%99%E3%82%8B-cff1429c-5af6-41ec-a5b4-74f2c278e98c

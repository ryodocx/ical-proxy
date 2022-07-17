# ical-proxy

A proxy server convert to iCal from any sources

![architecture](docs/architecture.drawio.png)

## Feature

* Input
  * [x] [Redmine Issue](https://www.redmine.org/projects/redmine/wiki/Rest_Issues)
  * [ ] [Redmine Version](https://www.redmine.org/projects/redmine/wiki/Rest_Versions)
  * [ ] [External iCalendar]()
  * [ ] [RSS 1.0]()
  * [ ] [RSS 2.0]()
  * [ ] [Atom]()
* Output
  * [x] [iCalendar(VEVENT)](https://datatracker.ietf.org/doc/html/rfc5545)

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
$ curl http://localhost:8080
```

## Subscribe

* macOS
  * https://support.apple.com/ja-jp/guide/calendar/icl1022/mac
* Outlook
  * https://support.microsoft.com/ja-jp/office/outlook-com-%E3%81%A7%E4%BA%88%E5%AE%9A%E8%A1%A8%E3%82%92%E3%82%A4%E3%83%B3%E3%83%9D%E3%83%BC%E3%83%88%E3%81%BE%E3%81%9F%E3%81%AF%E8%B3%BC%E8%AA%AD%E3%81%99%E3%82%8B-cff1429c-5af6-41ec-a5b4-74f2c278e98c

## Convert & Filtering

* Convert to iCal from any feed via [OPA/Rego](https://www.openpolicyagent.org/)
* OPA output is read as single VEVENT
* VEVENT Spec: https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.1
* example: [./configs](./configs/)

expected data after through OPA
```json
{
    "allowed": true, // if false, event is ignored
    "event": {
        "UID": "<UniqID>",
        "SUMMARY": "<title>",
        "DTSTART;VALUE=DATE": "YYYYMMDD"
        ︙
    }
}
```

iCal output example
```ics
BEGIN:VCALENDAR
BEGIN:VEVENT
UID:https://redmine.example.com/issues/1
SUMMARY:subject1
DESCRIPTION:description1
DTSTART;VALUE=DATE:20220717
TRANSP:TRANSPARENT
END:VEVENT
BEGIN:VEVENT
UID:https://redmine.example.com/issues/2
SUMMARY:subject2
DESCRIPTION:description2
DTSTART;VALUE=DATE:20220718
TRANSP:TRANSPARENT
END:VEVENT
END:VCALENDAR
```

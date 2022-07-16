RFC: https://datatracker.ietf.org/doc/html/rfc5545

sample.ics
```
X-COMMENT:コメント行 https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.4
BEGIN:VCALENDAR

X-COMMENT:https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.1
BEGIN:VEVENT

X-COMMENT:https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.12
SUMMARY:iCal test

X-COMMENT:https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.5
X-COMMENT:\nで改行
DESCRIPTION:hello\nworld

X-COMMENT:https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.7
LOCATION:tokyo

X-COMMENT:https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.4
DTSTART;VALUE="DATE":20220715

X-COMMENT:https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.1
X-COMMENT:ATTENDEEPropertyはExchange(Office365)では表示されない
ATTENDEE;CN="ひとりめ";ROLE="REQ-PARTICIPANT";PARTSTAT="ACCEPTED":mailto:email1@example.com
ATTENDEE;CN="ふたりめ";ROLE="OPT-PARTICIPANT";PARTSTAT="DECLINED":mailto:email2@example.com

X-COMMENT:https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.6
X-COMMENT:URLPropertyはExchange(Office365)では表示されない
URL:https://ja.wikipedia.org/wiki/ICalendar

END:VEVENT
END:VCALENDAR
```

package ical.simple

due_date_ns = time.parse_ns("2006-01-02", input.due_date)

url = sprintf("https://redmine.example.com/issues/%d", [input.id])

allowed {
	abs(due_date_ns - time.now_ns()) < ((365 * 24) * 3600) * 1000000000 # 1年以内の期日である
}

event["SUMMARY"] = input.subject

event["DESCRIPTION"] = sprintf("URL: %s\n\nDescription:\n%s", [url, replace(input.description, "\r", "")])

event["DTSTART;VALUE=DATE"] = sprintf("%04d%02d%02d", time.date(due_date_ns))

event["URL"] = url

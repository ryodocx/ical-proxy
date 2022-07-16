package ical.simple

due_date_ns = time.parse_ns("2006-01-02", input.due_date)

default enabled = false

enabled {
	# 1年以内の期日である
	abs(due_date_ns - time.now_ns()) < ((365 * 24) * 3600) * 1000000000
}

event["SUMMARY"] = input.subject

event["DESCRIPTION"] = replace(input.description, "\r", "")

event["DTSTART;VALUE=DATE"] = sprintf("%04d%02d%02d", time.date(due_date_ns))

event["URL"] = sprintf("https://redmine.tokyo.optim.co.jp/sre/issues/%d", [input.id])

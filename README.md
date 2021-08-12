# Simple Time API

A simple REST API to get the current time for different timezones.
* Directly accessing `/api/time` without any query parameter returns current UTC time.
* Use standard time zone names to get current time in one or more time zones. E.g. `/api/time?tz=Asia/Kolkata` or `/api/time?tz=Asia/Kolkata,Asia/Shanghai,America/New_York`
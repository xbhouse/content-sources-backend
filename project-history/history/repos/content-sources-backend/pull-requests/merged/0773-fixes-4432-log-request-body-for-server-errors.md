---
type: pull_request
number: 773
title: "Fixes 4432: log request body for server errors"
state: merged
author: dominikvagner
created: 2024-08-14T09:59:28Z
updated: 2024-08-19T16:20:46Z
closed: 2024-08-19T16:20:45Z
merged: 2024-08-19T16:20:45Z
base_branch: main
head_branch: 4432
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/773
---

# Pull Request #773: Fixes 4432: log request body for server errors

**Author**: @dominikvagner
**Created**: August 14, 2024 at 09:59 AM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4432`

## Description

## Summary
Added a middleware which stores the body of incoming requests (cutoff to the limit of 1000B) and logs the request body when the response error code is of the 500 level (i.e. Server Error).

Wrote automated tests for checking the body storing it's recovery, cutoff limit and logging. Also tested this manually and checked for `request-id` retention and logging.

## Testing steps
### 1. Pick a temporary sacrificial POST endpoint like [this one](https://github.com/content-services/content-sources-backend/blob/f64f4872fecbc401b4bf9cf9833cfc4be31d2fb6/pkg/handler/rpms.go#L49C23-L49C38).
### 2. Comment out the body of it and put this line in the handler.
```
return ce.NewErrorResponse(http.StatusInternalServerError, "Test Error", "5xx testing error")
```
### 3. Start the server and send a request to it.
example request:
```
http http://localhost:8000/api/content-sources/v1.0/rpms/names "$( ./scripts/header.sh acme 1111)" limit:=100 uuids:='["cdb499cd-fc3b-44f9-9d4b-81824a418090"]' search='YfbrYLBvVcnY' x-Rh-Insights-Request-Id:9d229d34-7219-405d-ba3d-8f99cf7c36ae -v
```
### 4. Look at logs for an error message with the `request body` and `request-id` printed out.
![image](https://github.com/user-attachments/assets/a47dd594-7396-43ae-8bdf-0990b8a8db47)

---

## Discussion

### Comment by @jlsherrill on August 14, 2024 at 10:00 AM UTC

https://issues.redhat.com/browse/HMS-4432

### Comment by @xbhouse on August 15, 2024 at 12:01 PM UTC

overall i think this looks great and works well! @jlsherrill can you take a look in case i missed something?

### Comment by @jlsherrill on August 15, 2024 at 12:51 PM UTC

one comment, otherwise everything else looks great!

---

## Reviews

### Review by @xbhouse - Commented on August 15, 2024 at 11:58 AM UTC

### Review by @dominikvagner - Commented on August 15, 2024 at 12:16 PM UTC

### Review by @jlsherrill - Commented on August 15, 2024 at 12:51 PM UTC

### Review by @dominikvagner - Commented on August 15, 2024 at 01:44 PM UTC

### Review by @dominikvagner - Commented on August 16, 2024 at 12:13 PM UTC

### Review by @jlsherrill - Commented on August 16, 2024 at 12:38 PM UTC

### Review by @dominikvagner - Commented on August 16, 2024 at 01:56 PM UTC

### Review by @jlsherrill - Commented on August 16, 2024 at 01:57 PM UTC

### Review by @dominikvagner - Commented on August 16, 2024 at 03:26 PM UTC

### Review by @dominikvagner - Commented on August 16, 2024 at 04:34 PM UTC

### Review by @dominikvagner - Commented on August 16, 2024 at 04:37 PM UTC

### Review by @jlsherrill - Dismissed on August 16, 2024 at 05:39 PM UTC

Nice work!!

### Review by @jlsherrill - Commented on August 16, 2024 at 05:39 PM UTC

### Review by @dominikvagner - Commented on August 19, 2024 at 06:59 AM UTC

### Review by @xbhouse - Approved on August 19, 2024 at 11:29 AM UTC

awesome work!! 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/773*

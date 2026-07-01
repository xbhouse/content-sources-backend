---
type: pull_request
number: 351
title: "Fixes 2273: handle bad uuids on validate api"
state: merged
author: jlsherrill
created: 2023-08-03T20:20:07Z
updated: 2023-08-10T16:00:23Z
closed: 2023-08-10T15:33:15Z
merged: 2023-08-10T15:33:15Z
base_branch: main
head_branch: 2273
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/351
---

# Pull Request #351: Fixes 2273: handle bad uuids on validate api

**Author**: @jlsherrill
**Created**: August 03, 2023 at 08:20 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `2273`

## Description

## Summary

The validation api didn't handle bad uuids well.  In addition, the error was returned to the user but was not logged.  This adds that logging

## Testing steps
Perform this request:

```
POST http://localhost:8000/api/content-sources/v1.0/repository_parameters/validate/
Content-Type: application/json
x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJqZG9lIn0sImludGVybmFsIjp7Im9yZ19pZCI6IjEyMyJ9fX0K

[{"url":  "http://yum.theforeman.org/pulpcore/3.16/el8/x86_64/", "uuid":  "'response.write(731,493*173,328)'", "snapshot":  false}]
```

---

## Discussion

### Comment by @jlsherrill on August 03, 2023 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-2273

### Comment by @jlsherrill on August 03, 2023 at 08:50 PM UTC

I also realized that the log i added in the first commit didn't have a request id in it.  So now i'm passing the echo context all the way down in the second commit.  This gets you a nice request_id (this is without the query fix):

```
4:37PM ERR pkg/errors/handler_error.go:103 > Internal Server Error Generated error="ERROR: invalid input syntax for type uuid: \"'response.write(731,493*173,328)'\" (SQLSTATE 22P02)" request_id=kPDNggre2IrWRCE5YwxARTarPe08mXFt
4:37PM INF ../../go/pkg/mod/github.com/ziflex/lecho/v3@v3.5.0/middleware.go:165 > bytes_in=131 bytes_out=100 host=localhost:8000 latency=1.37264 latency_human=1.37264ms method=POST referer= remote_ip=127.0.0.1 request_id=kPDNggre2IrWRCE5YwxARTarPe08mXFt status=500 uri=/api/content-sources/v1.0/repository_parameters/validate/ user_agent="Apache-HttpClient/4.5.13 (Java/11.0.14.1)"
```

### Comment by @jlsherrill on August 04, 2023 at 01:31 PM UTC

/retest

### Comment by @jlsherrill on August 09, 2023 at 04:50 PM UTC

/retest

### Comment by @swadeley on August 10, 2023 at 03:33 PM UTC

Hi

tested by writing new sections for the `test_validate_repo_parameters`  test,  and running that test.
Details in the associated Jira ticket.

Thank you.

---

## Reviews

### Review by @rverdile - Commented on August 08, 2023 at 04:52 PM UTC

### Review by @jlsherrill - Commented on August 08, 2023 at 07:38 PM UTC

### Review by @rverdile - Commented on August 09, 2023 at 11:32 AM UTC

### Review by @jlsherrill - Commented on August 09, 2023 at 03:53 PM UTC

### Review by @rverdile - Approved on August 09, 2023 at 04:48 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/351*

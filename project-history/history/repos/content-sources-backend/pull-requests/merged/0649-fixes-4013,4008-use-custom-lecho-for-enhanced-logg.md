---
type: pull_request
number: 649
title: "Fixes 4013,4008: Use custom lecho for enhanced logging"
state: merged
author: jlsherrill
created: 2024-04-25T18:02:13Z
updated: 2024-04-26T20:01:46Z
closed: 2024-04-26T20:01:46Z
merged: 2024-04-26T20:01:46Z
base_branch: main
head_branch: custom_lecho
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/649
---

# Pull Request #649: Fixes 4013,4008: Use custom lecho for enhanced logging

**Author**: @jlsherrill
**Created**: April 25, 2024 at 06:02 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `custom_lecho`

## Description

## Summary

This uses a custom lecho fork to accomplish two things:
1.  log slow requests at a higher log level
2.  do not treat non-500 errors as needing error level logging

2.  was kinda tricky.  We need another middleware that runs around lecho to properly parse the error into the correct struct, extract the correct status, and set that on the response.  Then lecho has to re-lookup the response to get the proper status code.  I will work on contributing 1. back to lecho, but i'm not sure if the changes for 2 make sense to?

## Testing steps

to test 1:
add in a time.sleep(2*time.Second)  somewhere in a handler (like listRepositories).  Watch requests get logged at warn

to test 2:

in listRepositories add some code after the filter is parsed:

```
 if filterData.Search == "foo" {
         return ce.NewErrorResponse(500, "Error listing repositories", "ISE")
 }
```

now you can make requests to:

```
GET http://localhost:8000/api/content-sources/v1.0/repositories/?search=foo
```
to trigger an ISE.  You can try to create the same repo 2x to trigger a 400.  ISEs should still be logged at error level, but 400s should be logged at Info like normal.


## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on April 25, 2024 at 06:02 PM UTC

custom lecho PR: https://github.com/content-services/lecho/pull/1/files

### Comment by @jlsherrill on April 25, 2024 at 06:30 PM UTC

https://issues.redhat.com/browse/HMS-4013

### Comment by @jlsherrill on April 25, 2024 at 06:30 PM UTC

https://issues.redhat.com/browse/HMS-4008

### Comment by @swadeley on April 26, 2024 at 06:34 PM UTC

/retest

### Comment by @swadeley on April 26, 2024 at 06:40 PM UTC

  Error: undefined: lecho (typecheck)

### Comment by @xbhouse on April 26, 2024 at 06:46 PM UTC

testing 1 looks good, seeing slow requests get logged as warn: `2:26PM WRN ../lecho/middleware.go:171 > bytes_in=0 bytes_out=894 host=localhost:8000 latency=2017.993208 latency_human=2.017993208s method=GET referer= remote_ip=::1 request_id=OxeJrkomMgWkoHffKhCvzlIBaMqPfEIw status=200 uri=/api/content-sources/v1.0/repositories/ user_agent=PostmanRuntime/7.37.3`

testing 2 looks good too, <500 errors logged as info: `2:32PM INF ../lecho/middleware.go:171 > error="error: code=400, title=Error creating repository, detail=Repository with this name already belongs to organization \n" bytes_in=227 bytes_out=0 host=localhost:8000 latency=5.938958 latency_human=5.938958ms method=POST referer= remote_ip=::1 request_id=qpSQNnmMxgiaGLKJOEAIDHQxxFnmiyym status=200 uri=/api/content-sources/v1.0/repositories/ user_agent=PostmanRuntime/7.37.3`


---

## Reviews

### Review by @xbhouse - Approved on April 26, 2024 at 07:33 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/649*

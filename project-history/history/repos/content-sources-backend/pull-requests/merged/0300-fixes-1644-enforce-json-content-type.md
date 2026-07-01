---
type: pull_request
number: 300
title: "Fixes 1644: Enforce JSON Content Type"
state: merged
author: dpang314
created: 2023-05-31T20:38:55Z
updated: 2023-06-05T17:30:18Z
closed: 2023-06-05T17:23:13Z
merged: 2023-06-05T17:23:13Z
base_branch: main
head_branch: HMS-1644-enforce-content-type
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/300
---

# Pull Request #300: Fixes 1644: Enforce JSON Content Type

**Author**: @dpang314
**Created**: May 31, 2023 at 08:38 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-1644-enforce-content-type`

## Description

## Summary

Adds middleware `EnforceJSONContentType` that returns an error with status code 415 "Unsupported Media Type" if the Content-Type header is missing, invalid, or not `application/json`. If the Content-Type has optional parameters (e.g. `application/json; parameter=value`), it is still accepted.

The middleware is added at root level and applied to any request that has a body

It affects the following endpoints: 

1. /rpms/names [post] - searchRpmByName
2. /repository_parameters/external_gpg_key/ [post] - fetchGpgKey
3. /repository_parameters/validate/ [post] - validate
4. /repositories [post] - createRepository
5. /repositories/bulk_create/ [post] - bulkCreateRepositories
6. /repositories/{uuid} [put] - fullUpdate
7. /repositories/{uuid} [patch] - partialUpdate

/repositories/{uuid}/introspect/ [post] accepts an optional JSON body, so the middleware is only applied when the body is passed in.

## Testing steps

Make a curl request to any of the 7 endpoints without setting the Content-Type header or with a Content-Type header that isn't `application/json`

 `curl -H "$( ./scripts/header.sh 9999 1111 )" "http://localhost:8000/api/content-sources/v1.0/rpms/names" -X POST -d '{ "urls": [ "https://omaciel.fedorapeople.org/fakerepo01/", "https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/" ],   "search": "3" } '`

Curl seems to set the content type to `application/x-www-form-urlencoded` by default so if no content type is specified, the error title returned will be "Incorrect content type" 

To get "Error parsing content type" make a request with an empty Content-Type or a value that doesn't follow the MIME type format

 `curl -H "$( ./scripts/header.sh 9999 1111 )" "http://localhost:8000/api/content-sources/v1.0/rpms/names" -X POST -d '{ "urls": [ "https://omaciel.fedorapeople.org/fakerepo01/", "https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/" ],   "search": "3" } ' -H "Content-Type:"`

 `curl -H "$( ./scripts/header.sh 9999 1111 )" "http://localhost:8000/api/content-sources/v1.0/rpms/names" -X POST -d '{ "urls": [ "https://omaciel.fedorapeople.org/fakerepo01/", "https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/" ],   "search": "3" } ' -H "Content-Type: not a valid type"`

---

## Discussion

### Comment by @app-sre-bot on May 31, 2023 at 08:39 PM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on May 31, 2023 at 08:45 PM UTC

Would it be possible to implement a [middleware](https://echo.labstack.com/middleware/) to check for this?  I think you could assume that any PUT/POST/PATCH request should expect this content-type, and you could just ignore the check for GET/DELETE/anything else.

EDIT: the benefit of doing it in middleware is that each individual handler wouldn't' have to check, and anyone writing a new api wouldn't have to remember

### Comment by @dpang314 on May 31, 2023 at 08:52 PM UTC

Middleware would definitely be a much cleaner solution. `POST /repositories/{uuid}/introspect/` currently has an optional body, but adding the middleware would require setting the content type to `application/json` even if you didn't want to specify a body. Is this fine?

### Comment by @jlsherrill on May 31, 2023 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-1644

### Comment by @jlsherrill on May 31, 2023 at 09:11 PM UTC

I think that's okay.  At least i'm okay with it.  Curious if @rverdile has any thoughts.  we could also ignore certain paths, which can be annoying but also would assure we dont' forget to add a check

### Comment by @rverdile on May 31, 2023 at 09:17 PM UTC

Yeah if the middleware can blanket apply to all the PUT/POST/PATCH endpoints I think that makes sense. The goal would be to register it in one place and have it apply to everything. 

I think it also makes to explicitly ignore endpoints where the body is not required. I'm assuming that won't happen too often.

### Comment by @jlsherrill on June 02, 2023 at 06:13 PM UTC

oktotest

### Comment by @jlsherrill on June 02, 2023 at 06:32 PM UTC

/retest

### Comment by @jlsherrill on June 02, 2023 at 06:36 PM UTC

[test]

### Comment by @jlsherrill on June 02, 2023 at 07:50 PM UTC

nice work @dpang314 !!   I've approved this and now it waits for QE to verify

### Comment by @swadeley on June 05, 2023 at 05:23 PM UTC

Hi

`application/json; parameter=value` still works

if you omit the content type:
`{"errors":[{"status":415,"title":"Incorrect content type","detail":"Content-Type must be application/json"}]}`

If you pass `-H "Content-Type: Not a valid type"`:
`{"errors":[{"status":415,"title":"Error parsing content type","detail":"mime: expected slash after first token"}]}`

If you pass `-H "Content-Type:"`:
`{"errors":[{"status":415,"title":"Error parsing content type","detail":"mime: no media type"}]}`





---

## Reviews

### Review by @jlsherrill - Approved on June 02, 2023 at 07:49 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/300*

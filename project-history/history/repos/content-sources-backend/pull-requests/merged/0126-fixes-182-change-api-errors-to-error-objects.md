---
type: pull_request
number: 126
title: "Fixes 182: change api errors to error objects"
state: merged
author: rverdile
created: 2022-10-20T18:37:09Z
updated: 2022-11-23T15:00:48Z
closed: 2022-11-21T16:16:37Z
merged: 2022-11-21T16:16:37Z
base_branch: main
head_branch: error-objects
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/126
---

# Pull Request #126: Fixes 182: change api errors to error objects

**Author**: @rverdile
**Created**: October 20, 2022 at 06:37 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `error-objects`

## Description

Adds error object that returns on API errors to satisfy [this standard](https://jsonapi.org/format/#errors). It looks like:
```
errors: [
   {
      status: 400,
      title: "error creating repo",
      detail: "url already exists"
   }
]
```

Also specifies error object in openAPI spec for 400,404,500 for all endpoints.

This replaces the default echo error handler with `customHTTPErrorHandler` in `pkg/config/config.go`. It also moves custom errors (`DaoError` and `HandlerError`) to a package named `errors`. This way they can be imported by everything, including `pkg/config.go`

TODO

- [x] Tests (I'll fix the tests once direction of PR is okay)
- [x] Return most common error code as error code for bulk create. See "[generically applicable](https://jsonapi.org/format/#errors-processing)".

To test:

This changes the error response for all API errors, so cause errors and check that the response is as expected. Also check the API spec has been updated to correctly reflect the changes.

---

## Discussion

### Comment by @jlsherrill on October 20, 2022 at 07:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-182

### Comment by @rverdile on October 24, 2022 at 07:20 PM UTC

I got a new idea for this I am trying where ErrorResponse is 
```
type ErrorResponse struct {
    Errors []Error{ Status, Detail, Title }
}
```
This way returning multiple errors works more directly with the handler. Instead of creating a new response like in BulkCreate.

### Comment by @rverdile on October 26, 2022 at 08:08 PM UTC

Changed `ErrorResponse` struct to have a slice of `HandlerError`, this way error handling more directly supports returning multiple errors like with `BulkCreate`.

To support this change I created a new `errors` package. Otherwise, I could not find a way to do this without causing circular dependencies. This way allows these errors to be imported by everything, including `pkg/config.go`.

### Comment by @jlsherrill on October 27, 2022 at 07:29 PM UTC

overall +1 to the structural changes. 

### Comment by @rverdile on November 15, 2022 at 04:07 PM UTC

updated with `TestNewErrorResponseFromEchoErrorOtherType` and `TestDBErrorToApi` suggestions. Thanks @avisiedo 

### Comment by @avisiedo on November 15, 2022 at 04:25 PM UTC

lgtm

### Comment by @swadeley on November 21, 2022 at 04:16 PM UTC

Tested, all API test pass.

### Comment by @swadeley on November 22, 2022 at 09:37 AM UTC

> Tested, all API test pass.

but after merging and rebuilding test EE, I get `ApiRepositoryBulkCreateResponse has no attribute 'uuid' at ['received_data'][0]['uuid']`.

---

## Reviews

### Review by @avisiedo - Changes Requested on October 24, 2022 at 10:59 AM UTC

Nice pr, I have learned about custom error handlers in echo.

A few updates and suggestions.

I will make a second round.

### Review by @rverdile - Commented on October 25, 2022 at 01:53 PM UTC

### Review by @rverdile - Commented on October 25, 2022 at 01:55 PM UTC

### Review by @rverdile - Commented on October 25, 2022 at 01:58 PM UTC

### Review by @avisiedo - Commented on October 25, 2022 at 04:25 PM UTC

### Review by @rverdile - Commented on October 25, 2022 at 04:42 PM UTC

### Review by @rverdile - Commented on October 25, 2022 at 04:42 PM UTC

### Review by @rverdile - Commented on October 26, 2022 at 08:03 PM UTC

### Review by @rverdile - Commented on October 27, 2022 at 01:49 PM UTC

### Review by @jlsherrill - Commented on October 27, 2022 at 07:25 PM UTC

### Review by @avisiedo - Commented on November 02, 2022 at 11:16 AM UTC

I have only one suggestion that I could be wrong.

thanks for the changes and the direction of the pr!

### Review by @rverdile - Commented on November 02, 2022 at 08:05 PM UTC

### Review by @avisiedo - Changes Requested on November 10, 2022 at 12:16 PM UTC

nice updates!
only a few updates for the unit tests detected from the ide.
current unit tests are passing :+1:

### Review by @rverdile - Commented on November 10, 2022 at 03:38 PM UTC

### Review by @rverdile - Commented on November 10, 2022 at 03:59 PM UTC

### Review by @rverdile - Commented on November 10, 2022 at 04:05 PM UTC

### Review by @rverdile - Commented on November 10, 2022 at 04:06 PM UTC

### Review by @rverdile - Commented on November 10, 2022 at 04:06 PM UTC

### Review by @rverdile - Commented on November 10, 2022 at 04:07 PM UTC

### Review by @rverdile - Commented on November 11, 2022 at 12:37 PM UTC

### Review by @avisiedo - Changes Requested on November 15, 2022 at 08:54 AM UTC

Nice change! Only one last suggestion! after it, it lgtm

### Review by @rverdile - Commented on November 15, 2022 at 04:05 PM UTC

### Review by @avisiedo - Approved on November 15, 2022 at 04:25 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/126*

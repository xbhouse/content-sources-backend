---
type: pull_request
number: 19
title: "CONTENT-40: Create content source"
state: merged
author: rverdile
created: 2022-05-26T17:43:34Z
updated: 2022-06-02T15:02:21Z
closed: 2022-06-02T15:02:20Z
merged: 2022-06-02T15:02:20Z
base_branch: main
head_branch: create-endpoint
labels: []
url: https://github.com/content-services/content-sources-backend/pull/19
---

# Pull Request #19: CONTENT-40: Create content source

**Author**: @rverdile
**Created**: May 26, 2022 at 05:43 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `create-endpoint`

## Description

Adds create endpoint for content source. Adds a 'dao' layer and 'api' package.

The 'api' package defines all the API request/response structs. For example, with create, when the handler gets a request, it binds it to the 'CreateRepository' struct in the 'api' package. Then, that struct is passed to the 'dao' package which interacts with the models and the database to the necessary items to the database.

To test you need to use the "x-hr-identity" header. You want a base64 encoded JSON that looks like: 

> "identity": {
        "account_number": "12345",
        "org_id": "54321"
}

Handling the account_number and org_id relies on this common package: https://github.com/RedHatInsights/platform-go-middlewares/tree/master/identity

TODO:

- [x]  Create endpoint with OpenAPI comments
- [x]  Unique constraint between org ID and URL
- [x]  dao layer
- [x]  tests
- [x]  address swag param type.


---

## Discussion

### Comment by @rverdile on May 27, 2022 at 01:56 PM UTC

@jlsherrill I think I did what you're describing, can you take a look?

### Comment by @rverdile on May 27, 2022 at 08:25 PM UTC

Moved the other structs to 'api' package. OpenAPI not working yet for some reason.

---

## Reviews

### Review by @jlsherrill - Commented on May 26, 2022 at 06:10 PM UTC

### Review by @jlsherrill - Commented on May 26, 2022 at 06:17 PM UTC

### Review by @jlsherrill - Commented on May 26, 2022 at 06:19 PM UTC

### Review by @jlsherrill - Commented on May 27, 2022 at 03:18 PM UTC

### Review by @jlsherrill - Commented on May 27, 2022 at 03:20 PM UTC

### Review by @rverdile - Commented on May 27, 2022 at 03:56 PM UTC

### Review by @jlsherrill - Commented on May 27, 2022 at 04:08 PM UTC

### Review by @jlsherrill - Commented on May 27, 2022 at 04:09 PM UTC

### Review by @rverdile - Commented on May 27, 2022 at 04:31 PM UTC

### Review by @jlsherrill - Commented on May 27, 2022 at 05:12 PM UTC

### Review by @jlsherrill - Commented on May 27, 2022 at 08:29 PM UTC

### Review by @jlsherrill - Commented on May 27, 2022 at 08:30 PM UTC

### Review by @jlsherrill - Commented on May 27, 2022 at 08:32 PM UTC

### Review by @jlsherrill - Commented on May 27, 2022 at 08:40 PM UTC

### Review by @jlsherrill - Commented on May 27, 2022 at 08:54 PM UTC

### Review by @rverdile - Commented on June 01, 2022 at 03:41 PM UTC

### Review by @rverdile - Commented on June 01, 2022 at 03:57 PM UTC

### Review by @jlsherrill - Commented on June 01, 2022 at 06:21 PM UTC

### Review by @jlsherrill - Commented on June 01, 2022 at 06:22 PM UTC

### Review by @rverdile - Commented on June 01, 2022 at 06:35 PM UTC

### Review by @jlsherrill - Commented on June 01, 2022 at 07:42 PM UTC

### Review by @jlsherrill - Commented on June 01, 2022 at 07:43 PM UTC

### Review by @jlsherrill - Commented on June 01, 2022 at 07:44 PM UTC

### Review by @jlsherrill - Commented on June 02, 2022 at 01:21 PM UTC

### Review by @jlsherrill - Commented on June 02, 2022 at 01:22 PM UTC

### Review by @jlsherrill - Commented on June 02, 2022 at 01:23 PM UTC

### Review by @jlsherrill - Commented on June 02, 2022 at 02:36 PM UTC

### Review by @rverdile - Commented on June 02, 2022 at 02:37 PM UTC

### Review by @jlsherrill - Approved on June 02, 2022 at 02:58 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/19*

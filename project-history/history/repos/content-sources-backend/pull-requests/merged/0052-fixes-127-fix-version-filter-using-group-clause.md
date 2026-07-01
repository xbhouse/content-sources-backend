---
type: pull_request
number: 52
title: "Fixes 127: fix version filter using group clause"
state: merged
author: jlsherrill
created: 2022-07-15T18:56:59Z
updated: 2022-07-25T07:50:55Z
closed: 2022-07-25T07:50:55Z
merged: 2022-07-25T07:50:55Z
base_branch: main
head_branch: drop_tables
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/52
---

# Pull Request #52: Fixes 127: fix version filter using group clause

**Author**: @jlsherrill
**Created**: July 15, 2022 at 06:56 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `drop_tables`

## Description

Does the following:

* version filter on repo list api was improperly constructing a query, it was doing:
   org_id = X AND version = 7 OR version = 9   which would improperly return all
   repos with version = 9.  The new query should properly be structured like:
   org_id = X AND (version = 7 OR version = 9)
* Stops dropping all tables as part of each test, this speeds up the tests greatly when
   there is seeded data
* Starts using random Org ids for tests

reproducer steps:

in a fresh database:




Create a repo in org 3 with el7 & el9
```
$ curl -X POST http://localhost:8000/api/content_sources/v1/repositories/   -H "`scripts/header.sh 3 3`"  -d '{"url": "http://foo.com/bar", "distribution_versions": ["el7", "el9"], "name": "foo"}'  -H "Content-Type: application/json" 
```



Create a repo in org 1 with el7 and el9
```
curl -X POST http://localhost:8000/api/content_sources/v1/repositories/   -H "`scripts/header.sh 1 1`"  -d '{"url": "http://foo.com/bar", "distribution_versions": ["el7", "el9"], "name": "foo"}'  -H "Content-Type: application/json"
```
query repos in org 1 with just el7 or el9, should only return a single entry, but we get two:



```
curl -X GET "http://localhost:8000/api/content_sources/v1/repositories/?version=el7,el9" -H "`scripts/header.sh 1 1`" 
{
   "data":[
      {
         "uuid":"6092ce16-b2ec-4d09-a4c4-c1eb35d9dea6",
         "name":"foo",
         "url":"http://foo.com/bar",
         "distribution_versions":[
            "el7",
            "el9"
         ],
         "distribution_arch":"",
         "account_id":"3",
         "org_id":"3"
      },
      {
         "uuid":"c86a011a-d688-4f9f-abe6-7af33d8ca20c",
         "name":"foo",
         "url":"http://foo.com/bar",
         "distribution_versions":[
            "el7",
            "el9"
         ],
         "distribution_arch":"",
         "account_id":"1",
         "org_id":"1"
      }
   ],
   "meta":{
      "limit":100,
      "offset":0,
      "count":2
   },
   "links":{
      "first":"/api/content_sources/v1/repositories/?limit=100\u0026offset=0\u0026version=el7,el9",
      "last":"/api/content_sources/v1/repositories/?limit=100\u0026offset=0\u0026version=el7,el9"
   }
}
```




---

## Discussion

### Comment by @jlsherrill on July 15, 2022 at 06:57 PM UTC

Note this is built ontop of https://github.com/content-services/content-sources-backend/pull/45   will rebase once that goes in

### Comment by @jlsherrill on July 15, 2022 at 07:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-127

### Comment by @jlsherrill on July 18, 2022 at 06:34 PM UTC

Built and pushed image for 4acf7e8f4c8a04dc0b8b6bd2ff120a0c2096c1fe

### Comment by @jlsherrill on July 20, 2022 at 03:46 PM UTC

Built and pushed image for f215900a1f98d099859b5bfdb3c8e517aeac7668

---

## Reviews

### Review by @Andrewgdewar - Commented on July 15, 2022 at 07:28 PM UTC

### Review by @Andrewgdewar - Commented on July 15, 2022 at 07:29 PM UTC

### Review by @jlsherrill - Commented on July 15, 2022 at 07:37 PM UTC

### Review by @rverdile - Approved on July 20, 2022 at 06:43 PM UTC

lgtm!

### Review by @swadeley - Approved on July 25, 2022 at 07:50 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/52*

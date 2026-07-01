---
type: pull_request
number: 228
title: "Fixes 1412: remove url validation on event reading"
state: merged
author: jlsherrill
created: 2023-03-09T18:20:00Z
updated: 2023-03-15T12:26:17Z
closed: 2023-03-15T12:26:14Z
merged: 2023-03-15T12:26:14Z
base_branch: main
head_branch: 1412
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/228
---

# Pull Request #228: Fixes 1412: remove url validation on event reading

**Author**: @jlsherrill
**Created**: March 09, 2023 at 06:20 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `1412`

## Description

## Summary
Previously bad urls were treated as a 'kafka' error, and not an introspection error.  This led to bad stats when trying to look at metrics.

## Testing steps
run the following query:

```.rest
### Create a single repository
POST http://{{host}}/api/content-sources/v1.0/repositories/
Content-Type: application/json
x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiQXNzb2NpYXRlIiwiYWNjb3VudF9udW1iZXIiOiIxIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMSJ9fX0K

{
  "name": "azure clizz2",
  "url": "/fo1",
  "distribution_versions": [
    "8"
  ],
  "distribution_arch": "x86_64"
},

```

then check metrics:
```
curl localhost:9000/metrics | grep kafka_message_result_total
```

you should see

```
content_sources_kafka_message_result_total{state="success"} 1
```

which is correct, the failure was in introspection, not in message parsing/handling 



---

## Discussion

### Comment by @jlsherrill on March 09, 2023 at 06:30 PM UTC

https://issues.redhat.com/browse/HMS-1412

### Comment by @jlsherrill on March 13, 2023 at 01:35 PM UTC

/retest

### Comment by @jlsherrill on March 14, 2023 at 01:06 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on March 14, 2023 at 08:33 PM UTC

looks good!!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/228*

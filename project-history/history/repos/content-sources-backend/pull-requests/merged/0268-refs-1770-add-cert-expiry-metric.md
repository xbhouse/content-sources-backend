---
type: pull_request
number: 268
title: "Refs 1770: add cert expiry metric"
state: merged
author: jlsherrill
created: 2023-05-09T19:53:14Z
updated: 2023-05-16T17:57:45Z
closed: 2023-05-16T17:57:42Z
merged: 2023-05-16T17:57:42Z
base_branch: main
head_branch: 1770
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/268
---

# Pull Request #268: Refs 1770: add cert expiry metric

**Author**: @jlsherrill
**Created**: May 09, 2023 at 07:53 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `1770`

## Description


## Summary
this adds a new metric "content_sources_rh_cert_expiry_days", that signifies the number of days until the current RH cert expires

## Testing steps
Ensure you have a cdn cert configured in your config.yaml, run the server, then run:
```
curl localhost:9000/metrics | grep expi
```

you'll see something like:

```
# HELP content_sources_rh_cert_expiry_days Number of days until the Red Hat client certificate expires
# TYPE content_sources_rh_cert_expiry_days gauge
content_sources_rh_cert_expiry_days 23
```

---

## Discussion

### Comment by @jlsherrill on May 09, 2023 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-1770

### Comment by @jlsherrill on May 10, 2023 at 04:13 PM UTC

[test]

### Comment by @jlsherrill on May 10, 2023 at 04:14 PM UTC

/retest

### Comment by @swadeley on May 16, 2023 at 04:06 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Commented on May 09, 2023 at 08:15 PM UTC

### Review by @rverdile - Approved on May 10, 2023 at 04:12 PM UTC

tested and lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/268*

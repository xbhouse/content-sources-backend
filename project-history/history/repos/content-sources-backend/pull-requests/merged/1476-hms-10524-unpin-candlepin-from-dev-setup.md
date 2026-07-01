---
type: pull_request
number: 1476
title: "HMS-10524: unpin candlepin from dev setup"
state: merged
author: xbhouse
created: 2026-04-22T20:06:55Z
updated: 2026-05-06T18:45:42Z
closed: 2026-05-06T18:45:42Z
merged: 2026-05-06T18:45:42Z
base_branch: main
head_branch: unpin-candlepin
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1476
---

# Pull Request #1476: HMS-10524: unpin candlepin from dev setup

**Author**: @xbhouse
**Created**: April 22, 2026 at 08:06 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `unpin-candlepin`

## Description

## Summary

- Reverts the dev candlepin tag back to dev-latest
- Adds environment variable required for our setup with the new image

## Testing steps

Tests should pass when a fixed image is pushed

---

## Discussion

### Comment by @xbhouse on April 22, 2026 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-10524

### Comment by @rverdile on May 06, 2026 at 01:50 PM UTC

@xbhouse I noticed there's a new dev-latest build, but the PR might be caching the image it used when it was first opened. I've seen something like that before.

### Comment by @xbhouse on May 06, 2026 at 01:55 PM UTC

oh @rverdile good point, i'll try opening this again. i am still seeing errors with the latest image locally though, is this the same error from before?

```
9:47AM FTL Could not create org error="couldn't create org: undefined response type: <!doctype html><html lang=\"en\"><head><title>HTTP Status 404 – Not Found</title><style type=\"text/css\">body {font-family:Tahoma,Arial,sans-serif;} h1, h2, h3, b {color:white;background-color:#525D76;} h1 {font-size:22px;} h2 {font-size:16px;} h3 {font-size:14px;} p {font-size:12px;} a {color:black;} .line {height:1px;background-color:#525D76;border:none;}</style></head><body><h1>HTTP Status 404 – Not Found</h1><hr class=\"line\" /><p><b>Type</b> Status Report</p><p><b>Message</b> The requested resource [&#47;candlepin&#47;owners] is not available</p><p><b>Description</b> The origin server did not find a current representation for the target resource or is not willing to disclose that one exists.</p><hr class=\"line\" /><h3>Apache Tomcat/9.0.110</h3></body></html>" severity=fatal
```

### Comment by @rverdile on May 06, 2026 at 02:24 PM UTC

@xbhouse That's the same symptom but it may not be the same root cause. You'd have to check the candlepin container logs. It was failing first because of a missing password, and so you may have to set `JPA_CONFIG_HIBERNATE_PASSWORD` in the docker-compose.yaml. They may not have fixed that. And then it failed because the image did not contain the right certs - which should be what they fixed.

---

## Reviews

### Review by @rverdile - Approved on May 06, 2026 at 06:18 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1476*

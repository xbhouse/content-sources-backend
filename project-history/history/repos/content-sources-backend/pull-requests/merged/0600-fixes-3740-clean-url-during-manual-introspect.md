---
type: pull_request
number: 600
title: "Fixes 3740: clean url during manual introspect"
state: merged
author: jlsherrill
created: 2024-03-07T22:45:02Z
updated: 2024-03-12T19:30:26Z
closed: 2024-03-12T19:10:24Z
merged: 2024-03-12T19:10:24Z
base_branch: main
head_branch: 3740
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/600
---

# Pull Request #600: Fixes 3740: clean url during manual introspect

**Author**: @jlsherrill
**Created**: March 07, 2024 at 10:45 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3740`

## Description

## Summary

previously we would 'cleanup' the url when you manually trigger introspection, but this was broken during a refactor

## Testing steps
```
make repos-import
 go run cmd/external-repos/main.go introspect --force https://cdn.redhat.com/content/dist/layered/rhel8/x86_64/ansible/2/os
```
Without this pr, the command will do nothing, with you should see:
```
5:44PM DBG Introspecting https://cdn.redhat.com/content/dist/layered/rhel8/x86_64/ansible/2/os/
```

also during ephemeral deployments you should see the this ansible repo get introspected, without this, its always pending

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on March 07, 2024 at 11:00 PM UTC

https://issues.redhat.com/browse/HMS-3740

### Comment by @swadeley on March 11, 2024 at 08:34 AM UTC

/retest

### Comment by @swadeley on March 11, 2024 at 10:20 AM UTC

Hi

I deployed with `OPTIONS_REPOSITORY_IMPORT_FILTER=small` so I just get `https://cdn.redhat.com/content/dist/layered/rhel8/x86_64/ansible/2/os/` in ephemeral.

I created repo with `https://cdn.redhat.com/content/dist/layered/rhel8/aarch64/ansible/2/os/`
NOTE: aarch not x86

I manually triggered introspection.

It introspects quickly enough, less than a minute, but it fails to make a snapshot. I think that is because its a RHEL repo:
` '[SSL: CERTIFICATE_VERIFY_FAILED] certificate verify failed: self signed certificate in certificate chain (_ssl.c:1129)')].`.

OK, lgtm. Merge on ack.

### Comment by @xbhouse on March 12, 2024 at 04:02 PM UTC

lgtm! 

---

## Reviews

### Review by @xbhouse - Approved on March 12, 2024 at 04:02 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/600*

---
type: pull_request
number: 442
title: "Fixes 2616: add gpg key file endpoint"
state: merged
author: rverdile
created: 2023-10-24T15:04:44Z
updated: 2023-11-24T05:00:40Z
closed: 2023-11-23T00:36:11Z
merged: 2023-11-23T00:36:11Z
base_branch: main
head_branch: gpg-key-file
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/442
---

# Pull Request #442: Fixes 2616: add gpg key file endpoint

**Author**: @rverdile
**Created**: October 24, 2023 at 03:04 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `gpg-key-file`

## Description

## Summary
Adds an endpoint that returns the GPG key of a repository as a file.

Also adds updates the fields in the config.repo file that is returned by the config.repo file endpoint.

## Testing steps
1. Create a repository with and without a GPG Key
2. Query the `repositories/:uuid/snapshots/:snapshot_uuid/config.repo` and `repositories/:uuid/gpg_key`
3. Verify the correct returns.


---

## Discussion

### Comment by @jlsherrill on October 24, 2023 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-2616

### Comment by @jlsherrill on October 27, 2023 at 07:41 PM UTC

it looks like the gpg_key route is still requiring the identity header:

```
$ curl  http://192.168.122.1:8000/api/content-sources/repositories/2c1413eb-119d-4b6f-ab54-c9af4f48c3e6/gpg_key
Bad Request: missing x-rh-identity header
```

Take a look at the enforce_identity.go, but it may also need to skip rbac if its not

### Comment by @jlsherrill on November 07, 2023 at 04:13 PM UTC

Couple of issues:

1. I'm still seeing:  
```
$ curl http://192.168.122.1:8000/api/content-sources/repositories/e68aeb70-6153-4503-8ea2-2ef3af594156/gpg_key/
Bad Request: missing x-rh-identity header
```
2.  When i fetch a repo.config, the gpg key url has double slashes:

```
GET http://localhost:8000/api/content-sources/v1.0/repositories/e68aeb70-6153-4503-8ea2-2ef3af594156/snapshots/26b98649-5602-4233-ac09-1ea0ddc54f4d/config.repo

HTTP/1.1 200 OK
Content-Type: text/plain; charset=UTF-8
X-Rh-Identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJmb28ifSwiYWNjb3VudF9udW1iZXIiOiIxMTQ4NjU4OSIsImludGVybmFsIjp7Im9yZ19pZCI6IjE2Nzg4NjE2In19fQo=
X-Rh-Insights-Request-Id: qlRKsyEHNUJeheQBbygwZpwbnhRHDFyB
Date: Tue, 07 Nov 2023 16:05:27 GMT
Content-Length: 296

[azure]
name=azure
baseurl=http://pulp_content:24816/pulp/content/7c419c57/e68aeb70-6153-4503-8ea2-2ef3af594156/891faaf8-f37f-4f07-a04a-527565bda02c/
gpgcheck=1
repo_gpgcheck=1
enabled=1
gpgkey=http://localhost:8000//api/content-sources/repositories/e68aeb70-6153-4503-8ea2-2ef3af594156/gpg_key/
```

### Comment by @Andrewgdewar on November 07, 2023 at 06:40 PM UTC

/retest

### Comment by @jlsherrill on November 08, 2023 at 03:12 PM UTC

/retest

### Comment by @jlsherrill on November 08, 2023 at 03:23 PM UTC

Overall this is working great!  One issue i did not forsee.  Right now pulp is republishing the yum metadata, which means the metadata signature is not present, so you get an error:

```
Error: Failed to download metadata for repo 'azure': Cannot download repomd.xml: Cannot download repodata/repomd.xml: All mirrors were tried
```
from dnf.  There are a couple routes we can go in the future to have pulp either sign the metadata, or preserve the signature, but we need to discuss it a bit.  For now, i think we should always set repo_gpgcheck=0.

I'll file a task to figure out what to do about this.

### Comment by @jlsherrill on November 13, 2023 at 02:48 PM UTC

this needs a rebase now 

### Comment by @swadeley on November 16, 2023 at 01:39 AM UTC

/retest

### Comment by @swadeley on November 21, 2023 at 01:32 AM UTC

Hi @rverdile , please rebase again.

---

## Reviews

### Review by @jlsherrill - Commented on October 27, 2023 at 06:27 PM UTC

### Review by @jlsherrill - Commented on October 27, 2023 at 07:31 PM UTC

### Review by @jlsherrill - Approved on November 14, 2023 at 12:21 AM UTC

ACK just needs a rebase still

### Review by @jlsherrill - Commented on November 22, 2023 at 02:12 PM UTC

### Review by @rverdile - Commented on November 22, 2023 at 02:47 PM UTC

### Review by @swadeley - Commented on November 23, 2023 at 12:35 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/442*

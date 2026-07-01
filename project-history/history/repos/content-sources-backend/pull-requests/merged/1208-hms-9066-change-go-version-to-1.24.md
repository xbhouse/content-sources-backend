---
type: pull_request
number: 1208
title: "HMS-9066: Change go version to 1.24"
state: merged
author: mayurilahane
created: 2025-09-18T04:06:03Z
updated: 2025-10-08T18:24:06Z
closed: 2025-10-08T18:24:06Z
merged: 2025-10-08T18:24:06Z
base_branch: main
head_branch: mlahane/HMS-9066
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1208
---

# Pull Request #1208: HMS-9066: Change go version to 1.24

**Author**: @mayurilahane
**Created**: September 18, 2025 at 04:06 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `mlahane/HMS-9066`

## Description

## Summary
update: Change go version to 1.24



---

## Discussion

### Comment by @jlsherrill on September 18, 2025 at 04:30 AM UTC

https://issues.redhat.com/browse/HMS-9066

### Comment by @mayurilahane on September 24, 2025 at 02:51 AM UTC

/retest

### Comment by @mayurilahane on September 24, 2025 at 01:35 PM UTC

> I would do a dependency bump here: `go get -u ./...`

Yeah i did that at very first but after rebasing with the main branch i think it changed the go.sum and go.mod file 
thanks for the reminder to do it again.

### Comment by @rverdile on September 24, 2025 at 01:56 PM UTC

Thanks, that looks good. Before we merge this one, I'm thinking we should merge the other PRs first so that we can update those dependencies as part of this PR as well. 

### Comment by @xbhouse on September 24, 2025 at 03:23 PM UTC

do we need to also update the build image in build/Dockerfile to `quay.io/projectquay/golang:1.24`? or is `registry.redhat.io/ubi9/go-toolset:latest` ok? 

### Comment by @mayurilahane on September 24, 2025 at 04:58 PM UTC

> do we need to also update the build image in build/Dockerfile to `quay.io/projectquay/golang:1.24`? or is `registry.redhat.io/ubi9/go-toolset:latest` ok?

i was thinking to update that too but docker file says it uses latest quay image so i did not bother it
https://github.com/content-services/content-sources-backend/blob/main/build/Dockerfile i think now its using 1.24.7 
 and i have mentioned 1.24 so it will be able to handle all the subversion /?

https://github.com/quay/claircore/actions/runs/17966798205/job/51100862528#step:3:26



### Comment by @swadeley on September 25, 2025 at 08:20 AM UTC

Hi, Konflux check failed `tests/test_repository_api_only.py::test_bulk_import` which is not relevant to this PR.

### Comment by @marusak on September 25, 2025 at 11:37 AM UTC

> i was thinking to update that too but docker file says it uses latest quay image so i did not bother it
https://github.com/content-services/content-sources-backend/blob/main/build/Dockerfile i think now its using 1.24.7
and i have mentioned 1.24 so it will be able to handle all the subversion /?

I believe @jlsherrill mentioned the other day in standup that we should change it to 1.24 to make sure it is locked to the same version. Isn't that so? :thinking: 

### Comment by @jlsherrill on September 25, 2025 at 02:19 PM UTC

yep!  all it is is changing the 'latest' tag here:

https://github.com/content-services/content-sources-backend/blob/main/build/Dockerfile#L7

to `1.24`

### Comment by @mayurilahane on October 08, 2025 at 01:42 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Commented on September 24, 2025 at 11:59 AM UTC

I would do a dependency bump here: `go get -u ./...`

### Review by @rverdile - Approved on October 08, 2025 at 04:05 PM UTC

looks good, nice job!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1208*

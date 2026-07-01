---
type: pull_request
number: 303
title: "Fixes 602: Update schema output and dataSchema typo"
state: merged
author: Andrewgdewar
created: 2023-06-02T20:58:30Z
updated: 2023-06-05T16:24:35Z
closed: 2023-06-05T16:24:35Z
merged: 2023-06-05T16:24:35Z
base_branch: main
head_branch: HMS-602-1
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/303
---

# Pull Request #303: Fixes 602: Update schema output and dataSchema typo

**Author**: @Andrewgdewar
**Created**: June 02, 2023 at 08:58 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `HMS-602-1`

## Description

## Summary

This should fix a few small issues with the cloud schema consuming the notifications from our app. 

## Testing steps


---

## Discussion

### Comment by @jlsherrill on June 02, 2023 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-602

### Comment by @jlsherrill on June 02, 2023 at 09:15 PM UTC

as part of this can you remove or turn to debug some of the logging that was added:  

https://github.com/content-services/content-sources-backend/commit/02353c7e07b28e9b2b6df68d069005af7cf17d7b
https://github.com/content-services/content-sources-backend/commit/3bc8a1f8e80af8eddaed0691f4e5ce4b6d950fa9
https://github.com/content-services/content-sources-backend/commit/c9b4ab5b6b77088ae40f3288880f9815a54f340e#diff-a3d824da3c42420cd5cbb0a4a2c0e7b5bfddd819652788a0596d195dc6e31fa5R370
https://github.com/content-services/content-sources-backend/commit/a64ae5206ddec96d7d011c309c4b58bc2d15f108#diff-a3d824da3c42420cd5cbb0a4a2c0e7b5bfddd819652788a0596d195dc6e31fa5R385


### Comment by @Andrewgdewar on June 05, 2023 at 03:41 PM UTC

> as part of this can you remove or turn to debug some of the logging that was added:
> 
> [02353c7](https://github.com/content-services/content-sources-backend/commit/02353c7e07b28e9b2b6df68d069005af7cf17d7b) [3bc8a1f](https://github.com/content-services/content-sources-backend/commit/3bc8a1f8e80af8eddaed0691f4e5ce4b6d950fa9) [c9b4ab5#diff-a3d824da3c42420cd5cbb0a4a2c0e7b5bfddd819652788a0596d195dc6e31fa5R370](https://github.com/content-services/content-sources-backend/commit/c9b4ab5b6b77088ae40f3288880f9815a54f340e#diff-a3d824da3c42420cd5cbb0a4a2c0e7b5bfddd819652788a0596d195dc6e31fa5R370) [a64ae52#diff-a3d824da3c42420cd5cbb0a4a2c0e7b5bfddd819652788a0596d195dc6e31fa5R385](https://github.com/content-services/content-sources-backend/commit/a64ae5206ddec96d7d011c309c4b58bc2d15f108#diff-a3d824da3c42420cd5cbb0a4a2c0e7b5bfddd819652788a0596d195dc6e31fa5R385)

^^ will create a follow-up ticket shortly for this. I want to confirm current changes are working in stage first.


### Comment by @jlsherrill on June 05, 2023 at 03:54 PM UTC

there are some linting issues that will need to be fixed

---

## Reviews

### Review by @Andrewgdewar - Commented on June 02, 2023 at 08:59 PM UTC

### Review by @jlsherrill - Commented on June 02, 2023 at 09:11 PM UTC

### Review by @jlsherrill - Commented on June 02, 2023 at 09:11 PM UTC

### Review by @jlsherrill - Commented on June 02, 2023 at 09:12 PM UTC

### Review by @jlsherrill - Commented on June 02, 2023 at 09:13 PM UTC

### Review by @Andrewgdewar - Commented on June 05, 2023 at 03:10 PM UTC

### Review by @jlsherrill - Approved on June 05, 2023 at 03:53 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/303*

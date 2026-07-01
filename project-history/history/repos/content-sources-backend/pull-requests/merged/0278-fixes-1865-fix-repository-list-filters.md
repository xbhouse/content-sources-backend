---
type: pull_request
number: 278
title: "Fixes 1865: fix repository list filters"
state: merged
author: croissanne
created: 2023-05-23T13:51:14Z
updated: 2023-06-06T09:30:14Z
closed: 2023-06-06T09:20:48Z
merged: 2023-06-06T09:20:48Z
base_branch: main
head_branch: fix-repos-filters
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/278
---

# Pull Request #278: Fixes 1865: fix repository list filters

**Author**: @croissanne
**Created**: May 23, 2023 at 01:51 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `fix-repos-filters`

## Description

## Summary

Takes into account 'any' in AvailableForArch/AvailableForVersion filters.

## Testing steps

1.  create a repository with arch 'x86_64'
2. create a 2nd repository with arch any
3.  curl  /api/content-sources/v1/repositories/?available_for_arch=x86_64

you should get both repositories.
Repeat a similar test for distribution_version and available_for_version


---

## Discussion

### Comment by @app-sre-bot on May 23, 2023 at 01:52 PM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on May 23, 2023 at 01:54 PM UTC

oktotest

### Comment by @jlsherrill on May 31, 2023 at 02:00 PM UTC

https://issues.redhat.com/browse/HMS-1865

### Comment by @swadeley on June 06, 2023 at 07:44 AM UTC

/retest

### Comment by @swadeley on June 06, 2023 at 08:58 AM UTC

Hi

testing in a shell

I created two repos

In [10]: repo1 = dict(
    ...:         name='repo1-Arch-86',
    ...:         url='https://alpha.example.com',
    ...:         distribution_arch='x86_64',
    ...:         distribution_versions=['7']
    ...:      )
    ...: 
    ...: repo2 = dict(
    ...:         name='repo2-Arch-any',
    ...:         url='https://beta.example.com',
    ...:         distribution_arch='any',
    ...:         distribution_versions=['any']
    ...:      )


Re available_for_arch test:

`app.content_sources.rest_client.repositories_api.list_repositories(available_for_arch="any")` returns one repo.
`app.content_sources.rest_client.repositories_api.list_repositories(available_for_arch="x86_64")` returns two repos.


Re. distro version test:

`app.content_sources.rest_client.repositories_api.list_repositories(available_for_version="any")` returns one repo.
`app.content_sources.rest_client.repositories_api.list_repositories(available_for_version="7")` returns both repos.


---

## Reviews

### Review by @jlsherrill - Commented on May 31, 2023 at 02:20 PM UTC

### Review by @jlsherrill - Commented on May 31, 2023 at 02:21 PM UTC

### Review by @croissanne - Commented on May 31, 2023 at 04:16 PM UTC

### Review by @jlsherrill - Approved on June 01, 2023 at 07:22 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/278*

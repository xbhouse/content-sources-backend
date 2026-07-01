---
type: pull_request
number: 14
title: "Fixes 1693: support .xz and .zst compression types"
state: merged
author: rverdile
created: 2023-05-16T20:18:39Z
updated: 2023-05-17T18:57:37Z
closed: 2023-05-17T18:57:21Z
merged: 2023-05-17T18:57:21Z
base_branch: master
head_branch: add-compression-types
labels: []
url: https://github.com/content-services/yummy/pull/14
---

# Pull Request #14: Fixes 1693: support .xz and .zst compression types

**Author**: @rverdile
**Created**: May 16, 2023 at 08:18 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `add-compression-types`

## Description

adds support for compressed .xz and .zst primary xml files

To create your own zstd compressed file:
1. Clone this repo https://github.com/facebook/zstd
2. run `make install` and `make check`
3. navigate into the examples folder and use the simple_compression executable to compress a file. for example `./simple_compression filename`

You can create a repository using the .xz compression type like this:
`createrepo_c    --general-compress-type=xz  ./path_to_repo`

Or an individual file: `xz -kz ./path_to_file`

---

## Discussion

### Comment by @jlsherrill on May 17, 2023 at 05:00 PM UTC

closes #11 

---

## Reviews

### Review by @jlsherrill - Approved on May 17, 2023 at 04:58 PM UTC

---

*Archived from: https://github.com/content-services/yummy/pull/14*

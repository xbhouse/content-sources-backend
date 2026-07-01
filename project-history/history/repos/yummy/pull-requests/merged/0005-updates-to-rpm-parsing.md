---
type: pull_request
number: 5
title: "Updates to rpm parsing"
state: merged
author: jlsherrill
created: 2022-07-06T16:16:47Z
updated: 2022-07-12T01:31:11Z
closed: 2022-07-12T01:31:11Z
merged: 2022-07-12T01:31:11Z
base_branch: master
head_branch: more
labels: []
url: https://github.com/content-services/yummy/pull/5
---

# Pull Request #5: Updates to rpm parsing

**Author**: @jlsherrill
**Created**: July 06, 2022 at 04:16 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `more`

## Description

* parses summary and checksum
* treats epoch as an int32 which is how rpm treats it
* Requires the user to pass in an http client.  This allows for
  client certs and other options to be set on the client

---

## Reviews

### Review by @Andrewgdewar - Approved on July 11, 2022 at 08:06 PM UTC

Moar data..TLDR.

---

*Archived from: https://github.com/content-services/yummy/pull/5*

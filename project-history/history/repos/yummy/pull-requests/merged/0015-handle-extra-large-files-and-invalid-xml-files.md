---
type: pull_request
number: 15
title: "handle extra large files and invalid xml files"
state: merged
author: jlsherrill
created: 2023-09-26T01:06:11Z
updated: 2023-09-26T16:31:41Z
closed: 2023-09-26T16:31:37Z
merged: 2023-09-26T16:31:37Z
base_branch: master
head_branch: bad_data_2
labels: []
url: https://github.com/content-services/yummy/pull/15
---

# Pull Request #15: handle extra large files and invalid xml files

**Author**: @jlsherrill
**Created**: September 26, 2023 at 01:06 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `bad_data_2`

## Description

This:
* refactors ParseCompressedXMLData to not use ReadAll.  This should reduce memory consumption (but did not fully solve the issue of having an extra large compressed file).
* uses a LimitReader to limit the amount of xml data we load.  This defaults to a limit of 500 mb, which seems to handle all the rhel repos, fedora 39 everything, epel 9
* Upgrades to go 1.20

---

## Discussion

### Comment by @jlsherrill on September 26, 2023 at 01:25 AM UTC

feel free to test with https://github.com/lpardoRH/custom-repo/blob/zip_bomb2/repodata/22803d047a64ae434d13c9adef3db936042544058ed3c03f12c6cb17bafff766-primary.xml.gz  locally

---

## Reviews

### Review by @jlsherrill - Commented on September 26, 2023 at 01:24 AM UTC

### Review by @rverdile - Commented on September 26, 2023 at 03:38 PM UTC

### Review by @rverdile - Approved on September 26, 2023 at 03:38 PM UTC

tested and lgtm!

### Review by @rverdile - Commented on September 26, 2023 at 03:49 PM UTC

### Review by @jlsherrill - Commented on September 26, 2023 at 04:31 PM UTC

---

*Archived from: https://github.com/content-services/yummy/pull/15*

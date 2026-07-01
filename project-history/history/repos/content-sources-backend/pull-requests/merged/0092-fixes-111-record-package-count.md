---
type: pull_request
number: 92
title: "Fixes 111: record package count"
state: merged
author: jlsherrill
created: 2022-08-30T20:13:13Z
updated: 2022-09-28T13:00:31Z
closed: 2022-09-28T12:48:45Z
merged: 2022-09-28T12:48:45Z
base_branch: main
head_branch: package_count
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/92
---

# Pull Request #92: Fixes 111: record package count

**Author**: @jlsherrill
**Created**: August 30, 2022 at 08:13 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `package_count`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on August 30, 2022 at 08:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-111

### Comment by @avisiedo on September 02, 2022 at 05:00 PM UTC

At `pkg/dao/repositories.go` method `Update` needs the last line below:

```go
	repo.URL = repoIn.URL
	repo.Revision = repoIn.Revision
	repo.PackageCount = repoIn.PackageCount
```

### Comment by @jlsherrill on September 12, 2022 at 03:41 PM UTC

This is ready for more review

### Comment by @avisiedo on September 14, 2022 at 04:15 PM UTC

I launched the below successfully and I got the package_count into the database:

```raw
make db-clean clean build db-up repos-import
./release/external-repos introspect-all
make db-cli-connect
select url, package_count from repositories;
```
branch need rebase; no further comments

### Comment by @jlsherrill on September 15, 2022 at 04:07 PM UTC

note, this is currently broken by https://issues.redhat.com/browse/HMSCONTENT-185

### Comment by @jlsherrill on September 21, 2022 at 05:40 PM UTC

need to update this to use a migration!

### Comment by @jlsherrill on September 23, 2022 at 05:53 PM UTC

This should be good for re-review @avisiedo 

### Comment by @mshriver on September 28, 2022 at 12:48 PM UTC

Merging before heavy QE given priorities and tasking. 

---

## Reviews

### Review by @rverdile - Commented on September 08, 2022 at 06:59 PM UTC

DeepCopyInto may also need an update

### Review by @avisiedo - Commented on September 26, 2022 at 09:33 AM UTC

### Review by @avisiedo - Changes Requested on September 26, 2022 at 10:07 AM UTC

Only a few changes at unit tests, and a question if creating a new ticket for some previous change that I have realized reviewing this pr.

the rest lgtm :+1: 

### Review by @jlsherrill - Commented on September 26, 2022 at 04:30 PM UTC

### Review by @avisiedo - Dismissed on September 27, 2022 at 05:54 AM UTC

### Review by @avisiedo - Commented on September 27, 2022 at 06:01 AM UTC

lgtm

I only see the bot is raising some errors

```raw
ApiRepositoryResponse has no attribute 'package_count' at ['received_data']['package_count']
```

I didn't realize about the qe-testing-needed label.

### Review by @avisiedo - Approved on September 27, 2022 at 02:21 PM UTC

lgtm

### Review by @Andrewgdewar - Approved on September 27, 2022 at 08:57 PM UTC

![Screen Shot 2022-09-27 at 2 57 31 PM](https://user-images.githubusercontent.com/38083295/192633952-448981eb-0507-4877-9497-1cbb93cd559f.png)


---

*Archived from: https://github.com/content-services/content-sources-backend/pull/92*

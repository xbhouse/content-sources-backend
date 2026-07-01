---
type: pull_request
number: 99
title: "Fixes 165: Set list default sort order (name asc)"
state: merged
author: Andrewgdewar
created: 2022-09-06T18:13:33Z
updated: 2022-09-21T21:08:31Z
closed: 2022-09-21T20:04:05Z
merged: 2022-09-21T20:04:05Z
base_branch: main
head_branch: CONTENT-165
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/99
---

# Pull Request #99: Fixes 165: Set list default sort order (name asc)

**Author**: @Andrewgdewar
**Created**: September 06, 2022 at 06:13 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `CONTENT-165`

## Description

Adds some logic for ordering the repository list response. 

Current sort_by parameters: 
- 	name  // Default
- 	url 
- 	distribution_arch 
- 	distribution_versions
-	package_count
-	status

All parameters can optionally add either ":asc" or ":desc"

A sort_by list is possible by using the following template when constructing the url params: 

.../content-sources/v1/repositories/?sort_by[]=status&sort_by[]=url:asc&sort_by[]=name:desc

this will construct a sql query as follows: status asc, url asc, name desc

For more information: [console doc sorting standardization](https://consoledot.pages.redhat.com/docs/dev/developer-references/rest/sorting.html)

---

## Discussion

### Comment by @jlsherrill on September 06, 2022 at 06:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-165

---

## Reviews

### Review by @Andrewgdewar - Commented on September 06, 2022 at 06:15 PM UTC

### Review by @Andrewgdewar - Commented on September 06, 2022 at 06:21 PM UTC

### Review by @Andrewgdewar - Commented on September 06, 2022 at 08:26 PM UTC

### Review by @Andrewgdewar - Commented on September 06, 2022 at 08:28 PM UTC

### Review by @jlsherrill - Commented on September 14, 2022 at 02:04 PM UTC

### Review by @jlsherrill - Commented on September 14, 2022 at 02:10 PM UTC

### Review by @jlsherrill - Commented on September 14, 2022 at 03:47 PM UTC

### Review by @Andrewgdewar - Commented on September 14, 2022 at 05:59 PM UTC

### Review by @Andrewgdewar - Commented on September 14, 2022 at 06:22 PM UTC

### Review by @swadeley - Commented on September 14, 2022 at 07:39 PM UTC

### Review by @Andrewgdewar - Commented on September 14, 2022 at 08:54 PM UTC

### Review by @Andrewgdewar - Commented on September 14, 2022 at 08:54 PM UTC

### Review by @Andrewgdewar - Commented on September 14, 2022 at 09:47 PM UTC

### Review by @jlsherrill - Changes Requested on September 15, 2022 at 05:00 PM UTC

### Review by @Andrewgdewar - Commented on September 15, 2022 at 09:08 PM UTC

### Review by @Andrewgdewar - Commented on September 15, 2022 at 09:23 PM UTC

### Review by @Andrewgdewar - Commented on September 15, 2022 at 09:24 PM UTC

### Review by @Andrewgdewar - Commented on September 16, 2022 at 07:13 PM UTC

### Review by @Andrewgdewar - Commented on September 16, 2022 at 07:25 PM UTC

### Review by @jlsherrill - Approved on September 20, 2022 at 02:29 PM UTC

### Review by @swadeley - Commented on September 21, 2022 at 06:06 PM UTC

Typos

### Review by @Andrewgdewar - Commented on September 21, 2022 at 07:26 PM UTC

### Review by @Andrewgdewar - Commented on September 21, 2022 at 07:30 PM UTC

### Review by @swadeley - Commented on September 21, 2022 at 08:03 PM UTC

lgtm

### Review by @Andrewgdewar - Commented on September 21, 2022 at 09:08 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/99*

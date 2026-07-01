---
type: pull_request
number: 988
title: "HMS-5522: cleanup snapshots even if partially deleted"
state: merged
author: jlsherrill
created: 2025-02-19T22:58:23Z
updated: 2025-02-27T19:21:51Z
closed: 2025-02-27T19:17:28Z
merged: 2025-02-27T19:17:28Z
base_branch: main
head_branch: 5522
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/988
---

# Pull Request #988: HMS-5522: cleanup snapshots even if partially deleted

**Author**: @jlsherrill
**Created**: February 19, 2025 at 10:58 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `5522`

## Description

## Summary

* Changes snapshot SoftDelete to not error if the given snapshot was already soft deleted
* Adds the ability to list snapshots that are also soft deleted and uses this in cleanup
* makes pulp repo version delete and distribution delete return a string pointer instead of a string to make it obvious that 404s will not return a task href
* modified it so that if an error occurred for one repository, it would log it and still try to do the rest.

## Testing steps

create a repository and get its repo config uuid

use this to create old snapshots, run in the terminal:

```
rcuuid="abcde"
created="2024-10-02 00:55:30.953994+00"
snapUUID=`uuidgen`

query="insert into snapshots (uuid,created_at, repository_configuration_uuid, version_href, publication_href, distribution_href, distribution_path) VALUES  ('$snapUUID','$created', '$rcuuid', '', '', '/$RANDOM/$RANDOM/', '/$RANDOM/$RANDOM');"
psql "sslmode=disable dbname=content user=content host=localhost port=5433 password=content" -c "$query"
```
do this a few times with different dates
then run ```OPTIONS_SNAPSHOT_RETAIN_DAYS_LIMIT=30 go run cmd/external-repos/main.go process-repos```

check the db/api and see the older snapshots get deleted.


---

## Discussion

### Comment by @jlsherrill on February 19, 2025 at 11:00 PM UTC

https://issues.redhat.com/browse/HMS-5522

### Comment by @jlsherrill on February 27, 2025 at 06:26 PM UTC

Updated!

---

## Reviews

### Review by @rverdile - Commented on February 26, 2025 at 08:06 PM UTC

### Review by @rverdile - Commented on February 26, 2025 at 08:07 PM UTC

### Review by @rverdile - Commented on February 26, 2025 at 08:08 PM UTC

### Review by @rverdile - Commented on February 26, 2025 at 08:08 PM UTC

everything looks and works good. a couple very small comments. 

### Review by @dominikvagner - Commented on February 27, 2025 at 12:48 PM UTC

### Review by @dominikvagner - Commented on February 27, 2025 at 12:51 PM UTC

looks good and works, thanks a lot for the assist on this one 🫶
just one small nit 🤏🏼

### Review by @jlsherrill - Commented on February 27, 2025 at 05:46 PM UTC

### Review by @rverdile - Approved on February 27, 2025 at 06:42 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/988*

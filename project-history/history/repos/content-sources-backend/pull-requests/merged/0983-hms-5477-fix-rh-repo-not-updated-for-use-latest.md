---
type: pull_request
number: 983
title: "HMS-5477: fix rh repo not updated for use-latest"
state: merged
author: rverdile
created: 2025-02-17T21:44:55Z
updated: 2025-02-18T19:38:12Z
closed: 2025-02-18T19:38:10Z
merged: 2025-02-18T19:38:10Z
base_branch: main
head_branch: rh-repo-bug
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/983
---

# Pull Request #983: HMS-5477: fix rh repo not updated for use-latest

**Author**: @rverdile
**Created**: February 17, 2025 at 09:44 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `rh-repo-bug`

## Description

## Summary
Red Hat repos within "use-latest" templates were not being updated on new snapshots because the template List query filters by org ID. When red hat repos are snapshotted nightly, an org ID of "-1" is used, so no templates would be returned. This adds a second query specifically to get all the templates containing the updated red hat repo.

## Testing steps
Testing this manually would be hard. The integration test covers it. I think the best you could do outside of stage is mimic the integration test manually

---

## Discussion

### Comment by @jlsherrill on February 17, 2025 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-5477

### Comment by @jlsherrill on February 18, 2025 at 12:36 AM UTC

i actually tested this by:

* commenting the CA option out on remote creation
*  editing the red_hat.json file to only include 1repo and set the url to a repo I can control
* load the repos, snapshot the url
* create a template set to latest with this 'red hat' repo in it
* curl the template url and see the initial content
* update my yum repo to have more packages
* clear the last_snapshot_task_uuid on the repo config & re-snapshot the repo 
* curl the template repo url again see the new content
it worked great!

just one comment and it looks like maybe an issue with the integration tests

### Comment by @rverdile on February 18, 2025 at 03:06 PM UTC

> i actually tested this by:

good idea!

### Comment by @rverdile on February 18, 2025 at 04:14 PM UTC

added the rpm check

---

## Reviews

### Review by @jlsherrill - Commented on February 18, 2025 at 12:26 AM UTC

### Review by @jlsherrill - Approved on February 18, 2025 at 06:37 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/983*

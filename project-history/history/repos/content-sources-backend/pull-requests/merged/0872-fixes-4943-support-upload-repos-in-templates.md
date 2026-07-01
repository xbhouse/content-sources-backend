---
type: pull_request
number: 872
title: "Fixes 4943: support upload repos in templates"
state: merged
author: jlsherrill
created: 2024-11-04T14:33:12Z
updated: 2024-11-11T21:30:30Z
closed: 2024-11-11T21:19:04Z
merged: 2024-11-11T21:19:04Z
base_branch: main
head_branch: 4943
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/872
---

# Pull Request #872: Fixes 4943: support upload repos in templates

**Author**: @jlsherrill
**Created**: November 04, 2024 at 02:33 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4943`

## Description

## Summary

Upload repos were not being included in template's candlepin environment, and the content sets did not have URLs

Note that this doesn't attempt to fix existing content sets, which i think at this stage is okay.

## Testing steps


1. On this PR   run:
```
make repos-import-rhel9
go run cmd/external-repos/main.go nightly-jobs 
```
start the server and let the RHEL repos snapshot
3. Create at least 1 upload repo and upload at least 1 rpm
4. Create a content template, assign rhel9 and upload repo
5. follow https://github.com/content-services/content-sources-backend/blob/main/docs/register_client.md to register a client and assign it to the template
6. confirm that the client can install content form the redhat repos and upload repo


---

## Discussion

### Comment by @jlsherrill on November 04, 2024 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-4943

### Comment by @jlsherrill on November 05, 2024 at 07:23 PM UTC

[test]

### Comment by @swadeley on November 07, 2024 at 06:14 AM UTC

/retest

### Comment by @jlsherrill on November 08, 2024 at 02:43 PM UTC

[test]

### Comment by @swadeley on November 11, 2024 at 10:04 AM UTC

/retest

### Comment by @swadeley on November 11, 2024 at 11:03 AM UTC

/retest

### Comment by @swadeley on November 11, 2024 at 03:17 PM UTC

/retest

### Comment by @swadeley on November 11, 2024 at 06:40 PM UTC

/retest

### Comment by @swadeley on November 11, 2024 at 09:19 PM UTC

Hi

I can create template with upload repo in the UI:

![image](https://github.com/user-attachments/assets/eab00527-1917-4717-bd98-272c20b96d41)

Testing installing a packge from the download repo will have to be done in stage.

Thank you


---

## Reviews

### Review by @dominikvagner - Approved on November 07, 2024 at 11:23 AM UTC

everything looks good, verified the testing steps and all works as expected! 🎉 
good job, approved! 👍🏼😄 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/872*

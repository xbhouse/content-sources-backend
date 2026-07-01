---
type: pull_request
number: 298
title: "Refs 1470: turn on new tasking system in stage"
state: merged
author: rverdile
created: 2023-05-30T20:42:21Z
updated: 2023-05-31T20:29:30Z
closed: 2023-05-31T20:29:26Z
merged: 2023-05-31T20:29:26Z
base_branch: main
head_branch: tasking-deploy
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/298
---

# Pull Request #298: Refs 1470: turn on new tasking system in stage

**Author**: @rverdile
**Created**: May 30, 2023 at 08:42 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `tasking-deploy`

## Description

## Summary
This is to turn on the new tasking system in stage.

Still trying to figure out how to completely test this in ephemeral. Can see the env var set to "true", but that's it so far. Logging doesn't seem to work and I need to somehow create a user to trigger introspect.

## Testing steps
Should be able to see this change in ephemeral. 

---

## Discussion

### Comment by @jlsherrill on May 30, 2023 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-1470

### Comment by @jlsherrill on May 31, 2023 at 12:18 PM UTC

any idea why this causes automation test failures?

---

## Reviews

### Review by @jlsherrill - Approved on May 31, 2023 at 08:23 PM UTC

Tested in ephemeral, works great!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/298*

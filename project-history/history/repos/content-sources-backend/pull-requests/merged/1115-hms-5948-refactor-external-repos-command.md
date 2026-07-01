---
type: pull_request
number: 1115
title: "HMS-5948: refactor external repos command"
state: merged
author: rverdile
created: 2025-05-21T20:03:13Z
updated: 2025-06-03T15:08:10Z
closed: 2025-06-03T15:08:04Z
merged: 2025-06-03T15:08:04Z
base_branch: main
head_branch: cleanup-cron
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1115
---

# Pull Request #1115: HMS-5948: refactor external repos command

**Author**: @rverdile
**Created**: May 21, 2025 at 08:03 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `cleanup-cron`

## Description

## Summary
Introduces cli library for external-repos commands, and moves the command implementations to a separate package.

Consolidates the cleanup jobs into the same cron. Pulp orphan cleanup is left on a separate weekly cadence, so pulp workers are not overwhelmed.

## Testing steps
1. Test all the external repos commands
2. `make build`
3. `./release/external-repos cleanup` to test the cleanup command. Use `cleanup --type snapshot,task` to specify a specific type. Or use `cleanup --exclude upload,repository` to exclude specific types.
4. Any commands that had arguments now use flags, e.g. `./external-repos snapshot --url <url>` instead of `./external-repos snapshot <url>`
5. Use `./external-repos --help` to find all the commands and their flags.


---

## Discussion

### Comment by @jlsherrill on May 21, 2025 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-5948

### Comment by @rverdile on May 29, 2025 at 07:27 PM UTC

done!

---

## Reviews

### Review by @jlsherrill - Commented on May 22, 2025 at 07:26 PM UTC

### Review by @rverdile - Commented on May 22, 2025 at 07:34 PM UTC

### Review by @xbhouse - Commented on May 27, 2025 at 08:23 PM UTC

it looks like the playwright api test for snapshot cleanup is failing due to the command format change, do you mind updating that test too? i think you'd just need to change [this line](https://github.com/content-services/content-sources-backend/blob/main/_playwright-tests/tests/API/Snapshots.spec.ts#L141)

### Review by @xbhouse - Commented on May 29, 2025 at 06:17 PM UTC

do you mind also updating the commands in the [introspection](https://github.com/content-services/content-sources-backend/blob/main/docs/workflows/introspection.md) and [snapshotting](https://github.com/content-services/content-sources-backend/blob/main/docs/workflows/snapshotting.md) docs? otherwise this looks really nice! great job :smile: 

### Review by @xbhouse - Approved on May 29, 2025 at 08:54 PM UTC

awesome job!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1115*

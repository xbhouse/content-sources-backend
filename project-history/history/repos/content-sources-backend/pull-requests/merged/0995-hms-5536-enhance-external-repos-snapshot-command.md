---
type: pull_request
number: 995
title: "HMS-5536: enhance external-repos snapshot command"
state: merged
author: dominikvagner
created: 2025-02-26T11:01:56Z
updated: 2025-03-03T08:59:58Z
closed: 2025-03-03T08:59:58Z
merged: 2025-03-03T08:59:58Z
base_branch: main
head_branch: 5536
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/995
---

# Pull Request #995: HMS-5536: enhance external-repos snapshot command

**Author**: @dominikvagner
**Created**: February 26, 2025 at 11:01 AM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `5536`

## Description

## Summary
This PR enhances the external-repos snapshot command in few ways:
- By adding a `--force` flag to it, which when used will enqueue those tasks even if they were recently snapshoted
- Allowing it to be run on even on non-RH repos
- Making the command clean the the inputted URLs, so that the ending slash isn't required
- Making it enqueue an `introspect` task along with the `snapshot` and `update-latest-snapshot` tasks

## Testing steps
1. Import the RH Ansible repo (`OPTIONS_REPOSITORY_IMPORT_FILTER=small go run cmd/external-repos/main.go import`)
2. Test that you can snasphot it with the `snapshot-command` command (`go run cmd/external-repos/main.go snapshot https://cdn.redhat.com/content/dist/layered/rhel8/x86_64/ansible/2/os/`). You should see all 3 tasks being enqueued.
3. Do the same with the `--force` flag and another 3 tasks should be enqueued. Try this with and without the slash at the end of the URL.
4. Create a custom repo and verify it also work with it's URL.

---

## Discussion

### Comment by @jlsherrill on February 26, 2025 at 11:30 AM UTC

https://issues.redhat.com/browse/HMS-5536

---

## Reviews

### Review by @xbhouse - Approved on February 27, 2025 at 04:24 PM UTC

lgtm! nice work!!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/995*

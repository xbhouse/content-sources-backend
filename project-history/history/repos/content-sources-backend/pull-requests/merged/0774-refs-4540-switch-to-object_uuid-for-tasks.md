---
type: pull_request
number: 774
title: "Refs 4540: switch to object_uuid for tasks"
state: merged
author: jlsherrill
created: 2024-08-14T18:56:16Z
updated: 2024-08-27T17:50:39Z
closed: 2024-08-27T17:49:38Z
merged: 2024-08-27T17:49:38Z
base_branch: main
head_branch: object_id
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/774
---

# Pull Request #774: Refs 4540: switch to object_uuid for tasks

**Author**: @jlsherrill
**Created**: August 14, 2024 at 06:56 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `object_id`

## Description

## Summary

Previously we had a 'repository_uuid' column on the tasks table.  This changes it to an 'object_uuid' column and adds an 'object_type' column to track the type.  This will allow for tracking different object types to the task (such as templates).

this will allow for cancellation of template tasks

Note that this creates the new columns but does not delete the old column.  this will happen in a followup PR to prevent errors during rollout.  This will need to be promoted to production before the followup PR is merged.

## Testing steps

automation passing 
admin_tasks still works




---

## Discussion

### Comment by @jlsherrill on August 14, 2024 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-4313

### Comment by @jlsherrill on August 14, 2024 at 07:00 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @swadeley on August 15, 2024 at 07:53 AM UTC

/retest

### Comment by @swadeley on August 15, 2024 at 09:03 AM UTC

Hi @jlsherrill , I don't see the Admin tab. 
I deployed backend=pr-774-7975280 together with  frontend=pr-318-518a3b3 .

EDIT: hang on, i think this is user related, i.e. which user I log in with.
EDIT EDIT: yep, its ok now

### Comment by @swadeley on August 15, 2024 at 02:32 PM UTC

All tests pass, and I checked Admin tasks ok too.

### Comment by @jlsherrill on August 15, 2024 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-4540

### Comment by @jlsherrill on August 26, 2024 at 05:43 PM UTC

updated!

### Comment by @jlsherrill on August 27, 2024 at 05:22 PM UTC

[test]

---

## Reviews

### Review by @rverdile - Commented on August 20, 2024 at 02:20 PM UTC

### Review by @rverdile - Commented on August 20, 2024 at 02:22 PM UTC

### Review by @jlsherrill - Commented on August 26, 2024 at 03:16 PM UTC

### Review by @rverdile - Commented on August 26, 2024 at 03:21 PM UTC

### Review by @rverdile - Approved on August 27, 2024 at 03:31 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/774*

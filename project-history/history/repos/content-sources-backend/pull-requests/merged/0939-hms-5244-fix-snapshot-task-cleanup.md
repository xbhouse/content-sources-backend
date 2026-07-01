---
type: pull_request
number: 939
title: "HMS-5244: fix snapshot task cleanup"
state: merged
author: dominikvagner
created: 2025-01-14T12:04:47Z
updated: 2025-01-31T12:09:50Z
closed: 2025-01-31T12:09:50Z
merged: 2025-01-31T12:09:50Z
base_branch: main
head_branch: 5244
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/939
---

# Pull Request #939: HMS-5244: fix snapshot task cleanup

**Author**: @dominikvagner
**Created**: January 14, 2025 at 12:04 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `5244`

## Description

## Summary
This PR fixes up the snapshot task's cleanup that happens on context cancellation. This caused errors that were being reported by glitchtip. 
(`error deleting rpm repository versions: 400 Bad Request: [\"The repository version cannot be deleted because it (or its publications) are currently being used to distribute content. Please update the necessary distributions first.\"]`) 

## Testing steps
1. Create a repo and let is snapshot.
2. Apply this git patch:
```
diff --git a/pkg/tasks/snapshot_helper.go b/pkg/tasks/snapshot_helper.go
--- a/pkg/tasks/snapshot_helper.go	(revision 8c9b3c2e8f371d28878423f47acd3f19039f46cb)
+++ b/pkg/tasks/snapshot_helper.go	(date 1736855224959)
@@ -111,7 +111,7 @@
 		return fmt.Errorf("unable to save distribution task href: %w", err)
 	}
 
-	return nil
+	return context.Canceled
 }
```
3. Update the created repo and create one more.
4. Let them snapshot and verify there are no errors in the logs and that the first repo has only one snapshot.
5. Revert the change made in the 2. step.
6. Trigger snapshots on the 2 created repos and verify that after the snapshot tasks complete their status is valid.



---

## Discussion

### Comment by @jlsherrill on January 14, 2025 at 12:30 PM UTC

https://issues.redhat.com/browse/HMS-5224

### Comment by @jlsherrill on January 15, 2025 at 07:00 AM UTC

https://issues.redhat.com/browse/HMS-5244

---

## Reviews

### Review by @rverdile - Commented on January 16, 2025 at 07:44 PM UTC

### Review by @rverdile - Commented on January 16, 2025 at 07:54 PM UTC

### Review by @dominikvagner - Commented on January 23, 2025 at 09:14 AM UTC

### Review by @dominikvagner - Commented on January 23, 2025 at 09:14 AM UTC

### Review by @rverdile - Commented on January 27, 2025 at 07:03 PM UTC

looks good but needs rebase!

### Review by @rverdile - Approved on January 30, 2025 at 06:42 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/939*

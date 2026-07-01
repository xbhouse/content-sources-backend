---
type: pull_request
number: 338
title: "Fixes 2083: log when a task fails to complete"
state: merged
author: rverdile
created: 2023-07-21T19:38:20Z
updated: 2023-08-11T14:36:56Z
closed: 2023-08-11T14:36:45Z
merged: 2023-08-11T14:36:45Z
base_branch: main
head_branch: bad-url-log
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/338
---

# Pull Request #338: Fixes 2083: log when a task fails to complete

**Author**: @rverdile
**Created**: July 21, 2023 at 07:38 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `bad-url-log`

## Description

## Summary
PR #333 should merge first. This PR is built on that PR.

To make it clearer that the task failed, this changes logging to log the error of the failed task.

This PR in combination with this [commit to zest](https://github.com/content-services/zest/commit/ee55f31f4234ca312ca13fc860b8cd0ca923b2cb) should fully address the JIRA issue. The change to zest will not be used until we update our app to a newer version of zest. 

## Testing steps
1. Create a repo with bad url (i.e. `http://example.com`), with snapshot enabled. 
2. You should see the error `json: cannot unmarshal string into Go struct field TaskResponse.error of type map[string]interface {}"` for the sync task specifically. 
3. To test the change to zest, update your go.mod to use the latest zest version. There will be some syntax errors in the pulp_client package that you will have to manually fix to use the latest version. The functions that end in `Api` have to be changed to `API`. 
4. Now you should see a different error for the sync task.


---

## Discussion

### Comment by @rverdile on July 21, 2023 at 07:38 PM UTC

Do not merge before PR #333 

### Comment by @jlsherrill on July 21, 2023 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-2083

### Comment by @jlsherrill on August 08, 2023 at 02:59 PM UTC

I still see 
```
json: cannot unmarshal string into Go struct field TaskResponse.error of type map[string]interface {}
```
in the task error field.  That should be fixed now right?

### Comment by @rverdile on August 08, 2023 at 05:36 PM UTC

> I still see
> 
> ```
> json: cannot unmarshal string into Go struct field TaskResponse.error of type map[string]interface {}
> ```
> 
> in the task error field. That should be fixed now right?

It's still fixed for me. Are you on a recent zest version?

### Comment by @jlsherrill on August 08, 2023 at 07:30 PM UTC

ah, i must have had an older version.  Looking good.  I am seeing the full traceback in the error, but we probably don't want to expose that to the user.  I'll file a bug!

---

## Reviews

### Review by @jlsherrill - Commented on August 04, 2023 at 08:57 PM UTC

### Review by @jlsherrill - Commented on August 04, 2023 at 08:58 PM UTC

### Review by @rverdile - Commented on August 07, 2023 at 02:03 PM UTC

### Review by @jlsherrill - Dismissed on August 08, 2023 at 02:27 PM UTC

### Review by @jlsherrill - Approved on August 08, 2023 at 07:30 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/338*

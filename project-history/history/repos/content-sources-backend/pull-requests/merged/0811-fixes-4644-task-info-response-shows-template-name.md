---
type: pull_request
number: 811
title: "Fixes 4644: task info response shows template name and uuid"
state: merged
author: dominikvagner
created: 2024-09-11T14:49:07Z
updated: 2024-09-16T14:47:10Z
closed: 2024-09-16T14:47:10Z
merged: 2024-09-16T14:47:10Z
base_branch: main
head_branch: 4644
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/811
---

# Pull Request #811: Fixes 4644: task info response shows template name and uuid

**Author**: @dominikvagner
**Created**: September 11, 2024 at 02:49 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4644`

## Description

## Summary
For tasks relating to templates it doesn't make sense to use the `repository_name` and `repository_uuid` fields, so this PR changes those to `object_name` and `object_uuid` and a new field `object_type` which distinguishes to which type of object (`"repository" || "template"`) they are associated to.

## Testing steps
1. Create a repository and a template.
2. List tasks with `GET /api/content-sources/v1.0/tasks/`.
3. Verify the `object_*` fields are present and correct.


---

## Discussion

### Comment by @jlsherrill on September 11, 2024 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-4644

### Comment by @swadeley on September 13, 2024 at 08:14 AM UTC

Hi @dominikvagner , please rebase.

### Comment by @swadeley on September 13, 2024 at 11:15 AM UTC

Hi @dominikvagner , I get error:

```
iqe_content_sources_api.exceptions.ApiTypeError: Got an unexpected parameter 'object_uuid' to method `list_tasks` 
````




### Comment by @dominikvagner on September 13, 2024 at 11:34 AM UTC

> Hi @dominikvagner , I get error:
> 
> ```
> iqe_content_sources_api.exceptions.ApiTypeError: Got an unexpected parameter 'object_uuid' to method `list_tasks` 
> ```

hi @swadeley, 
rebased 🫡 \
and as this PR changes the API response for tasks there might be some changes needed in the QE API bindings and maybe some search and replace 🤔😅  

### Comment by @swadeley on September 13, 2024 at 11:47 AM UTC

@dominikvagner , I did do a search and replace in tests like so:
```
:%s/repository_uuid/object_uuid/
:%s/repository_name/object_name/
```

 but 'list_tasks' is in the generated files:
```
iqe-content-sources-plugin/iqe_content_sources_api/api/tasks_api.py:223: in __list_tasks
```

### Comment by @dominikvagner on September 13, 2024 at 02:00 PM UTC

@swadeley 
@xbhouse helped me found the problem and a bug, which I fixed in the latest commit and the QE tests seem to failing because with the search and replace also the query param got replaced, which should have stayed as `repository_uuid` and a new `template_uuid` filter was added 😅 

### Comment by @xbhouse on September 13, 2024 at 06:09 PM UTC

can you add a unit test for filtering the tasks by template uuid? i don't see one but may have missed it 

### Comment by @swadeley on September 13, 2024 at 06:25 PM UTC

Hi, I have merged the updates to API bindings and to the IQE tests. So pr_checks should pass now.

### Comment by @swadeley on September 13, 2024 at 06:26 PM UTC

/retest

### Comment by @xbhouse on September 13, 2024 at 06:51 PM UTC

looks good and works well! just a few small comments 

---

## Reviews

### Review by @xbhouse - Commented on September 13, 2024 at 06:50 PM UTC

### Review by @xbhouse - Commented on September 13, 2024 at 06:50 PM UTC

### Review by @dominikvagner - Commented on September 16, 2024 at 11:46 AM UTC

### Review by @xbhouse - Approved on September 16, 2024 at 02:44 PM UTC

nice job!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/811*

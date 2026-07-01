---
type: pull_request
number: 535
title: "Fixes 2967: Delete content templates"
state: merged
author: xbhouse
created: 2024-01-15T20:53:37Z
updated: 2024-02-13T15:16:21Z
closed: 2024-01-23T16:21:00Z
merged: 2024-01-23T16:21:00Z
base_branch: main
head_branch: delete-templates
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/535
---

# Pull Request #535: Fixes 2967: Delete content templates

**Author**: @xbhouse
**Created**: January 15, 2024 at 08:53 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `delete-templates`

## Description

## Summary

Adds DELETE endpoint and deletion task for content templates 

## Testing steps

- Create a repository
- Create a content template using that repository UUID and grab the template UUID from the response
- Delete content template using that template UUID (`DELETE /templates/:uuid`)
- Response should be 204 No Content and template should be soft deleted, then removed from database once task is finished
- If the deletion task fails, the `deleted_at` field is cleared and record can still be fetched 
- Also adds a test case for the delete method on RepositoryConfig and updates that method to check for a record first (without this it looks like repo configs can be deleted even with an invalid org ID)

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on January 15, 2024 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-2967

### Comment by @xbhouse on January 16, 2024 at 05:40 PM UTC

if the enqueue fails, where should i handle clearing the `deleted_at` field ? should i return an error from the task in the handler and handle that case there or do it directly in the task? or somewhere else?

### Comment by @rverdile on January 18, 2024 at 02:42 PM UTC

> if the enqueue fails, where should i handle clearing the `deleted_at` field ? should i return an error from the task in the handler and handle that case there or do it directly in the task? or somewhere else?

If enqueue fails we can assume there's no task to run, so I think it makes sense to return an error from the handler after clearing the `deleted_at` field.

### Comment by @swadeley on January 22, 2024 at 03:03 PM UTC

/retest

### Comment by @swadeley on January 23, 2024 at 03:44 PM UTC

Hi, works as described:

```
In [2]: print(app.content_sources.rest_client.repositories_api.get_repo_configuration_file('c58d65bc-a46e-4813-890c-928781efae78','81a7418d-ac54-43ce-b6ff-95a15d182b83'))
2024-01-23 15:11:59.125 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, <snip>, params=[]
[test]
name=test
baseurl=https://<snip>.openshiftapps.com/pulp/content/9f97fa1d/c58d65bc-a46e-4813-890c-928781efae78/9837e3c4-8372-409a-8af5-3e1cba20b0b5/
module_hotfixes=0
gpgcheck=0
repo_gpgcheck=0
enabled=1
gpgkey=


In [3]: app.content_sources.rest_client.templates_api.list_templates(sort_by='name')
2024-01-23 15:25:04.836 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, <snip>, params=[('sort_by', 'name')]
Out[3]: 
{'data': [],
 'links': {'first': '/api/content-sources/v1/templates/?limit=100&offset=0&sort_by=name',
           'last': '/api/content-sources/v1/templates/?limit=100&offset=0&sort_by=name'},
 'meta': {'count': 0, 'limit': 100, 'offset': 0}}

In [4]:  app.content_sources.rest_client.templates_api.create_template(dict(
   ...:     ...:   name="my template one",
   ...:     ...:   repository_uuids=["c58d65bc-a46e-4813-890c-928781efae78"],
   ...:     ...:   description="my new template",
   ...:     ...:   arch="x86_64",
   ...:     ...:   version="7"
   ...:     ...: ))
2024-01-23 15:26:04.110 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, <snip>, params=[]
Out[4]: 
{'arch': 'x86_64',
 'date': '0001-01-01T00:00:00Z',
 'description': 'my new template',
 'name': 'my template one',
 'org_id': '3340851',
 'uuid': '5bc75f16-88a1-4f26-9f7a-6be4346456ff',
 'version': '7'}

In [5]: 

In [5]: app.content_sources.rest_client.templates_api.delete_template('5bc75f16-88a1-4f26-9f7a-6be4346456ff')
2024-01-23 15:39:09.972 [    INFO] [iqe.base.rest_client] REST: METHOD=DELETE, <snip>, params=[]

In [6]: app.content_sources.rest_client.templates_api.list_templates(sort_by='name')
2024-01-23 15:39:22.053 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, <snip>, params=[('sort_by', 'name')]
Out[6]: 
{'data': [],
 'links': {'first': '/api/content-sources/v1/templates/?limit=100&offset=0&sort_by=name',
           'last': '/api/content-sources/v1/templates/?limit=100&offset=0&sort_by=name'},
 'meta': {'count': 0, 'limit': 100, 'offset': 0}}

In [7]: 
```



---

## Reviews

### Review by @xbhouse - Commented on January 16, 2024 at 05:38 PM UTC

### Review by @rverdile - Commented on January 18, 2024 at 02:35 PM UTC

### Review by @rverdile - Commented on January 18, 2024 at 05:25 PM UTC

### Review by @rverdile - Commented on January 18, 2024 at 05:37 PM UTC

### Review by @xbhouse - Commented on January 18, 2024 at 05:50 PM UTC

### Review by @xbhouse - Commented on January 18, 2024 at 06:40 PM UTC

### Review by @rverdile - Commented on January 18, 2024 at 07:18 PM UTC

### Review by @rverdile - Commented on January 18, 2024 at 07:29 PM UTC

### Review by @xbhouse - Commented on January 18, 2024 at 07:38 PM UTC

### Review by @xbhouse - Commented on January 19, 2024 at 03:57 PM UTC

### Review by @rverdile - Approved on January 23, 2024 at 03:24 PM UTC

nice job!!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/535*

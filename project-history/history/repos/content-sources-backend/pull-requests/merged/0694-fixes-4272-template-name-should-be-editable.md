---
type: pull_request
number: 694
title: "Fixes 4272: template name should be editable"
state: merged
author: xbhouse
created: 2024-06-06T19:31:24Z
updated: 2024-06-13T09:30:21Z
closed: 2024-06-13T09:04:07Z
merged: 2024-06-13T09:04:07Z
base_branch: main
head_branch: 4272
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/694
---

# Pull Request #694: Fixes 4272: template name should be editable

**Author**: @xbhouse
**Created**: June 06, 2024 at 07:31 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4272`

## Description

## Summary

- Adds support for updating a template name 
- This doesn't have the candlepin integration in place yet (updating a template name will not update the environment name). Ticket for that [here](https://issues.redhat.com/browse/HMS-4290)

## Testing steps

- Add a repository, let it snapshot, then create a template
- Edit the template name. This should work via both the API and UI 

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate

---

## Discussion

### Comment by @jlsherrill on June 06, 2024 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-4272

### Comment by @swadeley on June 07, 2024 at 09:41 AM UTC

Hi

Editing name in UI now works.

I still see the date in the template being one day in the past when you open up to edit, but now I think this is due to browser time versus system time:

```
In [9]: print(datetime.datetime.now())
2024-06-07 10:30:25.916290

In [10]: app.content_sources.rest_client.templates_api.list_templates()
2024-06-07 10:30:38.395 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[]
Out[10]: 
{'data': [{'arch': 'x86_64',
           'date': '2024-06-06T23:00:00Z',
           'description': ' pr-694-5a61abf',
           'name': 'test1 edit1',
           'org_id': '3340851',
           'repository_uuids': ['44616d92-74df-4134-8f5a-992cd2493d78',
                                '972d29ad-b5b4-4fed-b109-477930c085db'],
           'rhsm_environment_id': 'd847708fc1dd43a4a9dd7e0c4e2f3bb1',
           'uuid': 'd847708f-c1dd-43a4-a9dd-7e0c4e2f3bb1',
           'version': '8'}],
 'links': {'first': '/api/content-sources/v1/templates/?limit=100&offset=0',
           'last': '/api/content-sources/v1/templates/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

In [11]: 
```

![image](https://github.com/content-services/content-sources-backend/assets/11247237/5f538866-2c44-41b0-8f7a-8ef501975d1c)



### Comment by @jlsherrill on June 07, 2024 at 01:26 PM UTC

we probably want to add support to the update_repository task that is added in https://github.com/content-services/content-sources-backend/pull/682 to update the name in candlepin.   I think i can handle that there to support this even before this PR goes in.  I'll test them together

### Comment by @jlsherrill on June 07, 2024 at 01:27 PM UTC

code looks great though!

### Comment by @xbhouse on June 07, 2024 at 02:49 PM UTC

> we probably want to add support to the update_repository task that is added in #682 to update the name in candlepin. I think i can handle that there to support this even before this PR goes in. I'll test them together

oh yea good thinking! thanks :) 

### Comment by @jlsherrill on June 07, 2024 at 02:59 PM UTC

oh silly me, my task was for repo update and this is template name changing!   My wires got crossed.  Do you mind updating the update_template_content task to update the template name?   I think you'd just wanna update the code here: https://github.com/content-services/content-sources-backend/blob/main/pkg/tasks/update_template_content.go#L460-L479  to update the environment too

### Comment by @xbhouse on June 07, 2024 at 03:04 PM UTC

@jlsherrill oh XD i thought that update_repository task handled updating a template name for some reason too. yep i can do that!

### Comment by @xbhouse on June 07, 2024 at 07:50 PM UTC

@jlsherrill it looks like candlepin doesn't support updating an environment yet, but it may get added in sometime next week. so i think this PR will need to wait on that. see Ryan's thread [here](https://redhat-internal.slack.com/archives/C04H33UF5U6/p1717788188453389)

### Comment by @swadeley on June 10, 2024 at 02:09 PM UTC

Hi

Please rebase to get fix in https://github.com/content-services/content-sources-backend/pull/697

### Comment by @swadeley on June 13, 2024 at 09:04 AM UTC

Hi

same test results as before for UI.


For API:

I created repo:

```
Out[1]: 
{'data': [{'arch': 'x86_64',
           'created_at': '2024-06-13T08:50:09.859328Z',
           'created_by': 'ephemeral-user',
           'date': '2024-06-12T23:00:00Z',
           'description': '',
           'last_updated_by': 'ephemeral-user',
           'name': 'test-temp',
           'org_id': '3340851',
           'repository_uuids': ['8dc1b279-7ef0-4e0a-87d8-b9c10a2292cb',
                                'fe4f6d45-6328-4874-8653-4c123db4c87c'],
           'rhsm_environment_id': '3327d7ae1bff47c2aa2bce74f0234e47',
           'updated_at': '2024-06-13T08:50:26.486057Z',
           'uuid': '3327d7ae-1bff-47c2-aa2b-ce74f0234e47',
           'version': '8'}],
 'links': {'first': '/api/content-sources/v1/templates/?limit=100&offset=0',
           'last': '/api/content-sources/v1/templates/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}
```

Updated the name:
```
In [3]: template_update_dict = dict(name="new-name el8 ",repository_uuids=['8dc1b279-7ef0-4e0a-87d8-b9c10a2292cb','fe4f6d45-6328-4874-8653-4c123db4c87c'],description="my 1st edit")

In [4]: app.content_sources.rest_client.templates_api.full_update_template('3327d7ae-1bff-47c2-aa2b-ce74f0234e47',template_update_dict)
2024-06-13 10:17:31.610 [    INFO] [iqe.base.rest_client] REST: METHOD=PUT, request_id=1<>, params=[]
Out[4]: 
{'arch': 'x86_64',
 'created_at': '2024-06-13T08:50:09.859328Z',
 'created_by': 'ephemeral-user',
 'date': '2024-06-13T09:17:31.558082Z',
 'description': 'my 1st edit',
 'last_updated_by': 'ephemeral-user',
 'name': 'new-name el8 ',
 'org_id': '3340851',
 'repository_uuids': ['8dc1b279-7ef0-4e0a-87d8-b9c10a2292cb',
                      'fe4f6d45-6328-4874-8653-4c123db4c87c'],
 'rhsm_environment_id': '3327d7ae1bff47c2aa2bce74f0234e47',
 'updated_at': '2024-06-13T09:17:31.560776Z',
 'uuid': '3327d7ae-1bff-47c2-aa2b-ce74f0234e47',
 'version': '8'}
```

Thank you

---

## Reviews

### Review by @jlsherrill - Approved on June 12, 2024 at 03:11 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/694*

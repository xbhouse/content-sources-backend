---
type: pull_request
number: 743
title: "Fixes 4307: add use_latest parameter to template"
state: merged
author: rverdile
created: 2024-07-12T16:44:25Z
updated: 2024-07-19T11:00:19Z
closed: 2024-07-19T10:40:39Z
merged: 2024-07-19T10:40:39Z
base_branch: main
head_branch: use_latest_param
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/743
---

# Pull Request #743: Fixes 4307: add use_latest parameter to template

**Author**: @rverdile
**Created**: July 12, 2024 at 04:44 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `use_latest_param`

## Description

## Summary
- Adds the use_latest parameter to templates
- Adds validation on template update
- Adds validation on template create and update to prevent a date being set when use_latest is true

## Testing steps
1. Create a template with the use_latest parameter set to true
2. If you create with a date specified, and the parameter set to true, there should be a validation error.
3. Fetching that template, you should see the template set to true
4. Update the template to change the use_latest value should also work.
5. Validation now runs on template update, so you should not be able to update a template with a blank name, invalid arch, or invalid version.

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on July 12, 2024 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-4307

### Comment by @xbhouse on July 16, 2024 at 05:22 PM UTC

i think because of [this](https://github.com/content-services/content-sources-backend/blob/81926101684bbacbbb76c149f59ec7d35f5585af/pkg/api/templates.go#L74), when i fully update (PUT) a template to change use_latest from false to true and i don't provide a date, it sets the date to the current time. seeing the same with partial update (PATCH). i think on updating a template from use_latest=false to true, maybe we should make sure the date is set back to zero here too? 

### Comment by @xbhouse on July 16, 2024 at 05:32 PM UTC

when creating a template with use_latest=true and the date set to an empty string, i get an error binding params i think because of the date type.

```
{
    "status": 400,
    "title": "Error binding params",
    "detail": "code=400, message=parsing time \"\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"\" as \"2006\", internal=parsing time \"\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"\" as \"2006\""
}
```        

removing the date from the request or setting the date to zero works: `"date": "0001-01-01T00:00:00Z"`. i'm not sure how to get around this without filling it with that zero date if it's set to an empty string. what do you think?

### Comment by @rverdile on July 16, 2024 at 06:48 PM UTC

> i think because of [this](https://github.com/content-services/content-sources-backend/blob/81926101684bbacbbb76c149f59ec7d35f5585af/pkg/api/templates.go#L74), when i fully update (PUT) a template to change use_latest from false to true and i don't provide a date, it sets the date to the current time. seeing the same with partial update (PATCH). i think on updating a template from use_latest=false to true, maybe we should make sure the date is set back to zero here too?

Good catch! I'll work on updating that

### Comment by @rverdile on July 16, 2024 at 06:51 PM UTC

> removing the date from the request or setting the date to zero works: `"date": "0001-01-01T00:00:00Z"`. i'm not sure how to get around this without filling it with that zero date if it's set to an empty string. what do you think?

If you set the date to empty string, I think I would expect an error. In this case null is really the zero value, right? Or not specifying date at all.

### Comment by @xbhouse on July 16, 2024 at 07:13 PM UTC

> > removing the date from the request or setting the date to zero works: `"date": "0001-01-01T00:00:00Z"`. i'm not sure how to get around this without filling it with that zero date if it's set to an empty string. what do you think?
> 
> If you set the date to empty string, I think I would expect an error. In this case null is really the zero value, right? Or not specifying date at all.

agreed! i was thinking about whether the UI would ever send a request with an empty string, but we could just handle that there so that never happens

### Comment by @swadeley on July 19, 2024 at 10:40 AM UTC

Hi

I created a repo, then a template:

```
In [3]: template_dict = dict(name="api-template el8 ",repository_uuids=['1f3f9318-7758-466f-8fe0-49eb7291143d'],description="my api template",arch="x86_64",version="8",use_latest=True)

In [5]: app.content_sources.rest_client.templates_api.create_template(template_dict)
2024-07-19 11:26:03.576 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=d6117854-d052-4071-a389-6f22de1e6cd2, params=[]
Out[5]: 
{'arch': 'x86_64',
 'created_at': '2024-07-19T10:26:03.527277602Z',
 'created_by': 'ephemeral-user',
 'date': '0001-01-01T00:00:00Z',
 'description': 'my api template',
 'last_updated_by': 'ephemeral-user',
 'name': 'api-template el8 ',
 'org_id': '3340851',
 'repository_uuids': ['1f3f9318-7758-466f-8fe0-49eb7291143d'],
 'rhsm_environment_id': 'd2269a8c7bee4bbb86c4b60572cd2c54',
 'updated_at': '2024-07-19T10:26:03.527277602Z',
 'use_latest': True,
 'uuid': 'd2269a8c-7bee-4bbb-86c4-b60572cd2c54',
 'version': '8'}

In [6]: app.content_sources.rest_client.templates_api.get_template('d2269a8c-7bee-4bbb-86c4-b60572cd2c54')
2024-07-19 11:29:02.461 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[]
Out[6]: 
{'arch': 'x86_64',
 'created_at': '2024-07-19T10:26:03.527277Z',
 'created_by': 'ephemeral-user',
 'date': '0001-01-01T00:00:00Z',
 'description': 'my api template',
 'last_updated_by': 'ephemeral-user',
 'name': 'api-template el8 ',
 'org_id': '3340851',
 'repository_uuids': ['1f3f9318-7758-466f-8fe0-49eb7291143d'],
 'rhsm_environment_id': 'd2269a8c7bee4bbb86c4b60572cd2c54',
 'updated_at': '2024-07-19T10:26:03.527277Z',
 'use_latest': True,
 'uuid': 'd2269a8c-7bee-4bbb-86c4-b60572cd2c54',
 'version': '8'}
```

Notice date format above, in the UI it shows as "31 Dec 0000"
If I open to edit the template in the UI, then the date is shown as "0001-01-01"

I presume that will be fixed in the frontend PR to follow.

Try update the `use_latest` to `False`:
```
In [7]: app.content_sources.rest_client.templates_api.partial_update_template('d2269a8c-7bee-4bbb-86c4-b60572cd2c54', {"use_latest":False})
2024-07-19 11:31:57.481 [    INFO] [iqe.base.rest_client] REST: METHOD=PATCH, request_id=<>, params=[]
Out[7]: 
{'arch': 'x86_64',
 'created_at': '2024-07-19T10:26:03.527277Z',
 'created_by': 'ephemeral-user',
 'date': '0001-01-01T00:00:00Z',
 'description': 'my api template',
 'last_updated_by': 'ephemeral-user',
 'name': 'api-template el8 ',
 'org_id': '3340851',
 'repository_uuids': ['1f3f9318-7758-466f-8fe0-49eb7291143d'],
 'rhsm_environment_id': 'd2269a8c7bee4bbb86c4b60572cd2c54',
 'updated_at': '2024-07-19T10:31:57.43104Z',
 'use_latest': False,
 'uuid': 'd2269a8c-7bee-4bbb-86c4-b60572cd2c54',
 'version': '8'}
```


Try change name:
```
In [8]: app.content_sources.rest_client.templates_api.partial_update_template('d2269a8c-7bee-4bbb-86c4-b60572cd2c54', {"name":''})
2024-07-19 11:32:26.481 [    INFO] [iqe.base.rest_client] REST: METHOD=PATCH, request_id=<>, params=[]
<snip>
HTTP response body: {"errors":[{"status":400,"title":"Error updating template","detail":"Name cannot be blank."}]}
```

lgtm

---

## Reviews

### Review by @xbhouse - Approved on July 18, 2024 at 03:16 PM UTC

looks good!! 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/743*

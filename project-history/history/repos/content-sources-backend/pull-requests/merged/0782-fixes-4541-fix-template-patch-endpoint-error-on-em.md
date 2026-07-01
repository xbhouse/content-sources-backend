---
type: pull_request
number: 782
title: "Fixes 4541: fix template patch endpoint error on empty date"
state: merged
author: dominikvagner
created: 2024-08-20T12:05:35Z
updated: 2024-08-21T09:30:31Z
closed: 2024-08-21T09:25:41Z
merged: 2024-08-21T09:25:41Z
base_branch: main
head_branch: 4541
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/782
---

# Pull Request #782: Fixes 4541: fix template patch endpoint error on empty date

**Author**: @dominikvagner
**Created**: August 20, 2024 at 12:05 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4541`

## Description

## Summary
There is a problem when sending a `PATCH` to update a template and supplying only the `use_latest=true` and not the date (supplying a null or an empty string), that errors out right now, but it should zero out the date instead.
This PR fixes that, by changing the way the date is decoded/unmarshaled from the JSON request body by wrapping it in a custom type (it's necessary to it this way, so that the parsing doesn't quit after first binding error, can't just silence the error returned by `Bind()` as that would ignore any following data).
Accepting `date=""` as zeroed date was also propagated to accompanying template `PUT` and `POST` endpoints.

## Testing steps
#### 1. Create a new template
```bash
http :8000/api/content-sources/v1.0/templates/ "$( ./scripts/header.sh acme 1111)" arch='x86_64' version='9' repository_uuids:='[]' description='testing patch date latest' name='test' use_latest:=false date="2024-07-20T11:15:17+02:00"
```

#### 2. Update `use_latest` to true with `PATCH` and _don't supply_ the date.
```bash
http PATCH :8000/api/content-sources/v1.0/templates/{_NEW TEMPLATE ID_} "$( ./scripts/header.sh acme 1111)" use_latest:=true
```
Should pass and zero out the date. (`"date": "0001-01-01T00:57:44+00:57"`)

#### 3. Revert the update
```bash
http PATCH :8000/api/content-sources/v1.0/templates/{_NEW TEMPLATE ID_} "$( ./scripts/header.sh acme 1111)" use_latest:=false date="2024-07-20T11:15:17+02:00"
```

#### 4. Update `use_latest` to true with `PATCH` and _supply_ the date as empty string.
```bash
http PATCH :8000/api/content-sources/v1.0/templates/{_NEW TEMPLATE ID_} "$( ./scripts/header.sh acme 1111)" use_latest:=true date="" name="renamed"
```
Should pass and zero out the date. (`"date": "0001-01-01T00:57:44+00:57"`) and the name should also change to "renamed" as the data following the date is also correctly parsed.


---

## Discussion

### Comment by @jlsherrill on August 20, 2024 at 12:30 PM UTC

https://issues.redhat.com/browse/HMS-4541

### Comment by @swadeley on August 21, 2024 at 09:25 AM UTC

Hi

lgtm

```

In [3]: app.content_sources.rest_client.templates_api.create_template(dict(arch='x86_64', version='8', repository_uuids=[], description='testing patch date latest', name='test', use_latest=False, date="2024-07-21
   ...: T10:00:00+01:00"))
2024-08-21 10:01:36.753 [    INFO] [iqe.base.rest_client] REST: POST https://ee-jjno106p.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/ with query params [] and x-rh-insights-request-id=57729aab-6478-458e-8fd4-a1df982da8f9
Out[3]: 
{'arch': 'x86_64',
 'created_at': '2024-08-21T09:01:36.701299271Z',
 'created_by': 'ephemeral-user',
 'date': '2024-07-21T09:00:00Z',
 'description': 'testing patch date latest',
 'last_updated_by': 'ephemeral-user',
 'name': 'test',
 'org_id': '3340851',
 'repository_uuids': [],
 'rhsm_environment_id': 'cc495b0b9b9d4efa822a802c04989d19',
 'updated_at': '2024-08-21T09:01:36.701299271Z',
 'use_latest': False,
 'uuid': 'cc495b0b-9b9d-4efa-822a-802c04989d19',
 'version': '8'}


In [5]: app.content_sources.rest_client.templates_api.partial_update_template('cc495b0b-9b9d-4efa-822a-802c04989d19', {"use_latest":True})
2024-08-21 10:14:35.828 [    INFO] [iqe.base.rest_client] REST: PATCH https://ee-jjno106p.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/cc495b0b-9b9d-4efa-822a-802c04989d19 with query params [] and x-rh-insights-request-id=<snip>
Out[5]: 
{'arch': 'x86_64',
 'created_at': '2024-08-21T09:01:36.701299Z',
 'created_by': 'ephemeral-user',
 'date': '0001-01-01T00:00:00Z',
 'description': 'testing patch date latest',
 'last_updated_by': 'ephemeral-user',
 'name': 'test',
 'org_id': '3340851',
 'repository_uuids': [],
 'rhsm_environment_id': 'cc495b0b9b9d4efa822a802c04989d19',
 'updated_at': '2024-08-21T09:14:35.777383Z',
 'use_latest': True,
 'uuid': 'cc495b0b-9b9d-4efa-822a-802c04989d19',
 'version': '8'}

In [6]: app.content_sources.rest_client.templates_api.partial_update_template('cc495b0b-9b9d-4efa-822a-802c04989d19', {"use_latest":False, "date":'2024-07-21T09:00:00Z'})
2024-08-21 10:19:26.878 [    INFO] [iqe.base.rest_client] REST: PATCH https://ee-jjno106p.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/cc495b0b-9b9d-4efa-822a-802c04989d19 with query params [] and x-rh-insights-request-id=<snip>
Out[6]: 
{'arch': 'x86_64',
 'created_at': '2024-08-21T09:01:36.701299Z',
 'created_by': 'ephemeral-user',
 'date': '2024-07-21T09:00:00Z',
 'description': 'testing patch date latest',
 'last_updated_by': 'ephemeral-user',
 'name': 'test',
 'org_id': '3340851',
 'repository_uuids': [],
 'rhsm_environment_id': 'cc495b0b9b9d4efa822a802c04989d19',
 'updated_at': '2024-08-21T09:19:26.826899Z',
 'use_latest': False,
 'uuid': 'cc495b0b-9b9d-4efa-822a-802c04989d19',
 'version': '8'}

In [7]: app.content_sources.rest_client.templates_api.partial_update_template('cc495b0b-9b9d-4efa-822a-802c04989d19', {"use_latest":True, "date":"", "name":"renamed"})
2024-08-21 10:22:49.284 [    INFO] [iqe.base.rest_client] REST: PATCH https://ee-jjno106p.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/cc495b0b-9b9d-4efa-822a-802c04989d19 with query params [] and x-rh-insights-request-id=<snip>
Out[7]: 
{'arch': 'x86_64',
 'created_at': '2024-08-21T09:01:36.701299Z',
 'created_by': 'ephemeral-user',
 'date': '0001-01-01T00:00:00Z',
 'description': 'testing patch date latest',
 'last_updated_by': 'ephemeral-user',
 'name': 'renamed',
 'org_id': '3340851',
 'repository_uuids': [],
 'rhsm_environment_id': 'cc495b0b9b9d4efa822a802c04989d19',
 'updated_at': '2024-08-21T09:22:49.238088Z',
 'use_latest': True,
 'uuid': 'cc495b0b-9b9d-4efa-822a-802c04989d19',
 'version': '8'}

In [8]: 

```

---

## Reviews

### Review by @xbhouse - Approved on August 20, 2024 at 09:06 PM UTC

working well and looks good! nice job! :) 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/782*

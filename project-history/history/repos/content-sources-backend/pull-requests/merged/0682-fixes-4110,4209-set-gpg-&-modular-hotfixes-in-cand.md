---
type: pull_request
number: 682
title: "Fixes 4110,4209: set gpg & modular hotfixes in candlepin"
state: merged
author: jlsherrill
created: 2024-05-30T21:20:20Z
updated: 2024-06-07T20:30:25Z
closed: 2024-06-07T20:02:07Z
merged: 2024-06-07T20:02:07Z
base_branch: main
head_branch: GPG
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/682
---

# Pull Request #682: Fixes 4110,4209: set gpg & modular hotfixes in candlepin

**Author**: @jlsherrill
**Created**: May 30, 2024 at 09:20 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `GPG`

## Description

## Summary

Adds support for gpg URL (On Content) & Modular hotfixes (as an override) in candlepin
Adds an update repository task to update these settings

## Testing steps

1.  Register a system following https://github.com/content-services/content-sources-backend/blob/main/docs/register_client.md
3. Create a custom repository with no GPG key and module filtering enabled (defaults)
4. Create a content template with the repository
5. associate the system to the template
6. On the client,  run   `subscription-manager refresh && subscription-manager repo-override && yum repolist`  check the /etc/yum.repos.d/redhat.repo file.  Verify that gpgcheck=0  and module_hotfixes=1 is not present.
7. Update the repository, add a GPG Key, toggle modularity filtering
8. On the client, run   `subscription-manager refresh && subscription-manager repo-override && yum repolist`  check the /etc/yum.repos.d/redhat.repo file.  Verify that gpgcheck=1, gpgkey=$URL  and module_hotfixes=1 is present.
9. Undo 7, and redo 8, and verify the settings reverted back to their normal behavior.
## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on May 30, 2024 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-4110

### Comment by @jlsherrill on June 05, 2024 at 08:56 PM UTC

oh and @swadeley i don't think there's anything to test here yet, but will be in the future when candlepin integration is turned on

### Comment by @jlsherrill on June 06, 2024 at 01:00 PM UTC

https://issues.redhat.com/browse/HMS-4209

### Comment by @rverdile on June 06, 2024 at 06:42 PM UTC

Is QE needed now that 4209 is included? The automation seems confused

### Comment by @jlsherrill on June 06, 2024 at 07:10 PM UTC

oh yeah, it probably is!  let me change it

### Comment by @swadeley on June 07, 2024 at 01:52 PM UTC

Hi @jlsherrill 

Re: HMS-4209 API - Modify template API to accept a list of repo_uuids

I think we need an apidocs update:

```
In [12]: app.content_sources.rest_client.templates_api.list_templates.attribute_map
Out[12]: 
{'offset': 'offset',
 'limit': 'limit',
 'version': 'version',
 'arch': 'arch',
 'name': 'name',
 'sort_by': 'sort_by'}

```

### Comment by @swadeley on June 07, 2024 at 08:01 PM UTC

Hi

Looks good:

```
In [8]: app.content_sources.rest_client.templates_api.create_template(template_dict)
2024-06-07 20:49:41.180 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=8<>, params=[]
Out[8]: 
{'arch': 'x86_64',
 'date': '0001-01-01T00:00:00Z',
 'description': 'my 2nd template',
 'name': 'test2 el8',
 'org_id': '3340851',
 'repository_uuids': ['c29491cb-fde9-4d3e-8948-97d8bdc09cd8'],
 'rhsm_environment_id': '2a6d25116b584dc3bf62a83989d2f76c',
 'uuid': '2a6d2511-6b58-4dc3-bf62-a83989d2f76c',
 'version': '8'}

In [9]: template_dict = dict(name="test1 el8",repository_uuids=[created_repo1.uuid],description="my 1st template",arch="x86_64",version="8")

In [10]: app.content_sources.rest_client.templates_api.create_template(template_dict)
2024-06-07 20:49:55.331 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[10]: 
{'arch': 'x86_64',
 'date': '0001-01-01T00:00:00Z',
 'description': 'my 1st template',
 'name': 'test1 el8',
 'org_id': '3340851',
 'repository_uuids': ['663b20a8-bfe8-49b1-9290-e1dba054294c'],
 'rhsm_environment_id': '8cb56d03502d4add8127cfca3a624cc5',
 'uuid': '8cb56d03-502d-4add-8127-cfca3a624cc5',
 'version': '8'}

In [11]:
```

at last:

```
In [21]: app.content_sources.rest_client.templates_api.list_templates(repository_uuids="'663b20a8-bfe8-49b1-9290-e1dba054294c','c29491cb-fde9-4d3e-8948-97d8bdc09cd8'")
2024-06-07 21:00:51.481 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<>, params=[('repository_uuids', "'663b20a8-bfe8-49b1-9290-e1dba054294c','c29491cb-fde9-4d3e-8948-97d8bdc09cd8'")]
Out[21]: 
{'data': [{'arch': 'x86_64',
           'date': '0001-01-01T00:00:00Z',
           'description': 'my 1st template',
           'name': 'test1 el8',
           'org_id': '3340851',
           'repository_uuids': ['663b20a8-bfe8-49b1-9290-e1dba054294c'],
           'rhsm_environment_id': '8cb56d03502d4add8127cfca3a624cc5',
           'uuid': '8cb56d03-502d-4add-8127-cfca3a624cc5',
           'version': '8'},
          {'arch': 'x86_64',
           'date': '0001-01-01T00:00:00Z',
           'description': 'my 2nd template',
           'name': 'test2 el8',
           'org_id': '3340851',
           'repository_uuids': ['c29491cb-fde9-4d3e-8948-97d8bdc09cd8'],
           'rhsm_environment_id': '2a6d25116b584dc3bf62a83989d2f76c',
           'uuid': '2a6d2511-6b58-4dc3-bf62-a83989d2f76c',
           'version': '8'}],
 'links': {'first': "/api/content-sources/v1/templates/?limit=100&offset=0&repository_uuids='663b20a8-bfe8-49b1-9290-e1dba054294c','c29491cb-fde9-4d3e-8948-97d8bdc09cd8'",
           'last': "/api/content-sources/v1/templates/?limit=100&offset=0&repository_uuids='663b20a8-bfe8-49b1-9290-e1dba054294c','c29491cb-fde9-4d3e-8948-97d8bdc09cd8'"},
 'meta': {'count': 2, 'limit': 100, 'offset': 0}}

In [22]: 
```

---

## Reviews

### Review by @rverdile - Commented on June 05, 2024 at 05:34 PM UTC

### Review by @rverdile - Commented on June 05, 2024 at 05:52 PM UTC

### Review by @rverdile - Commented on June 05, 2024 at 05:52 PM UTC

### Review by @rverdile - Commented on June 05, 2024 at 05:55 PM UTC

### Review by @jlsherrill - Commented on June 06, 2024 at 05:12 PM UTC

### Review by @jlsherrill - Commented on June 06, 2024 at 05:14 PM UTC

### Review by @rverdile - Approved on June 06, 2024 at 06:42 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/682*

---
type: pull_request
number: 866
title: "Fixes 4897: env_created field reset after updating template"
state: merged
author: xbhouse
created: 2024-10-25T14:59:00Z
updated: 2025-08-05T13:48:16Z
closed: 2024-10-29T09:12:47Z
merged: 2024-10-29T09:12:47Z
base_branch: main
head_branch: fix-query
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/866
---

# Pull Request #866: Fixes 4897: env_created field reset after updating template

**Author**: @xbhouse
**Created**: October 25, 2024 at 02:59 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `fix-query`

## Description

## Summary

Fixes query that sets the `rhsm_environment_created` field. Without this change, any subsequent template update task would set all the other templates' `rhsm_environment_created` fields to false.

(This is not testable in ephemeral because it depends on candlepin)

## Testing steps

1. Create a template. `rhsm_environment_created` should be false initially in the response
2. Fetch the template. `rhsm_environment_created` should be true
3. Create another template and then list them after a couple seconds to let the update-template-content task complete
4. Both templates should have their `rhsm_environment_created` fields set to true



---

## Discussion

### Comment by @jlsherrill on October 25, 2024 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-4897

### Comment by @xbhouse on October 25, 2024 at 03:26 PM UTC

test failing with expired gpg key:

```
 Error Trace:    /Users/bhouse/Repos/content-sources-backend/pkg/dao/repository_configs_test.go:2277
                Error:          Should be true
                Error Trace:    /Users/bhouse/Repos/content-sources-backend/pkg/dao/repository_configs_test.go:2278
                Error:          Not equal: 
                                actual  : "Error validating signature: openpgp: key expired. Is this the correct GPG Key?"
                                +Error validating signature: openpgp: key expired. Is this the correct GPG Key?
```

i'm guessing this is the gpg key in pkg/test/gpg/gpgkey.pub?

### Comment by @jlsherrill on October 25, 2024 at 03:31 PM UTC

yeah, i think so?
```
$ gpg --show-keys --with-fingerprint pkg/test/gpg/gpgkey.pub 
pub   dsa1024 2007-03-08 [SC]
      4CCA 1EAF 950C EE4A B839  76DC A040 830F 7FAC 5991
uid                      Google, Inc. Linux Package Signing Key <linux-packages-keymaster@google.com>
sub   elg2048 2007-03-08 [E]

pub   rsa4096 2016-04-12 [SC]
      EB4C 1BFD 4F04 2F6D DDCC  EC91 7721 F63B D38B 4796
uid                      Google Inc. (Linux Packages Signing Authority) <linux-packages-keymaster@google.com>
sub   rsa4096 2016-04-12 [S] [expired: 2019-04-12]
sub   rsa4096 2017-01-24 [S] [expired: 2020-01-24]
sub   rsa4096 2019-07-22 [S] [expired: 2022-07-21]
sub   rsa4096 2021-10-26 [S] [expired: 2024-10-25]
sub   rsa4096 2023-02-15 [S] [expires: 2026-02-14]
```

### Comment by @jlsherrill on October 25, 2024 at 03:33 PM UTC

doesn't look like we're using it in a test to actually validate a repo, try replacing it with this one? https://packages.microsoft.com/keys/microsoft.asc

### Comment by @xbhouse on October 25, 2024 at 03:36 PM UTC

hmmm with that one i see this error in the same test: `Error validating signature: openpgp: signature made by unknown entity. Is this the correct GPG Key?`

### Comment by @jlsherrill on October 25, 2024 at 03:41 PM UTC

oh then maybe we are testing that, it looks like maybe its a google repo?  try updating it to https://www.google.com/linuxrepositories/

### Comment by @xbhouse on October 25, 2024 at 03:51 PM UTC

yep looking at that one now, looks like that's the one we had that was expired. or well i get the same failure (`openpgp: key expired`), but the output here does look a bit different:

```
bhouse@mac:content-sources-backend $ gpg --show-keys --with-fingerprint pkg/test/gpg/gpgkey.pub 
pub   dsa1024 2007-03-08 [SC]
      4CCA 1EAF 950C EE4A B839  76DC A040 830F 7FAC 5991
uid                      Google, Inc. Linux Package Signing Key <linux-packages-keymaster@google.com>
sub   elg2048 2007-03-08 [E]

pub   rsa4096 2016-04-12 [SC]
      EB4C 1BFD 4F04 2F6D DDCC  EC91 7721 F63B D38B 4796
uid                      Google Inc. (Linux Packages Signing Authority) <linux-packages-keymaster@google.com>
sub   rsa4096 2016-04-12 [S] [expired: 2019-04-12]
sub   rsa4096 2017-01-24 [S] [expired: 2020-01-24]
sub   rsa4096 2019-07-22 [S] [expired: 2022-07-21]
sub   rsa4096 2021-10-26 [S] [expired: 2024-10-25]
sub   rsa4096 2023-02-15 [S] [expires: 2026-02-14]
sub   rsa4096 2024-01-30 [S] [expires: 2027-01-29]
```

### Comment by @swadeley on October 28, 2024 at 11:06 AM UTC

Hi @xbhouse 

I deployed `content-sources-backend=pr-866-7232dbb`

I created a repo

I created a template and then fetched it. In both cases `rhsm_environment_created` was False:

```
In [2]: template_dict = dict(arch='x86_64', version='8', repository_uuids=['ef361385-3470-45ae-b99e-799ae4a14db1'], name='test template', use_latest=True)

In [3]:  app.content_sources.rest_client.templates_api.create_template(template_dict)
2024-10-28 10:25:50.829 [    INFO] [iqe.base.rest_client] REST: POST https://ee-2woml2fz.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/ with query params [] and x-rh-insights-request-id=7475312b-0959-45f4-bd97-cd956aec3af8
Out[3]: 
{'arch': 'x86_64',
 'created_at': '2024-10-28T10:25:50.7602466Z',
 'created_by': 'ephemeral-user',
 'date': '0001-01-01T00:00:00Z',
 'description': '',
 'last_update_snapshot_error': '',
 'last_updated_by': 'ephemeral-user',
 'name': 'test template',
 'org_id': '3340851',
 'repository_uuids': ['ef361385-3470-45ae-b99e-799ae4a14db1'],
 'rhsm_environment_created': False,
 'rhsm_environment_id': 'db774ff31d084963bd476287b7f9a57e',
 'updated_at': '2024-10-28T10:25:50.7602466Z',
 'use_latest': True,
 'uuid': 'db774ff3-1d08-4963-bd47-6287b7f9a57e',
 'version': '8'}

In [4]: app.content_sources.rest_client.templates_api.get_template('db774ff3-1d08-4963-bd47-6287b7f9a57e')
2024-10-28 11:03:28.049 [    INFO] [iqe.base.rest_client] REST: GET https://ee-2woml2fz.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/db774ff3-1d08-4963-bd47-6287b7f9a57e with query params [] and x-rh-insights-request-id=eeb8696c-e18a-4d3d-a62f-85f9bf5aed4a
Out[4]: 
{'arch': 'x86_64',
 'created_at': '2024-10-28T10:25:50.760246Z',
 'created_by': 'ephemeral-user',
 'date': '0001-01-01T00:00:00Z',
 'description': '',
 'last_update_snapshot_error': '',
 'last_updated_by': 'ephemeral-user',
 'name': 'test template',
 'org_id': '3340851',
 'repository_uuids': ['ef361385-3470-45ae-b99e-799ae4a14db1'],
 'rhsm_environment_created': False,
 'rhsm_environment_id': 'db774ff31d084963bd476287b7f9a57e',
 'updated_at': '2024-10-28T10:25:50.760246Z',
 'use_latest': True,
 'uuid': 'db774ff3-1d08-4963-bd47-6287b7f9a57e',
 'version': '8'}

In [5]: 

```

Thank you


### Comment by @xbhouse on October 28, 2024 at 05:39 PM UTC

@swadeley ah yep this isn't testable in ephemeral. this field only ever gets set to true if candlepin is enabled. i think what we've been doing in these cases is just testing in stage

### Comment by @swadeley on October 29, 2024 at 09:12 AM UTC

> @swadeley ah yep this isn't testable in ephemeral. this field only ever gets set to true if candlepin is enabled. i think what we've been doing in these cases is just testing in stage

OK, will do.

### Comment by @swadeley on October 29, 2024 at 10:20 AM UTC

Hi, looks good:

```
In [4]: template_dict = dict(arch='x86_64', version='8', repository_uuids=['301efeeb-f1fa-4c73-90f0-8537d500d4d3'], name='test template', use_latest=True)

In [5]:  app.content_sources.rest_client.templates_api.create_template(template_dict)
2024-10-29 10:13:44.274 [    INFO] [iqe.base.rest_client] REST: POST https://console.stage.redhat.com/api/content-sources/v1/templates/ with query params [] and x-rh-insights-request-id=<>, <>
Out[5]: 
{'arch': 'x86_64',
 'created_at': '2024-10-29T10:13:44.056670252Z',
 'created_by': 'hms-qa-sandbox',
 'date': '0001-01-01T00:00:00Z',
 'description': '',
 'last_update_snapshot_error': '',
 'last_updated_by': 'hms-qa-sandbox',
 'name': 'test template',
 'org_id': '17799759',
 'repository_uuids': ['301efeeb-f1fa-4c73-90f0-8537d500d4d3'],
 'rhsm_environment_created': False,
 'rhsm_environment_id': '<2e7c6ac0adf46bc9cfd00c91a74897>',
 'updated_at': '2024-10-29T10:13:44.056670252Z',
 'use_latest': True,
 'uuid': 'd2e7c6ac-0adf-46bc-9cfd-00c91a748977',
 'version': '8'}

In [6]: app.content_sources.rest_client.templates_api.get_template('d2e7c6ac-0adf-46bc-9cfd-00c91a748977')
2024-10-29 10:14:14.841 [    INFO] [iqe.base.rest_client] REST: GET https://console.stage.redhat.com/api/content-sources/v1/templates/d2e7c6ac-0adf-46bc-9cfd-00c91a748977 with query params [] and x-rh-insights-request-id=<>, <>
Out[6]: 
{'arch': 'x86_64',
 'created_at': '2024-10-29T10:13:44.05667Z',
 'created_by': 'hms-qa-sandbox',
 'date': '0001-01-01T00:00:00Z',
 'description': '',
 'last_update_snapshot_error': '',
 'last_update_task': {'created_at': '2024-10-29T10:13:44Z',
                      'ended_at': '2024-10-29T10:13:47Z',
                      'error': '',
                      'object_name': 'test template',
                      'object_type': 'template',
                      'object_uuid': 'd2e7c6ac-0adf-46bc-9cfd-00c91a748977',
                      'org_id': '17799759',
                      'status': 'completed',
                      'type': 'update-template-content',
                      'uuid': 'cff0b0a1-8d1f-47ad-b910-5c305489a756'},
 'last_update_task_uuid': 'cff0b0a1-8d1f-47ad-b910-5c305489a756',
 'last_updated_by': 'hms-qa-sandbox',
 'name': 'test template',
 'org_id': '17799759',
 'repository_uuids': ['301efeeb-f1fa-4c73-90f0-8537d500d4d3'],
 'rhsm_environment_created': True,
 'rhsm_environment_id': '<2e7c6ac0adf46bc9cfd00c91a74897>',
 'updated_at': '2024-10-29T10:13:44.05667Z',
 'use_latest': True,
 'uuid': 'd2e7c6ac-0adf-46bc-9cfd-00c91a748977',
 'version': '8'}

In [7]: 

```

### Comment by @xbhouse on October 29, 2024 at 01:06 PM UTC

thank you @swadeley 

---

## Reviews

### Review by @jlsherrill - Approved on October 25, 2024 at 03:19 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/866*

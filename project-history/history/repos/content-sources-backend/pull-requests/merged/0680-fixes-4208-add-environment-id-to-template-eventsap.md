---
type: pull_request
number: 680
title: "Fixes 4208: add environment ID to template events/api"
state: merged
author: jlsherrill
created: 2024-05-28T18:19:54Z
updated: 2024-05-30T08:30:22Z
closed: 2024-05-30T08:09:22Z
merged: 2024-05-30T08:09:22Z
base_branch: main
head_branch: 4208
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/680
---

# Pull Request #680: Fixes 4208: add environment ID to template events/api

**Author**: @jlsherrill
**Created**: May 28, 2024 at 06:19 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4208`

## Description

## Summary

adds a new parameter 'client_environment_id' to kafka template events & api responses

## Testing steps

in local setup:  
```
KAFKA_TOPICS=platform.content-sources.template   make kafka-topic-consume  
```

now create a content template via the api/ui and watch the result:

```
Partition:0	content-type:application/cloudevents+json	null	{"specversion":"1.0","id":"14ebf3a8-1c8b-497a-9228-92a826452d81","source":"urn:redhat:source:console:app:repositories","type":"com.redhat.console.repositories.template-created","subject":"urn:redhat:subject:console:rhel:template-created","datacontenttype":"application/json","time":"2024-05-28T18:18:07.558203259Z","data":[{"uuid":"77e526a0-0ee9-4284-8d32-3eebe4ad9e84","name":"test","org_id":"9","description":"","arch":"","version":"","date":"2024-05-02T10:04:05-04:00","repository_uuids":[],"client_environment_id":"77e526a00ee942848d323eebe4ad9e84"}],"redhatorgid":"9"}
```

notice: `"client_environment_id":"77e526a00ee942848d323eebe4ad9e84"`

for api testing,   GET /templates/   GET /templates/UUID  & POST /templates/   should all return client_environment_id as part of the response


## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on May 28, 2024 at 06:31 PM UTC

https://issues.redhat.com/browse/HMS-4208

### Comment by @swadeley on May 29, 2024 at 10:54 AM UTC

Hi

I built API docs and tested:

```
In [10]: template_dict = dict(name="test",repository_uuids=['fd888d63-7aaf-4870-8515-034ba0e81f53'],description="my not so new template",arch="x86_64",version="7")

In [11]: app.content_sources.rest_client.templates_api.create_template(template_dict)
2024-05-29 11:45:56.825 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]


In [12]: app.content_sources.rest_client.templates_api.create_template(template_dict)
2024-05-29 11:46:25.957 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[12]: 
{'arch': 'x86_64',
 'client_environment_id': '189ce5d99d7d401d986c5aedfe4b7965',
 'date': '0001-01-01T00:00:00Z',
 'description': 'my not so new template',
 'name': 'test',
 'org_id': '3340851',
 'repository_uuids': ['fd888d63-7aaf-4870-8515-034ba0e81f53'],
 'uuid': '189ce5d9-9d7d-401d-986c-5aedfe4b7965',
 'version': '7'}
```

Looks correct.



### Comment by @swadeley on May 29, 2024 at 12:33 PM UTC

Hi

Client ID is not the same. Is that expected?:

```
n [6]: app.content_sources.rest_client.templates_api.list_templates()
2024-05-29 13:32:39.643 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=c17e130a-d962-4370-955c-075f8022aea4, params=[]
Out[6]: 
{'data': [{'arch': 'x86_64',
           'client_environment_id': '189ce5d99d7d401d986c5aedfe4b7965',
           'date': '0001-01-01T00:00:00Z',
           'description': 'my not so new template',
           'name': 'test',
           'org_id': '3340851',
           'repository_uuids': ['fd888d63-7aaf-4870-8515-034ba0e81f53'],
           'uuid': '189ce5d9-9d7d-401d-986c-5aedfe4b7965',
           'version': '7'},
          {'arch': 'x86_64',
           'client_environment_id': '0c3b94eb579c4527bdd1222d1ab559f8',
           'date': '0001-01-01T00:00:00Z',
           'description': 'my 2nd template',
           'name': 'test2',
           'org_id': '3340851',
           'repository_uuids': ['fd888d63-7aaf-4870-8515-034ba0e81f53'],
           'uuid': '0c3b94eb-579c-4527-bdd1-222d1ab559f8',
           'version': '7'},
          {'arch': 'x86_64',
           'client_environment_id': '2c795e1532ca4c72a0a33782d527a9dc',
           'date': '0001-01-01T00:00:00Z',
           'description': 'my 3rd template',
           'name': 'test3',
           'org_id': '3340851',
           'repository_uuids': ['fd888d63-7aaf-4870-8515-034ba0e81f53'],
           'uuid': '2c795e15-32ca-4c72-a0a3-3782d527a9dc',
           'version': '7'}],
 'links': {'first': '/api/content-sources/v1/templates/?limit=100&offset=0',
           'last': '/api/content-sources/v1/templates/?limit=100&offset=0'},
 'meta': {'count': 3, 'limit': 100, 'offset': 0}}

In [7]: 


```

### Comment by @swadeley on May 29, 2024 at 01:43 PM UTC

Hi

If `client_environment_id` is per template and not per user, then user is not the client.
The name is not clear, why would you have a different environment for each template even if user and repo is the same?
Would these be better?

rhsm-client-id ?
rhsm_environment_id?
rhsm_client_environment_id?

If the user cannot use these for anything and is not expected to know what "client" refers to, then maybe chose a name that is good to use in support tickets or in talking to Candlepin team (to avoid confusion).

### Comment by @rverdile on May 29, 2024 at 03:24 PM UTC

I like anything with "environment" in it, with slight preference towards "rhsm_environment_id".

### Comment by @jlsherrill on May 29, 2024 at 06:11 PM UTC

changed to rhsm_environment_id

### Comment by @swadeley on May 30, 2024 at 08:05 AM UTC

Hi

```
In [3]: template_dict = dict(name="test2 el8",repository_uuids=[created_repo.uuid],description="my 2nd template",arch="x86_64",version="8")

In [4]: app.content_sources.rest_client.templates_api.create_template(template_dict)
2024-05-30 09:03:42.511 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[4]: 
{'arch': 'x86_64',
 'date': '0001-01-01T00:00:00Z',
 'description': 'my 2nd template',
 'name': 'test2 el8',
 'org_id': '3340851',
 'repository_uuids': ['67c122c3-9310-4863-ad32-fd0ac2ab107e'],
 'rhsm_environment_id': 'df94f91f3f834c51b985bee90b56e07d',
 'uuid': 'df94f91f-3f83-4c51-b985-bee90b56e07d',
 'version': '8'}

In [5]: 
```

---

## Reviews

### Review by @rverdile - Commented on May 29, 2024 at 07:11 PM UTC

### Review by @rverdile - Commented on May 29, 2024 at 07:13 PM UTC

### Review by @swadeley - Commented on May 29, 2024 at 07:33 PM UTC

### Review by @jlsherrill - Commented on May 29, 2024 at 07:39 PM UTC

### Review by @rverdile - Approved on May 29, 2024 at 07:59 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/680*

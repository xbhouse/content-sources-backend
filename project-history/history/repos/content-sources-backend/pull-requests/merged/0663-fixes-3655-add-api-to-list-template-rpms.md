---
type: pull_request
number: 663
title: "Fixes 3655: Add api to list template rpms"
state: merged
author: xbhouse
created: 2024-05-08T19:31:08Z
updated: 2024-05-20T14:01:33Z
closed: 2024-05-20T13:34:34Z
merged: 2024-05-20T13:34:34Z
base_branch: main
head_branch: list-template-rpms
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/663
---

# Pull Request #663: Fixes 3655: Add api to list template rpms

**Author**: @xbhouse
**Created**: May 08, 2024 at 07:31 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `list-template-rpms`

## Description

## Summary

- Adds API for listing content template rpms: `/templates/:uuid/rpms` (similar to the list rpms endpoints for repositories and snapshots)
- Depends on [this](https://github.com/content-services/content-sources-backend/pull/659), will rebase once this is merged

## Testing steps

- Add a custom repository and a RH repository and let them snapshot (this will snapshot a small RH repo: `OPTIONS_REPOSITORY_IMPORT_FILTER=small go run cmd/external-repos/main.go import && go run cmd/external-repos/main.go nightly-jobs`)
- Create a template with those repositories, selecting the latest date 
- Test the `/templates/:uuid/rpms` endpoint with the template uuid and you should see the template's rpms (should include rpms from both of the snapshots in the template) in a response similar to this:
```
{
    "data": [
        {
            "name": "ansible",
            "arch": "noarch",
            "version": "2.8.0",
            "release": "1.el8ae",
            "epoch": "0",
            "summary": "SSH-based configuration management, deployment, and task execution system"
        },
        ...
```        

- To test that the correct snapshot's rpms are shown, you can update the custom repository with a different URL and let it snapshot
- Edit the template date to include this latest snapshot and view the rpms again. The rpms list should reflect the latest snapshot's rpms
- Edit the template to remove the custom repository and then view the rpms again, you should only see the rpms from the RH repository's snapshot

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on May 08, 2024 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-3655

### Comment by @jlsherrill on May 15, 2024 at 02:40 PM UTC

oh and this needs a rebase now

### Comment by @swadeley on May 16, 2024 at 06:52 PM UTC

@jlsherrill this looks like it is ready for testing, if so, please move the Jira card to set the flag.

Thank you

### Comment by @swadeley on May 20, 2024 at 12:46 PM UTC

/retest

### Comment by @swadeley on May 20, 2024 at 01:16 PM UTC

Hi

here we can see three fake packages from the custom repo and the ansible package from the ansible RHEL repo:

```
In [4]: app.content_sources.rest_client.templates_api.list_templates()
2024-05-20 14:12:17.937 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<snip>, params=[]
Out[4]: 
{'data': [{'arch': 'x86_64',
           'date': '2024-05-19T23:00:00Z',
           'description': 'pr-663',
           'name': 'test-template',
           'org_id': '3340851',
           'repository_uuids': ['19756213-a479-429e-adc6-8334a4a7a077',
                                '7b1386dd-06e6-47dd-a1ff-7e73967dd9ee'],
           'uuid': '6a7c931a-be08-4645-89d7-2793e12dd574',
           'version': '8'}],
 'links': {'first': '/api/content-sources/v1/templates/?limit=100&offset=0',
           'last': '/api/content-sources/v1/templates/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

In [6]: app.content_sources.rest_client.templates_api.list_template_rpms('6a7c931a-be08-4645-89d7-2793e12dd574',limit=4)
2024-05-20 14:13:36.891 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<snip>, params=[('limit', 4)]
Out[6]: 
{'data': [{'arch': 'noarch',
           'epoch': '0',
           'name': 'acme-package',
           'release': '1.elfake',
           'summary': 'acme-package package',
           'version': '1.0.1'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'acme-package',
           'release': '1.elfake',
           'summary': 'acme-package package',
           'version': '1.0.2'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'acme-package',
           'release': '1.elfake',
           'summary': 'acme-package package',
           'version': '1.1.2'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'ansible',
           'release': '1.el8ae',
           'summary': 'SSH-based configuration management, deployment, and '
                      'task execution system',
           'version': '2.8.0'}],
 'links': {'first': '/api/content-sources/v1/templates/6a7c931a-be08-4645-89d7-2793e12dd574/rpms?limit=4&offset=0',
           'last': '/api/content-sources/v1/templates/6a7c931a-be08-4645-89d7-2793e12dd574/rpms?limit=4&offset=64',
           'next': '/api/content-sources/v1/templates/6a7c931a-be08-4645-89d7-2793e12dd574/rpms?limit=4&offset=4'},
 'meta': {'count': 65, 'limit': 4, 'offset': 0}}

In [7]: 
```


### Comment by @swadeley on May 20, 2024 at 01:23 PM UTC

Hi

I changed the custom repo:

```
In [7]: app.content_sources.rest_client.templates_api.list_templates()
2024-05-20 14:19:28.943 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<snip>, params=[]
Out[7]: 
{'data': [{'arch': 'x86_64',
           'date': '2024-05-19T23:00:00Z',
           'description': 'pr-663 edited',
           'name': 'test-template',
           'org_id': '3340851',
           'repository_uuids': ['19756213-a479-429e-adc6-8334a4a7a077',
                                '461c1b4c-89b1-49aa-abcb-feb391007a1d'],
           'uuid': '6a7c931a-be08-4645-89d7-2793e12dd574',
           'version': '8'}],
 'links': {'first': '/api/content-sources/v1/templates/?limit=100&offset=0',
           'last': '/api/content-sources/v1/templates/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}
```

At the end of the list we see the one new fake package:

```
In [11]: app.content_sources.rest_client.templates_api.list_template_rpms('6a7c931a-be08-4645-89d7-2793e12dd574',limit=4,offset=56)
2024-05-20 14:21:28.771 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<snip>, params=[('limit', 4), ('offset', 56)]
Out[11]: 
{'data': [{'arch': 'noarch',
           'epoch': '0',
           'name': 'ansible-test',
           'release': '1.el8ae',
           'summary': 'Tool for testing ansible plugin and module code',
           'version': '2.9.9'},
          {'arch': 'noarch',
           'epoch': '0',
           'name': 'cockateel',
           'release': '1',
           'summary': 'A dummy package of cockateel',
           'version': '3.1'},
          {'arch': 'x86_64',
           'epoch': '0',
           'name': 'sshpass',
           'release': '3.el8ae',
           'summary': 'Non-interactive SSH authentication utility',
           'version': '1.06'}],
 'links': {'first': '/api/content-sources/v1/templates/6a7c931a-be08-4645-89d7-2793e12dd574/rpms?limit=4&offset=0',
           'last': '/api/content-sources/v1/templates/6a7c931a-be08-4645-89d7-2793e12dd574/rpms?limit=4&offset=56',
           'prev': '/api/content-sources/v1/templates/6a7c931a-be08-4645-89d7-2793e12dd574/rpms?limit=4&offset=52'},
 'meta': {'count': 59, 'limit': 4, 'offset': 56}}

In [12]: 
```

### Comment by @swadeley on May 20, 2024 at 01:34 PM UTC

Hi

removed the custom repo (notice only one uuid):
```
In [12]: app.content_sources.rest_client.templates_api.list_templates()
2024-05-20 14:31:59.599 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=<snip>, params=[]
Out[12]: 
{'data': [{'arch': 'x86_64',
           'date': '2024-05-19T23:00:00Z',
           'description': 'pr-663 edited again',
           'name': 'test-template',
           'org_id': '3340851',
           'repository_uuids': ['19756213-a479-429e-adc6-8334a4a7a077'],
           'uuid': '6a7c931a-be08-4645-89d7-2793e12dd574',
           'version': '8'}],
 'links': {'first': '/api/content-sources/v1/templates/?limit=100&offset=0',
           'last': '/api/content-sources/v1/templates/?limit=100&offset=0'},
 'meta': {'count': 1, 'limit': 100, 'offset': 0}}

In [13]: app.content_sources.rest_client.templates_api.list_template_rpms('6a7c931a-be08-4645-89d7-2793e12dd574',limit=4,offset=56)
2024-05-20 14:32:06.422 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=e<snip>, params=[('limit', 4), ('offset', 56)]
Out[13]: 
{'data': [{'arch': 'noarch',
           'epoch': '0',
           'name': 'ansible-test',
           'release': '1.el8ae',
           'summary': 'Tool for testing ansible plugin and module code',
           'version': '2.9.9'},
          {'arch': 'x86_64',
           'epoch': '0',
           'name': 'sshpass',
           'release': '3.el8ae',
           'summary': 'Non-interactive SSH authentication utility',
           'version': '1.06'}],
 'links': {'first': '/api/content-sources/v1/templates/6a7c931a-be08-4645-89d7-2793e12dd574/rpms?limit=4&offset=0',
           'last': '/api/content-sources/v1/templates/6a7c931a-be08-4645-89d7-2793e12dd574/rpms?limit=4&offset=56',
           'prev': '/api/content-sources/v1/templates/6a7c931a-be08-4645-89d7-2793e12dd574/rpms?limit=4&offset=52'},
 'meta': {'count': 58, 'limit': 4, 'offset': 56}}

In [14]: 
```

Now the 'A dummy package of cockateel' is gone.

---

## Reviews

### Review by @jlsherrill - Commented on May 13, 2024 at 06:51 PM UTC

### Review by @jlsherrill - Commented on May 13, 2024 at 06:52 PM UTC

### Review by @xbhouse - Commented on May 14, 2024 at 03:07 PM UTC

### Review by @jlsherrill - Commented on May 15, 2024 at 02:40 PM UTC

### Review by @jlsherrill - Approved on May 16, 2024 at 07:52 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/663*

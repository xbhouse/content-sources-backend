---
type: pull_request
number: 911
title: "Fixes 5023: api to fetch template repo configs"
state: merged
author: xbhouse
created: 2024-12-02T17:40:46Z
updated: 2024-12-13T16:30:34Z
closed: 2024-12-13T16:14:13Z
merged: 2024-12-13T16:14:13Z
base_branch: main
head_branch: 5023
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/911
---

# Pull Request #911: Fixes 5023: api to fetch template repo configs

**Author**: @xbhouse
**Created**: December 02, 2024 at 05:40 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `5023`

## Description

## Summary

Adds an endpoint to fetch the repo configs for a template

## Testing steps

1. Create a couple repos and let them snapshot
2. Create a template with those repos
3. Fetch the repo config files in the template via this endpoint: `GET /templates/:template_uuid/config.repo`
4. You should see a response similar to this:

```
[test1]
name=test1
baseurl=http://pulp.content:8081/pulp/content/cs-8997773847/templates/1708416b-d0c6-429a-88ba-2c3201c12f66/4a13e189-4b4f-46e8-af93-59f7e32b28ad
module_hotfixes=0
gpgcheck=1
repo_gpgcheck=0
enabled=1
gpgkey=http://pulp.content:8000/api/content-sources/v1.0/repository_gpg_key/4a13e189-4b4f-46e8-af93-59f7e32b28ad
sslclientcert=/etc/pki/consumer/cert.pem
sslclientkey=/etc/pki/consumer/key.pem

[ansible-2-for-rhel-8-x86_64-rpms]
name=Red Hat Ansible Engine 2 for RHEL 8 x86_64 (RPMs)
baseurl=http://pulp.content:8081/pulp/content/cs-redhat/templates/1708416b-d0c6-429a-88ba-2c3201c12f66/content/dist/layered/rhel8/x86_64/ansible/2/os/
module_hotfixes=0
gpgcheck=1
repo_gpgcheck=0
enabled=1
gpgkey=http://pulp.content:8000/api/content-sources/v1.0/repository_gpg_key/3b487063-79ec-4226-b91f-4982a6192ec5
sslclientcert=/etc/pki/consumer/cert.pem
sslclientkey=/etc/pki/consumer/key.pem
```

5. Verify the distribution paths are correct
6. Verify that a client and normal users can access the endpoint. You can set your identity header to be of type System to test the client auth locally: `{"identity":{"type":"System","system":{},"account_number":"11491786","internal":{"org_id":"17684632"}}}`

---

## Discussion

### Comment by @jlsherrill on December 02, 2024 at 06:00 PM UTC

https://issues.redhat.com/browse/HMS-5023

### Comment by @swadeley on December 03, 2024 at 12:46 PM UTC

Hi @xbhouse 

When I try to apply your APIdoc changes it is causing "searchUploads" section in APIdocs to be deleted. I did try a rebase on your checked out branch but that did not help. Please rebase for me. Thank you.

### Comment by @xbhouse on December 03, 2024 at 03:12 PM UTC

hi @swadeley, i just rebased but i think the [searchUploads](https://github.com/content-services/content-sources-backend/pull/909) endpoint isn't in main yet. i'll rebase again once that's merged 

EDIT: that PR was closed, so i think we won't actually have a searchUploads endpoint. the createUpload endpoint will instead be modified

### Comment by @swadeley on December 06, 2024 at 07:04 AM UTC

Hi

API doc works:
```
n [13]:  app.content_sources.rest_client.templates_api.get_template('a3caa32e-48b5-4a0f-841f-7db19c48e801')
2024-12-06 15:02:49.876 [    INFO] [iqe.base.rest_client] REST: GET https://ee-8osedbe9.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/a3caa32e-48b5-4a0f-841f-7db19c48e801 with query params [] and x-rh-insights-request-id=<>
Out[13]: 
{'arch': 'x86_64',
 'created_at': '2024-12-06T07:01:11.887626Z',
 'created_by': 'ephemeral-user',
 'date': '0001-01-01T00:00:00Z',
 'description': 'pr911',
 'last_update_snapshot_error': '',
 'last_updated_by': 'ephemeral-user',
 'name': 'test-template',
 'org_id': '3340851',
 'repository_uuids': ['1de2f582-8420-4258-8089-7b4db3cba68a',
                      '6d668de9-f09d-4fed-8e35-f917c4fb2aaf',
                      '885bcac9-c693-4843-9498-7c641731ab12',
                      'd6d32fb8-0361-4d60-9e38-ea2ebe4ebeda'],
 'rhsm_environment_created': False,
 'rhsm_environment_id': 'a3caa32e48b54a0f841f7db19c48e801',
 'snapshots': [{'added_counts': {'rpm.advisory': 32,
                                 'rpm.package': 58,
                                 'rpm.repo_metadata_file': 1},
                'content_counts': {'rpm.advisory': 32,
                                   'rpm.package': 58,
                                   'rpm.repo_metadata_file': 1},
                'created_at': '2024-12-06T05:16:54.634121Z',
                'removed_counts': {},
                'repository_name': 'Red Hat Ansible Engine 2 for RHEL 8 x86_64 '
                                   '(RPMs)',
                'repository_path': 'cs-redhat/1de2f582-8420-4258-8089-7b4db3cba68a/b9cab487-1089-45cc-a388-23ab3947ed8d',
                'repository_uuid': '1de2f582-8420-4258-8089-7b4db3cba68a',
                'url': 'http://pulp-content:8000/api/pulp-content/cs-redhat/1de2f582-8420-4258-8089-7b4db3cba68a/b9cab487-1089-45cc-a388-23ab3947ed8d/',
                'uuid': '76cdd656-321f-4156-965f-4320f0323383'},
               {'added_counts': {'rpm.package': 1},
                'content_counts': {'rpm.package': 1},
                'created_at': '2024-12-06T06:59:24.407318Z',
                'removed_counts': {},
                'repository_name': 'Bulk3-UrudSdjB',
                'repository_path': 'cs-da3fff3c6c/6d668de9-f09d-4fed-8e35-f917c4fb2aaf/5f8c71af-4ab1-4cfc-9304-7345f6830386',
                'repository_uuid': '6d668de9-f09d-4fed-8e35-f917c4fb2aaf',
                'url': 'http://pulp-content:8000/api/pulp-content/cs-da3fff3c6c/6d668de9-f09d-4fed-8e35-f917c4fb2aaf/5f8c71af-4ab1-4cfc-9304-7345f6830386/',
                'uuid': '7c357aaf-f919-4f31-acc2-6e1d9678f892'},
               {'added_counts': {'rpm.package': 1},
                'content_counts': {'rpm.package': 1},
                'created_at': '2024-12-06T06:59:22.45057Z',
                'removed_counts': {},
                'repository_name': 'Bulk1-NxjIIeGp',
                'repository_path': 'cs-da3fff3c6c/885bcac9-c693-4843-9498-7c641731ab12/1dcb6c00-4ab4-44cd-8c00-d78115e389bd',
                'repository_uuid': '885bcac9-c693-4843-9498-7c641731ab12',
                'url': 'http://pulp-content:8000/api/pulp-content/cs-da3fff3c6c/885bcac9-c693-4843-9498-7c641731ab12/1dcb6c00-4ab4-44cd-8c00-d78115e389bd/',
                'uuid': '6a9e5e8f-d4c4-4854-b2f3-1ff2ec60383b'},
               {'added_counts': {'rpm.package': 1},
                'content_counts': {'rpm.package': 1},
                'created_at': '2024-12-06T06:59:50.387938Z',
                'removed_counts': {},
                'repository_name': 'Bulk2-DlAaqjTa',
                'repository_path': 'cs-da3fff3c6c/d6d32fb8-0361-4d60-9e38-ea2ebe4ebeda/4e9e7b03-f72c-4ae5-849f-2d53dc330bcf',
                'repository_uuid': 'd6d32fb8-0361-4d60-9e38-ea2ebe4ebeda',
                'url': 'http://pulp-content:8000/api/pulp-content/cs-da3fff3c6c/d6d32fb8-0361-4d60-9e38-ea2ebe4ebeda/4e9e7b03-f72c-4ae5-849f-2d53dc330bcf/',
                'uuid': 'd03635c5-c3fc-4180-bdc8-f7850e3cfd5f'}],
 'updated_at': '2024-12-06T07:01:11.887626Z',
 'use_latest': True,
 'uuid': 'a3caa32e-48b5-4a0f-841f-7db19c48e801',
 'version': '8'}

In [14]: 

```

### Comment by @swadeley on December 06, 2024 at 07:21 AM UTC

From inside a POD:
```
sh-5.1$ curl -L http://pulp-content:8000/api/pulp-content/cs-da3fff3c6c/d6d32fb8-0361-4d60-9e38-ea2ebe4ebeda/4e9e7b03-f72c-4ae5-849f-2d53dc330bcf/Packa/Packages/c/

<html>
<head><title>Index of /api/pulp-content/cs-da3fff3c6c/d6d32fb8-0361-4d60-9e38-ea2ebe4ebeda/4e9e7b03-f72c-4ae5-849f-2d53dc330bcf/Packages/c/</title></head>
<body bgcolor="white">
<h1>Index of /api/pulp-content/cs-da3fff3c6c/d6d32fb8-0361-4d60-9e38-ea2ebe4ebeda/4e9e7b03-f72c-4ae5-849f-2d53dc330bcf/Packages/c/</h1>
<hr><pre><a href="../">../</a>
<a href="cockateel-3.1-1.noarch.rpm">cockateel-3.1-1.noarch.rpm</a>                                                                          06-Dec-2024 06:59  2.5 kB
</pre><hr></body>
</html>sh-5.1$ 
```

### Comment by @swadeley on December 10, 2024 at 03:37 AM UTC

/retest

### Comment by @xbhouse on December 10, 2024 at 08:22 PM UTC

> This isn't caused by your PR, but we might fix this here real quick 😅
> 
> One issue I'm seeing is if I make a request as a system to an endpoint that doesn't skip RBAC, there is a panic
> 
> I think we need to check [here](https://github.com/content-services/content-sources-backend/blob/main/pkg/cache/redis.go#L70) if user is not nil before trying to set the value in the cache. We should return an error if no user is set.

ah right i meant to fix this, will push that up shortly!

### Comment by @swadeley on December 11, 2024 at 03:18 AM UTC

Hi @xbhouse 

As you have not made any API docs changes since I last tested my tests are still valid[1], but I just want to ask about "Verify the distribution paths are correct". 

When I GET the template, there is a list of snapshots, and each snapshot has `repository_path` which is used to compile the `url` I tested [2].

Is that the correct path or do I need to check some other path?

Thank you

[1] https://github.com/content-services/content-sources-backend/pull/911#issuecomment-2522283515
[2] https://github.com/content-services/content-sources-backend/pull/911#issuecomment-2522306026

### Comment by @xbhouse on December 11, 2024 at 02:16 PM UTC

> Hi @xbhouse
> 
> As you have not made any API docs changes since I last tested my tests are still valid[1], but I just want to ask about "Verify the distribution paths are correct".
> 
> When I GET the template, there is a list of snapshots, and each snapshot has `repository_path` which is used to compile the `url` I tested [2].
> 
> Is that the correct path or do I need to check some other path?
> 
> Thank you
> 
> [1] [#911 (comment)](https://github.com/content-services/content-sources-backend/pull/911#issuecomment-2522283515) [2] [#911 (comment)](https://github.com/content-services/content-sources-backend/pull/911#issuecomment-2522306026)

hi @swadeley sorry for the delay!

you could check that the content in each `baseurl` returned in the response from `/templates/:template_uuid/config.repo` shows the same content as the snapshot url (the url you're checking [here](https://github.com/content-services/content-sources-backend/pull/911#issuecomment-2522306026)) from the response when fetching a template:

```
curl http://pulp.content:8081/pulp/content/cs-b3b69ddd79/templates/51bc58ba-6951-48dc-a6a1-ff1f2a235013/33ca3979-5d17-46a0-80af-4b906a0c4b4e/Packages/

<html>
<head><title>Index of /pulp/content/cs-b3b69ddd79/templates/51bc58ba-6951-48dc-a6a1-ff1f2a235013/33ca3979-5d17-46a0-80af-4b906a0c4b4e/Packages/</title></head>
<body bgcolor="white">
<h1>Index of /pulp/content/cs-b3b69ddd79/templates/51bc58ba-6951-48dc-a6a1-ff1f2a235013/33ca3979-5d17-46a0-80af-4b906a0c4b4e/Packages/</h1>
<hr><pre><a href="../">../</a>
<a href="c/">c/</a>                                                                                                  11-Dec-2024 14:06  2.2 kB
<a href="e/">e/</a>                                                                                                  11-Dec-2024 14:06  2.2 kB
<a href="g/">g/</a>                                                                                                  11-Dec-2024 14:06  2.2 kB
<a href="l/">l/</a>                                                                                                  11-Dec-2024 14:06  2.2 kB
<a href="m/">m/</a>                                                                                                  11-Dec-2024 14:06  2.2 kB
<a href="p/">p/</a>                                                                                                  11-Dec-2024 14:06  2.2 kB
<a href="s/">s/</a>                                                                                                  11-Dec-2024 14:06  2.2 kB
<a href="w/">w/</a>                                                                                                  11-Dec-2024 14:06  2.2 kB
</pre><hr></body>
</html>%                                                                                                                                                                 

curl http://pulp.content:8081/pulp/content/cs-b3b69ddd79/33ca3979-5d17-46a0-80af-4b906a0c4b4e/e8172fa3-db7f-4e23-a829-d1a9664409bb/Packages/

<html>
<head><title>Index of /pulp/content/cs-b3b69ddd79/33ca3979-5d17-46a0-80af-4b906a0c4b4e/e8172fa3-db7f-4e23-a829-d1a9664409bb/Packages/</title></head>
<body bgcolor="white">
<h1>Index of /pulp/content/cs-b3b69ddd79/33ca3979-5d17-46a0-80af-4b906a0c4b4e/e8172fa3-db7f-4e23-a829-d1a9664409bb/Packages/</h1>
<hr><pre><a href="../">../</a>
<a href="c/">c/</a>                                                                                                  11-Dec-2024 14:06  2.2 kB
<a href="e/">e/</a>                                                                                                  11-Dec-2024 14:06  2.2 kB
<a href="g/">g/</a>                                                                                                  11-Dec-2024 14:06  2.2 kB
<a href="l/">l/</a>                                                                                                  11-Dec-2024 14:06  2.2 kB
<a href="m/">m/</a>                                                                                                  11-Dec-2024 14:06  2.2 kB
<a href="p/">p/</a>                                                                                                  11-Dec-2024 14:06  2.2 kB
<a href="s/">s/</a>                                                                                                  11-Dec-2024 14:06  2.2 kB
<a href="w/">w/</a>                                                                                                  11-Dec-2024 14:06  2.2 kB
</pre><hr></body>
</html>%    
```

### Comment by @swadeley on December 12, 2024 at 02:30 PM UTC

Thank you @jlsherrill 

Now I found I can get repo config like so:
`In [9]: print(app.content_sources.rest_client.repositories_api.get_template_repo_configuration_files('756bf617-12eb-4343-be3e-b4a68713ab49'))
`
get the base URL and then
`
sh-5.1$ curl -L "http://pulp-content:8000/api/pulp-content/cs-563d8adcab/templates/756bf617-12eb-4343-be3e-b4a68713ab49/4dbf0e46-28a0-4abc-8710-d1db31082338"
404: Not Found
`



### Comment by @swadeley on December 13, 2024 at 03:39 PM UTC

Hi

new path works:
`In [2]: app.content_sources.rest_client.templates_api.get_template_repo_configuration_file`


```

In [4]:             template_dict = dict(
   ...:                 name="template-test",
   ...:                 repository_uuids=['ba38949c-9901-4b4e-93b7-0d5e8763779f'],
   ...:                 description="Test template",
   ...:                 arch="x86_64",
   ...:                 version="8",
   ...:                 use_latest=True,
   ...:             )

In [5]: app.content_sources.rest_client.templates_api.create_template(template_dict)
2024-12-13 23:32:56.922 [    INFO] [iqe.base.rest_client] REST: POST https://ee-imqrt7wl.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/ with query params [] and x-rh-insights-request-id=<-56ae-420d-9f0d->
Out[5]: 
{'arch': 'x86_64',
 'created_at': '2024-12-13T15:32:56.813505136Z',
 'created_by': 'ephemeral-user',
 'date': '0001-01-01T00:00:00Z',
 'description': 'Test template',
 'last_update_snapshot_error': '',
 'last_updated_by': 'ephemeral-user',
 'name': 'template-test',
 'org_id': '3340851',
 'repository_uuids': ['ba38949c-9901-4b4e-93b7-0d5e8763779f'],
 'rhsm_environment_created': False,
 'rhsm_environment_id': '484bf6190e094688af048af45e95918c',
 'updated_at': '2024-12-13T15:32:56.813505136Z',
 'use_latest': True,
 'uuid': '484bf619-0e09-4688-af04-8af45e95918c',
 'version': '8'}



In [7]: print(app.content_sources.rest_client.templates_api.get_template_repo_configuration_file('484bf619-0e09-4688-af04-8af45e95918c'))
2024-12-13 23:34:40.253 [    INFO] [iqe.base.rest_client] REST: GET https://ee-imqrt7wl.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1/templates/484bf619-0e09-4688-af04-8af45e95918c/config.repo with query params [] and x-rh-insights-request-id=37744096-5e8e-47ca-9109-795bd7134fa6
[ansible-2-for-rhel-8-x86_64-rpms]
name=Red Hat Ansible Engine 2 for RHEL 8 x86_64 (RPMs)
baseurl=http://pulp-content:8000/api/pulp-content/cs-redhat/templates/484bf619-0e09-4688-af04-8af45e95918c/content/dist/layered/rhel8/x86_64/ansible/2/os/
module_hotfixes=0
gpgcheck=1
repo_gpgcheck=0
enabled=1
gpgkey=https://env-ephemeral-jhdtpy.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1.0/repository_gpg_key/ba38949c-9901-4b4e-93b7-0d5e8763779f
sslclientcert=/etc/pki/consumer/cert.pem
sslclientkey=/etc/pki/consumer/key.pem

In [8]: 

```
`sh-5.1$ curl -L   "http://pulp-content:8000/api/pulp-content/cs-redhat/templates/484bf619-0e09-4688-af04-8af45e95918c/content/dist/layered/rhel8/"
404: Not Foundsh-5.1$ 
`


---

## Reviews

### Review by @rverdile - Commented on December 03, 2024 at 09:15 PM UTC

### Review by @rverdile - Commented on December 03, 2024 at 09:17 PM UTC

### Review by @xbhouse - Commented on December 03, 2024 at 09:27 PM UTC

### Review by @xbhouse - Commented on December 03, 2024 at 09:29 PM UTC

### Review by @rverdile - Commented on December 04, 2024 at 09:04 PM UTC

### Review by @rverdile - Commented on December 04, 2024 at 09:11 PM UTC

### Review by @xbhouse - Commented on December 05, 2024 at 01:36 PM UTC

### Review by @rverdile - Commented on December 05, 2024 at 08:10 PM UTC

### Review by @rverdile - Commented on December 05, 2024 at 09:26 PM UTC

One other thing here. When using the identity with the system type, we need to skip rbac for this API

If you turn on mock_rbac, and then test with the system type, you can see it error. In this case it will panic, but it shouldn't do that either.

### Review by @rverdile - Commented on December 09, 2024 at 05:03 PM UTC

### Review by @rverdile - Commented on December 09, 2024 at 05:03 PM UTC

### Review by @xbhouse - Commented on December 10, 2024 at 03:25 PM UTC

### Review by @xbhouse - Commented on December 10, 2024 at 03:45 PM UTC

### Review by @xbhouse - Commented on December 10, 2024 at 04:05 PM UTC

### Review by @rverdile - Commented on December 10, 2024 at 06:56 PM UTC

### Review by @rverdile - Commented on December 10, 2024 at 07:00 PM UTC

This isn't caused by your PR, but we might fix this here real quick :sweat_smile: 

One issue I'm seeing is if I make a request as a system to an endpoint that doesn't skip RBAC, there is a panic

I think we need to check [here](https://github.com/content-services/content-sources-backend/blob/main/pkg/cache/redis.go#L70) if user is not nil before trying to set the value in the cache. We should return an error if no user is set.



### Review by @xbhouse - Commented on December 10, 2024 at 08:32 PM UTC

### Review by @rverdile - Commented on December 10, 2024 at 09:03 PM UTC

### Review by @rverdile - Approved on December 10, 2024 at 09:07 PM UTC

looks good!

### Review by @jlsherrill - Commented on December 12, 2024 at 02:12 PM UTC

### Review by @jlsherrill - Commented on December 12, 2024 at 02:14 PM UTC

### Review by @xbhouse - Commented on December 13, 2024 at 02:26 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/911*

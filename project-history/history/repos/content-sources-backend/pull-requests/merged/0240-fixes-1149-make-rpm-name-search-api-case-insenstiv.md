---
type: pull_request
number: 240
title: "Fixes 1149: make rpm name search api case insenstive"
state: merged
author: jlsherrill
created: 2023-03-28T20:34:46Z
updated: 2024-04-17T13:10:44Z
closed: 2023-03-29T12:38:22Z
merged: 2023-03-29T12:38:22Z
base_branch: main
head_branch: 1149
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/240
---

# Pull Request #240: Fixes 1149: make rpm name search api case insenstive

**Author**: @jlsherrill
**Created**: March 28, 2023 at 08:34 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `1149`

## Description

## Summary
IB is requesting that rpm name search be case insensitive 
## Testing steps

Run these requests:
```
### Create a single repository
POST http://{{host}}/api/content-sources/v1.0/repositories/
Content-Type: application/json
x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiQXNzb2NpYXRlIiwiYWNjb3VudF9udW1iZXIiOiIxIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMSJ9fX0K

{
  "name": "needed2",
  "url": "https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/",
  "distribution_versions": [
    "8"
  ],
  "distribution_arch": "x86_64"
}


### Rpm search
POST  http://{{host}}/api/content-sources/v1.0/rpms/names
Content-Type: application/json
x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiQXNzb2NpYXRlIiwiYWNjb3VudF9udW1iZXIiOiIxIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMSJ9fX0K

{
  "urls": [
    "https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/"
  ],
  "search": "Walrus"
}
```

you should get results for 'walrus'

---

## Discussion

### Comment by @jlsherrill on March 28, 2023 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-1149

### Comment by @swadeley on March 29, 2023 at 08:52 AM UTC

Hi, this is not working for me (I tested in webUI).

I used this tag when I deployed:
`--set-image-tag quay.io/cloudservices/content-sources-backend=pr-240-latest`

### Comment by @swadeley on March 29, 2023 at 12:38 PM UTC

Hi, OK, with API call you posted it does work:
`$ curl --insecure  -X POST  -u <admin_name>:<redacted> https://env-ephemeral-<redacted>.openshiftapps.com/api/content-sources/v1.0/rpms/names -H 'Content-Type: application/json' -d '{
  "urls": [
    "https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/"
  ],
  "search": "Walrus"
}'
[{"package_name":"walrus","summary":"A dummy package of walrus"}]``

### Comment by @swadeley on April 17, 2024 at 01:10 PM UTC

For IQE:
```
In [6]:  app.content_sources.rest_client.repositories_api.search_rpm(
   ...:         dict(
   ...:             search="bear",
   ...:             urls=[
   ...:                 'https://jlsherrill.fedorapeople.org/fake-repos/revision/two/',
   ...:             ],
   ...:         )
   ...:     )
   ...: 
   ...: 
2024-04-17 14:05:29.539 [    INFO] [iqe.base.rest_client] REST: METHOD=POST, request_id=<>, params=[]
Out[6]: [{'package_name': 'bear', 'summary': 'A dummy package of bear'}]
```


---

## Reviews

### Review by @Andrewgdewar - Approved on March 28, 2023 at 11:07 PM UTC

![and my ack!](https://user-images.githubusercontent.com/38083295/228386703-d3b1a4c1-9382-44cc-9e87-831b8e024ea1.gif)


---

*Archived from: https://github.com/content-services/content-sources-backend/pull/240*

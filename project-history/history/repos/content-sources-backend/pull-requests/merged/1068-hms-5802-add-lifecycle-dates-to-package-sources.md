---
type: pull_request
number: 1068
title: "HMS-5802: add lifecycle dates to package sources"
state: merged
author: rverdile
created: 2025-04-08T19:08:54Z
updated: 2025-07-07T20:09:35Z
closed: 2025-04-14T15:03:43Z
merged: 2025-04-14T15:03:43Z
base_branch: main
head_branch: modules-end-dates
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1068
---

# Pull Request #1068: HMS-5802: add lifecycle dates to package sources

**Author**: @rverdile
**Created**: April 08, 2025 at 07:08 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `modules-end-dates`

## Description

## Summary
Adds the lifecycle dates from the roadmap API to the package sources in RPM search.

Fixes issue where we were returning older versions of module streams in package sources. Only the latest version is needed.

## Testing steps
1. run `make repos-import` and `go run cmd/external-repos/main.go process-repos` to make sure you have the RHEL 8 and RHEL 9 appstreams repos snapshotted
2. In your config.yaml, set the new fields under client
```
  roadmap:
    server: https://console.stage.redhat.com/api/roadmap/v1
    username:
    password:
    proxy:
```
4. Search RPMs (/rpms/names) and RPM snapshots (/snapshots/rpms/names)
```
RPMs:
{
  "search": "maven",
  "urls": ["https://cdn.redhat.com/content/dist/rhel9/9/x86_64/appstream/os/"],
  "include_package_sources": true
}

Snapshot:
{
  "search": "maven",
  "uuids": ["snapshot-uuid"],
  "include_package_sources": true
}

```
5. The response should include the "start_date" and "end_date" of the module streams listed in package_sources




---

## Discussion

### Comment by @jlsherrill on April 08, 2025 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-5802

### Comment by @swadeley on April 09, 2025 at 01:26 PM UTC

Hi

I see the start and end date keys.

```
In [1]: app.content_sources.rest_client.rpms_api.search_rpm(dict(urls=["https://rverdile.fedorapeople.org/dummy-repos/modules/repo1/"], search='kangaroo',  include_package_sources=True))
<snip>
Out[1]: 
[{'package_name': 'kangaroo',
  'package_sources': [{'arch': 'noarch',
                       'context': 'deadbeef',
                       'description': 'A module for the kangaroo 0.3 package',
                       'end_date': '',
                       'name': 'kangaroo',
                       'start_date': '',
                       'stream': '0',
                       'type': 'module',
                       'version': '20180730223407'}],
  'summary': 'hop like a kangaroo in Australia'}]
```

```
In [2]: app.content_sources.rest_client.rpms_api.search_snapshot_rpms(dict(uuids=['be907504-4caa-4ae3-8276-1dd0c52b2f6e'], search='kangaroo',  include_package_sources=True))
<snip>
Out[2]: 
[{'package_name': 'kangaroo',
  'package_sources': [{'arch': 'noarch',
                       'context': 'deadbeef',
                       'description': 'A module for the kangaroo 0.3 package',
                       'end_date': '',
                       'name': 'kangaroo',
                       'start_date': '',
                       'stream': '0',
                       'type': 'module',
                       'version': '20180730223407'}],
  'summary': 'hop like a kangaroo in Australia'}]

In [6]: 

```

### Comment by @rverdile on April 09, 2025 at 05:07 PM UTC

It's only going to be populated for RHEL 8 and RHEL 9 because that is the scope of the roadmap api. I guess I can omit them if they're empty

### Comment by @xbhouse on April 09, 2025 at 05:22 PM UTC

> It's only going to be populated for RHEL 8 and RHEL 9 because that is the scope of the roadmap api. I guess I can omit them if they're empty

hrmmm that's what i was testing with and they're empty 😕 

```
{
  "search": "maven",
  "urls": ["https://cdn.redhat.com/content/dist/rhel9/9/x86_64/appstream/os/"],
  "include_package_sources": true
}

[
    {
        "package_name": "maven",
        "summary": "Java project management and project comprehension tool",
        "package_sources": [
            {
                "type": "module",
                "name": "maven",
                "stream": "3.8",
                "context": "470dcefd",
                "arch": "x86_64",
                "version": "9040020240210002822",
                "description": "Maven is a software project management and comprehension tool. Based on the concept of a project object model (POM), Maven can manage a project's build, reporting and documentation from a central piece of information.",
                "start_date": "",
                "end_date": ""
            }
        ]
    },
    {
        "package_name": "maven-lib",
        "summary": "Core part of Maven",
        "package_sources": [
            {
                "type": "module",
                "name": "maven",
                "stream": "3.8",
                "context": "470dcefd",
                "arch": "x86_64",
                "version": "9040020240210002822",
                "description": "Maven is a software project management and comprehension tool. Based on the concept of a project object model (POM), Maven can manage a project's build, reporting and documentation from a central piece of information.",
                "start_date": "",
                "end_date": ""
            }
        ]
    },
    ...
]
```    

### Comment by @rverdile on April 09, 2025 at 06:14 PM UTC

> hrmmm that's what i was testing with and they're empty 😕

It seems the output from the roadmap API has changed significantly since yesterday :) so I'll have to go back and fix this. Going to put back into draft temporarily

### Comment by @rverdile on April 09, 2025 at 07:36 PM UTC

Should be working now

### Comment by @swadeley on April 14, 2025 at 10:07 AM UTC

/retest

### Comment by @swadeley on April 14, 2025 at 03:03 PM UTC

Hi

no more redundant keys with empty values:
```
In [2]: app.content_sources.rest_client.rpms_api.search_snapshot_rpms(dict(uuids=['3b84362c-4a0a-4078-85ef-efe5287fc4ae'], search='kangaroo',  include_package_sources=True))
<snip>
Out[2]: 
[{'package_name': 'kangaroo',
  'package_sources': [{'arch': 'noarch',
                       'context': 'deadbeef',
                       'description': 'A module for the kangaroo 0.3 package',
                       'name': 'kangaroo',
                       'stream': '0',
                       'type': 'module',
                       'version': '20180730223407'}],
  'summary': 'A dummy package of kangaroo'}]
````

Testing RHEL repo will have to be in stage.

Thank you.

---

## Reviews

### Review by @xbhouse - Commented on April 09, 2025 at 03:06 PM UTC

### Review by @xbhouse - Commented on April 09, 2025 at 03:12 PM UTC

### Review by @rverdile - Commented on April 09, 2025 at 04:03 PM UTC

### Review by @xbhouse - Commented on April 09, 2025 at 04:42 PM UTC

similar to Stephen's output, i'm seeing empty strings for start_date and end_date. maybe i'm testing with the wrong repos/rpms, did you see those populated?

### Review by @xbhouse - Commented on April 09, 2025 at 05:24 PM UTC

### Review by @rverdile - Commented on April 09, 2025 at 06:13 PM UTC

### Review by @xbhouse - Approved on April 10, 2025 at 04:57 PM UTC

works well and lgtm! nice job 😃 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1068*

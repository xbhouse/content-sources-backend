---
type: pull_request
number: 461
title: "Fixes 2443: API support for list / search package groups"
state: merged
author: xbhouse
created: 2023-11-07T22:00:59Z
updated: 2024-04-17T13:24:02Z
closed: 2023-12-05T00:52:10Z
merged: 2023-12-05T00:52:10Z
base_branch: main
head_branch: pgroups-api-support
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/461
---

# Pull Request #461: Fixes 2443: API support for list / search package groups

**Author**: @xbhouse
**Created**: November 07, 2023 at 10:00 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `pgroups-api-support`

## Description

## Summary

This adds:

- `repositories/{uuid}/package_groups` endpoint for listing, returns ID, Name, Description, and Package List metadata
- `package_groups/names` endpoint for searching, returns Name, Description, and Package List of all groups found across the repositories requested
- database support to store package group metadata
- groups are de-duplicated based on combination of name/ID/package list
- search API returns full list of packages that are associated with the search term and repositories

## Testing steps

Endpoint: `$BASE_URL/package_groups/names`

Example request: `{
  "urls": [
    "https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/"
  ],
  "search": "xfce"
}`

- add a repository that has a comps.xml file (e.g. https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/) and make a request to `repositories/{uuid}/package_groups`, should list package groups associated with that repo
- make a request to `package_groups/names` with a valid repo URL or repo UUID and a search term, should return the package groups with that search term in the name
** for example, the epel 8 and epel 9 repos have some package groups (one is `xfce`) with different package lists. if you add both of these repos and include both repos in the request, the search API should return the full list of packages from both repos like this:
`[
    {
        "package_group_name": "Xfce",
        "description": "A lightweight desktop environment that works well on low end machines.",
        "package_list": [
            "xfce4-screensaver",
            "tumbler",
            "NetworkManager-gnome",
            "openssh-askpass",
            "gdm",
            "xfce4-terminal",
            "thunar-volman",
            "xfce4-panel",
            "xfce4-appfinder",
            "network-manager-applet",
            "xfwm4",
            "xfce4-settings",
            "xfce4-pulseaudio-plugin",
            "pinentry-gtk",
            "mousepad",
            "xfdesktop",
            "xfconf",
            "Thunar",
            "xfce-polkit",
            "thunar-archive-plugin",
            "xfce4-session",
            "pinentry-gnome3",
            "xfce4-power-manager"
        ]
    }
]`

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on November 09, 2023 at 06:00 PM UTC

https://issues.redhat.com/browse/HMS-2443

### Comment by @xbhouse on November 21, 2023 at 09:27 PM UTC

> I think there may be some places where we can create shared functions or structs between rpm, package groups, and environments. And in the future, snapshot rpms, package groups, and environments. All 6 of these things have a similar flow, so it probably makes sense to try and de-duplicate where possible.
> 
> I left comments where I think might be potential places to create shared code. With that said, it probably make sense to leave it as-is here, and make changes to the environments PR instead. This way you can see all 3 at the same time and easily spot any duplication.

that makes sense! i will make those changes in the environments PR 

### Comment by @swadeley on December 04, 2023 at 12:33 AM UTC

/retest

### Comment by @swadeley on December 04, 2023 at 03:12 AM UTC

Hi, I have a repo:


```
Out[2]: 
{'data': [{'account_id': '0369233',
           'content_type': 'rpm',
           'distribution_arch': 'x86_64',
           'distribution_versions': ['8'],
           'failed_introspections_count': 0,
           'gpg_key': '',
           'last_introspection_error': '',
           'last_introspection_time': '2023-12-04T02:47:22Z',
           'last_success_introspection_time': '2023-12-04T02:47:22Z',
           'last_update_introspection_time': '2023-12-04T01:36:24Z',
           'metadata_verification': False,
           'name': 'epel8',
           'org_id': '3340851',
           'origin': 'external',
           'package_count': 9947,
           'snapshot': False,
           'status': 'Valid',
           'url': 'https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/',
           'uuid': '72e59643-5839-456e-b45e-9968a5676a97'}],
```
 
This works:

`In [14]: app.content_sources.rest_client.repositories_api.list_repositories_package_groups('72e59643-5839
Out[14]: 
{'data': [{'description': '',
           'id': 'base-x',
           'name': 'base-x',
           'packagelist': ['glx-utils'],
           'uuid': '0012491e-866f-4658-8787-e027c265f48a'},
          {'description': '',
           'id': 'core',
           'name': 'core',
           'packagelist': ['bash'],
           'uuid': '5e5106fa-55dc-44d5-bc08-0e1f45db6fe3'},
          {'description': 'A set of packages that provide the Critical Path '
                          'functionality for the KDE desktop',
           'id': 'critical-path-kde',
           'name': 'Critical Path (KDE)',
           'packagelist': ['kdelibs', 'sddm'],
           'uuid': 'd5b67286-b88c-498a-9e11-b42c2bb395f8'},`

But this does not work:
`In [20]: app.content_sources.rest_client.repositories_api.search_package_group('Xfce')
ApiTypeError: Invalid type for variable 'api_search_package_group_request'. Required value type is ApiSearchPackageGroupRequest and passed type was str at ['api_search_package_group_request']
`

### Comment by @xbhouse on December 04, 2023 at 03:11 PM UTC

hi @swadeley. the request should be formed exactly like what you are passing to search_rpm() [here](https://gitlab.cee.redhat.com/insights-qe/iqe-content-sources-plugin/-/blob/master/iqe_content_sources/tests/test_repository_api_only.py?ref_type=heads#L556). so instead of a string it would be a dictionary. i'll update the testing steps to clarify this too

### Comment by @swadeley on December 05, 2023 at 12:45 AM UTC

Hi @xbhouse 

Thank you

```
In [7]: app.content_sources.rest_client.repositories_api.search_package_group({"search": "Xfce", "urls": ["https://dl.fedoraproject.org/pub/epel/8
   ...: /Everything/x86_64/"]})
Out[7]: 
[{'description': 'A lightweight desktop environment that works well on low end '
                 'machines.',
  'package_group_name': 'Xfce',
  'package_list': ['xfce4-screensaver',
                   'tumbler',
                   'NetworkManager-gnome',
                   'openssh-askpass',
                   'gdm',
                   'xfce4-terminal',
                   'thunar-volman',
                   'xfce4-panel',
                   'xfce4-appfinder',
                   'xfwm4',
                   'xfce4-settings',
                   'xfce4-pulseaudio-plugin',
                   'pinentry-gtk',
                   'mousepad',
                   'xfdesktop',
                   'xfconf',
                   'Thunar',
                   'xfce-polkit',
                   'thunar-archive-plugin',
                   'xfce4-session',
                   'xfce4-power-manager']}]
```

---

## Reviews

### Review by @rverdile - Commented on November 21, 2023 at 08:31 PM UTC

I think there may be some places where we can create shared functions or structs between rpm, package groups, and environments. And in the future, snapshot rpms, package groups, and environments. All 6 of these things have a similar flow, so it probably makes sense to try and de-duplicate where possible.

I left comments where I think might be potential places to create shared code. With that said, it probably make sense to leave it as-is here, and make changes to the environments PR instead. This way you can see all 3 at the same time and easily spot any duplication.

### Review by @rverdile - Commented on November 21, 2023 at 09:10 PM UTC

### Review by @rverdile - Commented on November 21, 2023 at 09:10 PM UTC

### Review by @xbhouse - Commented on November 21, 2023 at 09:25 PM UTC

### Review by @rverdile - Commented on November 22, 2023 at 03:33 PM UTC

### Review by @rverdile - Commented on November 22, 2023 at 03:49 PM UTC

### Review by @rverdile - Commented on November 22, 2023 at 05:01 PM UTC

This looks good, just needs a rebase and potentially that one test

### Review by @xbhouse - Commented on November 29, 2023 at 09:02 PM UTC

### Review by @rverdile - Approved on November 30, 2023 at 07:23 PM UTC

nice job!! :tada: :tada: :tada: 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/461*

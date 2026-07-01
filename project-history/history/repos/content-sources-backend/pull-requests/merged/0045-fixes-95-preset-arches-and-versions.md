---
type: pull_request
number: 45
title: "Fixes 95: preset arches and versions"
state: merged
author: jlsherrill
created: 2022-06-24T18:27:19Z
updated: 2022-07-19T20:20:39Z
closed: 2022-07-19T15:37:32Z
merged: 2022-07-19T15:37:32Z
base_branch: main
head_branch: params
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/45
---

# Pull Request #45: Fixes 95: preset arches and versions

**Author**: @jlsherrill
**Created**: June 24, 2022 at 06:27 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `params`

## Description

* App api /repository_parameters to return a list of versions and arches
* Validate repository_configurations are saved with a valid version and arch

---

## Discussion

### Comment by @jlsherrill on June 24, 2022 at 06:27 PM UTC

TODO:
* [x] add validation
* [x] tests

### Comment by @jlsherrill on June 24, 2022 at 06:37 PM UTC

Example:  
```
$ curl localhost:8000/api/content_sources/v1.0/repository_parameters/ | python -m json.tool
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   281  100   281    0     0   115k      0 --:--:-- --:--:-- --:--:--  137k
{
    "distribution_versions": [
        {
            "name": "el7",
            "label": "el7"
        },
        {
            "name": "el8",
            "label": "el8"
        },
        {
            "name": "el9",
            "label": "el9"
        }
    ],
    "distribution_arches": [
        {
            "name": "x86_64",
            "label": "x86_64"
        },
        {
            "name": "s390x",
            "label": "s390x"
        },
        {
            "name": "ppc64le",
            "label": "ppc64le"
        },
        {
            "name": "aarch64",
            "label": "aarch64"
        }
    ]
}
```

### Comment by @jlsherrill on June 28, 2022 at 08:20 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-95

### Comment by @jlsherrill on July 05, 2022 at 06:23 PM UTC

rebased, but still have to address comment

### Comment by @jlsherrill on July 11, 2022 at 04:16 PM UTC

rebased and updated

### Comment by @jlsherrill on July 14, 2022 at 06:28 PM UTC

> When I seeded data, I got this for versions: `{el7,el7,el8,el7}`. Should it be possible to have el7 3 times? Or just once?

I've fixed this during seeding, but we maybe should add a check for this.  i'll open a bug to add a validation

### Comment by @jlsherrill on July 18, 2022 at 06:43 PM UTC

Built and pushed image for cc5b7fe126b507ba9bb9cf0ec27c2da31b6f0c08

### Comment by @jlsherrill on July 19, 2022 at 08:20 PM UTC

https://issues.redhat.com/browse/14702683

---

## Reviews

### Review by @jlsherrill - Commented on June 24, 2022 at 06:30 PM UTC

### Review by @jlsherrill - Commented on June 24, 2022 at 06:33 PM UTC

### Review by @rverdile - Commented on June 27, 2022 at 08:07 PM UTC

### Review by @jlsherrill - Commented on June 27, 2022 at 08:17 PM UTC

### Review by @rverdile - Commented on June 27, 2022 at 08:20 PM UTC

### Review by @jlsherrill - Commented on June 27, 2022 at 08:32 PM UTC

### Review by @rverdile - Commented on June 28, 2022 at 02:13 PM UTC

### Review by @jlsherrill - Commented on June 29, 2022 at 03:00 PM UTC

### Review by @jlsherrill - Commented on July 06, 2022 at 07:50 PM UTC

### Review by @rverdile - Commented on July 07, 2022 at 07:51 PM UTC

### Review by @rverdile - Commented on July 07, 2022 at 07:51 PM UTC

### Review by @rverdile - Approved on July 07, 2022 at 07:52 PM UTC

two small comments but everything else working well

### Review by @rverdile - Commented on July 07, 2022 at 08:03 PM UTC

### Review by @swadeley - Commented on July 11, 2022 at 12:37 PM UTC

### Review by @jlsherrill - Commented on July 11, 2022 at 03:42 PM UTC

### Review by @jlsherrill - Commented on July 11, 2022 at 03:50 PM UTC

### Review by @jlsherrill - Commented on July 11, 2022 at 03:51 PM UTC

### Review by @rverdile - Changes Requested on July 14, 2022 at 03:03 PM UTC

When I seeded data, I got this for versions: `{el7,el7,el8,el7}`. Should it be possible to have el7 3 times? Or just once?

### Review by @rverdile - Commented on July 14, 2022 at 05:00 PM UTC

### Review by @rverdile - Commented on July 14, 2022 at 05:02 PM UTC

### Review by @jlsherrill - Commented on July 14, 2022 at 06:27 PM UTC

### Review by @rverdile - Approved on July 14, 2022 at 08:55 PM UTC

re-ack

### Review by @rverdile - Commented on July 15, 2022 at 01:06 PM UTC

### Review by @jlsherrill - Commented on July 15, 2022 at 01:26 PM UTC

### Review by @swadeley - Approved on July 19, 2022 at 03:32 PM UTC

ACK

### Review by @swadeley - Approved on July 19, 2022 at 03:37 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/45*

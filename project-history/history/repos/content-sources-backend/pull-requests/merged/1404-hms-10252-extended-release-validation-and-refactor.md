---
type: pull_request
number: 1404
title: "HMS-10252: extended release validation and refactor repo params"
state: merged
author: rverdile
created: 2026-03-04T20:20:56Z
updated: 2026-03-24T14:00:21Z
closed: 2026-03-24T14:00:12Z
merged: 2026-03-24T14:00:12Z
base_branch: main
head_branch: template-validation
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1404
---

# Pull Request #1404: HMS-10252: extended release validation and refactor repo params

**Author**: @rverdile
**Created**: March 04, 2026 at 08:20 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `template-validation`

## Description

## Summary
Refactors ExtendedReleaseFeature struct to separate label (eus/e4s) from feature name (RHEL-EUS-x86_64) and add architecture field. Implements validation for ExtendedRelease and ExtendedReleaseVersion fields in both Template and RepositoryConfiguration models.

Refactors the repository_parameters API to make it easier to parse and remove the need to update the frontend when adding new release streams.

## Testing steps
1. Repositories are tested by the unit test. There is no way to set these fields as a user.
2. For templates, try creating a template with an invalid extended release type and version. Previously, you'd get a 500 error.
3. GET `/repository_parameters/` and the new distribution_minor_versions and extended_release_streams fields should look something like this

```
"distribution_minor_versions": [
    {
      "name": "RHEL 8.6",
      "label": "8.6",
      "major": "8",
      "extended_release_stream": [
        "e4s"
      ]
    },
```
```
"extended_release_streams": [
    {
      "name": "Extended Update Support (EUS)",
      "label": "eus",
      "architectures": [
        {
          "name": "x86_64",
          "label": "x86_64",
          "entitled": true
        },
        {
          "name": "aarch64",
          "label": "aarch64",
          "entitled": false
        }
      ]
    },
```


---

## Discussion

### Comment by @xbhouse on March 04, 2026 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-10252

### Comment by @rverdile on March 23, 2026 at 06:56 PM UTC

/retest

---

## Reviews

### Review by @katarinazaprazna - Approved on March 24, 2026 at 11:37 AM UTC

Looks great to me! Thanks for the PR :)

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1404*

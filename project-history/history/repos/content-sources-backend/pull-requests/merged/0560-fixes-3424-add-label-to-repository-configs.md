---
type: pull_request
number: 560
title: "Fixes 3424: add label to repository configs"
state: merged
author: jlsherrill
created: 2024-02-02T21:33:51Z
updated: 2024-02-06T18:30:26Z
closed: 2024-02-06T18:07:41Z
merged: 2024-02-06T18:07:41Z
base_branch: main
head_branch: 3424
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/560
---

# Pull Request #560: Fixes 3424: add label to repository configs

**Author**: @jlsherrill
**Created**: February 02, 2024 at 09:33 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3424`

## Description

## Summary

Adds a label column to the repository configurations table, and returns it from the repositories api.  
For custom repos this will currently be blank (in the future we will likely set this).
For red hat repos, this will be populated
For the snapshot config.repo api, this label is used (if present). 

## Testing steps

* check out main and run 'make repos-import'.  GET /repositories?origin=external,red_hat and see that there is no label attribute
* checkout this pr, db migrate, run `make repos-import` again
* refetch the repositories via the same api and confirm that there is now a label.
* Create a custom repo and see that its label is blank
* trigger a snapshot of a redhat repo (go run cmd/external-repos/main.go snapshot $URL)  
* fetch the config.repo file and confirm that the correct label is there:   GET /repositories/UUID/snapshots/UUID/config.repo

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on February 02, 2024 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-3424

### Comment by @swadeley on February 05, 2024 at 12:18 PM UTC

/retest

### Comment by @swadeley on February 05, 2024 at 01:18 PM UTC

Hi

I get  introspection failure  missing comps file in repo[1]. Tool tip shows:

`error getting comps: error parsing comps.xml: invalid file type: must be gzip, xz, or zstd`

I do not see any label in the Red Hat Ansible Engine repo config. Should I?

[1] http://yum.theforeman.org/pulpcore/3.16/el7/x86_64/


### Comment by @swadeley on February 05, 2024 at 02:41 PM UTC

Hi, I raised Jira for the comps file detection problem:
 https://issues.redhat.com/browse/HMS-3539

### Comment by @jlsherrill on February 05, 2024 at 03:27 PM UTC

the label is in the square brackets:

for example, needed_errata here:

```
[needed_errata]
name=needed errata
baseurl=http://pulp_content:24816/pulp/content/7f584fe5/9d7cac21-c81c-4b98-9adf-0539ad2a9a68/ac7891a4-0560-4af2-b1e5-e7586fe89e03/
module_hotfixes=0
gpgcheck=0
repo_gpgcheck=0
enabled=1
gpgkey=
```

### Comment by @swadeley on February 05, 2024 at 07:28 PM UTC

Hi, OK, sorry

[ansible-2-for-rhel-8-x86_64-rpms]
name=Red Hat Ansible Engine 2 for RHEL 8 x86_64 (RPMs)
baseurl=https://<>.openshiftapps.com/api/pulp-content/7d531f1b/3a77a11a-583e-438e-bc40-34dbbb899c2e/53ffc4b1-4b2c-47dd-a55f-8fbdd96d3eb2/
module_hotfixes=0
gpgcheck=1
repo_gpgcheck=0
enabled=1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-redhat-release

### Comment by @swadeley on February 06, 2024 at 05:07 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Commented on February 05, 2024 at 09:43 PM UTC

### Review by @jlsherrill - Commented on February 06, 2024 at 01:36 PM UTC

### Review by @jlsherrill - Commented on February 06, 2024 at 01:37 PM UTC

### Review by @rverdile - Approved on February 06, 2024 at 03:53 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/560*

---
type: pull_request
number: 215
title: "Fixes 1252: switch to new epel urls"
state: merged
author: jlsherrill
created: 2023-02-16T13:26:31Z
updated: 2023-02-20T18:03:40Z
closed: 2023-02-20T18:02:17Z
merged: 2023-02-20T18:02:17Z
base_branch: main
head_branch: epel
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/215
---

# Pull Request #215: Fixes 1252: switch to new epel urls

**Author**: @jlsherrill
**Created**: February 16, 2023 at 01:26 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `epel`

## Description

and drop epel8 modularity.  Also fixes a db-health-check issue


## Testing steps

```testing.http
### Create test repos
POST http://localhost:8000/api/content-sources/v1.0/repositories/bulk_create/
Content-Type: application/json
x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiQXNzb2NpYXRlIiwiYWNjb3VudF9udW1iZXIiOiIxIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMSJ9fX0K

[
  {
   "name":"old_epel9",
    "url": "https://download-i2.fedoraproject.org/pub/epel/9/Everything/x86_64/"
  },
  {
    "name": "old_epel8",
    "url": "https://download-i2.fedoraproject.org/pub/epel/8/Everything/x86_64/"
  },
  {
    "name": "new_epel8",
    "url": "https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/"
  }
]

### List repos

GET http://localhost:8000/api/content-sources/v1.0/repositories/
Content-Type: application/json
x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiQXNzb2NpYXRlIiwiYWNjb3VudF9udW1iZXIiOiIxIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMSJ9fX0K

```

1. On main, execute the first request to create an old epel 9 repo, and the old epel 8 and new epel8 repos
2. switch to this PR, and run `make db-migrate-up`
3. run the list repos command, and notice that the old_epel9 repo has been updated, but both epel8 repos remain as they were


For testing in the ephemeral env:
1.  spin up current latest image (not this pr)
2. create a series of repos:
  a. Current "epel 9 Everything" from popular repos (https://download-i2.fedoraproject.org/pub/epel/9/Everything/x86_64/)
  b. current "epel 8 everything" from popular repos (https://download-i2.fedoraproject.org/pub/epel/8/Everything/x86_64/)
  c. New "epel 8 everything repository" (https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/)
  d. current "epel 8 modular repository"  (https://download-i2.fedoraproject.org/pub/epel/8/Modular/x86_64/)
  e.  current "epel 7" from popular repos (https://download-i2.fedoraproject.org/pub/epel/7/x86_64/)
3.  deploy new image to your Ephemeral environment
4. examine UI, repos should be updated:
  a. should be updated to the new epel 9 url: https://dl.fedoraproject.org/pub/epel/9/Everything/x86_64/
  b. should not be updated (because a repo with the new url already exists)
  c.  should not be updated (because its already the new url)
  d.  should not be updated (we're removing it from the popular repos list, but don't delete existing user copies)
  e. should be updated to the new epel 7 url: https://dl.fedoraproject.org/pub/epel/7/x86_64/

---

## Discussion

### Comment by @jlsherrill on February 16, 2023 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-1252

### Comment by @swadeley on February 20, 2023 at 03:02 PM UTC

/retest

### Comment by @swadeley on February 20, 2023 at 03:28 PM UTC

/retest

### Comment by @swadeley on February 20, 2023 at 06:01 PM UTC

Hi

I have tested as described above.
Works as expected, but note that the presence of a new version of the same repo, e.g.
"c. New "epel 8 everything repository" (https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/)"
means you will have to remove the old one. i.e. current "epel 8 everything"

---

## Reviews

### Review by @Andrewgdewar - Commented on February 16, 2023 at 03:11 PM UTC

### Review by @jlsherrill - Commented on February 16, 2023 at 03:15 PM UTC

### Review by @Andrewgdewar - Commented on February 16, 2023 at 03:24 PM UTC

### Review by @Andrewgdewar - Commented on February 16, 2023 at 05:17 PM UTC

### Review by @Andrewgdewar - Commented on February 16, 2023 at 05:18 PM UTC

### Review by @jlsherrill - Commented on February 16, 2023 at 06:22 PM UTC

### Review by @Andrewgdewar - Approved on February 17, 2023 at 03:57 PM UTC

Ack!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/215*

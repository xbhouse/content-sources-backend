---
type: pull_request
number: 241
title: "Fixes 1564: Update migration to prevent status ui bug"
state: merged
author: Andrewgdewar
created: 2023-04-03T21:31:38Z
updated: 2023-04-04T11:30:14Z
closed: 2023-04-04T11:00:29Z
merged: 2023-04-04T11:00:29Z
base_branch: main
head_branch: HMS-1564
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/241
---

# Pull Request #241: Fixes 1564: Update migration to prevent status ui bug

**Author**: @Andrewgdewar
**Created**: April 03, 2023 at 09:31 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-1564`

## Description

## Summary
This fixes an issue where in dev/ephemeral environments adding a "Popular Repository" would cause the API to return a Status of "" (empty string). 
This prevented the UI from updating the repository list the frontend looks at a repository's "Status" to determine whether or not to poll the API.

## Testing steps
Add a popular repository that needs to introspect.
One you haven't added before, or clean out your database.

After adding quickly switch to the repository list view you should see a blank status.

Result after fixing: 
It should say Pending, or any state, if introspection has finished.

---

## Discussion

### Comment by @jlsherrill on April 03, 2023 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-1564

### Comment by @swadeley on April 04, 2023 at 11:00 AM UTC

Hi

I first confirmed that before this change there was blank status immediately after adding Popular repo. Now I see "In progress" with a swirlling icon to the left, changing to Valid after introspection.

thank you

---

## Reviews

### Review by @jlsherrill - Approved on April 04, 2023 at 12:49 AM UTC

ACK, normally we wouldn't update a migration, but in this case it actually makes sense.  The problem seen here would never actually be seen in an existing running environment, only for new environments

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/241*

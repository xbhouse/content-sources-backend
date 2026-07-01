---
type: pull_request
number: 918
title: "HMS-5125: Add support for resumable downloads"
state: merged
author: Andrewgdewar
created: 2024-12-11T21:29:55Z
updated: 2024-12-31T18:08:32Z
closed: 2024-12-31T18:08:31Z
merged: 2024-12-31T18:08:31Z
base_branch: main
head_branch: HMS-5125
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/918
---

# Pull Request #918: HMS-5125: Add support for resumable downloads

**Author**: @Andrewgdewar
**Created**: December 11, 2024 at 09:29 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `HMS-5125`

## Description

## Summary
Adds support for resumable downloads. 

Can test with [this](https://github.com/content-services/content-sources-frontend/pull/397) front-end PR.


## Testing steps

- Create an upoad repository.
- Select an item for upload, allow it to upload around halfway and then immediately hit refresh.
- After refresh, choose the same item, you should be able to see that item continue where you left off. 
- After this item has been completed, upload it, and create another upload repository. 
- Choose the same file, it should instantly be available.



---

## Discussion

### Comment by @jlsherrill on December 11, 2024 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-5125

### Comment by @jlsherrill on December 18, 2024 at 06:12 PM UTC

https://issues.redhat.com/browse/Fixes 5125

---

## Reviews

### Review by @dominikvagner - Changes Requested on December 20, 2024 at 12:21 PM UTC

small change needed in the API spec, otherwise looks good, nice job! 😄👏🏼  
everything works as expected 👍🏼 

_(tiny nit about style/formatting: in Go the `if err != nil` check usually follow the err assignment without a new line in between them, makes it more clear what error they are checking)_

### Review by @Andrewgdewar - Commented on December 20, 2024 at 05:51 PM UTC

### Review by @dominikvagner - Approved on December 20, 2024 at 06:01 PM UTC

### Review by @jlsherrill - Commented on December 20, 2024 at 08:52 PM UTC

### Review by @Andrewgdewar - Commented on December 27, 2024 at 05:54 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/918*

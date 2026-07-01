---
type: pull_request
number: 202
title: "Fixes 576: add package dependency graph"
state: merged
author: rverdile
created: 2023-02-08T16:16:42Z
updated: 2023-02-15T14:01:32Z
closed: 2023-02-15T14:01:26Z
merged: 2023-02-15T14:01:26Z
base_branch: main
head_branch: pkg-graph
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/202
---

# Pull Request #202: Fixes 576: add package dependency graph

**Author**: @rverdile
**Created**: February 08, 2023 at 04:16 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `pkg-graph`

## Description

## Summary
Adds a graph visualization of our app's internal dependencies and `make pkg-graph-generate` to auto-generate a new image.

Maybe I should put this in the README or Architecture doc?

ToDo
- [ ] Add to README.md  or Architecture doc?

## Testing steps
1. Look for any problems with the graph
2. Run the make command


---

## Discussion

### Comment by @rverdile on February 08, 2023 at 04:23 PM UTC

There are two ways to graph this: with and without showing the dependencies used by test files. I'm not sure which one we want (maybe both). I started up uploading just the one without the test files, but I'm open to suggestions.

Showing dependencies without test files can be misleading, for example `dao tests` imports `seeds`, but the graph does not show that `dao` imports `seeds`. However, showing test dependencies is really messy.

**No tests**
![pkg-graph-no-tests](https://user-images.githubusercontent.com/51457421/217588462-0f17ce6e-e0f4-44e5-9ff3-ed17b40aadc1.png)


**Tests**
![pkg-graph-tests](https://user-images.githubusercontent.com/51457421/217588639-d024fbfa-352f-4969-b3d9-778892151724.png)


### Comment by @jlsherrill on February 08, 2023 at 04:30 PM UTC

https://issues.redhat.com/browse/HMS-576

### Comment by @jlsherrill on February 09, 2023 at 05:34 PM UTC

Overall looks good, i'd not include the one with tests, i don't think its super helpful

---

## Reviews

### Review by @jlsherrill - Approved on February 14, 2023 at 08:04 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/202*

---
type: pull_request
number: 13
title: "Fixes 4740: Change base query to use inner joins"
state: merged
author: jlsherrill
created: 2024-09-25T15:53:47Z
updated: 2024-10-02T22:17:55Z
closed: 2024-10-02T22:16:46Z
merged: 2024-10-02T22:16:46Z
base_branch: main
head_branch: query_speedup
labels: []
url: https://github.com/content-services/tang/pull/13
---

# Pull Request #13: Fixes 4740: Change base query to use inner joins

**Author**: @jlsherrill
**Created**: September 25, 2024 at 03:53 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `query_speedup`

## Description

I saw improvements with package queries by around 2x.  Errata query times were very similar to before.

To test:

Load a bunch of repos and sync them (make repos-import && go run cmd/external-repos/main.go nightly-jobs).  You might add some large fedora repos too.


in content-sources-backend go.mod add this under `go 1.23`:

```
replace github.com/content-services/tang => /home/jlsherri/git/tang
```

(replace with path to your checkout).
```
go get ./...
```

(Note that when switching branches, or modifying code in tang, you must  `rm  release/content-sources`  before you `make run` to pick up the changes

* Confirm that listing rpms and advisories for templates works
* compare the counts before and after this change
* To see a time difference, you can enable query logging in postgres:
```
psql --host localhost -U pulp
ALTER SYSTEM SET log_min_duration_statement = 25;
select pg_reload_conf();
```

Now after fetching some set of content, you can run ```podman logs --tail 1000 cs_postgres_1``` and see the time it takes:

```
LOG:  duration: 208.538 ms
```

You can switch back and forth between main and this branch to see the changes.  You may want to compare the times (should be about half).



---

## Discussion

### Comment by @rverdile on October 01, 2024 at 01:14 PM UTC

also seeing ~2x speedup!

---

## Reviews

### Review by @rverdile - Approved on October 01, 2024 at 01:14 PM UTC

---

*Archived from: https://github.com/content-services/tang/pull/13*

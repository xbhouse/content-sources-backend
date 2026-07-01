---
type: pull_request
number: 652
title: "Fixes 3930: speed up slow rpm searching"
state: merged
author: xbhouse
created: 2024-04-29T13:58:57Z
updated: 2024-04-30T13:00:23Z
closed: 2024-04-30T12:41:18Z
merged: 2024-04-30T12:41:18Z
base_branch: main
head_branch: fix-slow-rpm-search
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/652
---

# Pull Request #652: Fixes 3930: speed up slow rpm searching

**Author**: @xbhouse
**Created**: April 29, 2024 at 01:58 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `fix-slow-rpm-search`

## Description

## Summary

Adds pg_trgm extension and new index on rpm names to speed up slow search rpm queries.

It looks like this solution makes the search rpms query about 10x faster (6ms vs 60ms locally) when there is a more selective filter on rpm name. 

Some caveats: 
- Index creation took a bit of time for me locally (fluctuated between 1 and 7 seconds)
- Index size is twice the size of the b-tree index that already exists on rpms.name
- Less selective filters on an rpm name (1 or 2 characters) are less efficient than the original search (70ms vs 60ms)

References:
https://www.postgresql.org/docs/current/pgtrgm.html#PGTRGM-INDEX
https://www.depesz.com/2011/02/19/waiting-for-9-1-faster-likeilike/

## Testing steps

- Run the migration to add the new extension and index (for QE, migration happens on deployment)
- Make a request to /rpms/names with a search term like `3proxy`, should be faster than before
- Verify introspection takes about the same time with or without the index (for me it was about a 1 ms difference)

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on April 29, 2024 at 02:00 PM UTC

https://issues.redhat.com/browse/HMS-3930

### Comment by @mayurilahane on April 30, 2024 at 11:47 AM UTC

/retest

### Comment by @mayurilahane on April 30, 2024 at 12:28 PM UTC

lgtm !
deployment is successfully done!

`2024-04-30 08:18:16 [    INFO] [          MainThread] successfully deployed to namespace ephemeral-derewx
2024-04-30 08:18:16 [    INFO] [          MainThread] successful deployment
`

and
I see the mentioned pod running in it 

`content-sources-backend-task-worker-6d85c65c46-gm2g7             1/1     Running     0             13m
`

Rpm search will be tested after this is merged
https://issues.redhat.com/browse/HMS-4045



### Comment by @mayurilahane on April 30, 2024 at 12:32 PM UTC

Logs of the same pod also confirm that migration is done successfully.

```
(new_env) [mlahane@ibm-p8-kvm-01-fsp iqe-content-sources-plugin]$ oc logs content-sources-backend-task-worker-6d85c65c46-gm2g7 -c db-migrate-init 
{"level":"warn","time":"2024-04-30T12:13:38Z","message":"config.yaml file not loaded: Config File \"config.yaml\" Not Found in \"[/configs]\""}
{"level":"debug","platform.content-sources.template":"platform.content-sources.template","time":"2024-04-30T12:13:38Z","message":"internalToReal"}
{"level":"debug","platform.content-sources.template":"platform.content-sources.template","time":"2024-04-30T12:13:38Z","message":"realToInternal"}
{"level":"debug","platform.notifications.ingress":"platform.notifications.ingress","time":"2024-04-30T12:13:38Z","message":"internalToReal"}
{"level":"debug","platform.notifications.ingress":"platform.notifications.ingress","time":"2024-04-30T12:13:38Z","message":"realToInternal"}
{"level":"error","error":"no RDS available","time":"2024-04-30T12:13:38Z","message":"Cannot read RDS CA cert"}
{"level":"warn","time":"2024-04-30T12:13:38Z","message":"Bucket name: content-sources-central-pulp-s3"}
0xc000396570
{"level":"debug","time":"2024-04-30T12:13:39Z","message":"Successfully migrated up"}

```

---

## Reviews

### Review by @jlsherrill - Approved on April 29, 2024 at 08:15 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/652*

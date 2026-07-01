---
type: pull_request
number: 619
title: "Fixes 3857: allow searching for popular repos"
state: merged
author: jlsherrill
created: 2024-03-28T18:26:50Z
updated: 2024-04-09T13:09:33Z
closed: 2024-04-09T13:09:33Z
merged: 2024-04-09T13:09:33Z
base_branch: main
head_branch: POPULAR
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/619
---

# Pull Request #619: Fixes 3857: allow searching for popular repos

**Author**: @jlsherrill
**Created**: March 28, 2024 at 06:26 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `POPULAR`

## Description

## Summary

Image builder wanted to be able to search popular repositories and suggest users add them.  This expands the rpm search api to support that.  It makes the assumption that at least one user in the environment has added the repo for introspection, which is a safe assumption at this point.  

## Testing steps

```
OPTIONS_REPOSITORY_IMPORT_FILTER=small go run cmd/external-repos/main.go import
go run cmd/external-repos/main.go nightly-job
```

then search for some repos (no need to add them):

```
POST http://localhost:8000/api/content-sources/v1.0/rpms/names
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiIxNjc4ODYxNiIsImludGVybmFsIjp7ImNyb3NzX2FjY2VzcyI6ZmFsc2UsIm9yZ19pZCI6IjE2Nzg4NjE2IiwiYXV0aF90aW1lIjowfSwidXNlciI6eyJpc19hY3RpdmUiOnRydWUsImVtYWlsIjoidGVzdEByZWRoYXQuY29tIiwiZmlyc3RfbmFtZSI6ImZpcnN0X25hbWUiLCJ1c2VybmFtZSI6ImNvbnRlbnQ5IiwiaXNfb3JnX2FkbWluIjp0cnVlLCJsb2NhbGUiOiJlbl91cyIsImlzX2ludGVybmFsIjpmYWxzZSwidXNlcl9pZCI6IjU1OTAyMTU3IiwibGFzdF9uYW1lIjoibGFzdF9uYW1lIn0sImF1dGhfdHlwZSI6Imp3dC1hdXRoIiwiYWNjb3VudF9udW1iZXIiOiIxMTQ4NjU4OSIsInR5cGUiOiJVc2VyIn0sImVudGl0bGVtZW50cyI6eyJyaG9kcyI6eyJpc19lbnRpdGxlZCI6ZmFsc2UsImlzX3RyaWFsIjpmYWxzZX0sInJob3NhayI6eyJpc19lbnRpdGxlZCI6ZmFsc2UsImlzX3RyaWFsIjpmYWxzZX0sInNldHRpbmdzIjp7ImlzX2VudGl0bGVkIjp0cnVlLCJpc190cmlhbCI6ZmFsc2V9LCJpbnRlcm5hbCI6eyJpc19lbnRpdGxlZCI6ZmFsc2UsImlzX3RyaWFsIjpmYWxzZX0sInVzZXJfcHJlZmVyZW5jZXMiOnsiaXNfZW50aXRsZWQiOnRydWUsImlzX3RyaWFsIjpmYWxzZX0sImluc2lnaHRzIjp7ImlzX2VudGl0bGVkIjp0cnVlLCJpc190cmlhbCI6ZmFsc2V9LCJhY3MiOnsiaXNfZW50aXRsZWQiOmZhbHNlLCJpc190cmlhbCI6ZmFsc2V9LCJzdWJzY3JpcHRpb25zIjp7ImlzX2VudGl0bGVkIjp0cnVlLCJpc190cmlhbCI6ZmFsc2V9LCJzbWFydF9tYW5hZ2VtZW50Ijp7ImlzX2VudGl0bGVkIjp0cnVlLCJpc190cmlhbCI6ZmFsc2V9LCJhbnNpYmxlIjp7ImlzX2VudGl0bGVkIjp0cnVlLCJpc190cmlhbCI6ZmFsc2V9LCJjb3N0X21hbmFnZW1lbnQiOnsiaXNfZW50aXRsZWQiOnRydWUsImlzX3RyaWFsIjpmYWxzZX0sIm1pZ3JhdGlvbnMiOnsiaXNfZW50aXRsZWQiOnRydWUsImlzX3RyaWFsIjpmYWxzZX0sIm9wZW5zaGlmdCI6eyJpc19lbnRpdGxlZCI6ZmFsc2UsImlzX3RyaWFsIjpmYWxzZX0sInJoZWwiOnsiaXNfZW50aXRsZWQiOnRydWUsImlzX3RyaWFsIjpmYWxzZX0sInJob2FtIjp7ImlzX2VudGl0bGVkIjpmYWxzZSwiaXNfdHJpYWwiOmZhbHNlfX19
Content-Type: application/json


{
	"search": "vim",
	"urls": [
		"https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/"
	]
}
```

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on March 28, 2024 at 06:30 PM UTC

https://issues.redhat.com/browse/HMS-3857

### Comment by @mayurilahane on March 28, 2024 at 07:04 PM UTC

/retest

### Comment by @xbhouse on April 01, 2024 at 07:58 PM UTC

similar to the other search APIs, should we add a check to ensure the repositories are found before searching? maybe this isn't necessary if we're assuming they should always have already been added for introspection

### Comment by @xbhouse on April 01, 2024 at 08:01 PM UTC

lgtm other than the one question :) 

### Comment by @swadeley on April 07, 2024 at 07:24 PM UTC

/retest

### Comment by @jlsherrill on April 09, 2024 at 12:37 PM UTC

@xbhouse i'm gonna file a bug for that, if that works?

### Comment by @jlsherrill on April 09, 2024 at 12:39 PM UTC

filed: https://issues.redhat.com/browse/HMS-3916

### Comment by @xbhouse on April 09, 2024 at 12:59 PM UTC

sounds good! ack!

---

## Reviews

### Review by @xbhouse - Approved on April 09, 2024 at 12:59 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/619*

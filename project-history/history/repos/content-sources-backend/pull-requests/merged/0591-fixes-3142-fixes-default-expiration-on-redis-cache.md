---
type: pull_request
number: 591
title: "Fixes 3142: fixes default expiration on redis cache"
state: merged
author: jlsherrill
created: 2024-02-29T20:23:02Z
updated: 2024-03-01T10:30:34Z
closed: 2024-03-01T10:01:27Z
merged: 2024-03-01T10:01:27Z
base_branch: main
head_branch: 3142
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/591
---

# Pull Request #591: Fixes 3142: fixes default expiration on redis cache

**Author**: @jlsherrill
**Created**: February 29, 2024 at 08:23 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3142`

## Description

## Summary
and resets cache keys

## Testing steps

to check the default, apply this patch:

```
diff --git a/cmd/external-repos/main.go b/cmd/external-repos/main.go
index 6eb17810..26065f99 100644
--- a/cmd/external-repos/main.go
+++ b/cmd/external-repos/main.go
@@ -27,6 +27,8 @@ func main() {
        config.Load()
        config.ConfigureLogging()
 
+       fmt.Printf("FOO: %v", config.Get().Clients.Redis.Expiration.Rbac)
+
        err := db.Connect()
        if err != nil {
                log.Panic().Err(err).Msg("Failed to connect to database")
```

then when you run: 
```
 go run cmd/external-repos/main.go 
{"level":"warn","time":"2024-02-29T15:11:32-05:00","message":"Caching is disabled."}
FOO: 0s
```
With this change, you should see the default of 60s.

Also you can look in redis itself.
Before:
```
 podman run -it docker.io/library/redis   redis-cli   -h 192.168.1.104  -n 1
192.168.1.104:6379[1]> KEYS *
1) "central-pulp-content-path"
2) "auth:new,1"
192.168.1.104:6379[1]> TTL "auth:new,1"
(integer) -1
```
After:
```
192.168.1.104:6379[1]> KEYS *
1) "authen:new,1"
2) "auth:new4,1"
3) "central-pulp-content-path"
4) "auth:new,1"
192.168.1.104:6379[1]> TTL "authen:new,1"
(integer) 47
```
## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on February 29, 2024 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-3142

### Comment by @jlsherrill on February 29, 2024 at 08:54 PM UTC

QE testing may be impossible in Ephemeral.  You need a user that has no repository permissions (or read only) in rbac, access the application with their user, then give them more permissions, wait ~60s, and see that the permissions take effect (by trying to create a repository).

---

## Reviews

### Review by @rverdile - Approved on February 29, 2024 at 09:40 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/591*

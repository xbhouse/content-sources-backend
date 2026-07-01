---
type: pull_request
number: 865
title: "Fixes 4822: rename domains with cs- prefix"
state: closed
author: jlsherrill
created: 2024-10-25T01:23:41Z
updated: 2024-11-25T14:15:57Z
closed: 2024-11-25T14:15:57Z
base_branch: main
head_branch: 4822_1
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/865
---

# Pull Request #865: Fixes 4822: rename domains with cs- prefix

**Author**: @jlsherrill
**Created**: October 25, 2024 at 01:23 AM UTC
**Status**: Closed
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4822_1`

## Description

## Summary

Changes template model to no longer contain RepositoryConfigurations directly, but instead TemplateRepositoryConfigurations, this was due to an observed bug with gorm.  Strangely this bug was not always reproducible, but it was observed that querying Templates with the default scope resulted in RepositoryConfigurations included that had been removed from the template. 

This moves all the jobs to one command  ./cmd/jobs/main.go   that takes the name of the job as an argument.

Deletes the repair_latest and cleanup_versions jobs

Adds a new job 'rename_domains' that renames the red hat domain to 'cs-redhat' if its not already that.  Renames all other domains to cs-PREVIOUS_NAME  if it doesn't already start with 'cs-'.  

The rename process involves:

1.  loading the old name from the db
2. updating the paths for all the templates in candlepin
3. updating the content overrides for all the templates in candlepin TODO
4. updating the domain name in pulp
5. updating our own domains table

5 is done last so that if the process errors at any point, it can be re-run from the start for that domain simply by re-running the job.

## Testing steps

1.  On main, revert this commit: https://github.com/content-services/content-sources-backend/commit/ebdbba4f3841284682633178ea8f779193bf93ad
2. run:
```
make repos-import-rhel9
go run cmd/external-repos/main.go nightly-jobs 
```
start the server and let the RHEL repos snapshot
3.  Create at least 1 custom repo with snapshotting enabled
4. Create a content template, assign rhel9 and custom repo(s)
5. follow https://github.com/content-services/content-sources-backend/blob/main/docs/register_client.md to register a client
6. confirm that the client can install content form the redhat repos and custom repos
7. rename the domains:  `go run cmd/jobs/main.go rename_domains`
8. On the client run  `subscription-manager refresh`   && `subscription-manager repo-override`  
9. verify that you can still install content from the red hat and custom repos
10. confirm that the templates urls now point to   cs-DOMAIN  & cs-redhat
11. confirm that the repo configs in the webUI for snapshots still work
12. If you still see the old domains in the redhat.repo file, then try re-registering



---

## Discussion

### Comment by @jlsherrill on October 25, 2024 at 01:30 AM UTC

https://issues.redhat.com/browse/HMS-4822

### Comment by @xbhouse on November 05, 2024 at 04:18 PM UTC

for existing repo snapshots, after running the rename-domains job i see the old domains when fetching the repos and so the urls point to a 404. also see the old domains in the config.repos from existing repo snapshots. i'm guessing it is expected to have to trigger new snapshots before we'll see the new domains in these places? 



### Comment by @xbhouse on November 05, 2024 at 05:28 PM UTC

everything works as described on the client 🎉 just a couple questions above and a nitpick, overall nice job!!

### Comment by @jlsherrill on November 05, 2024 at 07:24 PM UTC

> for existing repo snapshots, after running the rename-domains job i see the old domains when fetching the repos and so the urls point to a 404. also see the old domains in the config.repos from existing repo snapshots. i'm guessing it is expected to have to trigger new snapshots before we'll see the new domains in these places?

Hrmmmm this sounds like something didn't go as planned, I'll try to reproduce.

### Comment by @xbhouse on November 05, 2024 at 08:03 PM UTC

> > for existing repo snapshots, after running the rename-domains job i see the old domains when fetching the repos and so the urls point to a 404. also see the old domains in the config.repos from existing repo snapshots. i'm guessing it is expected to have to trigger new snapshots before we'll see the new domains in these places?
> 
> Hrmmmm this sounds like something didn't go as planned, I'll try to reproduce.

ah, also i'm seeing this error from the delete-repository-snapshots task when deleting those existing custom repos:
```
"error reading task: 404 Not Found: \n<!doctype html>\n<html lang=\"en\">\n<head>\n  <title>Not Found</title>\n</head>\n<body>\n  <h1>Not Found</h1><p>The requested resource was not found on this server.</p>\n</body>\n</html>\n"
```

to reproduce these things:

1. revert [this commit](https://github.com/content-services/content-sources-backend/commit/ebdbba4f3841284682633178ea8f779193bf93ad) in main
2. import a red hat repo and add a custom repo and wait for both to snapshot
3. checkout this pr and run rename_domains
4. list repos (snapshot urls show old domain)
5. copy the config.repo output in the UI (baseurl shows old domain)
6. delete the custom repo 

### Comment by @jlsherrill on November 08, 2024 at 02:33 PM UTC

[test]

### Comment by @jlsherrill on November 08, 2024 at 03:14 PM UTC

[test]

### Comment by @jlsherrill on November 08, 2024 at 07:13 PM UTC

[test]

### Comment by @jlsherrill on November 08, 2024 at 07:31 PM UTC

i did discover that i wasn't updating repository_path in the snapshots table, so i added that as part of https://github.com/content-services/content-sources-backend/pull/865/commits/1e5b4263cb64bf3d85a6af260aa89bb55fd0347d  and added an integration test

### Comment by @jlsherrill on November 11, 2024 at 01:04 AM UTC

[test]

### Comment by @xbhouse on November 11, 2024 at 04:44 PM UTC

repository paths look good now!

after running rename-domains, i'm still seeing the delete-repository-snapshots task fail on deleteRpmDistribution when deleting a repo created previously though:
```
"error reading task: 404 Not Found: \n<!doctype html>\n<html lang=\"en\">\n<head>\n  <title>Not Found</title>\n</head>\n<body>\n  <h1>Not Found</h1><p>The requested resource was not found on this server.</p>\n</body>\n</html>\n" 
```

also when fetching a snapshot task in the UI for existing repos (which uses /admin/tasks/:uuid):
```
"error: code=500, title=Error fetching task, detail=error: code=500, title=Error parsing task payload, detail=error reading task: 404 Not Found: \n<!doctype html>\n<html lang=\"en\">\n<head>\n  <title>Not Found</title>\n</head>\n<body>\n  <h1>Not Found</h1><p>The requested resource was not found on this server.</p>\n</body>\n</html>\n \n \n"
```

i think the pulp hrefs in our db need to be updated to the new domain too? i see the hrefs using the new domain in pulp:
```
/api/pulp/cs-65f27c79/api/v3/distributions/rpm/rpm/01931c0c-999b-7c2b-9654-cbd2719b949b/
/api/pulp/cs-65f27c79/api/v3/repositories/rpm/rpm/01931c0c-8a47-73b0-bab7-cc2f05e239e7/versions/1/
/api/pulp/cs-65f27c79/api/v3/publications/rpm/rpm/01931c0c-9518-77d1-a5d5-a308bcca9579/
```
but not in our db:
```
/api/pulp/65f27c79/api/v3/distributions/rpm/rpm/01931c0c-999b-7c2b-9654-cbd2719b949b/
/api/pulp/65f27c79/api/v3/repositories/rpm/rpm/01931c0c-8a47-73b0-bab7-cc2f05e239e7/versions/1/
/api/pulp/65f27c79/api/v3/publications/rpm/rpm/01931c0c-9518-77d1-a5d5-a308bcca9579/
```

### Comment by @swadeley on November 13, 2024 at 09:47 AM UTC

Hi @jlsherrill , please rebase.

### Comment by @xbhouse on November 20, 2024 at 03:58 PM UTC

ah i think the task hrefs need to be updated too 😅 seeing this when trying to view a snapshot task in the admin task ui (if the repo was snapshotted before switching over to the new domain):

```
"error: code=500, title=Error fetching task, detail=error: code=500, title=Error parsing task payload, detail=error reading task: 404 Not Found: \n<!doctype html>\n<html lang=\"en\">\n<head>\n  <title>Not Found</title>\n</head>\n<body>\n  <h1>Not Found</h1><p>The requested resource was not found on this server.</p>\n</body>\n</html>\n \n \n"
```

### Comment by @xbhouse on November 20, 2024 at 04:15 PM UTC

snapshot and template hrefs look good now, and installing content on the client still works as expected

### Comment by @jlsherrill on November 25, 2024 at 02:15 PM UTC

we've decided to not do this at this time.  Closing

---

## Reviews

### Review by @xbhouse - Commented on November 05, 2024 at 05:13 PM UTC

### Review by @xbhouse - Commented on November 05, 2024 at 05:26 PM UTC

### Review by @xbhouse - Commented on November 05, 2024 at 05:31 PM UTC

### Review by @jlsherrill - Commented on November 05, 2024 at 07:23 PM UTC

### Review by @rverdile - Commented on November 06, 2024 at 05:34 PM UTC

### Review by @jlsherrill - Commented on November 07, 2024 at 02:20 PM UTC

### Review by @rverdile - Commented on November 07, 2024 at 04:53 PM UTC

### Review by @jlsherrill - Commented on November 08, 2024 at 07:19 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/865*

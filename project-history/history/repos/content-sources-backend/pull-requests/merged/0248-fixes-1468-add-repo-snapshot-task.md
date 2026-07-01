---
type: pull_request
number: 248
title: "Fixes 1468: add repo snapshot task"
state: merged
author: jlsherrill
created: 2023-04-10T16:16:43Z
updated: 2023-05-10T19:29:34Z
closed: 2023-05-03T13:34:01Z
merged: 2023-05-03T13:34:01Z
base_branch: main
head_branch: 1468
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/248
---

# Pull Request #248: Fixes 1468: add repo snapshot task

**Author**: @jlsherrill
**Created**: April 10, 2023 at 04:16 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `1468`

## Description

## Summary
This adds a task to snapshot a repository.  It is meant to recover from failures due to interruptions after:
* remote creation
* repo creation

but not due to interuptions during:
* repo sync
* publication creation
* distribution creation

This resiliency will have to be added during integration with the tasking system.

Other changes:
* Moves dao mocks out of test/mocks into the dao package.  This makes more sense from a dependency perspective
* introduces a Dao registry, to simplify creating and passing around these dao objects, especially in tests.
* Adds a snapshot dao with creation supported

## Testing steps
create a repository via ui or api:

```
POST http://{{host}}/api/content-sources/v1.0/repositories/
Content-Type: application/json
x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiQXNzb2NpYXRlIiwiYWNjb3VudF9udW1iZXIiOiIxIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMSJ9fX0K

{
  "name": "needed",
  "url": "https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/",
  "distribution_versions": [
    "8"
  ],
  "distribution_arch": "x86_64"
}
```

run the 'task' manually:
```
go run cmd/external-repos/main.go pulp-create https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/ 1
```

Confirm that a distribution was created:
```
curl http://localhost:8080/pulp/api/v3/distributions/rpm/rpm/  -u admin:password | jq
```

and you can fetch from it:

```
curl http://localhost:8080/pulp/content/a63e6d0a-9df9-4553-8eaa-e3c19180e15d/a4300e93-4b62-4704-aba5-13bfae95f593/repodata/repomd.xml
```






## TODO:
* [x] Integration tests
* [x] Down migration
* [ ] 

---

## Discussion

### Comment by @jlsherrill on April 10, 2023 at 04:30 PM UTC

https://issues.redhat.com/browse/HMS-1468

### Comment by @rverdile on April 11, 2023 at 06:34 PM UTC

+1 to the dao registry change

### Comment by @jlsherrill on April 26, 2023 at 02:40 PM UTC

Right now that is expected.  What we're gonna need to do is save the task ids on the task itself somehow (to be done during integration) and then as part of the task check that the sync task/publish task is completed, and if not retry it.   

Since this isn't being done today, what happens is the sync finishes, the next run happens and sees that nothing has changed so no new snapshot is created.  With the above changes, it'd basically pickup the previous pulp task and continue with it.

---

## Reviews

### Review by @jlsherrill - Commented on April 17, 2023 at 05:59 PM UTC

### Review by @rverdile - Commented on April 19, 2023 at 05:42 PM UTC

### Review by @rverdile - Commented on April 19, 2023 at 07:08 PM UTC

### Review by @jlsherrill - Commented on April 24, 2023 at 07:57 PM UTC

### Review by @rverdile - Commented on April 25, 2023 at 07:45 PM UTC

I was trying to see what happens when things go wrong and I might be seeing an issue with creating a snapshot, i.e. `pulp-create`, if it fails to create previously. I don't know if it's an issue here or a pulp issue.

For example,
 1.  `pulp-create https://fixtures.pulpproject.org/rpm-unsigned/  1`, for example 
 2. `ctrl+c` exit the command mid-process 
 3. `pulp-create` again,  it tells me "successfully created". However, when I navigate to `http://localhost:8080/pulp/content` I don't actually see the content there. 

I think I saw similar behavior hard-coding a panic in `PollTask` or an error `SyncRpmRepository`, such that it returns an error.

I guess my expectation would be that I could run the same command and it would work the second time no matter how bad it failed the first time.

### Review by @rverdile - Commented on April 26, 2023 at 07:26 PM UTC

### Review by @rverdile - Commented on April 28, 2023 at 04:11 PM UTC

This is looking good to me. 

Based on the TODO comments, it seems like the gaps are meant to be addressed when integrating with the tasking system. I think it's doing everything it's supposed to in its current state. I think the way it's organized will also make it easy to integrate with the tasking system, so that's a big plus.

I'll think about this a little longer, and I did leave the one small comment, but so far I haven't found anything missing that's not addressed by a TODO comment.

### Review by @rverdile - Commented on May 01, 2023 at 07:18 PM UTC

### Review by @jlsherrill - Commented on May 01, 2023 at 08:12 PM UTC

### Review by @rverdile - Approved on May 02, 2023 at 04:48 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/248*

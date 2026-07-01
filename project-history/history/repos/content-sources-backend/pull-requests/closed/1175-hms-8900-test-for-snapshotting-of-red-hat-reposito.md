---
type: pull_request
number: 1175
title: "HMS-8900: Test for Snapshotting of Red Hat Repositories"
state: closed
author: mayurilahane
created: 2025-08-14T04:34:09Z
updated: 2025-09-02T04:32:46Z
closed: 2025-09-02T04:32:46Z
base_branch: main
head_branch: mlahane/HMS-8900
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1175
---

# Pull Request #1175: HMS-8900: Test for Snapshotting of Red Hat Repositories

**Author**: @mayurilahane
**Created**: August 14, 2025 at 04:34 AM UTC
**Status**: Closed
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `mlahane/HMS-8900`

## Description

Adds a Playwright API test that verifies public and Red Hat repo's introspection.



---

## Discussion

### Comment by @jlsherrill on August 14, 2025 at 05:00 AM UTC

https://issues.redhat.com/browse/HMS-8900

### Comment by @xbhouse on August 15, 2025 at 07:27 PM UTC

i may not be too clear on the goal of this test, but i think this might make more sense as an integration test to verify the repos in stage. 

right now  we're only introspecting / snapshotting the small RH repo and the EPEL 10 repo in CI, so this test would really only verify that small RH repo. 

also the public repos' statuses aren't `Invalid` but they aren't `Valid` either because they haven't been introspected, and i'm not sure that we should introspect all of them in CI :thinking: 

### Comment by @mayurilahane on August 15, 2025 at 08:55 PM UTC

> i may not be too clear on the goal of this test, but i think this might make more sense as an integration test to verify the stage repos.
> 
> right now we're only introspecting / snapshotting the small RH repo and the EPEL 10 repo in CI, so this test would really only verify that small RH repo.
> 
> also the public repos' statuses aren't `Invalid` but they aren't `Valid` either because they haven't been introspected, and i'm not sure that we should introspect all of them in CI 🤔

previous card for this test has this acceptance criteria 

```A user with the snapshotting feature can list red hat repositories that are being snapshotted in the api & UI ```

```A user with the snapshotting feature can list red hat repository snapshots in the api & UI```

https://issues.redhat.com/browse/HMS-2487

Hey @swadeley any thoughts ?


### Comment by @swadeley on August 20, 2025 at 02:56 PM UTC

> > i may not be too clear on the goal of this test, but i think this might make more sense as an integration test to verify the stage repos.
> > right now we're only introspecting / snapshotting the small RH repo and the EPEL 10 repo in CI, so this test would really only verify that small RH repo.
> > also the public repos' statuses aren't `Invalid` but they aren't `Valid` either because they haven't been introspected, and i'm not sure that we should introspect all of them in CI 🤔
> 
> previous card for this test has this acceptance criteria
> 
> `A user with the snapshotting feature can list red hat repositories that are being snapshotted in the api & UI `
> 
> `A user with the snapshotting feature can list red hat repository snapshots in the api & UI`
> 
> https://issues.redhat.com/browse/HMS-2487
> 
> Hey @swadeley any thoughts ?

Hi, the Jira says "A user with the snapshotting feature can list", so we are testing that the user has permissions and tools to do things rather than testing the repos themselves. 
We can test against stage as we don't mind where the repos are, we just need to test the user has the ability to use the API and UI created in HMS-2243 to list repo's snapshots. 
We only need to test one RHEL and one custom repo as we are not testing the validity of the repos or the backend code.
It makes sense to have the test in the Integration project so it can run nightly and just test repos that already exist in stage without per-test setup and to avoid introspection in CI. 

### Comment by @mayurilahane on September 02, 2025 at 04:32 AM UTC

closing this because raised it here - https://github.com/content-services/content-sources-frontend/pull/633


---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1175*

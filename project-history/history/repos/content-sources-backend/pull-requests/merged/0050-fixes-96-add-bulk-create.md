---
type: pull_request
number: 50
title: "Fixes 96: Add bulk create"
state: merged
author: rverdile
created: 2022-07-12T20:40:05Z
updated: 2022-07-25T07:54:38Z
closed: 2022-07-25T07:54:38Z
merged: 2022-07-25T07:54:38Z
base_branch: main
head_branch: bulk-create
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/50
---

# Pull Request #50: Fixes 96: Add bulk create

**Author**: @rverdile
**Created**: July 12, 2022 at 08:40 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `bulk-create`

## Description

Adds bulk create

**Example:**
POST `/api/content_sources/v1/repositories/bulk_create`

Body:
```
[
  {
    "name": "test1",
    "url": "test1",
  },
  {
    "name": "test2",
    "url": "test2",
  }
]
```

Response:
```
[
    {
       "error": nil,
       "repository": {...repo}
    },
    
    {
    
       "error":nil,
       "repository": {...repo}
    }, ...
 ]

OR when there's an error

[
  {
       "error": "error",
       "repository": nil,
    },
{
       "error": nil,
       "repository": nil,
    },

]
```

Still needs

- [x] Handler tests
- [x] Fix error response


---

## Discussion

### Comment by @jlsherrill on July 12, 2022 at 08:45 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-96

### Comment by @jlsherrill on July 12, 2022 at 09:14 PM UTC

The changes I made to the tests to delete the data from the 3 tables in my
rh repo load pr will probably solve the transaction errors.  The order of
deletion was wrong

On Tue, Jul 12, 2022, 4:40 PM rverdile ***@***.***> wrote:

> Works in practice, but can't get tests to work yet because of some weird
> db transaction errors. Gonna make a draft PR in case I need help fixing :)
>
> I'll update this description with more stuff once I fix that.
> ------------------------------
> You can view, comment on, or merge this pull request online at:
>
>   https://github.com/content-services/content-sources-backend/pull/50
> Commit Summary
>
>    - 6ec5a05
>    <https://github.com/content-services/content-sources-backend/pull/50/commits/6ec5a0582788eb95408c660fe5ad236f93531474>
>    Fixes 96: Add bulk create
>
> File Changes
>
> (7 files
> <https://github.com/content-services/content-sources-backend/pull/50/files>
> )
>
>    - *M* pkg/dao/error.go
>    <https://github.com/content-services/content-sources-backend/pull/50/files#diff-aab5a6ff6e60a8adb3c68abe7a6e5f77894b49722ac6e6e00ed31a54baa75fd6>
>    (6)
>    - *M* pkg/dao/interfaces.go
>    <https://github.com/content-services/content-sources-backend/pull/50/files#diff-2890cf291694326df603760e3664720654eedef0b24b2a2db407c6d1c8c4da9f>
>    (1)
>    - *M* pkg/dao/repository.go
>    <https://github.com/content-services/content-sources-backend/pull/50/files#diff-4613b72907d5407f2d396f8946b7744d5e95d1d8c57545a7fbcf4fecace55a82>
>    (107)
>    - *M* pkg/dao/repository_test.go
>    <https://github.com/content-services/content-sources-backend/pull/50/files#diff-d117c6646c5fd3154993cd1578e3335cd07dc6260dfd2fb36841c3ca400da1ce>
>    (39)
>    - *M* pkg/dao/suite_test.go
>    <https://github.com/content-services/content-sources-backend/pull/50/files#diff-3765cee86bf193a9d472329749d34f6c068829023e610e7da26837e441a703a1>
>    (4)
>    - *M* pkg/handler/repositories.go
>    <https://github.com/content-services/content-sources-backend/pull/50/files#diff-5ea0664036ce68d4fb37e1a290f05df4868fb8b9cdff39f03d0e4fad2e9e0375>
>    (45)
>    - *M* pkg/handler/repository_test.go
>    <https://github.com/content-services/content-sources-backend/pull/50/files#diff-a8dd18992b56849843b60650d1e482fa51dd8f739a6fab75fd3591b6e18fd603>
>    (10)
>
> Patch Links:
>
>    -
>    https://github.com/content-services/content-sources-backend/pull/50.patch
>    -
>    https://github.com/content-services/content-sources-backend/pull/50.diff
>
> —
> Reply to this email directly, view it on GitHub
> <https://github.com/content-services/content-sources-backend/pull/50>, or
> unsubscribe
> <https://github.com/notifications/unsubscribe-auth/AADAORNRILUCRWTUIN6UV4DVTXJ3DANCNFSM53MJNGQQ>
> .
> You are receiving this because you are subscribed to this thread.Message
> ID: ***@***.***>
>


### Comment by @rverdile on July 13, 2022 at 04:19 PM UTC

I fixed that here too. I think it has something to do with there being a transaction both in the test and in the dao layer.

### Comment by @jlsherrill on July 14, 2022 at 07:05 PM UTC

ahh, maybe try nested transactions? https://gorm.io/docs/transactions.html#Nested-Transactions

### Comment by @rverdile on July 15, 2022 at 07:00 PM UTC

Marking as ready for review, but I don't anticipate the transaction stuff is in its final form.

### Comment by @jlsherrill on July 18, 2022 at 03:31 PM UTC

One thing to note, is the card describes a response that looks like this:
{ "RepoName": {"error": "Error Found"}, "2nd repo name": {"error":nil}}

but the current pr doesn't follow that.  You might wanna check with @Andrewgdewar as he and i discussed this previously. 

We could also return the repo that was created as part of it within that structure:

{ "RepoName": {"error": "Error encountered", "repository": nil}, 
"2nd repo name": {"error":nil, "repository": {"uuid":"abcd"...}}}




### Comment by @rverdile on July 18, 2022 at 05:37 PM UTC

The example on the card shows one repo in the response, which is why I misunderstood. Anyway, it does make more sense to return all the errors at once. I'll fix the PR.

### Comment by @rverdile on July 18, 2022 at 05:46 PM UTC

Also, should it be an array? 
```
[
  {  "repoName1": 
   {
       "error": "error1",
        "repo": nil
    }
  },
  {  "repoName2": 
   {
       "error": nil,
        "repo": {repo...}
    }
  },
]
```

And to clarify, this would also return the successful repos on error? Not just the ones that errored?

### Comment by @jlsherrill on July 18, 2022 at 06:33 PM UTC

repo names are supposed to be unique, so i think just  just one big hash where the name is the key is fine, you may want to preemptively make sure that no duplicates are passed in though and reject the request outright.  IDK @Andrewgdewar your thoughts?

### Comment by @jlsherrill on July 18, 2022 at 06:42 PM UTC

Built and pushed image for 5f7208ad3d802c4c0fd34681b31421766e388de4

### Comment by @rverdile on July 19, 2022 at 03:47 PM UTC

ignore the push this doesn't work yet

### Comment by @jlsherrill on July 20, 2022 at 03:47 PM UTC

Built and pushed image for 7da566d8d8890af896b360f3324501eda3e47610

### Comment by @rverdile on July 20, 2022 at 04:03 PM UTC

should be all updated now

### Comment by @jlsherrill on July 20, 2022 at 04:09 PM UTC

Built and pushed image for 0c4064ad456845d0a94b94c5ee732a7357ca3fc4

### Comment by @rverdile on July 20, 2022 at 07:30 PM UTC

new response?

```
[
    {
       "error": nil,
       "repository": {...repo}
    },
    
    {
    
       "error":nil,
       "repository": {...repo}
    }, ...
 ]

OR when there's an error

[
  {
       "error": "error",
       "repository": nil,
    },
]

```
where arrays is in same order as requested

### Comment by @jlsherrill on July 21, 2022 at 04:42 PM UTC

Built and pushed image for 910fca4a69f1b209e4991d9ec263cab7d437d295

### Comment by @jlsherrill on July 22, 2022 at 08:21 PM UTC

Built and pushed image for 0b2b0940b480543cda033bb27cd468d521fd0913

---

## Reviews

### Review by @rverdile - Commented on July 15, 2022 at 06:56 PM UTC

### Review by @jlsherrill - Commented on July 15, 2022 at 07:34 PM UTC

### Review by @jlsherrill - Commented on July 15, 2022 at 07:40 PM UTC

### Review by @jlsherrill - Commented on July 15, 2022 at 07:40 PM UTC

### Review by @jlsherrill - Commented on July 21, 2022 at 05:26 PM UTC

### Review by @jlsherrill - Approved on July 21, 2022 at 05:26 PM UTC

### Review by @rverdile - Commented on July 21, 2022 at 08:38 PM UTC

### Review by @swadeley - Approved on July 25, 2022 at 07:54 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/50*

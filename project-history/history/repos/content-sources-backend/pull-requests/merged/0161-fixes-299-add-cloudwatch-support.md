---
type: pull_request
number: 161
title: "Fixes 299: add cloudwatch support"
state: merged
author: jlsherrill
created: 2022-12-21T20:46:30Z
updated: 2023-01-06T15:56:22Z
closed: 2023-01-06T15:56:19Z
merged: 2023-01-06T15:56:19Z
base_branch: main
head_branch: 299
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/161
---

# Pull Request #161: Fixes 299: add cloudwatch support

**Author**: @jlsherrill
**Created**: December 21, 2022 at 08:46 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `299`

## Description

TODO:

* [x] Integrate with app-interface/clowder


Note, I tested this with cloudwatch directly with:

```
CLOUDWATCH_REGION=us-east-1  \
	CLOUDWATCH_GROUP=ContentSources \
        CLOUDWATCH_STREAM=test \
	CLOUDWATCH_KEY="somekey" \
	CLOUDWATCH_SECRET=somesecret \
	go run cmd/content-sources/main.go api consumer
```

you may need to create a cloudwatch group & stream, as well as a user that has access with the given key and secret.

---

## Discussion

### Comment by @jlsherrill on December 21, 2022 at 09:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-299

### Comment by @jlsherrill on January 05, 2023 at 09:21 PM UTC

I don't think so, we'll have to confirm clowdet support is good once this
goes to stage

On Thu, Jan 5, 2023, 4:10 PM rverdile ***@***.***> wrote:

> ***@***.**** commented on this pull request.
>
> Looks like it works to me. I created a cloudwatch group and stream in my
> aws account, added the configuration, and then I could see the server
> requests being logged to the stream.
>
> Anything else to think about here?
>
> —
> Reply to this email directly, view it on GitHub
> <https://github.com/content-services/content-sources-backend/pull/161#pullrequestreview-1238052479>,
> or unsubscribe
> <https://github.com/notifications/unsubscribe-auth/AADAORNTS7IBHURH5B7O7RLWQ42EXANCNFSM6AAAAAATGATADM>
> .
> You are receiving this because you authored the thread.Message ID:
> <content-services/content-sources-backend/pull/161/review/1238052479@
> github.com>
>


---

## Reviews

### Review by @jlsherrill - Commented on January 04, 2023 at 05:56 PM UTC

### Review by @rverdile - Commented on January 05, 2023 at 09:10 PM UTC

Looks like it works to me. I created a cloudwatch group and stream in my aws account, added the configuration, and then I could see the server requests being logged to the stream.

Anything else to think about here?

### Review by @rverdile - Approved on January 05, 2023 at 09:53 PM UTC

ack

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/161*

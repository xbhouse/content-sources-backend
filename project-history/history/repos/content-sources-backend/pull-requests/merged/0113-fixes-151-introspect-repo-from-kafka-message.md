---
type: pull_request
number: 113
title: "Fixes 151: introspect repo from kafka message"
state: merged
author: avisiedo
created: 2022-09-26T13:05:25Z
updated: 2022-10-31T18:00:28Z
closed: 2022-10-31T17:32:23Z
merged: 2022-10-31T17:32:23Z
base_branch: main
head_branch: hmscontent-151-kafka-introspect
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/113
---

# Pull Request #113: Fixes 151: introspect repo from kafka message

**Author**: @avisiedo
**Created**: September 26, 2022 at 01:05 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `hmscontent-151-kafka-introspect`

## Description

- Add launch.json to debug easier from vscode.
- Add kafka-introspect command to consume kafka message to introspect repositories.
- Add initial kafka support.
- Generate code for the kafka message schemas: `make build` regenerate `.json` and `.gen.go` files when the schema change.
- Update dependencies.
- Update kafka make rules to use `kafkacat` to produce a test message.
- 

TODO:
* [x] Fix test failures.
* [x] Update documentation.
* [x] Clean-up code.
* [x] Try ephemeral default pool pool
* [x] Try ephemeral managed-kafka pool
* [x] Try ephemeral real-managed-kafka pool
* [x] Log information about the received message.
* [x] Fix multi arch base image for kafka broker and zookeeper.
* [x] Add unit tests.
* [x] Add comments to public function and methods
* [x] Move kafkaconfigmap configuration in an external file (not sure the reason I wrote this item; I don't see right now the need). The reason was to have a better control on the client library parameters; the intention is not provide all the parameters, but externalize them allow to quickly update config, restart infra, and try, if some client parameter is necessary to be customized. Added card for this particular item to address in the future: https://issues.redhat.com/browse/HMSCONTENT-236


---

## Discussion

### Comment by @jlsherrill on September 26, 2022 at 01:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-151

### Comment by @avisiedo on September 28, 2022 at 08:24 AM UTC

First success run in ephemeral environment (using default pool):

```raw
{"level":"debug","time":"2022-09-28T08:15:48Z","message":"OnMessage was called; Key=871a77ca-f216-4a76-af30-2d76e1da04fe"}
{"level":"debug","time":"2022-09-28T08:15:48Z","message":"Introspecting http://mirror.stream.centos.org/9-stream/AppStream/x86_64/os/"}
{"level":"debug","time":"2022-09-28T08:17:27Z","message":"IntrospectionUrl returned 13106 new packages"}
```

### Comment by @avisiedo on September 28, 2022 at 09:16 AM UTC

Similar result in pool `managed-kafka` in ephemeral environment after fix some typo:

```raw
{"level":"debug","time":"2022-09-28T09:12:45Z","message":"OnMessage was called; Key=5b707a6c-c4fa-41f0-b8a8-859da236f6d3"}
{"level":"debug","time":"2022-09-28T09:12:45Z","message":"Introspecting http://mirror.stream.centos.org/9-stream/AppStream/x86_64/os/"}
{"level":"debug","time":"2022-09-28T09:14:37Z","message":"IntrospectionUrl returned 13106 new packages"}
```

### Comment by @jlsherrill on October 03, 2022 at 08:29 PM UTC

* We'll need to work with bulkCreate and update as well (if url changes)
* When testing with the local UI, i get an error when trying to send the message:  
```
4:24PM WRN error producing kafka message: expected a value for 'X-Rh-Insights-Request-Id' http header
```

I'm guessing that in this particular case that header isn't present.  We should probably not make it required

### Comment by @rverdile on October 04, 2022 at 06:43 PM UTC

There's a lot of "TODO" comments, are you gonna address all of them on this PR?

### Comment by @avisiedo on October 05, 2022 at 12:00 PM UTC

thanks @rverdile, about the TODO comments I want to convert in cards; as I am updating with the reviews and unit tests I am adding them; before to change from draft to ready I will remove all them, but they are placeholders for later.

### Comment by @avisiedo on October 05, 2022 at 12:24 PM UTC

thanks @jlsherrill ! I have updated the code to produce the messages when `bulk_create` is called; making the changes for when the url is updated too.

### Comment by @avisiedo on October 05, 2022 at 01:07 PM UTC

Rebased and fix introspectURL invokation; currently forcing the introspection when the message is consumed.

### Comment by @avisiedo on October 10, 2022 at 11:23 AM UTC

/retest

### Comment by @jlsherrill on October 12, 2022 at 07:22 PM UTC

Overall i this is shaping up.  I think just some cleanup and resolving the TODOs  (either via jira cards, in which case we should remove the TODO comment, deciding they aren't worth doing, or actually doing them).  

### Comment by @avisiedo on October 18, 2022 at 09:55 AM UTC

/retest

### Comment by @avisiedo on October 20, 2022 at 08:35 AM UTC

update: code from http handlers has been moved to pkg/event/producer

reviewing and updating the review messages; thanks mates!

### Comment by @avisiedo on October 21, 2022 at 12:08 PM UTC

Several updates:

* Updated the pending comments.
* Updated the pending reviews.
* Reviewed the todos; some of them updated, other translated into new cards to be considered.

Thanks for all your comments and reviews: @rverdile @swadeley @jlsherrill 

### Comment by @avisiedo on October 25, 2022 at 08:44 AM UTC

update: rebased with the last changes from the `main` branch

### Comment by @avisiedo on October 25, 2022 at 04:00 PM UTC

Last todo in this PR is addressed on this card: https://issues.redhat.com/browse/HMSCONTENT-236 

Checking the todo item.

And rebasing main branch.

### Comment by @jlsherrill on October 26, 2022 at 06:16 PM UTC

[test]

### Comment by @jlsherrill on October 27, 2022 at 09:04 PM UTC

Question there about the replicas, and then i think we can go ahead and merge this.

---

## Reviews

### Review by @avisiedo - Commented on September 28, 2022 at 11:53 AM UTC

### Review by @avisiedo - Commented on September 28, 2022 at 11:55 AM UTC

### Review by @avisiedo - Commented on September 28, 2022 at 12:14 PM UTC

### Review by @avisiedo - Commented on September 28, 2022 at 04:28 PM UTC

### Review by @avisiedo - Commented on September 28, 2022 at 06:34 PM UTC

### Review by @rverdile - Commented on October 04, 2022 at 06:34 PM UTC

### Review by @rverdile - Commented on October 04, 2022 at 06:38 PM UTC

### Review by @avisiedo - Commented on October 05, 2022 at 12:00 PM UTC

### Review by @rverdile - Commented on October 05, 2022 at 03:21 PM UTC

### Review by @rverdile - Commented on October 05, 2022 at 04:03 PM UTC

### Review by @rverdile - Commented on October 05, 2022 at 04:06 PM UTC

### Review by @rverdile - Commented on October 05, 2022 at 05:52 PM UTC

### Review by @avisiedo - Commented on October 10, 2022 at 11:22 AM UTC

### Review by @rverdile - Commented on October 11, 2022 at 02:00 PM UTC

### Review by @avisiedo - Commented on October 11, 2022 at 02:15 PM UTC

### Review by @avisiedo - Commented on October 11, 2022 at 02:18 PM UTC

### Review by @avisiedo - Commented on October 11, 2022 at 02:19 PM UTC

### Review by @jlsherrill - Commented on October 12, 2022 at 03:49 PM UTC

### Review by @jlsherrill - Commented on October 12, 2022 at 04:04 PM UTC

### Review by @jlsherrill - Commented on October 12, 2022 at 04:05 PM UTC

### Review by @jlsherrill - Commented on October 12, 2022 at 04:44 PM UTC

### Review by @jlsherrill - Commented on October 12, 2022 at 05:02 PM UTC

### Review by @jlsherrill - Commented on October 12, 2022 at 05:06 PM UTC

### Review by @jlsherrill - Commented on October 12, 2022 at 05:06 PM UTC

### Review by @jlsherrill - Commented on October 12, 2022 at 05:06 PM UTC

### Review by @jlsherrill - Commented on October 12, 2022 at 05:11 PM UTC

### Review by @jlsherrill - Commented on October 12, 2022 at 05:35 PM UTC

### Review by @avisiedo - Commented on October 13, 2022 at 01:59 PM UTC

### Review by @avisiedo - Commented on October 13, 2022 at 02:04 PM UTC

### Review by @avisiedo - Commented on October 13, 2022 at 02:05 PM UTC

### Review by @avisiedo - Commented on October 13, 2022 at 04:10 PM UTC

### Review by @avisiedo - Commented on October 13, 2022 at 05:44 PM UTC

### Review by @avisiedo - Commented on October 13, 2022 at 05:48 PM UTC

### Review by @avisiedo - Commented on October 13, 2022 at 05:54 PM UTC

### Review by @avisiedo - Commented on October 18, 2022 at 07:47 PM UTC

### Review by @swadeley - Commented on October 19, 2022 at 01:58 PM UTC

some suggestions

### Review by @avisiedo - Commented on October 20, 2022 at 08:46 AM UTC

### Review by @avisiedo - Commented on October 20, 2022 at 09:32 AM UTC

### Review by @avisiedo - Commented on October 20, 2022 at 09:35 AM UTC

### Review by @avisiedo - Commented on October 20, 2022 at 10:03 AM UTC

### Review by @avisiedo - Commented on October 21, 2022 at 11:28 AM UTC

### Review by @avisiedo - Commented on October 21, 2022 at 11:33 AM UTC

### Review by @avisiedo - Commented on October 21, 2022 at 11:45 AM UTC

### Review by @avisiedo - Commented on October 21, 2022 at 12:07 PM UTC

### Review by @rverdile - Commented on October 21, 2022 at 03:41 PM UTC

### Review by @rverdile - Commented on October 21, 2022 at 03:41 PM UTC

### Review by @avisiedo - Commented on October 21, 2022 at 05:51 PM UTC

### Review by @avisiedo - Commented on October 21, 2022 at 05:51 PM UTC

### Review by @swadeley - Commented on October 24, 2022 at 08:23 AM UTC

### Review by @swadeley - Commented on October 24, 2022 at 08:25 AM UTC

### Review by @swadeley - Commented on October 24, 2022 at 08:26 AM UTC

### Review by @swadeley - Commented on October 24, 2022 at 08:27 AM UTC

### Review by @swadeley - Commented on October 24, 2022 at 08:28 AM UTC

### Review by @swadeley - Commented on October 24, 2022 at 08:29 AM UTC

### Review by @swadeley - Commented on October 24, 2022 at 08:30 AM UTC

### Review by @swadeley - Commented on October 24, 2022 at 08:30 AM UTC

### Review by @swadeley - Commented on October 24, 2022 at 08:32 AM UTC

### Review by @swadeley - Commented on October 24, 2022 at 08:33 AM UTC

### Review by @swadeley - Commented on October 24, 2022 at 08:35 AM UTC

### Review by @jlsherrill - Commented on October 27, 2022 at 01:28 PM UTC

### Review by @avisiedo - Commented on October 31, 2022 at 11:24 AM UTC

### Review by @jlsherrill - Approved on October 31, 2022 at 02:53 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/113*

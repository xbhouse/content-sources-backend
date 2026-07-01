---
type: pull_request
number: 219
title: "Fixes 360: use log level from config/env"
state: merged
author: rverdile
created: 2023-02-23T14:42:31Z
updated: 2023-02-28T14:41:35Z
closed: 2023-02-28T14:41:30Z
merged: 2023-02-28T14:41:30Z
base_branch: main
head_branch: log-level
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/219
---

# Pull Request #219: Fixes 360: use log level from config/env

**Author**: @rverdile
**Created**: February 23, 2023 at 02:42 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `log-level`

## Description

## Summary
The global log level was hard-coded as info, so the log-level set in config was not taking effect.

It looks like this bug came about from a misunderstanding on a previous pr (#152), that `SetGlobalLevel` is the default log level. I believe the bug in that PR was fixed by adding `LOGGING_LEVEL` to `deployment.yaml`, but not by hard-coding `SetGlobalLevel` to `info`.

## Testing steps
1. In your config and/or env, set the logging level to `info`.
2. `make run`, and you should see several messages logged to "info" in the terminal.
3.  Change your log level to `warn`.
4. `make run`, and you should no longer see those messages.
5. Double check my understanding of PR #152 and make sure I did not break what was fixed there. 

---

## Discussion

### Comment by @jlsherrill on February 23, 2023 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-360

---

## Reviews

### Review by @jlsherrill - Approved on February 23, 2023 at 08:14 PM UTC

ACK   Testing locally it all behaved as i'd expect.  Log level in config.yaml works, via env variable works as expected.  Tested in ephemeral env and see 404s properly

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/219*

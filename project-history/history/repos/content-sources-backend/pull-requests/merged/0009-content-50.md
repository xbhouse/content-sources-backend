---
type: pull_request
number: 9
title: "CONTENT-50"
state: merged
author: Andrewgdewar
created: 2022-05-18T17:40:04Z
updated: 2022-12-20T14:03:15Z
closed: 2022-05-26T17:26:24Z
merged: 2022-05-26T17:26:24Z
base_branch: main
head_branch: CONTENT-50
labels: []
url: https://github.com/content-services/content-sources-backend/pull/9
---

# Pull Request #9: CONTENT-50

**Author**: @Andrewgdewar
**Created**: May 18, 2022 at 05:40 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `CONTENT-50`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on May 25, 2022 at 02:47 PM UTC

can we somehow integrate in with https://github.com/rs/zerolog#pretty-logging  maybe via an Env variable?  

Running as is is really hard to read

### Comment by @Andrewgdewar on May 25, 2022 at 04:51 PM UTC

> can we somehow integrate in with https://github.com/rs/zerolog#pretty-logging maybe via an Env variable?
> 
> Running as is is really hard to read

I added the consoleWriter, there are a ton of config options there but left it with just a timestamp and the out of the box defaults.

Do we need the env variable? Would not whatever method of recording the logs server-side benefit from the added formatting?

### Comment by @jlsherrill on May 25, 2022 at 05:08 PM UTC

i guess its fine for now.  We can re-evaulate as we go along.  

I was getting a log for each request on your previous push, but now i'm not getting anything when curling /ping, any idea?

### Comment by @jlsherrill on May 25, 2022 at 05:18 PM UTC

You might also checkout the provisioning team's zerolog integration: https://github.com/RHEnVision/provisioning-backend/blob/c6f31b463defcd55861434ab42b74fdd8d6ddd91/internal/logging/zerolog.go

### Comment by @jlsherrill on May 25, 2022 at 05:27 PM UTC

and we should probably integrate with https://github.com/spf13/viper at some point, in a simialr way.  I'll add a card to do that to the board

### Comment by @Andrewgdewar on May 25, 2022 at 05:58 PM UTC

> 

I'm investigating now. 

---

## Reviews

### Review by @Andrewgdewar - Commented on May 25, 2022 at 02:37 PM UTC

### Review by @jlsherrill - Commented on May 25, 2022 at 02:47 PM UTC

### Review by @Andrewgdewar - Commented on May 25, 2022 at 04:54 PM UTC

### Review by @jlsherrill - Commented on May 26, 2022 at 01:14 PM UTC

### Review by @Andrewgdewar - Commented on May 26, 2022 at 02:49 PM UTC

### Review by @Andrewgdewar - Commented on May 26, 2022 at 03:13 PM UTC

### Review by @jlsherrill - Approved on May 26, 2022 at 04:47 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/9*

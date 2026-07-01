---
type: pull_request
number: 491
title: "Fixes 3189: change gpg key file endpoint"
state: merged
author: rverdile
created: 2023-11-29T20:03:33Z
updated: 2023-12-05T23:30:22Z
closed: 2023-12-05T23:10:10Z
merged: 2023-12-05T23:10:10Z
base_branch: main
head_branch: gpg-deploy
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/491
---

# Pull Request #491: Fixes 3189: change gpg key file endpoint

**Author**: @rverdile
**Created**: November 29, 2023 at 08:03 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `gpg-deploy`

## Description

## Summary
Changes the gpg key file endpoint to `repository_gpg_key/:uuid` so that it can be whitelisted i.e. skip auth
## Testing steps
1. Deploy PR in ephemeral
2. Create repository with GPG key and snapshot enabled
3. Use the GPG Key URL in the config.repo file to open a tab (in browser) to the GPG Key
4. Should be able to open the URL in your browser despite not having the correct auth

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on November 29, 2023 at 08:05 PM UTC

https://issues.redhat.com/browse/HMS-3189

### Comment by @jlsherrill on November 30, 2023 at 04:25 PM UTC

this still didn't seem to work in ephemeral: 

curl -v https://ee-9n5k9jye.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/api/content-sources/v1.0/repositories/be1c2024-f3de-42d3-bedd-4aeb7d4a1d17/gpg_key/

HTTP/1.1 403 Forbidden



### Comment by @rverdile on November 30, 2023 at 07:15 PM UTC

/retest

### Comment by @jlsherrill on December 04, 2023 at 02:11 PM UTC

/retest

### Comment by @jlsherrill on December 04, 2023 at 03:41 PM UTC

I'm getting 'Bad Request: missing x-rh-identity header'  I suspect our auth skipper middleware needs to be updated :)

But it looks like its getting based the gateway!

### Comment by @jlsherrill on December 04, 2023 at 03:42 PM UTC

also the URL in the config.repo file needs to be updated

### Comment by @jlsherrill on December 04, 2023 at 05:06 PM UTC

don't forget to regen the api doc :)

### Comment by @swadeley on December 05, 2023 at 07:17 AM UTC

Hi @rverdile , I put the setup steps in the Jira as its rather verbose.

I get 403 with the gpgkey URL.

### Comment by @swadeley on December 05, 2023 at 11:10 PM UTC

> Hi @rverdile , I put the setup steps in the Jira as its rather verbose.
> 
> I get 403 with the gpgkey URL.

Hi, as per @jlsherrill I ran deployment script again pointing to the local checked out PR branch to get the changes in deployment.yaml. Now it works.

---

## Reviews

### Review by @jlsherrill - Commented on November 30, 2023 at 05:08 PM UTC

### Review by @rverdile - Commented on November 30, 2023 at 05:32 PM UTC

### Review by @jlsherrill - Commented on December 04, 2023 at 02:11 PM UTC

### Review by @jlsherrill - Commented on December 04, 2023 at 04:21 PM UTC

### Review by @rverdile - Commented on December 04, 2023 at 04:39 PM UTC

### Review by @rverdile - Commented on December 04, 2023 at 04:39 PM UTC

### Review by @jlsherrill - Commented on December 04, 2023 at 04:43 PM UTC

### Review by @rverdile - Commented on December 04, 2023 at 04:46 PM UTC

### Review by @jlsherrill - Commented on December 04, 2023 at 04:49 PM UTC

### Review by @jlsherrill - Commented on December 04, 2023 at 04:49 PM UTC

### Review by @jlsherrill - Commented on December 04, 2023 at 04:49 PM UTC

### Review by @rverdile - Commented on December 04, 2023 at 04:51 PM UTC

### Review by @jlsherrill - Commented on December 04, 2023 at 05:23 PM UTC

### Review by @jlsherrill - Commented on December 04, 2023 at 05:24 PM UTC

### Review by @jlsherrill - Commented on December 04, 2023 at 05:25 PM UTC

### Review by @rverdile - Commented on December 04, 2023 at 05:47 PM UTC

### Review by @jlsherrill - Approved on December 04, 2023 at 05:49 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/491*

---
type: pull_request
number: 1230
title: "HMS-9382: update pulp client to support certificates"
state: merged
author: xbhouse
created: 2025-10-08T20:21:17Z
updated: 2025-10-21T17:45:21Z
closed: 2025-10-21T17:45:21Z
merged: 2025-10-21T17:45:21Z
base_branch: main
head_branch: pulp-with-certs
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1230
---

# Pull Request #1230: HMS-9382: update pulp client to support certificates

**Author**: @xbhouse
**Created**: October 08, 2025 at 08:21 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `pulp-with-certs`

## Description

## Summary

- Updates pulp client to use certs or basic auth if certs aren't specified
- Updates our local dev deployment to use an nginx proxy in front of the pulp_api service
- Adjusts the `make compose-up` command to generate the certs needed to test locally

## Testing steps

1. Run `make compose-clean` and `make compose-up` to generate the certs and start the proxy
2. Set `clients.pulp.server` to `https://localhost:8443`
3. Comment out `clients.pulp.username` and `clients.pulp.password`
4. Set `clients.pulp.client_cert_path`, `clients.pulp.client_key_path`, and `clients.pulp.ca_cert_path` to the paths to the generated certs in `compose_files/pulp/assets/certs/dev_certs`
5. Start the server and test snapshotting a repository multiple times, running the import and process-repos commands, creating and uploading content to an upload repository, and creating a template (basically test that things still work as expected and don't break :smile:)
6. Uncomment `clients.pulp.username` and `clients.pulp.password`
7. Restart the server, same functionality should still work


---

## Discussion

### Comment by @xbhouse on October 09, 2025 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-9382

### Comment by @jlsherrill on October 10, 2025 at 07:49 PM UTC

/retest

### Comment by @xbhouse on October 14, 2025 at 01:08 PM UTC

/retest

### Comment by @jlsherrill on October 17, 2025 at 01:19 PM UTC

/retest

### Comment by @xbhouse on October 20, 2025 at 12:28 PM UTC

/retest

### Comment by @xbhouse on October 20, 2025 at 07:48 PM UTC

/retest

### Comment by @xbhouse on October 21, 2025 at 12:31 PM UTC

/retest

### Comment by @xbhouse on October 21, 2025 at 02:28 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Commented on October 09, 2025 at 12:24 AM UTC

### Review by @jlsherrill - Commented on October 09, 2025 at 07:08 PM UTC

### Review by @xbhouse - Commented on October 09, 2025 at 07:23 PM UTC

### Review by @rverdile - Commented on October 13, 2025 at 04:36 PM UTC

### Review by @jlsherrill - Commented on October 15, 2025 at 03:20 PM UTC

### Review by @jlsherrill - Commented on October 15, 2025 at 03:35 PM UTC

### Review by @xbhouse - Commented on October 15, 2025 at 04:54 PM UTC

### Review by @jlsherrill - Commented on October 15, 2025 at 06:40 PM UTC

### Review by @xbhouse - Commented on October 15, 2025 at 06:46 PM UTC

### Review by @jlsherrill - Commented on October 15, 2025 at 07:05 PM UTC

### Review by @xbhouse - Commented on October 15, 2025 at 08:52 PM UTC

### Review by @jlsherrill - Commented on October 15, 2025 at 09:26 PM UTC

### Review by @jlsherrill - Approved on October 15, 2025 at 09:27 PM UTC

ACK, nice work!!

### Review by @jlsherrill - Commented on October 16, 2025 at 01:18 AM UTC

### Review by @xbhouse - Commented on October 16, 2025 at 03:28 PM UTC

### Review by @jlsherrill - Commented on October 16, 2025 at 06:05 PM UTC

### Review by @xbhouse - Commented on October 16, 2025 at 06:11 PM UTC

### Review by @jlsherrill - Commented on October 16, 2025 at 06:39 PM UTC

### Review by @jlsherrill - Commented on October 16, 2025 at 07:01 PM UTC

### Review by @xbhouse - Commented on October 16, 2025 at 07:58 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1230*

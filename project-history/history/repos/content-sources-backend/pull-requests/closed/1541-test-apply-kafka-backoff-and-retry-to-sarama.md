---
type: pull_request
number: 1541
title: "Test: Apply kafka backoff and retry to sarama"
state: closed
author: swadeley
created: 2026-06-17T11:08:46Z
updated: 2026-06-18T09:15:38Z
closed: 2026-06-18T09:15:38Z
base_branch: main
head_branch: swadeley/backoff_and_retry_in_sarama_config
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1541
---

# Pull Request #1541: Test: Apply kafka backoff and retry to sarama

**Author**: @swadeley
**Created**: June 17, 2026 at 11:08 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/backoff_and_retry_in_sarama_config`

## Description

## Summary
 Apply kafka backoff and retry to sarama

Kafka retry, backoff, and required-acks from config were defined but never applied in GetSaramaConfig(). This applies them so template and notification producers retry transient broker failures.

SendTemplateEvent now returns an error on send failure so callers can detect and report failed events instead of silently continuing.

## Testing steps

tests pass



---

## Discussion

### Comment by @swadeley on June 18, 2026 at 08:34 AM UTC

Hi, closing this until some Kafka producer experiences overloading.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1541*

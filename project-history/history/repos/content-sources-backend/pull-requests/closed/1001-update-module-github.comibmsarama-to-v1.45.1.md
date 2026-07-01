---
type: pull_request
number: 1001
title: "Update module github.com/IBM/sarama to v1.45.1"
state: closed
author: red-hat-konflux
created: 2025-03-02T18:27:49Z
updated: 2025-10-07T12:16:25Z
closed: 2025-03-03T13:48:26Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-ibm-sarama-1.x
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1001
---

# Pull Request #1001: Update module github.com/IBM/sarama to v1.45.1

**Author**: @red-hat-konflux
**Created**: March 02, 2025 at 06:27 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-ibm-sarama-1.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [github.com/IBM/sarama](https://redirect.github.com/IBM/sarama) | require | patch | `v1.45.0` -> `v1.45.1` |

---

### Release Notes

<details>
<summary>IBM/sarama (github.com/IBM/sarama)</summary>

### [`v1.45.1`](https://redirect.github.com/IBM/sarama/releases/tag/v1.45.1): Version 1.45.1 (2025-03-02)

[Compare Source](https://redirect.github.com/IBM/sarama/compare/v1.45.0...v1.45.1)

<!-- Release notes generated using configuration in .github/release.yaml at main -->

#### What's Changed

##### :tada: New Features / Improvements

-   feat(producer): add MaxBufferBytes to limit retry buffer size by [@&#8203;wanwenli](https://redirect.github.com/wanwenli) in [https://github.com/IBM/sarama/pull/3088](https://redirect.github.com/IBM/sarama/pull/3088)
-   feat(producer): add sync pool for channel reuse by [@&#8203;kasimtj](https://redirect.github.com/kasimtj) in [https://github.com/IBM/sarama/pull/3109](https://redirect.github.com/IBM/sarama/pull/3109)
-   feat: exponential backoff for clients (KIP-580) by [@&#8203;wanwenli](https://redirect.github.com/wanwenli) in [https://github.com/IBM/sarama/pull/3099](https://redirect.github.com/IBM/sarama/pull/3099)

##### :bug: Fixes

-   fix(sasl): add nilguard around token to prevent panic by [@&#8203;hoo47](https://redirect.github.com/hoo47) in [https://github.com/IBM/sarama/pull/3076](https://redirect.github.com/IBM/sarama/pull/3076)
-   fix(test): consumer group fetch request messages by [@&#8203;stsmurf](https://redirect.github.com/stsmurf) in [https://github.com/IBM/sarama/pull/3081](https://redirect.github.com/IBM/sarama/pull/3081)
-   fix: remove redundant nil check by [@&#8203;knbr13](https://redirect.github.com/knbr13) in [https://github.com/IBM/sarama/pull/3089](https://redirect.github.com/IBM/sarama/pull/3089)
-   fix(consumer): add recovery from no leader partitions by [@&#8203;liutao365](https://redirect.github.com/liutao365) in [https://github.com/IBM/sarama/pull/3101](https://redirect.github.com/IBM/sarama/pull/3101)
-   produce: set MaxTimestamp by [@&#8203;rockwotj](https://redirect.github.com/rockwotj) in [https://github.com/IBM/sarama/pull/3108](https://redirect.github.com/IBM/sarama/pull/3108)

##### :package: Dependency updates

-   chore(deps): bump go.opentelemetry.io/otel from 1.24.0 to 1.29.0 in /examples/interceptors by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/IBM/sarama/pull/3071](https://redirect.github.com/IBM/sarama/pull/3071)
-   chore(deps): bump the otel group across 1 directory with 2 updates by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/IBM/sarama/pull/3072](https://redirect.github.com/IBM/sarama/pull/3072)
-   chore(deps): bump the golang-x group across 1 directory with 2 updates by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/IBM/sarama/pull/3098](https://redirect.github.com/IBM/sarama/pull/3098)

##### :wrench: Maintenance

-   chore(deps): prevent otel upgrades for now by [@&#8203;dnwe](https://redirect.github.com/dnwe) in [https://github.com/IBM/sarama/pull/3069](https://redirect.github.com/IBM/sarama/pull/3069)
-   chore: add version constant for kafka 3.7.2 by [@&#8203;dnwe](https://redirect.github.com/dnwe) in [https://github.com/IBM/sarama/pull/3073](https://redirect.github.com/IBM/sarama/pull/3073)
-   chore(ci): fetch kafka 4.0 via tar.gz rather than git by [@&#8203;dnwe](https://redirect.github.com/dnwe) in [https://github.com/IBM/sarama/pull/3079](https://redirect.github.com/IBM/sarama/pull/3079)
-   fix(ci): tighten up github workflows by [@&#8203;dnwe](https://redirect.github.com/dnwe) in [https://github.com/IBM/sarama/pull/3080](https://redirect.github.com/IBM/sarama/pull/3080)
-   chore(ci): analyse actions in codeql by [@&#8203;dnwe](https://redirect.github.com/dnwe) in [https://github.com/IBM/sarama/pull/3085](https://redirect.github.com/IBM/sarama/pull/3085)
-   chore(ci): bump golangci-lint version to v1.63.4 by [@&#8203;dnwe](https://redirect.github.com/dnwe) in [https://github.com/IBM/sarama/pull/3090](https://redirect.github.com/IBM/sarama/pull/3090)
-   feat(ci): add dedicated staticcheck run by [@&#8203;dnwe](https://redirect.github.com/dnwe) in [https://github.com/IBM/sarama/pull/3091](https://redirect.github.com/IBM/sarama/pull/3091)

#### New Contributors

-   [@&#8203;hoo47](https://redirect.github.com/hoo47) made their first contribution in [https://github.com/IBM/sarama/pull/3076](https://redirect.github.com/IBM/sarama/pull/3076)
-   [@&#8203;stsmurf](https://redirect.github.com/stsmurf) made their first contribution in [https://github.com/IBM/sarama/pull/3081](https://redirect.github.com/IBM/sarama/pull/3081)
-   [@&#8203;knbr13](https://redirect.github.com/knbr13) made their first contribution in [https://github.com/IBM/sarama/pull/3089](https://redirect.github.com/IBM/sarama/pull/3089)
-   [@&#8203;liutao365](https://redirect.github.com/liutao365) made their first contribution in [https://github.com/IBM/sarama/pull/3101](https://redirect.github.com/IBM/sarama/pull/3101)
-   [@&#8203;rockwotj](https://redirect.github.com/rockwotj) made their first contribution in [https://github.com/IBM/sarama/pull/3108](https://redirect.github.com/IBM/sarama/pull/3108)
-   [@&#8203;kasimtj](https://redirect.github.com/kasimtj) made their first contribution in [https://github.com/IBM/sarama/pull/3109](https://redirect.github.com/IBM/sarama/pull/3109)

**Full Changelog**: https://github.com/IBM/sarama/compare/v1.45.0...v1.45.1

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "after 5am on sunday" (UTC), Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ1cGRhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ0YXJnZXRCcmFuY2giOiJtYWluIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @red-hat-konflux on March 02, 2025 at 06:27 PM UTC

### ℹ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 5 additional dependencies were updated


Details:


| **Package**           | **Change**             |
| :-------------------- | :--------------------- |
| `golang.org/x/sync`   | `v0.10.0` -> `v0.11.0` |
| `golang.org/x/crypto` | `v0.32.0` -> `v0.33.0` |
| `golang.org/x/net`    | `v0.34.0` -> `v0.35.0` |
| `golang.org/x/sys`    | `v0.29.0` -> `v0.30.0` |
| `golang.org/x/text`   | `v0.21.0` -> `v0.22.0` |

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1001*

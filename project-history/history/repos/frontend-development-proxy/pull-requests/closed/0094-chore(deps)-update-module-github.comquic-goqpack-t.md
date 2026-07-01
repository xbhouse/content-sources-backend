---
type: pull_request
number: 94
title: "chore(deps): update module github.com/quic-go/qpack to v0.6.0 - autoclosed"
state: closed
author: red-hat-konflux
created: 2025-12-07T20:32:35Z
updated: 2026-01-07T08:42:40Z
closed: 2026-01-07T08:42:38Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-quic-go-qpack-0.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/94
---

# Pull Request #94: chore(deps): update module github.com/quic-go/qpack to v0.6.0 - autoclosed

**Author**: @red-hat-konflux
**Created**: December 07, 2025 at 08:32 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-quic-go-qpack-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/quic-go/qpack](https://redirect.github.com/quic-go/qpack) | `v0.5.1` -> `v0.6.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fquic-go%2fqpack/v0.6.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fquic-go%2fqpack/v0.5.1/v0.6.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>quic-go/qpack (github.com/quic-go/qpack)</summary>

### [`v0.6.0`](https://redirect.github.com/quic-go/qpack/releases/tag/v0.6.0)

[Compare Source](https://redirect.github.com/quic-go/qpack/compare/v0.5.1...v0.6.0)

This release updates the `Decoder` API to decode header fields incrementally rather than emitting them all at once ([#&#8203;67](https://redirect.github.com/quic-go/qpack/issues/67)). This cuts allocations, improves performance, and enables consumers to process fields as they are decoded.

#### What's Changed

- ci: update golangci-lint to v2, run it on both Go versions by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;57](https://redirect.github.com/quic-go/qpack/pull/57)
- bump Go version to 1.23 in go.mod, update CI config by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;56](https://redirect.github.com/quic-go/qpack/pull/56)
- update to Go 1.25, drop Go 1.23 by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;59](https://redirect.github.com/quic-go/qpack/pull/59)
- Bump actions/setup-go from 5 to 6 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;60](https://redirect.github.com/quic-go/qpack/pull/60)
- Bump actions/checkout from 4 to 5 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;58](https://redirect.github.com/quic-go/qpack/pull/58)
- ci: remove 386 (32 bit x86) by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;61](https://redirect.github.com/quic-go/qpack/pull/61)
- ci: enable ClusterFuzzLite PR fuzzing by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;62](https://redirect.github.com/quic-go/qpack/pull/62)
- Bump golangci/golangci-lint-action from 8 to 9 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;63](https://redirect.github.com/quic-go/qpack/pull/63)
- ci: fix Codecov configuration by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;65](https://redirect.github.com/quic-go/qpack/pull/65)
- add benchmarks for the decoder by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;64](https://redirect.github.com/quic-go/qpack/pull/64)
- ci: disable fail-fast for all matrix jobs by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;68](https://redirect.github.com/quic-go/qpack/pull/68)
- example: fix input file parsing by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;69](https://redirect.github.com/quic-go/qpack/pull/69)
- interop: update qifs to [`da52cd9`](https://redirect.github.com/quic-go/qpack/commit/da52cd9) by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;70](https://redirect.github.com/quic-go/qpack/pull/70)
- simplify testing setup by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;71](https://redirect.github.com/quic-go/qpack/pull/71)
- rework Decoder API to decode header fields one by one by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;67](https://redirect.github.com/quic-go/qpack/pull/67)
- use io.ErrUnexpectedEOF error when input is too short by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;73](https://redirect.github.com/quic-go/qpack/pull/73)
- README: improve description and fix interop command by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;74](https://redirect.github.com/quic-go/qpack/pull/74)

**Full Changelog**: <https://github.com/quic-go/qpack/compare/v0.5.1...v0.6.0>

</details>

---

### Configuration

📅 **Schedule**: Branch creation - At any time (no schedule defined), Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/94*

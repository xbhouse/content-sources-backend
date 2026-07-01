---
type: pull_request
number: 2083
title: "Update module github.com/gin-gonic/gin to v1.12.0 - autoclosed"
state: closed
author: red-hat-konflux
created: 2026-03-02T05:22:59Z
updated: 2026-03-04T09:11:37Z
closed: 2026-03-04T09:11:34Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-gin-gonic-gin-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2083
---

# Pull Request #2083: Update module github.com/gin-gonic/gin to v1.12.0 - autoclosed

**Author**: @red-hat-konflux
**Created**: March 02, 2026 at 05:22 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-gin-gonic-gin-1.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/gin-gonic/gin](https://redirect.github.com/gin-gonic/gin) | `v1.11.0` -> `v1.12.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgin-gonic%2fgin/v1.12.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgin-gonic%2fgin/v1.11.0/v1.12.0?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>gin-gonic/gin (github.com/gin-gonic/gin)</summary>

### [`v1.12.0`](https://redirect.github.com/gin-gonic/gin/blob/HEAD/CHANGELOG.md#Gin-v1120)

[Compare Source](https://redirect.github.com/gin-gonic/gin/compare/v1.11.0...v1.12.0)

##### Features

- feat(render): add bson protocol ([#&#8203;4145](https://redirect.github.com/gin-gonic/gin/pull/4145))
- feat(context): add GetError and GetErrorSlice methods for error retrieval ([#&#8203;4502](https://redirect.github.com/gin-gonic/gin/pull/4502))
- feat(binding): add support for encoding.UnmarshalText in uri/query binding ([#&#8203;4203](https://redirect.github.com/gin-gonic/gin/pull/4203))
- feat(gin): add option to use escaped path ([#&#8203;4420](https://redirect.github.com/gin-gonic/gin/pull/4420))
- feat(context): add Protocol Buffers support to content negotiation ([#&#8203;4423](https://redirect.github.com/gin-gonic/gin/pull/4423))
- feat(context): implemented Delete method ([#&#8203;38e7651](https://redirect.github.com/gin-gonic/gin/commit/38e7651))
- feat(logger): color latency ([#&#8203;4146](https://redirect.github.com/gin-gonic/gin/pull/4146))

##### Enhancements

- perf(tree): reduce allocations in findCaseInsensitivePath ([#&#8203;4417](https://redirect.github.com/gin-gonic/gin/pull/4417))
- perf(recovery): optimize line reading in stack function ([#&#8203;4466](https://redirect.github.com/gin-gonic/gin/pull/4466))
- perf(path): replace regex with custom functions in redirectTrailingSlash ([#&#8203;4414](https://redirect.github.com/gin-gonic/gin/pull/4414))
- perf(tree): optimize path parsing using strings.Count ([#&#8203;4246](https://redirect.github.com/gin-gonic/gin/pull/4246))
- chore(logger): allow skipping query string output ([#&#8203;4547](https://redirect.github.com/gin-gonic/gin/pull/4547))
- chore(context): always trust xff headers from unix socket ([#&#8203;3359](https://redirect.github.com/gin-gonic/gin/pull/3359))
- chore(response): prevent Flush() panic when the underlying ResponseWriter does not implement `http.Flusher` ([#&#8203;4479](https://redirect.github.com/gin-gonic/gin/pull/4479))
- refactor(recovery): smart error comparison ([#&#8203;4142](https://redirect.github.com/gin-gonic/gin/pull/4142))
- refactor(context): replace hardcoded localhost IPs with constants ([#&#8203;4481](https://redirect.github.com/gin-gonic/gin/pull/4481))
- refactor(utils): move util functions to utils.go ([#&#8203;4467](https://redirect.github.com/gin-gonic/gin/pull/4467))
- refactor(binding): use maps.Copy for cleaner map handling ([#&#8203;4352](https://redirect.github.com/gin-gonic/gin/pull/4352))
- refactor(context): using maps.Clone ([#&#8203;4333](https://redirect.github.com/gin-gonic/gin/pull/4333))
- refactor(ginS): use sync.OnceValue to simplify engine function ([#&#8203;4314](https://redirect.github.com/gin-gonic/gin/pull/4314))
- refactor: replace magic numbers with named constants in bodyAllowedForStatus ([#&#8203;4529](https://redirect.github.com/gin-gonic/gin/pull/4529))
- refactor: for loop can be modernized using range over int ([#&#8203;4392](https://redirect.github.com/gin-gonic/gin/pull/4392))

##### Bug Fixes

- fix(tree): panic in findCaseInsensitivePathRec with RedirectFixedPath ([#&#8203;4535](https://redirect.github.com/gin-gonic/gin/pull/4535))
- fix(render): write content length in Data.Render ([#&#8203;4206](https://redirect.github.com/gin-gonic/gin/pull/4206))
- fix(context): ClientIP handling for multiple X-Forwarded-For header values ([#&#8203;4472](https://redirect.github.com/gin-gonic/gin/pull/4472))
- fix(binding): empty value error ([#&#8203;2169](https://redirect.github.com/gin-gonic/gin/pull/2169))
- fix(recover): suppress http.ErrAbortHandler in recover ([#&#8203;4336](https://redirect.github.com/gin-gonic/gin/pull/4336))
- fix(gin): literal colon routes not working with engine.Handler() ([#&#8203;4415](https://redirect.github.com/gin-gonic/gin/pull/4415))
- fix(gin): close os.File in RunFd to prevent resource leak ([#&#8203;4422](https://redirect.github.com/gin-gonic/gin/pull/4422))
- fix(response): refine hijack behavior for response lifecycle ([#&#8203;4373](https://redirect.github.com/gin-gonic/gin/pull/4373))
- fix(binding): improve empty slice/array handling in form binding ([#&#8203;4380](https://redirect.github.com/gin-gonic/gin/pull/4380))
- fix(debug): version mismatch ([#&#8203;4403](https://redirect.github.com/gin-gonic/gin/pull/4403))
- fix: correct typos, improve documentation clarity, and remove dead code ([#&#8203;4511](https://redirect.github.com/gin-gonic/gin/pull/4511))

##### Build process updates / CI

- ci: update Go version support to 1.25+ across CI and docs ([#&#8203;4550](https://redirect.github.com/gin-gonic/gin/pull/4550))
- chore(binding): upgrade bson dependency to mongo-driver v2 ([#&#8203;4549](https://redirect.github.com/gin-gonic/gin/pull/4549))

</details>

---

### Configuration

📅 **Schedule**: Branch creation - Between 03:00 AM and 10:59 AM, only on Monday ( * 3-10 * * 1 ) in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0Mi4yNi41LXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjI2LjUtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @red-hat-konflux on March 02, 2026 at 05:23 AM UTC

### ℹ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 1 additional dependency was updated
- The `go` directive was updated for compatibility reasons


Details:


| **Package**        | **Change**             |
| :----------------- | :--------------------- |
| `go`               | `1.24.10` -> `1.25.0`  |
| `golang.org/x/net` | `v0.49.0` -> `v0.51.0` |

### Comment by @github-actions on March 02, 2026 at 05:23 AM UTC

<!-- sc-environment-impact-check -->
## SC Environment Impact Assessment

**Overall Impact:** ⚪ **NONE**

No SC Environment-specific impacts detected in this PR.

<details>
<summary>What was checked</summary>

This PR was automatically scanned for:
- Database migrations
- ClowdApp configuration changes
- Kessel integration changes
- AWS service integrations (S3, RDS, ElastiCache)
- Kafka topic changes
- Secrets management changes
- External dependencies
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2083*

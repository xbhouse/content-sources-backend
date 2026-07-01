---
type: pull_request
number: 1623
title: "Update module google.golang.org/protobuf to v1.36.6"
state: merged
author: red-hat-konflux
created: 2025-03-30T20:24:37Z
updated: 2026-04-04T20:23:02Z
closed: 2025-04-14T09:20:09Z
merged: 2025-04-14T09:20:08Z
base_branch: master
head_branch: konflux/mintmaker/master/google.golang.org-protobuf-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1623
---

# Pull Request #1623: Update module google.golang.org/protobuf to v1.36.6

**Author**: @red-hat-konflux
**Created**: March 30, 2025 at 08:24 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/google.golang.org-protobuf-1.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [google.golang.org/protobuf](https://redirect.github.com/protocolbuffers/protobuf-go) | indirect | patch | `v1.36.2` -> `v1.36.6` |

---

### Release Notes

<details>
<summary>protocolbuffers/protobuf-go (google.golang.org/protobuf)</summary>

### [`v1.36.6`](https://redirect.github.com/protocolbuffers/protobuf-go/releases/tag/v1.36.6)

[Compare Source](https://redirect.github.com/protocolbuffers/protobuf-go/compare/v1.36.5...v1.36.6)

**Full Changelog**: https://github.com/protocolbuffers/protobuf-go/compare/v1.36.5...v1.36.6

User-visible changes:
[CL/657895](https://go-review.googlesource.com/c/protobuf/+/657895): internal_gengo: generate a const string literal for the raw descriptor
[CL/653536](https://go-review.googlesource.com/c/protobuf/+/653536): proto: Add CloneOf\[M Message]\(m M) M

Maintenance:
[CL/649135](https://go-review.googlesource.com/c/protobuf/+/649135): all: set Go language version to Go 1.22
[CL/654955](https://go-review.googlesource.com/c/protobuf/+/654955): types/descriptorpb: regenerate using latest protobuf v30 release

### [`v1.36.5`](https://redirect.github.com/protocolbuffers/protobuf-go/releases/tag/v1.36.5)

[Compare Source](https://redirect.github.com/protocolbuffers/protobuf-go/compare/v1.36.4...v1.36.5)

**Full Changelog**: https://github.com/protocolbuffers/protobuf-go/compare/v1.36.4...v1.36.5

Bug fixes:
[CL/644437](https://go-review.googlesource.com/c/protobuf/+/644437): protogen: fix name mangling for fields with identical GoCamelCase

Maintenance:
[CL/641655](https://go-review.googlesource.com/c/protobuf/+/641655): all: remove weak field support

### [`v1.36.4`](https://redirect.github.com/protocolbuffers/protobuf-go/releases/tag/v1.36.4)

[Compare Source](https://redirect.github.com/protocolbuffers/protobuf-go/compare/v1.36.3...v1.36.4)

**Full Changelog**: https://github.com/protocolbuffers/protobuf-go/compare/v1.36.3...v1.36.4

Bug fixes:
[CL/642975](https://go-review.googlesource.com/c/protobuf/+/642975): reflect/protodesc: fix panic when working with dynamicpb

Maintenance:
[CL/643276](https://go-review.googlesource.com/c/protobuf/+/643276): internal_gengo: avoid allocations in rawDescGZIP() accessors
[CL/642857](https://go-review.googlesource.com/c/protobuf/+/642857): internal_gengo: switch back from string literal to hex byte slice
[CL/642055](https://go-review.googlesource.com/c/protobuf/+/642055): internal_gengo: use unsafe.StringData() to avoid a descriptor copy
[CL/638135](https://go-review.googlesource.com/c/protobuf/+/638135): internal_gengo: store raw descriptor in .rodata section

### [`v1.36.3`](https://redirect.github.com/protocolbuffers/protobuf-go/releases/tag/v1.36.3)

[Compare Source](https://redirect.github.com/protocolbuffers/protobuf-go/compare/v1.36.2...v1.36.3)

**Full Changelog**: https://github.com/protocolbuffers/protobuf-go/compare/v1.36.2...v1.36.3

Bug fixes:
[CL/642575](https://go-review.googlesource.com/c/protobuf/+/642575): reflect/protodesc: fix panic when working with dynamicpb
[CL/641036](https://go-review.googlesource.com/c/protobuf/+/641036): cmd/protoc-gen-go: remove json struct tags from unexported fields

User-visible changes:
[CL/641876](https://go-review.googlesource.com/c/protobuf/+/641876): proto: add example for GetExtension, SetExtension
[CL/642015](https://go-review.googlesource.com/c/protobuf/+/642015): runtime/protolazy: replace internal doc link with external link

Maintenance:
[CL/641635](https://go-review.googlesource.com/c/protobuf/+/641635): all: split flags.ProtoLegacyWeak out of flags.ProtoLegacy
[CL/641019](https://go-review.googlesource.com/c/protobuf/+/641019): internal/impl: remove unused exporter parameter
[CL/641018](https://go-review.googlesource.com/c/protobuf/+/641018): internal/impl: switch to reflect.Value.IsZero
[CL/641035](https://go-review.googlesource.com/c/protobuf/+/641035): internal/impl: clean up unneeded Go<1.12 MapRange() alternative
[CL/641017](https://go-review.googlesource.com/c/protobuf/+/641017): types/dynamicpb: switch atomicExtFiles to atomic.Uint64 type

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "after 5am on sunday" in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ1cGRhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ0YXJnZXRCcmFuY2giOiJtYXN0ZXIiLCJsYWJlbHMiOltdfQ==-->


---

## Discussion

### Comment by @jira-linking on March 30, 2025 at 08:24 PM UTC

Commits missing Jira IDs:
9cd6916356774dd700df2ef321fd24fea86775e7


### Comment by @codecov-commenter on March 30, 2025 at 08:29 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1623?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 58.15%. Comparing base ([`d90de5c`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/d90de5c730ea2b0e4721ef52cb7d0b0298e3555c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`9cd6916`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/9cd6916356774dd700df2ef321fd24fea86775e7?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 902 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1623   +/-   ##
=======================================
  Coverage   58.15%   58.15%           
=======================================
  Files         146      146           
  Lines       11304    11304           
=======================================
  Hits         6574     6574           
  Misses       4143     4143           
  Partials      587      587           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1623/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1623/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.15% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1623?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1623*

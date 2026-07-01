---
type: pull_request
number: 1061
title: "Update module google.golang.org/protobuf to v1.36.6"
state: closed
author: red-hat-konflux
created: 2025-04-06T05:35:58Z
updated: 2025-04-08T17:56:02Z
closed: 2025-04-08T17:09:49Z
base_branch: main
head_branch: konflux/mintmaker/main/google.golang.org-protobuf-1.x
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1061
---

# Pull Request #1061: Update module google.golang.org/protobuf to v1.36.6

**Author**: @red-hat-konflux
**Created**: April 06, 2025 at 05:35 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/google.golang.org-protobuf-1.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [google.golang.org/protobuf](https://redirect.github.com/protocolbuffers/protobuf-go) | indirect | patch | `v1.36.1` -> `v1.36.6` |

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

### [`v1.36.2`](https://redirect.github.com/protocolbuffers/protobuf-go/releases/tag/v1.36.2)

[Compare Source](https://redirect.github.com/protocolbuffers/protobuf-go/compare/v1.36.1...v1.36.2)

**Full Changelog**: https://github.com/protocolbuffers/protobuf-go/compare/v1.36.1...v1.36.2

Bug fixes:
[CL/638515](https://go-review.googlesource.com/c/protobuf/+/638515): internal/impl: fix WhichOneof() to work with synthetic oneofs

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

### Comment by @red-hat-konflux on April 08, 2025 at 05:56 PM UTC

### Renovate Ignore Notification

Because you closed this PR without merging, Renovate will ignore this update (`v1.36.6`). You will get a PR once a newer version is released. To ignore this dependency forever, add it to the `ignoreDeps` array of your Renovate config.

If you accidentally closed this PR, or if you changed your mind: rename this PR to get a fresh replacement PR.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1061*

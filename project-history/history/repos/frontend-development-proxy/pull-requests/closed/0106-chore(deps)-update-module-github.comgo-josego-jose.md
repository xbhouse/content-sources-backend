---
type: pull_request
number: 106
title: "chore(deps): update module github.com/go-jose/go-jose/v3 to v4 - autoclosed"
state: closed
author: red-hat-konflux
created: 2026-01-07T12:38:25Z
updated: 2026-06-05T17:29:31Z
closed: 2026-06-05T17:29:31Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-go-jose-go-jose-v3-4.x
labels: ["dependencies"]
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/106
---

# Pull Request #106: chore(deps): update module github.com/go-jose/go-jose/v3 to v4 - autoclosed

**Author**: @red-hat-konflux
**Created**: January 07, 2026 at 12:38 PM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-go-jose-go-jose-v3-4.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/go-jose/go-jose/v3](https://redirect.github.com/go-jose/go-jose) | `v3.0.5` → `v4.1.4` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgo-jose%2fgo-jose%2fv3/v4.1.4?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgo-jose%2fgo-jose%2fv3/v3.0.5/v4.1.4?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>go-jose/go-jose (github.com/go-jose/go-jose/v3)</summary>

### [`v4.1.4`](https://redirect.github.com/go-jose/go-jose/compare/v4.1.3...v4.1.4)

[Compare Source](https://redirect.github.com/go-jose/go-jose/compare/v4.1.3...v4.1.4)

### [`v4.1.3`](https://redirect.github.com/go-jose/go-jose/releases/tag/v4.1.3)

[Compare Source](https://redirect.github.com/go-jose/go-jose/compare/v4.1.2...v4.1.3)

This release drops Go 1.23 support as that Go release is no longer supported. With that, we can drop `x/crypto` and no longer have any external dependencies in go-jose outside of the standard library!

This release fixes a bug where a critical b64 header was ignored if in an unprotected header. It is now rejected instead of ignored.

#### What's Changed

- Remove Go 1.23 support by [@&#8203;mcpherrinm](https://redirect.github.com/mcpherrinm) in [#&#8203;205](https://redirect.github.com/go-jose/go-jose/pull/205)
- Reject JWS with an unprotected critical b64 header by [@&#8203;mcpherrinm](https://redirect.github.com/mcpherrinm) in [#&#8203;210](https://redirect.github.com/go-jose/go-jose/pull/210)

**Full Changelog**: <https://github.com/go-jose/go-jose/compare/v4.1.2...v4.1.3>

### [`v4.1.2`](https://redirect.github.com/go-jose/go-jose/releases/tag/v4.1.2)

[Compare Source](https://redirect.github.com/go-jose/go-jose/compare/v4.1.1...v4.1.2)

#### What's Changed

go-jose v4.1.2 improves some documentation, errors, and removes the only 3rd-party dependency.

- Update go-jose documentation by [@&#8203;mcpherrinm](https://redirect.github.com/mcpherrinm) in [#&#8203;198](https://redirect.github.com/go-jose/go-jose/pull/198)
- Remove dependency on testify by [@&#8203;wardviaene](https://redirect.github.com/wardviaene) in [#&#8203;197](https://redirect.github.com/go-jose/go-jose/pull/197)
- Improve error message for invalid private keys by [@&#8203;ProjectMutilation](https://redirect.github.com/ProjectMutilation) in [#&#8203;195](https://redirect.github.com/go-jose/go-jose/pull/195)
- JWK unsupported error when unmarshalling by [@&#8203;fprojetto](https://redirect.github.com/fprojetto) in [#&#8203;191](https://redirect.github.com/go-jose/go-jose/pull/191)
- Add JSONWebKey type to makeJWERecipient by [@&#8203;alvarolivie](https://redirect.github.com/alvarolivie) in [#&#8203;200](https://redirect.github.com/go-jose/go-jose/pull/200)
- testutils/assert: remove True, Nil, NotNil by [@&#8203;jsha](https://redirect.github.com/jsha) in [#&#8203;202](https://redirect.github.com/go-jose/go-jose/pull/202)

#### New Contributors

- [@&#8203;wardviaene](https://redirect.github.com/wardviaene) made their first contribution in [#&#8203;197](https://redirect.github.com/go-jose/go-jose/pull/197)
- [@&#8203;fprojetto](https://redirect.github.com/fprojetto) made their first contribution in [#&#8203;191](https://redirect.github.com/go-jose/go-jose/pull/191)
- [@&#8203;alvarolivie](https://redirect.github.com/alvarolivie) made their first contribution in [#&#8203;200](https://redirect.github.com/go-jose/go-jose/pull/200)

**Full Changelog**: <https://github.com/go-jose/go-jose/compare/v4.1.1...v4.1.2>

### [`v4.1.1`](https://redirect.github.com/go-jose/go-jose/releases/tag/v4.1.1)

[Compare Source](https://redirect.github.com/go-jose/go-jose/compare/v4.1.0...v4.1.1)

#### What's Changed

- Drop go-cmp dependency by [@&#8203;mcpherrinm](https://redirect.github.com/mcpherrinm) in [#&#8203;186](https://redirect.github.com/go-jose/go-jose/pull/186)
- jws: improve performance and allocations for ParseSignedCompact by [@&#8203;drakkan](https://redirect.github.com/drakkan) in [#&#8203;188](https://redirect.github.com/go-jose/go-jose/pull/188)
- Add missing quote to unknown curve message [#&#8203;170](https://redirect.github.com/go-jose/go-jose/issues/170) by [@&#8203;sudhanvaghebbale](https://redirect.github.com/sudhanvaghebbale) in [#&#8203;189](https://redirect.github.com/go-jose/go-jose/pull/189)
- Fix incorrect validation by [@&#8203;ProjectMutilation](https://redirect.github.com/ProjectMutilation) in [#&#8203;192](https://redirect.github.com/go-jose/go-jose/pull/192)
- Restore Go 1.23 compatibility by [@&#8203;anuraaga](https://redirect.github.com/anuraaga) in [#&#8203;193](https://redirect.github.com/go-jose/go-jose/pull/193)

#### New Contributors

- [@&#8203;drakkan](https://redirect.github.com/drakkan) made their first contribution in [#&#8203;188](https://redirect.github.com/go-jose/go-jose/pull/188)
- [@&#8203;sudhanvaghebbale](https://redirect.github.com/sudhanvaghebbale) made their first contribution in [#&#8203;189](https://redirect.github.com/go-jose/go-jose/pull/189)
- [@&#8203;ProjectMutilation](https://redirect.github.com/ProjectMutilation) made their first contribution in [#&#8203;192](https://redirect.github.com/go-jose/go-jose/pull/192)
- [@&#8203;anuraaga](https://redirect.github.com/anuraaga) made their first contribution in [#&#8203;193](https://redirect.github.com/go-jose/go-jose/pull/193)

**Full Changelog**: <https://github.com/go-jose/go-jose/compare/v4.1.0...v4.1.1>

### [`v4.1.0`](https://redirect.github.com/go-jose/go-jose/releases/tag/v4.1.0)

[Compare Source](https://redirect.github.com/go-jose/go-jose/compare/v4.0.5...v4.1.0)

#### What's Changed

- Document `signatureAlgorithms` argument by [@&#8203;tgeoghegan](https://redirect.github.com/tgeoghegan) in [#&#8203;163](https://redirect.github.com/go-jose/go-jose/pull/163)
- Add custom error for unsupported JWS signature algorithms by [@&#8203;beautifulentropy](https://redirect.github.com/beautifulentropy) in [#&#8203;181](https://redirect.github.com/go-jose/go-jose/pull/181)
- use stdlib pbkdf2 package on go 1.24 by [@&#8203;kruskall](https://redirect.github.com/kruskall) in [#&#8203;180](https://redirect.github.com/go-jose/go-jose/pull/180)
- The minimum supported Go version is now 1.24

#### New Contributors

- [@&#8203;kruskall](https://redirect.github.com/kruskall) made their first contribution in [#&#8203;180](https://redirect.github.com/go-jose/go-jose/pull/180)

**Full Changelog**: <https://github.com/go-jose/go-jose/compare/v4.0.5...v4.1.0>

### [`v4.0.5`](https://redirect.github.com/go-jose/go-jose/releases/tag/v4.0.5)

[Compare Source](https://redirect.github.com/go-jose/go-jose/compare/v4.0.4...v4.0.5)

#### What's Changed

- Don't allow unbounded amounts of splits by [@&#8203;mcpherrinm](https://redirect.github.com/mcpherrinm) in [#&#8203;167](https://redirect.github.com/go-jose/go-jose/pull/167)

Fixes <https://github.com/go-jose/go-jose/security/advisories/GHSA-c6gw-w398-hv78>

Various other dependency updates, small fixes, and documentation updates in the full changelog

#### New Contributors

- [@&#8203;tgeoghegan](https://redirect.github.com/tgeoghegan) made their first contribution in [#&#8203;161](https://redirect.github.com/go-jose/go-jose/pull/161)

**Full Changelog**: <https://github.com/go-jose/go-jose/compare/v4.0.4...v4.0.5>

### [`v4.0.4`](https://redirect.github.com/go-jose/go-jose/releases/tag/v4.0.4): Version 4.0.4

[Compare Source](https://redirect.github.com/go-jose/go-jose/compare/v4.0.3...v4.0.4)

### Fixed

- Reverted "Allow unmarshalling JSONWebKeySets with unsupported key types" as a breaking change. See [#&#8203;136](https://redirect.github.com/go-jose/go-jose/issues/136) / [#&#8203;137](https://redirect.github.com/go-jose/go-jose/issues/137).

### [`v4.0.3`](https://redirect.github.com/go-jose/go-jose/releases/tag/v4.0.3): Version 4.0.3

[Compare Source](https://redirect.github.com/go-jose/go-jose/compare/v4.0.2...v4.0.3)

#### Changed

- Allow unmarshalling JSONWebKeySets with unsupported key types ([#&#8203;130](https://redirect.github.com/go-jose/go-jose/issues/130))
- Document that OpaqueKeyEncrypter can't be implemented (for now) ([#&#8203;129](https://redirect.github.com/go-jose/go-jose/issues/129))
- Dependency updates

### [`v4.0.2`](https://redirect.github.com/go-jose/go-jose/releases/tag/v4.0.2): Version 4.0.2

[Compare Source](https://redirect.github.com/go-jose/go-jose/compare/v4.0.1...v4.0.2)

#### What's Changed

- [Improved documentation](https://redirect.github.com/go-jose/go-jose/pull/104) of Verify() to note that JSONWebKeySet is a supported argument type
- [Defined exported error values](https://redirect.github.com/go-jose/go-jose/pull/117) for missing x5c header and unsupported elliptic curves error cases

#### New Contributors

- [@&#8203;mitar](https://redirect.github.com/mitar) made their first contribution in [#&#8203;104](https://redirect.github.com/go-jose/go-jose/pull/104)
- [@&#8203;milosgajdos](https://redirect.github.com/milosgajdos) made their first contribution in [#&#8203;117](https://redirect.github.com/go-jose/go-jose/pull/117)

**Full Changelog**: <https://github.com/go-jose/go-jose/compare/v4.0.1...v4.0.2>

### [`v4.0.1`](https://redirect.github.com/go-jose/go-jose/releases/tag/v4.0.1): Version 4.0.1

[Compare Source](https://redirect.github.com/go-jose/go-jose/compare/v4.0.0...v4.0.1)

#### Fixed

- An attacker could send a JWE containing compressed data that used large
  amounts of memory and CPU when decompressed by `Decrypt` or `DecryptMulti`.
  Those functions now return an error if the decompressed data would exceed
  250kB or 10x the compressed size (whichever is larger). Thanks to
  Enze Wang\@&#8203;Alioth and Jianjun Chen\@&#8203;Zhongguancun Lab ([@&#8203;zer0yu](https://redirect.github.com/zer0yu) and [@&#8203;chenjj](https://redirect.github.com/chenjj))
  for reporting.

### [`v4.0.0`](https://redirect.github.com/go-jose/go-jose/releases/tag/v4.0.0): Version 4.0.0

[Compare Source](https://redirect.github.com/go-jose/go-jose/compare/v3.0.5...v4.0.0)

This release makes some breaking changes in order to more thoroughly address the vulnerabilities discussed in [Three New Attacks Against JSON Web Tokens][1], "Sign/encrypt confusion", "Billion hash attack", and "Polyglot token".

#### Changed

- Limit JWT encryption types (exclude password or public key types) ([#&#8203;78](https://redirect.github.com/go-jose/go-jose/issues/78))
- Enforce minimum length for HMAC keys ([#&#8203;85](https://redirect.github.com/go-jose/go-jose/issues/85))
- jwt: match any audience in a list, rather than requiring all audiences ([#&#8203;81](https://redirect.github.com/go-jose/go-jose/issues/81))
- jwt: accept only Compact Serialization ([#&#8203;75](https://redirect.github.com/go-jose/go-jose/issues/75))
- jws: Add expected algorithms for signatures ([#&#8203;74](https://redirect.github.com/go-jose/go-jose/issues/74))
- Require specifying expected algorithms for ParseEncrypted,
  ParseSigned, ParseDetached, jwt.ParseEncrypted, jwt.ParseSigned,
  jwt.ParseSignedAndEncrypted ([#&#8203;69](https://redirect.github.com/go-jose/go-jose/issues/69), [#&#8203;74](https://redirect.github.com/go-jose/go-jose/issues/74))
  - Usually there is a small, known set of appropriate algorithms for a program to use and it's a mistake to allow unexpected algorithms. For instance the "billion hash attack" relies in part on programs accepting the PBES2 encryption algorithm and doing the necessary work even if they weren't specifically configured to allow PBES2.
- Revert "Strip padding off base64 strings" ([#&#8203;82](https://redirect.github.com/go-jose/go-jose/issues/82))
- The specs require base64url encoding without padding.
- Minimum supported Go version is now 1.21

#### Added

- ParseSignedCompact, ParseSignedJSON, ParseEncryptedCompact, ParseEncryptedJSON.
  - These allow parsing a specific serialization, as opposed to ParseSigned and ParseEncrypted, which try to automatically detect which serialization was provided. It's common to require a specific serialization for a specific protocol - for instance JWT requires Compact serialization.

[1]: https://i.blackhat.com/BH-US-23/Presentations/US-23-Tervoort-Three-New-Attacks-Against-JSON-Web-Tokens.pdf

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjk5LjAtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6WyJkZXBlbmRlbmNpZXMiXX0=-->


---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/106*

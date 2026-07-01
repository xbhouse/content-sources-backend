---
type: pull_request
number: 45
title: "Update module github.com/mholt/acmez/v3 to v3.1.4"
state: merged
author: red-hat-konflux
created: 2025-11-06T08:19:33Z
updated: 2025-11-11T14:31:59Z
closed: 2025-11-11T14:31:56Z
merged: 2025-11-11T14:31:56Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-mholt-acmez-v3-3.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/45
---

# Pull Request #45: Update module github.com/mholt/acmez/v3 to v3.1.4

**Author**: @red-hat-konflux
**Created**: November 06, 2025 at 08:19 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-mholt-acmez-v3-3.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/mholt/acmez/v3](https://redirect.github.com/mholt/acmez) | `v3.0.0` -> `v3.1.4` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fmholt%2facmez%2fv3/v3.1.4?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fmholt%2facmez%2fv3/v3.0.0/v3.1.4?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>mholt/acmez (github.com/mholt/acmez/v3)</summary>

### [`v3.1.4`](https://redirect.github.com/mholt/acmez/releases/tag/v3.1.4)

[Compare Source](https://redirect.github.com/mholt/acmez/compare/v3.1.3...v3.1.4)

A few fixes that could technically be breaking, but are actually bugs, for example, a function that was spelled "Acccount" instead of "Account", and the way the DNS01Account helper function worked has been corrected to match the intent/documentation.

We also now truncate NotBefore and NotAfter to second precision.

#### What's Changed

- Fixed typo by [@&#8203;huzpsb](https://redirect.github.com/huzpsb) in [#&#8203;44](https://redirect.github.com/mholt/acmez/pull/44)
- fix(challenge): use consistent structure by [@&#8203;Integralist](https://redirect.github.com/Integralist) in [#&#8203;45](https://redirect.github.com/mholt/acmez/pull/45)

#### New Contributors

- [@&#8203;huzpsb](https://redirect.github.com/huzpsb) made their first contribution in [#&#8203;44](https://redirect.github.com/mholt/acmez/pull/44)
- [@&#8203;Integralist](https://redirect.github.com/Integralist) made their first contribution in [#&#8203;45](https://redirect.github.com/mholt/acmez/pull/45)

**Full Changelog**: <https://github.com/mholt/acmez/compare/v3.1.3...v3.1.4>

### [`v3.1.3`](https://redirect.github.com/mholt/acmez/releases/tag/v3.1.3)

[Compare Source](https://redirect.github.com/mholt/acmez/compare/v3.1.2...v3.1.3)

Minor release that adds helpers for the dns-account-01 challenge, and an exported `Thumbprint()` method for ACME accounts.

#### What's Changed

- Fix typo by [@&#8203;cpach](https://redirect.github.com/cpach) in [#&#8203;41](https://redirect.github.com/mholt/acmez/pull/41)
- Add helpers for dns-account-01 challenge (close [#&#8203;42](https://redirect.github.com/mholt/acmez/issues/42)) by [@&#8203;mholt](https://redirect.github.com/mholt) in [#&#8203;43](https://redirect.github.com/mholt/acmez/pull/43)

#### New Contributors

- [@&#8203;cpach](https://redirect.github.com/cpach) made their first contribution in [#&#8203;41](https://redirect.github.com/mholt/acmez/pull/41)

**Full Changelog**: <https://github.com/mholt/acmez/compare/v3.1.2...v3.1.3>

### [`v3.1.2`](https://redirect.github.com/mholt/acmez/releases/tag/v3.1.2)

[Compare Source](https://redirect.github.com/mholt/acmez/compare/v3.1.1...v3.1.2)

Better support for invalid ARI windows, which should never happen anyway.

**Full Changelog**: <https://github.com/mholt/acmez/compare/v3.1.1...v3.1.2>

### [`v3.1.1`](https://redirect.github.com/mholt/acmez/releases/tag/v3.1.1)

[Compare Source](https://redirect.github.com/mholt/acmez/compare/v3.1.0...v3.1.1)

Very minor bug fix to ensure the TLS-ALPN challenge works for IDNs.

**Full Changelog**: <https://github.com/mholt/acmez/compare/v3.1.0...v3.1.1>

### [`v3.1.0`](https://redirect.github.com/mholt/acmez/releases/tag/v3.1.0)

[Compare Source](https://redirect.github.com/mholt/acmez/compare/v3.0.1...v3.1.0)

#### What's Changed

- Support for TNAuthlist identifier by [@&#8203;samuhvarta](https://redirect.github.com/samuhvarta) in [#&#8203;35](https://redirect.github.com/mholt/acmez/pull/35)
- \[fix]: use base64url encoding when reading TNAuthList from csr by [@&#8203;samuhvarta](https://redirect.github.com/samuhvarta) in [#&#8203;37](https://redirect.github.com/mholt/acmez/pull/37)

#### New Contributors

- [@&#8203;samuhvarta](https://redirect.github.com/samuhvarta) made their first contribution in [#&#8203;35](https://redirect.github.com/mholt/acmez/pull/35)

**Full Changelog**: <https://github.com/mholt/acmez/compare/v3.0.1...v3.1.0>

### [`v3.0.1`](https://redirect.github.com/mholt/acmez/compare/v3.0.0...v3.0.1)

[Compare Source](https://redirect.github.com/mholt/acmez/compare/v3.0.0...v3.0.1)

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

## Discussion

### Comment by @Hyperkid123 on November 11, 2025 at 02:17 PM UTC

/retest

---

## Reviews

### Review by @Hyperkid123 - Approved on November 11, 2025 at 02:31 PM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/45*

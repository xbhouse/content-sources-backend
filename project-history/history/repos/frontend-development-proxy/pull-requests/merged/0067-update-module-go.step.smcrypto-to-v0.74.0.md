---
type: pull_request
number: 67
title: "Update module go.step.sm/crypto to v0.74.0"
state: merged
author: red-hat-konflux
created: 2025-11-12T00:35:42Z
updated: 2025-11-21T09:49:11Z
closed: 2025-11-21T09:49:11Z
merged: 2025-11-21T09:49:11Z
base_branch: main
head_branch: konflux/mintmaker/main/go.step.sm-crypto-0.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/67
---

# Pull Request #67: Update module go.step.sm/crypto to v0.74.0

**Author**: @red-hat-konflux
**Created**: November 12, 2025 at 12:35 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/go.step.sm-crypto-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [go.step.sm/crypto](https://redirect.github.com/smallstep/crypto) | `v0.45.0` -> `v0.74.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/go.step.sm%2fcrypto/v0.74.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/go.step.sm%2fcrypto/v0.45.0/v0.74.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>smallstep/crypto (go.step.sm/crypto)</summary>

### [`v0.74.0`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.74.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.73.0...v0.74.0)

#### What's Changed

- Add support for providing `*tpm.TPM` instance when creating a `TPMKMS` by [@&#8203;maraino](https://redirect.github.com/maraino) in [#&#8203;885](https://redirect.github.com/smallstep/crypto/pull/885)

##### Dependencies

- Bump github.com/aws/aws-sdk-go-v2/config from 1.31.15 to 1.31.16 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;889](https://redirect.github.com/smallstep/crypto/pull/889)
- Bump modernc.org/sqlite from 1.39.1 to 1.40.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;888](https://redirect.github.com/smallstep/crypto/pull/888)
- Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.46.2 to 1.47.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;887](https://redirect.github.com/smallstep/crypto/pull/887)
- Bump google.golang.org/api from 0.253.0 to 0.254.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;886](https://redirect.github.com/smallstep/crypto/pull/886)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.73.0...v0.74.0>

### [`v0.73.0`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.73.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.72.0...v0.73.0)

#### What's Changed

- Add support for AES128 and AES256 YubiKey management keys by [@&#8203;hslatman](https://redirect.github.com/hslatman) in [#&#8203;884](https://redirect.github.com/smallstep/crypto/pull/884)

##### Dependencies

- chore(deps): Bump github.com/aws/aws-sdk-go-v2/config from 1.31.12 to 1.31.13 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;875](https://redirect.github.com/smallstep/crypto/pull/875)
- chore(deps): Bump google.golang.org/api from 0.251.0 to 0.252.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;877](https://redirect.github.com/smallstep/crypto/pull/877)
- chore(deps): Bump github.com/Azure/azure-sdk-for-go/sdk/azidentity from 1.12.0 to 1.13.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;879](https://redirect.github.com/smallstep/crypto/pull/879)
- chore(deps): Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.45.6 to 1.46.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;878](https://redirect.github.com/smallstep/crypto/pull/878)
- chore(deps): Bump cloud.google.com/go/kms from 1.23.1 to 1.23.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;876](https://redirect.github.com/smallstep/crypto/pull/876)
- chore(deps): Bump github.com/aws/aws-sdk-go-v2/config from 1.31.13 to 1.31.15 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;882](https://redirect.github.com/smallstep/crypto/pull/882)
- chore(deps): Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.46.0 to 1.46.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;880](https://redirect.github.com/smallstep/crypto/pull/880)
- chore(deps): Bump google.golang.org/api from 0.252.0 to 0.253.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;881](https://redirect.github.com/smallstep/crypto/pull/881)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.72.0...v0.73.0>

### [`v0.72.0`](https://redirect.github.com/smallstep/crypto/compare/v0.71.0...v0.72.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.71.0...v0.72.0)

### [`v0.71.0`](https://redirect.github.com/smallstep/crypto/compare/v0.70.0...v0.71.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.70.0...v0.71.0)

### [`v0.70.0`](https://redirect.github.com/smallstep/crypto/compare/v0.69.0...v0.70.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.69.0...v0.70.0)

### [`v0.69.0`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.69.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.68.0...v0.69.0)

#### What's Changed

- Accept nonRepudiation when unmarshalling keyUsage by [@&#8203;neverpanic](https://redirect.github.com/neverpanic) in [#&#8203;811](https://redirect.github.com/smallstep/crypto/pull/811)
- Fix usage of azure keyvault client-secret when specified in URI by [@&#8203;dellekappa](https://redirect.github.com/dellekappa) in [#&#8203;819](https://redirect.github.com/smallstep/crypto/pull/819)
- Correctly set the machine store flag when requested by [@&#8203;areed](https://redirect.github.com/areed) in [#&#8203;802](https://redirect.github.com/smallstep/crypto/pull/802)

#### Dependencies

- Bump google.golang.org/api from 0.240.0 to 0.244.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;816](https://redirect.github.com/smallstep/crypto/pull/816)
- Bump github.com/Azure/azure-sdk-for-go/sdk/azcore from 1.18.1 to 1.18.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;815](https://redirect.github.com/smallstep/crypto/pull/815)
- Bump modernc.org/sqlite from 1.38.0 to 1.38.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;812](https://redirect.github.com/smallstep/crypto/pull/812)
- Bump github.com/go-piv/piv-go/v2 from 2.3.0 to 2.4.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;813](https://redirect.github.com/smallstep/crypto/pull/813)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.29.17 to 1.30.3 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;817](https://redirect.github.com/smallstep/crypto/pull/817)
- Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.41.2 to 1.43.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;818](https://redirect.github.com/smallstep/crypto/pull/818)

#### New Contributors

- [@&#8203;neverpanic](https://redirect.github.com/neverpanic) made their first contribution in [#&#8203;811](https://redirect.github.com/smallstep/crypto/pull/811)
- [@&#8203;dellekappa](https://redirect.github.com/dellekappa) made their first contribution in [#&#8203;819](https://redirect.github.com/smallstep/crypto/pull/819)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.68.0...v0.69.0>

### [`v0.68.0`](https://redirect.github.com/smallstep/crypto/compare/v0.67.0...v0.68.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.67.0...v0.68.0)

### [`v0.67.0`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.67.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.66.0...v0.67.0)

#### What's Changed

- Add support for the Data Protection Keychain by [@&#8203;maraino](https://redirect.github.com/maraino) in [#&#8203;793](https://redirect.github.com/smallstep/crypto/pull/793)

#### Dependencies

- Bump google.golang.org/api from 0.234.0 to 0.235.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;780](https://redirect.github.com/smallstep/crypto/pull/780)
- Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.38.3 to 1.40.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;783](https://redirect.github.com/smallstep/crypto/pull/783)
- Bump modernc.org/sqlite from 1.37.1 to 1.38.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;786](https://redirect.github.com/smallstep/crypto/pull/786)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.29.14 to 1.29.15 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;785](https://redirect.github.com/smallstep/crypto/pull/785)
- Bump google.golang.org/grpc from 1.72.2 to 1.73.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;787](https://redirect.github.com/smallstep/crypto/pull/787)
- Bump golang.org/x/crypto from 0.38.0 to 0.39.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;784](https://redirect.github.com/smallstep/crypto/pull/784)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.29.15 to 1.29.16 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;788](https://redirect.github.com/smallstep/crypto/pull/788)
- Bump github.com/Azure/azure-sdk-for-go/sdk/azidentity from 1.10.0 to 1.10.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;790](https://redirect.github.com/smallstep/crypto/pull/790)
- Bump golang.org/x/net from 0.40.0 to 0.41.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;789](https://redirect.github.com/smallstep/crypto/pull/789)
- Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.40.0 to 1.41.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;792](https://redirect.github.com/smallstep/crypto/pull/792)
- Bump google.golang.org/api from 0.235.0 to 0.237.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;791](https://redirect.github.com/smallstep/crypto/pull/791)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.66.0...v0.67.0>

### [`v0.66.0`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.66.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.65.0...v0.66.0)

#### What's Changed

- Add new extension fields on certificate requests by [@&#8203;maraino](https://redirect.github.com/maraino) in [#&#8203;767](https://redirect.github.com/smallstep/crypto/pull/767)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.65.0...v0.66.0>

### [`v0.65.0`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.65.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.64.0...v0.65.0)

#### What's Changed

- Upgrade `github.com/smallstep/go-attestation` to `2306d5b4` by [@&#8203;hslatman](https://redirect.github.com/hslatman) in [#&#8203;775](https://redirect.github.com/smallstep/crypto/pull/775), fixing an issue with ECDSA key signatures on Windows.

Dependencies:

- Bump github.com/googleapis/gax-go/v2 from 2.14.1 to 2.14.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;774](https://redirect.github.com/smallstep/crypto/pull/774)
- Bump google.golang.org/grpc from 1.72.0 to 1.72.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;771](https://redirect.github.com/smallstep/crypto/pull/771)
- Bump github.com/Azure/azure-sdk-for-go/sdk/azidentity from 1.9.0 to 1.10.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;773](https://redirect.github.com/smallstep/crypto/pull/773)
- Bump google.golang.org/api from 0.232.0 to 0.234.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;778](https://redirect.github.com/smallstep/crypto/pull/778)
- Bump cloud.google.com/go/kms from 1.21.2 to 1.22.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;776](https://redirect.github.com/smallstep/crypto/pull/776)
- Bump modernc.org/sqlite from 1.37.0 to 1.37.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;777](https://redirect.github.com/smallstep/crypto/pull/777)
- Bump google.golang.org/grpc from 1.72.1 to 1.72.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;779](https://redirect.github.com/smallstep/crypto/pull/779)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.64.0...v0.65.0>

### [`v0.64.0`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.64.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.63.0...v0.64.0)

#### What's Changed

- Add support for 'DeleteKey' in CAPIKMS. by [@&#8203;joshdrake](https://redirect.github.com/joshdrake) in [#&#8203;770](https://redirect.github.com/smallstep/crypto/pull/770)

Dependencies:

- Bump golang.org/x/sys from 0.32.0 to 0.33.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;766](https://redirect.github.com/smallstep/crypto/pull/766)
- Bump golang.org/x/crypto from 0.37.0 to 0.38.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;763](https://redirect.github.com/smallstep/crypto/pull/763)
- Bump google.golang.org/api from 0.230.0 to 0.231.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;764](https://redirect.github.com/smallstep/crypto/pull/764)
- Bump golang.org/x/net from 0.39.0 to 0.40.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;765](https://redirect.github.com/smallstep/crypto/pull/765)
- Bump google.golang.org/api from 0.231.0 to 0.232.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;768](https://redirect.github.com/smallstep/crypto/pull/768)
- Bump github.com/google/go-tpm from 0.9.4 to 0.9.5 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;769](https://redirect.github.com/smallstep/crypto/pull/769)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.63.0...v0.64.0>

### [`v0.63.0`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.63.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.62.0...v0.63.0)

#### What's Changed

- Allow locating certificates using issuer + subject CN in CAPI KMS. by [@&#8203;joshdrake](https://redirect.github.com/joshdrake) in [#&#8203;696](https://redirect.github.com/smallstep/crypto/pull/696)
- Add support for rsa.PSSSaltLengthAuto on capi by [@&#8203;maraino](https://redirect.github.com/maraino) in [#&#8203;762](https://redirect.github.com/smallstep/crypto/pull/762)

Dependencies:

- Bump go.uber.org/mock from 0.5.1 to 0.5.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;760](https://redirect.github.com/smallstep/crypto/pull/760)
- Bump google.golang.org/api from 0.229.0 to 0.230.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;759](https://redirect.github.com/smallstep/crypto/pull/759)
- Bump github.com/google/go-tpm from 0.9.3 to 0.9.4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;761](https://redirect.github.com/smallstep/crypto/pull/761)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.62.0...v0.63.0>

### [`v0.62.0`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.62.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.61.0...v0.62.0)

#### What's Changed

- Extract public key from private by [@&#8203;maraino](https://redirect.github.com/maraino) in [#&#8203;758](https://redirect.github.com/smallstep/crypto/pull/758)

Dependencies:

- Bump google.golang.org/grpc from 1.71.1 to 1.72.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;756](https://redirect.github.com/smallstep/crypto/pull/756)
- Bump cloud.google.com/go/kms from 1.21.1 to 1.21.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;757](https://redirect.github.com/smallstep/crypto/pull/757)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.61.0...v0.62.0>

### [`v0.61.0`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.61.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.60.0...v0.61.0)

#### What's Changed

- Fix templates with policies with Go 1.24 by [@&#8203;maraino](https://redirect.github.com/maraino) in [#&#8203;744](https://redirect.github.com/smallstep/crypto/pull/744)

Dependencies:

- Bump golang.org/x/net from 0.37.0 to 0.38.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;742](https://redirect.github.com/smallstep/crypto/pull/742)
- Bump modernc.org/sqlite from 1.36.2 to 1.37.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;741](https://redirect.github.com/smallstep/crypto/pull/741)
- Bump google.golang.org/api from 0.227.0 to 0.228.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;740](https://redirect.github.com/smallstep/crypto/pull/740)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.29.10 to 1.29.12 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;739](https://redirect.github.com/smallstep/crypto/pull/739)
- Add a methods to verify Google's CloudHSM attestations by [@&#8203;maraino](https://redirect.github.com/maraino) in [#&#8203;720](https://redirect.github.com/smallstep/crypto/pull/720)
- Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.38.1 to 1.38.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;745](https://redirect.github.com/smallstep/crypto/pull/745)
- Bump go.uber.org/mock from 0.5.0 to 0.5.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;747](https://redirect.github.com/smallstep/crypto/pull/747)
- Bump google.golang.org/grpc from 1.71.0 to 1.71.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;749](https://redirect.github.com/smallstep/crypto/pull/749)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.29.12 to 1.29.13 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;750](https://redirect.github.com/smallstep/crypto/pull/750)
- Bump github.com/Azure/azure-sdk-for-go/sdk/azcore from 1.17.1 to 1.18.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;746](https://redirect.github.com/smallstep/crypto/pull/746)
- Bump golang.org/x/crypto from 0.36.0 to 0.37.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;748](https://redirect.github.com/smallstep/crypto/pull/748)
- Bump golang.org/x/net from 0.38.0 to 0.39.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;755](https://redirect.github.com/smallstep/crypto/pull/755)
- Bump github.com/Azure/azure-sdk-for-go/sdk/azidentity from 1.8.2 to 1.9.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;754](https://redirect.github.com/smallstep/crypto/pull/754)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.29.13 to 1.29.14 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;752](https://redirect.github.com/smallstep/crypto/pull/752)
- Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.38.2 to 1.38.3 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;751](https://redirect.github.com/smallstep/crypto/pull/751)
- Bump google.golang.org/api from 0.228.0 to 0.229.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;753](https://redirect.github.com/smallstep/crypto/pull/753)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.60.0...v0.61.0>

### [`v0.60.0`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.60.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.59.2...v0.60.0)

#### What's Changed

- Add support for tokens with no pins by [@&#8203;maraino](https://redirect.github.com/maraino) in [#&#8203;730](https://redirect.github.com/smallstep/crypto/pull/730)

Dependencies:

- Bump github.com/Azure/azure-sdk-for-go/sdk/azcore from 1.17.0 to 1.17.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;737](https://redirect.github.com/smallstep/crypto/pull/737)
- Bump modernc.org/sqlite from 1.36.1 to 1.36.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;736](https://redirect.github.com/smallstep/crypto/pull/736)
- Bump google.golang.org/protobuf from 1.36.5 to 1.36.6 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;733](https://redirect.github.com/smallstep/crypto/pull/733)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.29.9 to 1.29.10 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;734](https://redirect.github.com/smallstep/crypto/pull/734)
- Bump google.golang.org/api from 0.226.0 to 0.227.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;735](https://redirect.github.com/smallstep/crypto/pull/735)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.59.2...v0.60.0>

### [`v0.59.2`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.59.2)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.59.1...v0.59.2)

#### What's Changed

- Fix open files with a full path in softkms by [@&#8203;maraino](https://redirect.github.com/maraino) in [#&#8203;732](https://redirect.github.com/smallstep/crypto/pull/732)

##### Dependencies

- Bump golang.org/x/sys from 0.30.0 to 0.31.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;723](https://redirect.github.com/smallstep/crypto/pull/723)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.29.6 to 1.29.9 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;722](https://redirect.github.com/smallstep/crypto/pull/722)
- Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.37.18 to 1.38.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;725](https://redirect.github.com/smallstep/crypto/pull/725)
- Bump golang.org/x/crypto from 0.35.0 to 0.36.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;721](https://redirect.github.com/smallstep/crypto/pull/721)
- Bump golang.org/x/net from 0.35.0 to 0.37.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;724](https://redirect.github.com/smallstep/crypto/pull/724)
- Bump google.golang.org/grpc from 1.70.0 to 1.71.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;729](https://redirect.github.com/smallstep/crypto/pull/729)
- Bump google.golang.org/api from 0.223.0 to 0.226.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;728](https://redirect.github.com/smallstep/crypto/pull/728)
- Bump modernc.org/sqlite from 1.36.0 to 1.36.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;726](https://redirect.github.com/smallstep/crypto/pull/726)
- Bump cloud.google.com/go/kms from 1.21.0 to 1.21.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;727](https://redirect.github.com/smallstep/crypto/pull/727)
- Bump github.com/golang-jwt/jwt/v5 from 5.2.1 to 5.2.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;731](https://redirect.github.com/smallstep/crypto/pull/731)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.59.1...v0.59.2>

### [`v0.59.1`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.59.1)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.59.0...v0.59.1)

#### What's Changed

- Fix `uint32` casting on `linux/arm` by [@&#8203;hslatman](https://redirect.github.com/hslatman) in [#&#8203;714](https://redirect.github.com/smallstep/crypto/pull/714)
- Bump modernc.org/sqlite from 1.35.0 to 1.36.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;718](https://redirect.github.com/smallstep/crypto/pull/718)
- Bump github.com/google/go-tpm-tools from 0.4.4 to 0.4.5 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;719](https://redirect.github.com/smallstep/crypto/pull/719)
- Bump cloud.google.com/go/kms from 1.20.5 to 1.21.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;715](https://redirect.github.com/smallstep/crypto/pull/715)
- Bump google.golang.org/api from 0.221.0 to 0.223.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;716](https://redirect.github.com/smallstep/crypto/pull/716)
- Bump golang.org/x/crypto from 0.33.0 to 0.35.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;717](https://redirect.github.com/smallstep/crypto/pull/717)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.59.0...v0.59.1>

### [`v0.59.0`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.59.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.58.1...v0.59.0)

#### What's Changed

- Bump github.com/go-jose/go-jose/v3 from 3.0.3 to 3.0.4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;712](https://redirect.github.com/smallstep/crypto/pull/712)
- Add WithMinLenPasswordFile pemutil option by [@&#8203;dopey](https://redirect.github.com/dopey) in [#&#8203;711](https://redirect.github.com/smallstep/crypto/pull/711)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.58.1...v0.59.0>

### [`v0.58.1`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.58.1)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.58.0...v0.58.1)

#### What's Changed

- Use utf8 string for private key subject with non-printable characters by [@&#8203;areed](https://redirect.github.com/areed) in [#&#8203;710](https://redirect.github.com/smallstep/crypto/pull/710)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.58.0...v0.58.1>

### [`v0.58.0`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.58.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.57.1...v0.58.0)

#### What's Changed

- Add nssdb package by [@&#8203;areed](https://redirect.github.com/areed) in [#&#8203;690](https://redirect.github.com/smallstep/crypto/pull/690)
- Replace conflicting nssdb objects by [@&#8203;areed](https://redirect.github.com/areed) in [#&#8203;695](https://redirect.github.com/smallstep/crypto/pull/695)
- Fix `gosec` (and other) linter issues and remove `smallstep/assert` by [@&#8203;hslatman](https://redirect.github.com/hslatman) in [#&#8203;704](https://redirect.github.com/smallstep/crypto/pull/704)
- Use  SHA-256 for subject key id on FIPS 140-3 mode by [@&#8203;maraino](https://redirect.github.com/maraino) in [#&#8203;703](https://redirect.github.com/smallstep/crypto/pull/703)
- Add `CreateDecrypter` support to AWS KMS by [@&#8203;hslatman](https://redirect.github.com/hslatman) in [#&#8203;702](https://redirect.github.com/smallstep/crypto/pull/702)

##### Dependencies

- Bump google.golang.org/protobuf from 1.36.3 to 1.36.4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;694](https://redirect.github.com/smallstep/crypto/pull/694)
- Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.37.14 to 1.37.16 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;693](https://redirect.github.com/smallstep/crypto/pull/693)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.29.2 to 1.29.4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;691](https://redirect.github.com/smallstep/crypto/pull/691)
- Bump google.golang.org/api from 0.218.0 to 0.219.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;692](https://redirect.github.com/smallstep/crypto/pull/692)
- Bump golang.org/x/sys from 0.29.0 to 0.30.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;699](https://redirect.github.com/smallstep/crypto/pull/699)
- Bump golang.org/x/net from 0.34.0 to 0.35.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;697](https://redirect.github.com/smallstep/crypto/pull/697)
- Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.37.16 to 1.37.18 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;698](https://redirect.github.com/smallstep/crypto/pull/698)
- Bump google.golang.org/api from 0.219.0 to 0.220.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;700](https://redirect.github.com/smallstep/crypto/pull/700)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.29.4 to 1.29.6 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;709](https://redirect.github.com/smallstep/crypto/pull/709)
- Bump modernc.org/sqlite from 1.34.5 to 1.35.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;707](https://redirect.github.com/smallstep/crypto/pull/707)
- Bump github.com/Azure/azure-sdk-for-go/sdk/azidentity from 1.8.1 to 1.8.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;705](https://redirect.github.com/smallstep/crypto/pull/705)
- Bump google.golang.org/protobuf from 1.36.4 to 1.36.5 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;706](https://redirect.github.com/smallstep/crypto/pull/706)
- Bump google.golang.org/api from 0.220.0 to 0.221.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;708](https://redirect.github.com/smallstep/crypto/pull/708)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.57.1...v0.58.0>

### [`v0.57.1`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.57.1)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.57.0...v0.57.1)

#### What's Changed

- Bump google.golang.org/api from 0.216.0 to 0.217.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;679](https://redirect.github.com/smallstep/crypto/pull/679)
- Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.37.8 to 1.37.13 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;678](https://redirect.github.com/smallstep/crypto/pull/678)
- Bump google.golang.org/protobuf from 1.36.2 to 1.36.3 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;680](https://redirect.github.com/smallstep/crypto/pull/680)
- Bump github.com/Azure/azure-sdk-for-go/sdk/azidentity from 1.8.0 to 1.8.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;681](https://redirect.github.com/smallstep/crypto/pull/681)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.28.10 to 1.29.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;682](https://redirect.github.com/smallstep/crypto/pull/682)
- Fix flaky behavior for tests using sprig's `dateInZone` by [@&#8203;hslatman](https://redirect.github.com/hslatman) in [#&#8203;683](https://redirect.github.com/smallstep/crypto/pull/683)
- Bump cloud.google.com/go/kms from 1.20.4 to 1.20.5 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;688](https://redirect.github.com/smallstep/crypto/pull/688)
- Bump google.golang.org/grpc from 1.69.4 to 1.70.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;687](https://redirect.github.com/smallstep/crypto/pull/687)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.29.1 to 1.29.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;686](https://redirect.github.com/smallstep/crypto/pull/686)
- Bump google.golang.org/api from 0.217.0 to 0.218.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;685](https://redirect.github.com/smallstep/crypto/pull/685)
- Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.37.13 to 1.37.14 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;684](https://redirect.github.com/smallstep/crypto/pull/684)
- Extend jose.NewSigner support for other types by [@&#8203;maraino](https://redirect.github.com/maraino) in [#&#8203;689](https://redirect.github.com/smallstep/crypto/pull/689)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.57.0...v0.57.1>

### [`v0.57.0`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.57.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.56.0...v0.57.0)

#### What's Changed

- Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.37.7 to 1.37.8 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;661](https://redirect.github.com/smallstep/crypto/pull/661)
- Bump golang.org/x/net from 0.32.0 to 0.33.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;660](https://redirect.github.com/smallstep/crypto/pull/660)
- Bump github.com/google/go-tpm from 0.9.2 to 0.9.3 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;659](https://redirect.github.com/smallstep/crypto/pull/659)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.28.6 to 1.28.7 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;657](https://redirect.github.com/smallstep/crypto/pull/657)
- Bump github.com/googleapis/gax-go/v2 from 2.14.0 to 2.14.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;658](https://redirect.github.com/smallstep/crypto/pull/658)
- Bump google.golang.org/protobuf from 1.36.0 to 1.36.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;665](https://redirect.github.com/smallstep/crypto/pull/665)
- Bump cloud.google.com/go/kms from 1.20.2 to 1.20.3 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;662](https://redirect.github.com/smallstep/crypto/pull/662)
- Bump google.golang.org/api from 0.212.0 to 0.214.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;664](https://redirect.github.com/smallstep/crypto/pull/664)
- Bump google.golang.org/grpc from 1.69.0 to 1.69.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;663](https://redirect.github.com/smallstep/crypto/pull/663)
- Bump cloud.google.com/go/kms from 1.20.3 to 1.20.4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;666](https://redirect.github.com/smallstep/crypto/pull/666)
- Bump golang.org/x/sys from 0.28.0 to 0.29.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;669](https://redirect.github.com/smallstep/crypto/pull/669)
- Bump golang.org/x/crypto from 0.31.0 to 0.32.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;667](https://redirect.github.com/smallstep/crypto/pull/667)
- Bump golang.org/x/net from 0.33.0 to 0.34.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;668](https://redirect.github.com/smallstep/crypto/pull/668)
- Fix TPM manufacturer lookup for names with a space (`0x20`) character by [@&#8203;hslatman](https://redirect.github.com/hslatman) in [#&#8203;670](https://redirect.github.com/smallstep/crypto/pull/670)
- Bump google.golang.org/api from 0.214.0 to 0.216.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;671](https://redirect.github.com/smallstep/crypto/pull/671)
- Bump github.com/Azure/azure-sdk-for-go/sdk/azcore from 1.16.0 to 1.17.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;675](https://redirect.github.com/smallstep/crypto/pull/675)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.28.7 to 1.28.10 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;674](https://redirect.github.com/smallstep/crypto/pull/674)
- Bump google.golang.org/grpc from 1.69.2 to 1.69.4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;672](https://redirect.github.com/smallstep/crypto/pull/672)
- Bump google.golang.org/protobuf from 1.36.1 to 1.36.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;673](https://redirect.github.com/smallstep/crypto/pull/673)
- support for yubikey management key from a file by [@&#8203;maraino](https://redirect.github.com/maraino) in [#&#8203;676](https://redirect.github.com/smallstep/crypto/pull/676)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.56.0...v0.57.0>

### [`v0.56.0`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.56.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.55.0...v0.56.0)

#### What's Changed

- Bump github.com/stretchr/testify from 1.9.0 to 1.10.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;640](https://redirect.github.com/smallstep/crypto/pull/640)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.28.5 to 1.28.6 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;641](https://redirect.github.com/smallstep/crypto/pull/641)
- Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.37.6 to 1.37.7 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;642](https://redirect.github.com/smallstep/crypto/pull/642)
- Bump cloud.google.com/go/kms from 1.20.1 to 1.20.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;646](https://redirect.github.com/smallstep/crypto/pull/646)
- Bump golang.org/x/crypto from 0.29.0 to 0.30.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;648](https://redirect.github.com/smallstep/crypto/pull/648)
- Bump google.golang.org/api from 0.209.0 to 0.210.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;645](https://redirect.github.com/smallstep/crypto/pull/645)
- Bump google.golang.org/grpc from 1.67.1 to 1.68.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;647](https://redirect.github.com/smallstep/crypto/pull/647)
- Bump google.golang.org/protobuf from 1.35.2 to 1.36.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;653](https://redirect.github.com/smallstep/crypto/pull/653)
- Bump golang.org/x/crypto from 0.30.0 to 0.31.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;652](https://redirect.github.com/smallstep/crypto/pull/652)
- Bump google.golang.org/grpc from 1.68.1 to 1.69.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;651](https://redirect.github.com/smallstep/crypto/pull/651)
- Bump github.com/google/go-tpm from 0.9.1 to 0.9.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;654](https://redirect.github.com/smallstep/crypto/pull/654)
- Bump google.golang.org/api from 0.210.0 to 0.212.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;650](https://redirect.github.com/smallstep/crypto/pull/650)
- Add support for imported keys by [@&#8203;maraino](https://redirect.github.com/maraino) in [#&#8203;656](https://redirect.github.com/smallstep/crypto/pull/656)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.55.0...v0.56.0>

### [`v0.55.0`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.55.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.54.2...v0.55.0)

#### What's Changed

- Add go.uber.org/mock to tools.go, generate mocks using go run by [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) in [#&#8203;633](https://redirect.github.com/smallstep/crypto/pull/633)
- Add `NSING`, `WiseKey` and `SecEdge` to known manufacturers by [@&#8203;hslatman](https://redirect.github.com/hslatman) in [#&#8203;634](https://redirect.github.com/smallstep/crypto/pull/634)
- Bump google.golang.org/api from 0.205.0 to 0.209.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;637](https://redirect.github.com/smallstep/crypto/pull/637)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.28.3 to 1.28.5 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;638](https://redirect.github.com/smallstep/crypto/pull/638)
- Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.37.5 to 1.37.6 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;636](https://redirect.github.com/smallstep/crypto/pull/636)
- Add max-sessions parameter on pkcs11 kms by [@&#8203;maraino](https://redirect.github.com/maraino) in [#&#8203;639](https://redirect.github.com/smallstep/crypto/pull/639)

#### New Contributors

- [@&#8203;marten-seemann](https://redirect.github.com/marten-seemann) made their first contribution in [#&#8203;633](https://redirect.github.com/smallstep/crypto/pull/633)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.54.2...v0.55.0>

### [`v0.54.2`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.54.2)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.54.1...v0.54.2)

#### What's Changed

- Allow to retry with no tags when loading privatekeys by [@&#8203;maraino](https://redirect.github.com/maraino) in [#&#8203;631](https://redirect.github.com/smallstep/crypto/pull/631)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.54.1...v0.54.2>

### [`v0.54.1`](https://redirect.github.com/smallstep/crypto/releases/tag/v0.54.1)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.54.0...v0.54.1)

#### What's Changed

- Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.36.3 to 1.37.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;606](https://redirect.github.com/smallstep/crypto/pull/606)
- Bump golang.org/x/net from 0.29.0 to 0.30.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;607](https://redirect.github.com/smallstep/crypto/pull/607)
- Bump google.golang.org/protobuf from 1.34.2 to 1.35.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;609](https://redirect.github.com/smallstep/crypto/pull/609)
- Bump google.golang.org/grpc from 1.67.0 to 1.67.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;610](https://redirect.github.com/smallstep/crypto/pull/610)
- Bump google.golang.org/api from 0.199.0 to 0.200.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;615](https://redirect.github.com/smallstep/crypto/pull/615)
- Bump github.com/Azure/azure-sdk-for-go/sdk/azidentity from 1.7.0 to 1.8.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;611](https://redirect.github.com/smallstep/crypto/pull/611)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.27.39 to 1.27.43 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;612](https://redirect.github.com/smallstep/crypto/pull/612)
- Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.37.1 to 1.37.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;614](https://redirect.github.com/smallstep/crypto/pull/614)
- Bump github.com/Azure/azure-sdk-for-go/sdk/azcore from 1.14.0 to 1.15.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;613](https://redirect.github.com/smallstep/crypto/pull/613)
- Bump github.com/Azure/azure-sdk-for-go/sdk/azcore from 1.15.0 to 1.16.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;618](https://redirect.github.com/smallstep/crypto/pull/618)
- Bump google.golang.org/api from 0.200.0 to 0.201.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;617](https://redirect.github.com/smallstep/crypto/pull/617)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.27.43 to 1.28.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;616](https://redirect.github.com/smallstep/crypto/pull/616)
- Bump cloud.google.com/go/kms from 1.20.0 to 1.20.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;619](https://redirect.github.com/smallstep/crypto/pull/619)
- Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.37.2 to 1.37.3 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;620](https://redirect.github.com/smallstep/crypto/pull/620)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.28.0 to 1.28.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;622](https://redirect.github.com/smallstep/crypto/pull/622)
- Bump github.com/go-piv/piv-go/v2 from 2.2.0 to 2.3.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;623](https://redirect.github.com/smallstep/crypto/pull/623)
- Bump google.golang.org/api from 0.203.0 to 0.204.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;624](https://redirect.github.com/smallstep/crypto/pull/624)
- Bump github.com/aws/aws-sdk-go-v2/service/kms from 1.37.3 to 1.37.5 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;629](https://redirect.github.com/smallstep/crypto/pull/629)
- Bump golang.org/x/sys from 0.26.0 to 0.27.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;627](https://redirect.github.com/smallstep/crypto/pull/627)
- Bump github.com/aws/aws-sdk-go-v2/config from 1.28.1 to 1.28.3 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;625](https://redirect.github.com/smallstep/crypto/pull/625)
- Bump google.golang.org/api from 0.204.0 to 0.205.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;626](https://redirect.github.com/smallstep/crypto/pull/626)
- Fix cgo data conversion for bytes and strings by [@&#8203;maraino](https://redirect.github.com/maraino) in [#&#8203;630](https://redirect.github.com/smallstep/crypto/pull/630)

**Full Changelog**: <https://github.com/smallstep/crypto/compare/v0.54.0...v0.54.1>

### [`v0.54.0`](https://redirect.github.com/smallstep/crypto/compare/v0.53.0...v0.54.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.53.0...v0.54.0)

### [`v0.53.0`](https://redirect.github.com/smallstep/crypto/compare/v0.52.0...v0.53.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.52.0...v0.53.0)

### [`v0.52.0`](https://redirect.github.com/smallstep/crypto/compare/v0.51.2...v0.52.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.51.2...v0.52.0)

### [`v0.51.2`](https://redirect.github.com/smallstep/crypto/compare/v0.51.1...v0.51.2)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.51.1...v0.51.2)

### [`v0.51.1`](https://redirect.github.com/smallstep/crypto/compare/v0.51.0...v0.51.1)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.51.0...v0.51.1)

### [`v0.51.0`](https://redirect.github.com/smallstep/crypto/compare/v0.50.0...v0.51.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.50.0...v0.51.0)

### [`v0.50.0`](https://redirect.github.com/smallstep/crypto/compare/v0.49.0...v0.50.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.49.0...v0.50.0)

### [`v0.49.0`](https://redirect.github.com/smallstep/crypto/compare/v0.48.1...v0.49.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.48.1...v0.49.0)

### [`v0.48.1`](https://redirect.github.com/smallstep/crypto/compare/v0.48.0...v0.48.1)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.48.0...v0.48.1)

### [`v0.48.0`](https://redirect.github.com/smallstep/crypto/compare/v0.47.1...v0.48.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.47.1...v0.48.0)

### [`v0.47.1`](https://redirect.github.com/smallstep/crypto/compare/v0.47.0...v0.47.1)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.47.0...v0.47.1)

### [`v0.47.0`](https://redirect.github.com/smallstep/crypto/compare/v0.46.0...v0.47.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.46.0...v0.47.0)

### [`v0.46.0`](https://redirect.github.com/smallstep/crypto/compare/v0.45.1...v0.46.0)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.45.1...v0.46.0)

### [`v0.45.1`](https://redirect.github.com/smallstep/crypto/compare/v0.45.0...v0.45.1)

[Compare Source](https://redirect.github.com/smallstep/crypto/compare/v0.45.0...v0.45.1)

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

### Comment by @red-hat-konflux on November 12, 2025 at 12:35 AM UTC

### ℹ Artifact update notice

##### File name: rh_identity_transform/go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 2 additional dependencies were updated


Details:


| **Package**              | **Change**             |
| :----------------------- | :--------------------- |
| `golang.org/x/time`      | `v0.12.0` -> `v0.14.0` |
| `google.golang.org/grpc` | `v1.71.0` -> `v1.76.0` |

---

## Reviews

### Review by @Hyperkid123 - Approved on November 21, 2025 at 09:49 AM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/67*

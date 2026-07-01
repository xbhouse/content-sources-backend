---
type: pull_request
number: 1183
title: "fix: ensure openssl is installed in ubi8-micro"
state: merged
author: SteveHNH
created: 2023-03-02T15:28:30Z
updated: 2023-03-06T15:51:18Z
closed: 2023-03-06T15:51:18Z
merged: 2023-03-06T15:51:18Z
base_branch: master
head_branch: add_openssl
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1183
---

# Pull Request #1183: fix: ensure openssl is installed in ubi8-micro

**Author**: @SteveHNH
**Created**: March 02, 2023 at 03:28 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `add_openssl`

## Description

## Secure Coding Practices Checklist GitHub Link
- https://github.com/RedHatInsights/secure-coding-checklist

## Secure Coding Checklist
- [x] Input Validation
- [x] Output Encoding
- [x] Authentication and Password Management
- [x] Session Management
- [x] Access Control
- [x] Cryptographic Practices
- [x] Error Handling and Logging
- [x] Data Protection
- [x] Communication Security
- [x] System Configuration
- [x] Database Security
- [x] File Management
- [x] Memory Management
- [x] General Coding Practices


---

## Discussion

### Comment by @jira-linking on March 02, 2023 at 03:28 PM UTC

Commits missing Jira IDs:
0e9e0b909565606c29159f47a0e36c13ddf62fee
d624d173006c4092453ae8a5975a121cc5692137
dd15aa8297c598dfb3855818bd191ddf5c6b0b29
623af5c4769b8d1118c711aa562e34cfc60e8961
e6cbbe0335ae5400aa56c9f15d55a571cccf4683
276c907ccf2ae7dcf64c87dd4d8d21f034cda1b4


### Comment by @codecov-commenter on March 02, 2023 at 03:38 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1183?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`56.25`**% and project coverage change: **`-0.42`** :warning:
> Comparison is base [(`4999506`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/499950637e9720a44342710aa18ca5843aa44b6d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.21% compared to head [(`7f0a2e1`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1183?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.80%.

> :exclamation: Current head 7f0a2e1 differs from pull request most recent head 276c907. Consider uploading reports for the commit 276c907 to get more accurate results

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1183      +/-   ##
==========================================
- Coverage   63.21%   62.80%   -0.42%     
==========================================
  Files         102      103       +1     
  Lines        5878     5985     +107     
==========================================
+ Hits         3716     3759      +43     
- Misses       1696     1755      +59     
- Partials      466      471       +5     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `62.80% <56.25%> (-0.42%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1183?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisory\_systems\_v3.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1183?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zX3YzLmdv) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/baseline\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1183?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | `76.92% <ø> (ø)` | |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1183?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `75.00% <ø> (ø)` | |
| [manager/controllers/package\_versions.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1183?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3ZlcnNpb25zLmdv) | `69.04% <ø> (ø)` | |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1183?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `75.92% <ø> (ø)` | |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1183?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `73.84% <ø> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1183?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `70.45% <ø> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1183?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `78.94% <ø> (ø)` | |
| [manager/controllers/systemtags.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1183?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW10YWdzLmdv) | `80.55% <ø> (ø)` | |
| [manager/middlewares/authentication.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1183?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9hdXRoZW50aWNhdGlvbi5nbw==) | `20.28% <0.00%> (-0.30%)` | :arrow_down: |
| ... and [15 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1183?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1183?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @SteveHNH on March 02, 2023 at 08:41 PM UTC

Attempted to just copy the openssl binary and libraries but was still unable to overcome the error. Went back to the original idea of doing a dnf install against a temp rootfs and then copying that into the buildimg. Picked up that method from this dockerfile: https://github.com/fatherlinux/ubi-micro/blob/master/ubi8-micro-openssl

It gets the pod to run properly, but if someone wants to try using the ldd method, they're welcome to it. I think I've taken it as far as I can on my own with limited knowledge of the app itself.

### Comment by @MichaelMraka on March 06, 2023 at 12:40 PM UTC

Hello @SteveHNH , I think it's likely some configuration in /etc/ what's missing.
I can try it, is there a way how to check the image works in fedramp?

### Comment by @SteveHNH on March 06, 2023 at 02:14 PM UTC

@MichaelMraka sure! If you want to take a crack at it and ping me here or on slack, I can run the image in FR and see if it works there. As long as the pr builds an image for quay, we can use it.

### Comment by @syncrou on March 06, 2023 at 03:34 PM UTC

@MichaelMraka due to the tight deadline on this. Do you see a real issue with getting this PR merged now since @SteveHNH knows it works in our environment?

It would be great for us to be able to check another thing off our list so we can keep moving forward.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1183*

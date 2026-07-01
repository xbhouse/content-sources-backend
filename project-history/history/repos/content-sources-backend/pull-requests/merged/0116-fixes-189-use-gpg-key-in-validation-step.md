---
type: pull_request
number: 116
title: "Fixes 189: use gpg key in validation step"
state: merged
author: jlsherrill
created: 2022-09-29T19:21:36Z
updated: 2022-10-18T08:00:26Z
closed: 2022-10-18T07:53:48Z
merged: 2022-10-18T07:53:47Z
base_branch: main
head_branch: 189
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/116
---

# Pull Request #116: Fixes 189: use gpg key in validation step

**Author**: @jlsherrill
**Created**: September 29, 2022 at 07:21 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `189`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on September 29, 2022 at 07:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-189

### Comment by @jlsherrill on September 29, 2022 at 07:35 PM UTC

Testing:

1.  see that signature is present:
```
curl -X POST localhost:8000/api/content-sources/v1/repository_parameters/validate/    -H "`./scripts/header.sh 1`"  -H 'Content-Type: application/json'   -d '[{"url":"https://packages.microsoft.com/yumrepos/azure-cli/"}]'  | jq

```
2.  check one that isn't:
```
curl -X POST localhost:8000/api/content-sources/v1/repository_parameters/validate/    -H "`./scripts/header.sh 1`"  -H 'Content-Type: application/json'   -d '[{"url":"https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/"}]'  | jq
```


3.  Try the good case, with the correct gpg key:
 Create validation.json  file with contents:
```
[
  {"url":"https://packages.microsoft.com/yumrepos/azure-cli/",
    "metadata_verification": true,
    "gpg_key":"-----BEGIN PGP PUBLIC KEY BLOCK-----\nVersion: GnuPG v1.4.7 (GNU/Linux)\n\nmQENBFYxWIwBCADAKoZhZlJxGNGWzqV+1OG1xiQeoowKhssGAKvd+buXCGISZJwT\nLXZqIcIiLP7pqdcZWtE9bSc7yBY2MalDp9Liu0KekywQ6VVX1T72NPf5Ev6x6DLV\n7aVWsCzUAF+eb7DC9fPuFLEdxmOEYoPjzrQ7cCnSV4JQxAqhU4T6OjbvRazGl3ag\nOeizPXmRljMtUUttHQZnRhtlzkmwIrUivbfFPD+fEoHJ1+uIdfOzZX8/oKHKLe2j\nH632kvsNzJFlROVvGLYAk2WRcLu+RjjggixhwiB+Mu/A8Tf4V6b+YppS44q8EvVr\nM+QvY7LNSOffSO6Slsy9oisGTdfE39nC7pVRABEBAAG0N01pY3Jvc29mdCAoUmVs\nZWFzZSBzaWduaW5nKSA8Z3Bnc2VjdXJpdHlAbWljcm9zb2Z0LmNvbT6JATUEEwEC\nAB8FAlYxWIwCGwMGCwkIBwMCBBUCCAMDFgIBAh4BAheAAAoJEOs+lK2+EinPGpsH\n/32vKy29Hg51H9dfFJMx0/a/F+5vKeCeVqimvyTM04C+XENNuSbYZ3eRPHGHFLqe\nMNGxsfb7C7ZxEeW7J/vSzRgHxm7ZvESisUYRFq2sgkJ+HFERNrqfci45bdhmrUsy\n7SWw9ybxdFOkuQoyKD3tBmiGfONQMlBaOMWdAsic965rvJsd5zYaZZFI1UwTkFXV\nKJt3bp3Ngn1vEYXwijGTa+FXz6GLHueJwF0I7ug34DgUkAFvAs8Hacr2DRYxL5RJ\nXdNgj4Jd2/g6T9InmWT0hASljur+dJnzNiNCkbn9KbX7J/qK1IbR8y560yRmFsU+\nNdCFTW7wY0Fb1fWJ+/KTsC4=\n=J6gs\n-----END PGP PUBLIC KEY BLOCK-----\n"}
  ]
```
run:
```
curl -X POST localhost:8000/api/content-sources/v1/repository_parameters/validate/    -H "`./scripts/header.sh 1`"  -H 'Content-Type: application/json'   -d @validate.json | jq
```

4. try with the wrong gpg key:
Create a file  validate_wrong_key.json with contents:
```
[
  {"url":"https://packages.microsoft.com/yumrepos/azure-cli/",
    "metadata_verification": true,
    "gpg_key":"-----BEGIN PGP PUBLIC KEY BLOCK-----\nVersion: GnuPG v1\n\n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\n=fdzz\n-----END PGP PUBLIC KEY BLOCK-----\n"}
  ]
```
and run:
```
curl -X POST localhost:8000/api/content-sources/v1/repository_parameters/validate/    -H "`./scripts/header.sh 1`"  -H 'Content-Type: application/json'   -d @validate_wrong_key.json | jq 

```

5.  provide  an invalid key:
```
curl -X POST localhost:8000/api/content-sources/v1/repository_parameters/validate/    -H "`./scripts/header.sh 1`"  -H 'Content-Type: application/json'   -d '[{"url":"https://packages.microsoft.com/yumrepos/azure-cli/", "gpg_key":"asdf"}]'  | jq
```

---

## Reviews

### Review by @swadeley - Commented on September 30, 2022 at 08:32 AM UTC

### Review by @swadeley - Commented on September 30, 2022 at 08:32 AM UTC

### Review by @swadeley - Commented on September 30, 2022 at 08:33 AM UTC

### Review by @swadeley - Commented on September 30, 2022 at 08:34 AM UTC

### Review by @swadeley - Commented on September 30, 2022 at 08:40 AM UTC

### Review by @rverdile - Commented on October 11, 2022 at 07:29 PM UTC

### Review by @jlsherrill - Commented on October 11, 2022 at 07:34 PM UTC

### Review by @jlsherrill - Commented on October 11, 2022 at 07:46 PM UTC

### Review by @rverdile - Commented on October 11, 2022 at 08:29 PM UTC

### Review by @jlsherrill - Commented on October 11, 2022 at 08:50 PM UTC

### Review by @jlsherrill - Commented on October 12, 2022 at 02:47 PM UTC

### Review by @jlsherrill - Commented on October 12, 2022 at 02:48 PM UTC

### Review by @rverdile - Approved on October 12, 2022 at 03:14 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/116*

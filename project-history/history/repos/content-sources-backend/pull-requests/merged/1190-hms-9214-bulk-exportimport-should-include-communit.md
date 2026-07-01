---
type: pull_request
number: 1190
title: "HMS-9214: bulk export/import should include community and rh repos"
state: merged
author: xbhouse
created: 2025-09-04T21:34:13Z
updated: 2025-09-17T20:55:16Z
closed: 2025-09-17T20:55:16Z
merged: 2025-09-17T20:55:16Z
base_branch: main
head_branch: 9214
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1190
---

# Pull Request #1190: HMS-9214: bulk export/import should include community and rh repos

**Author**: @xbhouse
**Created**: September 04, 2025 at 09:34 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `9214`

## Description

## Summary

- Includes the community and Red Hat orgs when exporting and importing repos

## Testing steps

1. Import and snapshot a community, Red Hat, and custom repo

2. Use the bulk export API to export all these repos:
```
POST /repositories/bulk_export/
{
    "repository_uuids": [
        "<community_repo_uuid>", "<redhat_repo_uuid>", "<custom_repo_uuid>" 
    ]
}
```
You should see no errors and a response similar to this:
```
[
    {
        "name": "EPEL 10 Everything x86_64",
        "url": "https://dl.fedoraproject.org/pub/epel/10/Everything/x86_64/",
        "origin": "community",
        "distribution_versions": [
            "10"
        ],
        "distribution_arch": "x86_64",
        "gpg_key": "-----BEGIN PGP PUBLIC KEY BLOCK-----\n...\n=mhQZ\n-----END PGP PUBLIC KEY BLOCK-----",
        "metadata_verification": false,
        "module_hotfixes": false,
        "snapshot": true
    },
    {
        "name": "Red Hat CodeReady Linux Builder for RHEL 9 ARM 64 (RPMs)",
        "url": "https://cdn.redhat.com/content/dist/rhel9/9/aarch64/codeready-builder/os/",
        "origin": "red_hat",
        "distribution_versions": [
            "9"
        ],
        "distribution_arch": "aarch64",
        "gpg_key": "-----BEGIN PGP PUBLIC KEY BLOCK-----\n\nmQ...OfLHYjnefg==\n=UZd/\n-----END PGP PUBLIC KEY BLOCK-----\n\n",
        "metadata_verification": false,
        "module_hotfixes": false,
        "snapshot": true
    },
    {
        "name": "test-custom",
        "url": "https://content-services.github.io/fixtures/yum/comps-modules/v1/",
        "origin": "external",
        "distribution_versions": [
            "any"
        ],
        "distribution_arch": "any",
        "gpg_key": "",
        "metadata_verification": false,
        "module_hotfixes": false,
        "snapshot": true
    }
]
```

3. Use the bulk import API to import these repos in the same org, using the response from the bulk export API as the body:
```
POST /repositories/bulk_import/
<response from export as body of request>
```
You should see a response similar to what you see when repos are bulk created and warnings for each repo that the URL already exists. No repos should be created. 

4. Repeat step 3 with a different user/org in your identity header. You should still see warnings for the community and Red Hat repos, and the custom repo should be created.

---

## Discussion

### Comment by @jlsherrill on September 04, 2025 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-9214

### Comment by @swadeley on September 10, 2025 at 11:00 AM UTC

Hi @xbhouse 

After adding:
 ` community_repos:
    enabled: true
` 

to `configs/config.yaml`  (and doing another `make run`)

I was able to 
`curl -H "$( ./scripts/header.sh 9999 1111 )" "http://localhost:8000/api/content-sources/v1.0/repositories/?search=EPEL"`
to list the three repos:
```
""meta":{"limit":100,"offset":0,"count":3},"links":{"first":"/api/content-sources/v1.0/repositories/?limit=100\u0026offset=0\u0026search=EPEL","last":"/api/content-sources/v1.0/repositories/?limit=100\u0026offset=0\u0026search=EPEL"}}
"
```

then I could test:
```
sjw@t14:~/content-sources-backend$ curl -X POST -H "$( ./scripts/header.sh 9999 1111 )" -H "Content-Type: application/json" -d '{"repository_uuids": ["95821630-8b2f-4043-ac36-1083ee00923b", "ec2a9040-9cb8-47ae-ab5d-db5aae5f1548", "ac70bc7d-3812-4d66-99fe-fed84356e64c"]}' "http://localhost:8000/api/content-sources/v1.0/repositories/bulk_export/"
```

lgtm

### Comment by @xbhouse on September 10, 2025 at 12:05 PM UTC

thank you for the review @swadeley :) after talking with Ryan, we decided this would be better solved in image-builder-crc, so I'm closing this PR 

EDIT: reopened but with different behavior :smile: 

---

## Reviews

### Review by @swadeley - Dismissed on September 10, 2025 at 11:01 AM UTC

ACK

### Review by @rverdile - Approved on September 16, 2025 at 06:09 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1190*

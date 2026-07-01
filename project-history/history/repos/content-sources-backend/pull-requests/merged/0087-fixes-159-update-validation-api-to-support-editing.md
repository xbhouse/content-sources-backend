---
type: pull_request
number: 87
title: "Fixes 159: Update validation API to support editing"
state: merged
author: Andrewgdewar
created: 2022-08-24T19:26:38Z
updated: 2022-09-06T18:30:36Z
closed: 2022-09-06T18:17:06Z
merged: 2022-09-06T18:17:06Z
base_branch: main
head_branch: CONTENT-159
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/87
---

# Pull Request #87: Fixes 159: Update validation API to support editing

**Author**: @Andrewgdewar
**Created**: August 24, 2022 at 07:26 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `CONTENT-159`

## Description

This adds the ability for the validation API to EXCLUDE the sent UUIDs from its name/url uniqueness check, thus allowing the ability to use this API to validate against within the repository edit screen. 

This solves the problem of false positive (where the uniqueness check would fail) when validating a repo effectively against itself or others in the edit group.

An example: 

Previously this: 
```json
[
   {
        "name": "google",
        "url": "https://www.google.com"
        // "uuid": "a68feaa3-9746-4719-8e33-3ff4688fed85",
    },
    {
        "name": "apple",
        "url": "https://apple.com", 
        //  "uuid": "9d6fb177-ff10-41fc-81d4-ea12ae963042",
    }
]
```

would result in:
```json
[
    {
        "name": {
            "skipped": false,
            "valid": false,
            "error": "A repository with the name 'google' already exists."
        },
        "url": {
            "skipped": false,
            "valid": false,
            "error": "A repository with the URL 'https://www.google.com' already exists.",
            "http_code": 0,
            "metadata_present": false
        }
    },
    {
        "name": {
            "skipped": false,
            "valid": false,
            "error": "A repository with the name 'apple' already exists."
        },
        "url": {
            "skipped": false,
            "valid": false,
            "error": "A repository with the URL 'https://apple.com' already exists.",
            "http_code": 0,
            "metadata_present": false
        }
    }
]
```

Now well editing the above targeted repos, by adding in the commented UUIDs to exclude, we get the following: 
```json
[
    {
        "name": {
            "skipped": false,
            "valid": true,
            "error": ""
        },
        "url": {
            "skipped": false,
            "valid": true,
            "error": "",
            "http_code": 200,
            "metadata_present": true
        }
    },
    {
        "name": {
            "skipped": false,
            "valid": true,
            "error": ""
        },
        "url": {
            "skipped": false,
            "valid": true,
            "error": "",
            "http_code": 200,
            "metadata_present": true
        }
    }
]
```

Which is the expected outcome for validating when editing.


---

## Discussion

### Comment by @jlsherrill on August 24, 2022 at 07:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-159

### Comment by @Andrewgdewar on September 01, 2022 at 08:59 PM UTC

> can you add tests to `dao/repository_test.go` and and `handler/repository_parameters_test.go`, that cover the cases with uuid?

I may need some help with the handler/repository_parameters_test.go, I don't really get what that is supposed to be doing?

### Comment by @rverdile on September 02, 2022 at 01:01 PM UTC

So in `repository_parameters.go` you have this block of code:
```
excludedUUIDs := []string{}
	for i := 0; i < repoCount; i++ {
		if validationParams[i].UUID != nil {
			excludedUUIDs = append(excludedUUIDs, *validationParams[i].UUID)
		}
	}
```

In `repository_parameters_test.go` there's these 3 request:
```
requestBody := []api.RepositoryValidationRequest{
		{
			Name: pointy.String("myValidateRepo"),
		},
		{
			URL: pointy.String("http://myrepo.com"),
		},
		{},
	}
```

None of these requests have a UUID passed to them, so the code changes you've made related to the UUID kinda get skipped over. You might just add UUID to the first and/or second request so that part of the code gets hit.

---

## Reviews

### Review by @swadeley - Commented on August 25, 2022 at 09:45 AM UTC

Some suggestions

### Review by @Andrewgdewar - Commented on August 25, 2022 at 03:04 PM UTC

### Review by @rverdile - Commented on August 25, 2022 at 03:14 PM UTC

### Review by @rverdile - Commented on August 25, 2022 at 03:18 PM UTC

### Review by @Andrewgdewar - Commented on August 25, 2022 at 03:23 PM UTC

### Review by @rverdile - Commented on August 25, 2022 at 03:26 PM UTC

### Review by @Andrewgdewar - Commented on August 25, 2022 at 03:33 PM UTC

### Review by @Andrewgdewar - Commented on August 25, 2022 at 03:36 PM UTC

### Review by @rverdile - Commented on August 25, 2022 at 05:18 PM UTC

### Review by @Andrewgdewar - Commented on August 26, 2022 at 02:30 PM UTC

### Review by @rverdile - Commented on August 26, 2022 at 05:14 PM UTC

### Review by @rverdile - Commented on August 29, 2022 at 02:06 PM UTC

### Review by @Andrewgdewar - Commented on August 29, 2022 at 02:52 PM UTC

### Review by @rverdile - Commented on August 29, 2022 at 03:33 PM UTC

### Review by @rverdile - Commented on August 29, 2022 at 04:05 PM UTC

can you add tests to `dao/repository_test.go`  and and `handler/repository_parameters_test.go`, that cover the cases with uuid?

### Review by @rverdile - Commented on September 02, 2022 at 12:31 PM UTC

### Review by @Andrewgdewar - Commented on September 02, 2022 at 05:33 PM UTC

### Review by @rverdile - Approved on September 06, 2022 at 02:41 PM UTC

nice

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/87*

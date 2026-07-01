---
type: pull_request
number: 639
title: "Fixes 3599: Add candlepin content for templates"
state: merged
author: rverdile
created: 2024-04-16T12:49:01Z
updated: 2024-05-15T13:27:56Z
closed: 2024-05-15T13:27:40Z
merged: 2024-05-15T13:27:40Z
base_branch: main
head_branch: candlepin-content
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/639
---

# Pull Request #639: Fixes 3599: Add candlepin content for templates

**Author**: @rverdile
**Created**: April 16, 2024 at 12:49 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `candlepin-content`

## Description

## Summary
Adds candlepin product/pool/environment/content creation on template create. If updating a template, content is removed from environments if necessary.

## Testing steps
1. Create  a custom repository and a red hat repository.
2. Create a content template with these repositories included. Wait for task to finish.
3. The will be many things to check in candlepin: the new product created, the new pool created, the new environment created for the template, the new product content, the new content promoted to the environment.

Here is an example to list the products. Make sure you include the Authorization header. It is the base64 encoded username:password (admin:admin by default).
```
GET http://localhost:8181/candlepin/owners/content-sources-test/products
Accept: application/json
Authorization: Basic YWRtaW46YWRtaW4=
```
4. Refer to the [candlepin REST API spec](https://petstore.swagger.io/?url=https://raw.githubusercontent.com/candlepin/candlepin/main/api/candlepin-api-spec.yaml#/) to see how to query more endpoints.
5. Make sure the content URLs and labels are set correctly. Keep in mind, these will be different for custom vs Red Hat repositories.
6. Check the content URL against the base URL of the pulp distribution associated to the repository in the template.
7. Update the template to remove a repository. Wait for the task to finish.
8. The content set for that repository should no longer be in the environment.
9. Reference the JIRA card to make sure you're checking all the things.

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on April 16, 2024 at 01:00 PM UTC

https://issues.redhat.com/browse/HMS-3599

### Comment by @jlsherrill on April 18, 2024 at 05:28 PM UTC

Do you want to go ahead and add context passing and set the X-Correlation-ID header on candelpin requests similar to what i've done here: https://github.com/content-services/content-sources-backend/pull/637
or do it as a followup issue?



### Comment by @xbhouse on April 18, 2024 at 05:34 PM UTC

checking the products endpoint, i see the base url of the pulp distribution path (which i guess is also the repoconfig uuid?) in the content url for the custom repo, but not for the red hat repo. is this expected? 

red hat repo:
```
"contentUrl": "templates/34f84a07-60ee-4f19-b811-6aedbcc0f224/content/dist/layered/rhel8/x86_64/ansible/2/os",
"id": "aa19b2338be64ba3a342520979912602",
"label": "ansible-2-for-rhel-8-x86_64-rpms"
```

pulp distribution path: `aa19b233-8be6-4ba3-a342-520979912602/5275e346-a507-47a3-9ecd-0ef982aaece7`

repoconfig uuid: `aa19b233-8be6-4ba3-a342-520979912602`


custom repo:
```
"contentUrl": "templates/34f84a07-60ee-4f19-b811-6aedbcc0f224/d09b6122-5c2e-435f-b341-3a397c02b32f",
"id": "d09b61225c2e435fb3413a397c02b32f",
"label": "test1"
```

pulp distribution path: `d09b6122-5c2e-435f-b341-3a397c02b32f/d6de4b72-f1fb-4fb3-9a27-dbf37835e3b4`

repoconfig uuid: `d09b6122-5c2e-435f-b341-3a397c02b32f`

### Comment by @xbhouse on April 18, 2024 at 05:52 PM UTC

should the environment include the red hat repos? when querying environments i see only the custom repo referenced in environmentContent, not sure if that's expected. i used this endpoint `http://localhost:8181/candlepin/owners/content-sources-test/environments` 
```
[
    {
        "created": "2024-04-18T17:14:02+0000",
        "updated": "2024-04-18T17:14:02+0000",
        "id": "34f84a0760ee4f19b8116aedbcc0f224",
        "name": "34f84a0760ee4f19b8116aedbcc0f224",
        "type": null,
        "description": null,
        "contentPrefix": null,
        "owner": {
            "id": "2c9780868ef1fe49018ef1fe632a0002",
            "key": "content-sources-test",
            "displayName": "ContentSourcesTest",
            "href": "/owners/content-sources-test",
            "contentAccessMode": "org_environment"
        },
        "environmentContent": [
            {
                "contentId": "d09b61225c2e435fb3413a397c02b32f",
                "enabled": true
            }
        ]
    }
]
```        

EDIT: i do see the red hat repo content id referenced here when removing the custom repo from the template:
```
"environmentContent": [
            {
                "contentId": "aa19b2338be64ba3a342520979912602",
                "enabled": true
            }
        ] 
```        

### Comment by @rverdile on April 19, 2024 at 08:17 PM UTC

> Do you want to go ahead and add context passing and set the X-Correlation-ID header on candelpin requests similar to what i've done here: #637

Oh, I'll do it here.

### Comment by @rverdile on April 26, 2024 at 08:26 PM UTC

> should the environment include the red hat repos? when querying environments i see only the custom repo referenced in environmentContent

Yes, this was a bug and now it should be fixed. Good catch!

### Comment by @rverdile on April 26, 2024 at 08:30 PM UTC

> checking the products endpoint, i see the base url of the pulp distribution

It's expected that the base URLs in pulp are different. It sorta it is for candlepin too. The content URL of the custom repo content actually doesn't really matter because we're going to override that using the context overrides feature in Candlepin.

### Comment by @rverdile on April 26, 2024 at 08:30 PM UTC

Not sure why the test is failing here and not locally, need to investigate

### Comment by @jlsherrill on May 01, 2024 at 12:28 PM UTC

I'm in the middle of testing, but i noticed one thing.  For red hat content sets, it looks like we're creating them with a 'randomized' name because they already exists, but for red hat contnet, they should already exist, and we should just use them as is.

```
Content:
	Type: rpm
	Name: Red Hat Enterprise Linux 9 for x86_64 - BaseOS (RPMs)
	Label: rhel-9-for-x86_64-baseos-rpms_cbFyHWHJEp

```

### Comment by @jlsherrill on May 06, 2024 at 04:38 PM UTC

I noticed we weren't using the template name for the environment name, instead opting for the UUID, was there a reason for that?  These names will eventually show up on the client

### Comment by @rverdile on May 06, 2024 at 04:59 PM UTC

> I noticed we weren't using the template name for the environment name, instead opting for the UUID, was there a reason for that? These names will eventually show up on the client

No, I just wasn't thinking about it that way. I can change it to the name of the template.

### Comment by @jlsherrill on May 07, 2024 at 05:20 PM UTC

created a template and added the two base rhel9 repos and a custom repo and the update-template-content task failed with:

runtime error: index out of range [2] with length 1

when i removed the custom repo from the template it seemed to work fine

### Comment by @jlsherrill on May 14, 2024 at 12:39 PM UTC

oh one more small thing.  If i comment out the candlepin server URL it get an error on the task after creating a template.  This will be the case with stage for a bit.  We probably want to just not run the task if its not configured.

---

## Reviews

### Review by @rverdile - Commented on April 16, 2024 at 02:54 PM UTC

### Review by @jlsherrill - Commented on April 18, 2024 at 05:39 PM UTC

### Review by @jlsherrill - Commented on April 18, 2024 at 05:40 PM UTC

### Review by @jlsherrill - Commented on April 18, 2024 at 05:51 PM UTC

### Review by @jlsherrill - Commented on April 18, 2024 at 06:12 PM UTC

### Review by @rverdile - Commented on April 19, 2024 at 08:17 PM UTC

### Review by @jlsherrill - Commented on April 22, 2024 at 02:57 PM UTC

### Review by @jlsherrill - Commented on April 22, 2024 at 02:57 PM UTC

### Review by @jlsherrill - Commented on April 22, 2024 at 02:57 PM UTC

### Review by @jlsherrill - Commented on April 22, 2024 at 03:06 PM UTC

### Review by @rverdile - Commented on April 23, 2024 at 05:04 PM UTC

### Review by @rverdile - Commented on April 24, 2024 at 03:31 PM UTC

### Review by @rverdile - Commented on April 24, 2024 at 03:32 PM UTC

### Review by @jlsherrill - Commented on April 25, 2024 at 03:05 PM UTC

### Review by @rverdile - Commented on April 26, 2024 at 08:27 PM UTC

### Review by @rverdile - Commented on April 26, 2024 at 08:27 PM UTC

### Review by @rverdile - Commented on April 26, 2024 at 08:28 PM UTC

### Review by @jlsherrill - Commented on May 01, 2024 at 02:44 PM UTC

### Review by @rverdile - Commented on May 01, 2024 at 03:04 PM UTC

### Review by @rverdile - Commented on May 01, 2024 at 05:22 PM UTC

### Review by @jlsherrill - Commented on May 03, 2024 at 07:30 PM UTC

### Review by @jlsherrill - Commented on May 03, 2024 at 07:50 PM UTC

### Review by @jlsherrill - Commented on May 06, 2024 at 02:21 PM UTC

### Review by @rverdile - Commented on May 06, 2024 at 05:15 PM UTC

### Review by @jlsherrill - Commented on May 06, 2024 at 05:23 PM UTC

### Review by @jlsherrill - Commented on May 09, 2024 at 12:27 AM UTC

### Review by @jlsherrill - Commented on May 09, 2024 at 06:14 PM UTC

### Review by @jlsherrill - Commented on May 09, 2024 at 07:58 PM UTC

### Review by @jlsherrill - Commented on May 10, 2024 at 03:32 PM UTC

### Review by @jlsherrill - Commented on May 10, 2024 at 03:32 PM UTC

### Review by @jlsherrill - Commented on May 10, 2024 at 04:07 PM UTC

### Review by @rverdile - Commented on May 10, 2024 at 04:24 PM UTC

### Review by @jlsherrill - Commented on May 10, 2024 at 04:39 PM UTC

### Review by @rverdile - Commented on May 10, 2024 at 04:47 PM UTC

### Review by @rverdile - Commented on May 10, 2024 at 04:48 PM UTC

### Review by @rverdile - Commented on May 10, 2024 at 04:49 PM UTC

### Review by @rverdile - Commented on May 13, 2024 at 06:53 PM UTC

### Review by @jlsherrill - Commented on May 13, 2024 at 07:10 PM UTC

### Review by @jlsherrill - Commented on May 13, 2024 at 08:17 PM UTC

### Review by @jlsherrill - Commented on May 14, 2024 at 12:10 PM UTC

### Review by @jlsherrill - Commented on May 14, 2024 at 12:24 PM UTC

### Review by @jlsherrill - Approved on May 14, 2024 at 07:41 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/639*

---
type: pull_request
number: 870
title: "Fixes 3877: Add required fields for API spec requests"
state: merged
author: Andrewgdewar
created: 2024-10-31T20:31:37Z
updated: 2024-11-18T16:30:54Z
closed: 2024-11-18T15:53:36Z
merged: 2024-11-18T15:53:36Z
base_branch: main
head_branch: HMS-3877
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/870
---

# Pull Request #870: Fixes 3877: Add required fields for API spec requests

**Author**: @Andrewgdewar
**Created**: October 31, 2024 at 08:31 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-3877`

## Description

## Summary

- Add required fields for API spec requests.

## Testing steps



---

## Discussion

### Comment by @jlsherrill on October 31, 2024 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-3877

### Comment by @swadeley on November 01, 2024 at 04:38 PM UTC

Hi

How to test these changes?

If I take the first one:
```
   "definitions": {
        "api.AddUploadsRequest": {
            "type": "object",
            "required": [
                "artifacts",
                "uploads"
            ],
```
Looking at "add_upload" in ephemeral I do not see anything different to that in stage:
```
app.content_sources.rest_client.repositories_api.add_upload.
                                                                      allowed_values  call_with_http_info() headers_map           params_map           
                                                                      api_client            callable()            location_map          settings             
                                                                      attribute_map         collection_format_map openapi_types         validations        
```  


### Comment by @jlsherrill on November 05, 2024 at 09:58 PM UTC

I don't know that the changes need to be tested, but if you regenerate the api bindings you will see that code will likely need to be changed.  for example, this method signature changes:

```
@@ -102,9 +102,12 @@ class ApiContentUnitSearchRequest(ModelNormal):
     ])

     @convert_js_args_to_python_args
-    def __init__(self, *args, **kwargs):  # noqa: E501
+    def __init__(self, search, *args, **kwargs):  # noqa: E501
         """ApiContentUnitSearchRequest - a model defined in OpenAPI

+        Args:
+            search (str): Search string to search content unit names
+
```

So  if you are using the generated bindings its clearer that 'search' is required to be specified.   This doesn't actually change the behavior of the API at all (hence why the tests all pass without generating the bindings)

### Comment by @Andrewgdewar on November 12, 2024 at 05:39 PM UTC

Thanks for the review @rverdile @xbhouse! 


### Comment by @rverdile on November 12, 2024 at 08:47 PM UTC

there's some tests failing :)

you might have to rebase as well

### Comment by @Andrewgdewar on November 13, 2024 at 06:09 PM UTC

> I thought I saw them before, but TemplateRequest and RepositoryRequest are missing

Thanks for the feedback here, I did add to the template and repository request. Please confirm if those changes make sense.

### Comment by @rverdile on November 14, 2024 at 04:09 PM UTC

that's my last comment :)

### Comment by @Andrewgdewar on November 15, 2024 at 03:40 PM UTC

@rverdile @xbhouse Thank you for your thorough review here, I'm still learning 🍼 👶, I appreciate your patience!!

### Comment by @swadeley on November 15, 2024 at 08:45 PM UTC

/retest

### Comment by @swadeley on November 18, 2024 at 01:45 PM UTC

Hi

I found that I had to add `origin="external` to all repo create calls.
No changes to anything related to search this time (search calls had `origin="external`  from a while  back)

### Comment by @rverdile on November 18, 2024 at 02:15 PM UTC

Ah, it seems origin is actually not a required field. We should remove that then.

---

## Reviews

### Review by @rverdile - Commented on November 11, 2024 at 08:18 PM UTC

### Review by @rverdile - Commented on November 11, 2024 at 08:24 PM UTC

### Review by @rverdile - Commented on November 11, 2024 at 08:35 PM UTC

### Review by @rverdile - Commented on November 11, 2024 at 08:36 PM UTC

### Review by @xbhouse - Commented on November 12, 2024 at 02:23 PM UTC

### Review by @rverdile - Commented on November 12, 2024 at 08:49 PM UTC

### Review by @rverdile - Commented on November 12, 2024 at 08:54 PM UTC

I thought I saw them before, but TemplateRequest and RepositoryRequest are missing

### Review by @rverdile - Commented on November 12, 2024 at 08:55 PM UTC

### Review by @rverdile - Commented on November 14, 2024 at 04:06 PM UTC

### Review by @rverdile - Approved on November 15, 2024 at 06:17 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/870*

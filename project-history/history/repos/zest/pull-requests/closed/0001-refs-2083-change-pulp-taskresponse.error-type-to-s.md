---
type: pull_request
number: 1
title: "Refs 2083: change pulp TaskResponse.Error type to string"
state: closed
author: rverdile
created: 2023-07-18T19:30:53Z
updated: 2023-07-20T15:01:56Z
closed: 2023-07-20T15:01:56Z
base_branch: main
head_branch: error-type
labels: []
url: https://github.com/content-services/zest/pull/1
---

# Pull Request #1: Refs 2083: change pulp TaskResponse.Error type to string

**Author**: @rverdile
**Created**: July 18, 2023 at 07:30 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `error-type`

## Description

Manually change pulp_api.json to correct the type of the properties in TaskResponse.Error

Currently looks like:

```
 "TaskResponse":  {
...
    "error":  
       {
           "type": "object",
           "additionalProperties": {
                "type": "object"
            },
               "readOnly": true,
                "description": "A JSON Object of a fatal error encountered during the execution of this task."
      }
...
}
```

Should look like:
```
 "TaskResponse":  {
...
    "error":  
       {
           "type": "object",
           "additionalProperties": {
                "type": "string"
            },
               "readOnly": true,
                "description": "A JSON Object of a fatal error encountered during the execution of this task."
      }
...
}
```

Because the values of the fields in the error object are string:
```
"error": {
    "traceback": "  File \"/...",
    "description": "Cannot connect..."
  },
```

And according to [this]( https://swagger.io/docs/specification/data-models/dictionaries/) "The additionalProperties keyword specifies the type of values in the dictionary."



---

*Archived from: https://github.com/content-services/zest/pull/1*

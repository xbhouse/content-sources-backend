---
type: pull_request
number: 39
title: "Refs 47: db:add preliminar entity relationship diagram"
state: merged
author: avisiedo
created: 2022-06-21T17:19:29Z
updated: 2022-07-19T20:20:43Z
closed: 2022-07-14T21:19:08Z
merged: 2022-07-14T21:19:08Z
base_branch: main
head_branch: add-er-diagram
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/39
---

# Pull Request #39: Refs 47: db:add preliminar entity relationship diagram

**Author**: @avisiedo
**Created**: June 21, 2022 at 05:19 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `add-er-diagram`

## Description

This PR just add a draft for an entity-relationship to be reviewed.

Signed-off-by: Alejandro Visiedo <avisiedo@redhat.com>

---

## Discussion

### Comment by @jlsherrill on July 05, 2022 at 03:23 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-47

### Comment by @jlsherrill on July 05, 2022 at 03:26 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @jlsherrill on July 05, 2022 at 06:43 PM UTC

mentioned this on chat, but also wanted to record here, thoughts on using https://github.com/achiku/planter ?

### Comment by @avisiedo on July 08, 2022 at 02:35 PM UTC

I think is nice to use planter and keep the er-diagram in sync with the database schema! :)  I will add makefile rules to just run one command to get it! thanks

### Comment by @rverdile on July 08, 2022 at 02:55 PM UTC

+1, I like planter

### Comment by @jlsherrill on July 11, 2022 at 08:19 PM UTC

Would you want to add the rendered image to our docs?  we could either put it in our architecture doc:  https://github.com/content-services/content-sources-backend/blob/main/docs/architecture.md

or just link to the rendered image in the list here: https://github.com/content-services/content-sources-backend#more-info

### Comment by @jlsherrill on July 11, 2022 at 08:42 PM UTC

speaking of the rendered picture, wanna include that here too? :)

### Comment by @avisiedo on July 12, 2022 at 06:53 AM UTC

I think I have found a way that let to include the rendered diagram without to add the rendered image to the repository.

**Downsides**:
- It require the `.puml` file to be merged to render it.
- If the proxy service fails, the image is not rendered.

**Upside**: Allow to get the rendered diagram without upload the image file to the repository.

Below an example referencing the diagram on the branch to be merged:

```raw
![Database Model](https://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/avisiedo/content-sources-backend/add-er-diagram/docs/db-model.puml)
```

![Database Model](https://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/avisiedo/content-sources-backend/add-er-diagram/docs/db-model.puml)

### Comment by @avisiedo on July 12, 2022 at 09:58 AM UTC

I have updated a dynamic a reference that will render the diagram based on the `docs/db-model.puml` file at master branch; else, I can always add the `.png` file and link it.

It would render this: [dynamic image rendered](https://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/avisiedo/content-sources-backend/add-er-diagram/docs/db-model.puml).

Let me know your thoughts!

### Comment by @jlsherrill on July 19, 2022 at 08:20 PM UTC

https://issues.redhat.com/browse/14649800

---

## Reviews

### Review by @swadeley - Commented on June 29, 2022 at 01:48 PM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 04:04 PM UTC

### Review by @jlsherrill - Commented on July 11, 2022 at 08:15 PM UTC

### Review by @avisiedo - Commented on July 12, 2022 at 06:33 AM UTC

### Review by @jlsherrill - Approved on July 14, 2022 at 06:55 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/39*

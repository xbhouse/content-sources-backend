---
type: issue
number: 985
title: "Need unique IDs for Next/Back buttons to allow success tracking of Create patch template wizard"
state: open
author: andreasn
created: 2023-03-09T09:23:05Z
updated: 2023-08-01T13:54:09Z
labels: []
url: https://github.com/RedHatInsights/patchman-ui/issues/985
---

# Issue #985: Need unique IDs for Next/Back buttons to allow success tracking of Create patch template wizard

**Author**: @andreasn
**Created**: March 09, 2023 at 09:23 AM UTC
**Status**: Open
**Labels**: None

## Description

I'm a UX researcher working on RHEL/HMS on consoledot.
I want to learn if there are any step at in a wizard where users get stuck and abandons the flow. Once that step is identified as a pitfall, I would set up usability test sessions with a smaller group of people and figure out what goes wrong in this particular step. I've done it for the Image Builder and Provision wizards so far, but now I want to do it with the patch template wizard as well.
I will use Pendo for tracking, but I need some unique IDs for the buttons for Pendo to hook up to at each step in the wizard.
Something along the lines of:
```
wizard-newtemplate-next-button
wizard-newtemplate-back-button
wizard-systems-next-button
wizard-systems-back-button
wizard-review-next-button
wizard-review-back-button
```

---

## Discussion

### Comment by @andreasn on March 09, 2023 at 09:24 AM UTC

@croissanne has some code for this that might be reusable here

### Comment by @croissanne on March 09, 2023 at 09:28 AM UTC

We used custom buttons for this https://github.com/RedHatInsights/image-builder-frontend/blob/main/src/Components/CreateImageWizard/formComponents/CustomButtons.js.

### Comment by @ezr-ondrej on March 09, 2023 at 05:12 PM UTC

We did the same for provisioning in, for more inspiration https://github.com/RHEnVision/provisioning-frontend/pull/193 :)

### Comment by @leSamo on March 14, 2023 at 01:38 PM UTC

Hello, thank you for providing examples! We are in the process of reworking the Template wizard and we will include this change as well.

### Comment by @leSamo on August 01, 2023 at 01:54 PM UTC

This issue is blocked by https://github.com/data-driven-forms/react-forms/issues/1379

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/issues/985*

---
type: pull_request
number: 964
title: "Draft: investigate bundle size"
state: closed
author: jschuler
created: 2023-02-16T16:20:57Z
updated: 2023-04-18T20:36:43Z
closed: 2023-04-18T20:36:43Z
base_branch: master
head_branch: bundle-size
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/964
---

# Pull Request #964: Draft: investigate bundle size

**Author**: @jschuler
**Created**: February 16, 2023 at 04:20 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `bundle-size`

## Description

Try to bundle only esm packages. Before there was a mixture of js and esm.

Before with mixed dist/esm imports:
All: 3.56 MB
vendor: 3.36 MB

After only esm imports:
All: 3.18 MB
Vendor: 2.98 MB

However, the largest space offender appears to be when federated modules is used. In testing, after disabling `@redhat-cloud-services/frontend-components-config/federated-modules` plugin in prod.webpack.config.js, it can be seen:

Disable module federation:
All: 1.2 MB
Vendor: 1 MB

The issue appears to be that once module federation is enabled, dependencies are fully bundled (and not just the used imports).

I am not sure yet about a solution for this, but this issue is also being addressed here:
https://github.com/patternfly/patternfly-react/issues/8357

---

## Discussion

### Comment by @app-sre-bot on February 16, 2023 at 04:21 PM UTC

Can one of the admins verify this patch?

### Comment by @mkholjuraev on February 16, 2023 at 09:25 PM UTC

This seems very promising!  If we somehow get this PR merged, the bundle size decreases ~11%. When that new major solution gets finished, the bundle size drops! I am looking forward to seeing it.



### Comment by @mkholjuraev on February 20, 2023 at 11:41 AM UTC

@jschuler would you like to finish up the PR with only babel config changes? If we change the webpack configs, there are some issues that break the app. In case you are busy, let me know. I can finish this up for you.

 I think we can do something like this:
![image](https://user-images.githubusercontent.com/59481011/220096351-5d803e1b-1524-42e4-a59e-351a21f448ab.png)


### Comment by @jschuler on February 20, 2023 at 03:08 PM UTC

Hi @mkholjuraev feel free to close this PR and create a new one with the correct changes. This PR was only opened for demonstration purposes :) thank you

### Comment by @Hyperkid123 on February 21, 2023 at 02:48 PM UTC

@jschuler the issue is that module federation basically nukes tree-shaking.

When you mark a package as shared "@patternfly/react-core', webpack is unable to tree shake the module as it knows its supposed to be shared and can't remove any modules from it in case another container (apps) has the same module marked as shared. Image this scenario. Webpack would tree shake any module (like Button). Second containers that would consume already shared modules would however require the Button component. But in the shared PF package that Button component was removed by tree shaking and does not exist. And you would have a runtime error.

We have discussed this with the PF team and have a way how we can remove this problem. We have solved it already with our frontend components packages. Now you may ask why are we even sharing packages like PF. Well if a version require meant is met, the PF library gets loaded only once. So instead of loading PF modules 20x during a user session, we load them only 4x. (I don't have hard data just using these numbers as an example)
So the initial JS size is larger but during the user sessions, we actually save a lot of bandwidth.



---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/964*

---
type: pull_request
number: 1260
title: "Test: Stop deploying selenium image"
state: closed
author: swadeley
created: 2025-10-27T20:28:31Z
updated: 2026-01-13T06:53:46Z
closed: 2025-11-04T15:16:28Z
base_branch: main
head_branch: swadeley/test_removing_selenium_image
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1260
---

# Pull Request #1260: Test: Stop deploying selenium image

**Author**: @swadeley
**Created**: October 27, 2025 at 08:28 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/test_removing_selenium_image`

## Description

## Summary

  Part of phasing out of running IQE tests.

  All IQE UI testing has stopped.

## Testing steps

Check for impact in Konflux and automated merge requests to see if this breaks deployment.


---

## Discussion

### Comment by @swadeley on October 27, 2025 at 08:29 PM UTC

Hi @jlsherrill 

We can test removing IQE related things while not breaking deployment to ephemeral.

### Comment by @TenSt on October 31, 2025 at 01:57 PM UTC

@katarinazaprazna hey! I think the config you've edited looks good, let's add it to this PR or create a separate one and see how it will go.

### Comment by @swadeley on November 04, 2025 at 03:16 PM UTC

Hi

closing in favour of:  Test: Stop deploying IQE integration tests #1283 

---

## Reviews

### Review by @katarinazaprazna - Changes Requested on October 31, 2025 at 11:50 AM UTC

Great work on this PR! I noticed a few remaining artifacts from the UI testing stack that I believe we should clean up to prevent deployment issues.

I believe that setting the `IQE_SEL_IMAGE` parameter to an empty string while leaving the container that consumes it will cause the Kubernetes/OpenShift runtime to fail Pod creation. It will be unable to pull an image with a blank name.

Since UI tests are now fully deprecated, I believe that the goal should be to have only one container, the core IQE Test Runner.

I've put together a fix that removes all Selenium-related config. I would appreciate a senior colleague like @jlsherrill or @TenSt  reviewing the cleanup.

```
---
apiVersion: template.openshift.io/v1
kind: Template
metadata:
  name: content-sources-${NAME_STUB}
objects:
  - apiVersion: batch/v1
    kind: Job
    metadata:
      name: content-sources-${NAME_STUB}-${UID}
      annotations:
        "ignore-check.kube-linter.io/no-liveness-probe": "probes not required on Job pods"
        "ignore-check.kube-linter.io/no-readiness-probe": "probes not required on Job pods"
      labels:
        image-tag: ${IMAGE_TAG}
        iqe-image-tag: ${IQE_IMAGE_TAG}
        iqe-env: ${ENV_FOR_DYNACONF}
        iqe-plugin: ${IQE_PLUGINS}
    spec:
      backoffLimit: 0
      template:
        spec:
          imagePullSecrets:
            - name: quay-cloudservices-pull
          restartPolicy: Never
          containers:
            - name: content-sources-${NAME_STUB}-${IMAGE_TAG}-${UID}
              image: ${IQE_IMAGE}:${IQE_IMAGE_TAG}
              imagePullPolicy: Always
              args:
                - run
              env:
                - name: ENV_FOR_DYNACONF
                  value: ${ENV_FOR_DYNACONF}
                - name: DYNACONF_MAIN__use_beta
                  value: ${USE_BETA}
                - name: IQE_IBUTSU_SOURCE
                  value: content-sources-saas-${IMAGE_TAG}-${UID}
                - name: IQE_NETLOG
                  value: ${IQE_NETLOG}
                - name: IQE_PLUGINS
                  value: ${IQE_PLUGINS}
                - name: IQE_MARKER_EXPRESSION
                  value: ${IQE_MARKER_EXPRESSION}
                - name: IQE_FILTER_EXPRESSION
                  value: ${IQE_FILTER_EXPRESSION}
                - name: IQE_LOG_LEVEL
                  value: ${IQE_LOG_LEVEL}
                - name: IQE_REQUIREMENTS
                  value: ${IQE_REQUIREMENTS}
                - name: IQE_REQUIREMENTS_PRIORITY
                  value: ${IQE_REQUIREMENTS_PRIORITY}
                - name: IQE_TEST_IMPORTANCE
                  value: ${IQE_TEST_IMPORTANCE}
                - name: IQE_PARALLEL_ENABLED
                  value: ${IQE_PARALLEL_ENABLED}
                - name: DYNACONF_IQE_VAULT_LOADER_ENABLED
                  value: "true"
                - name: DYNACONF_IQE_VAULT_VERIFY
                  value: "true"
                - name: DYNACONF_IQE_VAULT_URL
                  valueFrom:
                    secretKeyRef:
                      key: url
                      name: iqe-vault
                      optional: true
                - name: DYNACONF_IQE_VAULT_MOUNT_POINT
                  valueFrom:
                    secretKeyRef:
                      key: mountPoint
                      name: iqe-vault
                      optional: true
                - name: DYNACONF_IQE_VAULT_ROLE_ID
                  valueFrom:
                    secretKeyRef:
                      key: roleId
                      name: iqe-vault
                      optional: true
                - name: DYNACONF_IQE_VAULT_SECRET_ID
                  valueFrom:
                    secretKeyRef:
                      key: secretId
                      name: iqe-vault
                      optional: true
              resources:
                limits:
                  cpu: "1"
                  memory: 1.5Gi
                requests:
                  cpu: 250m
                  memory: 512Mi
              terminationMessagePath: /dev/termination-log
              terminationMessagePolicy: File
parameters:
  - name: IMAGE_TAG
    value: ''
    required: false
    description: Optional app image tag for labeling
  - name: UID
    description: "Unique job name suffix"
    generate: expression
    from: "[a-z0-9]{6}"
  - name: IQE_IMAGE
    description: "container image path for the iqe plugin"
    value: quay.io/cloudservices/iqe-tests
  - name: IQE_IMAGE_TAG
    description: "container image tag for iqe plugin"
    value: 'insights-experiences'
  - name: ENV_FOR_DYNACONF
    value: stage_proxy
  - name: USE_BETA
    value: "true"
  - name: IQE_PLUGINS
    value: insights_experiences
  - name: IQE_MARKER_EXPRESSION
    value: 'content_sources'
  - name: IQE_FILTER_EXPRESSION
    value: ''
  - name: IQE_LOG_LEVEL
    value: info
  - name: IQE_REQUIREMENTS
    value: ''
  - name: IQE_REQUIREMENTS_PRIORITY
    value: ''
  - name: IQE_TEST_IMPORTANCE
    value: ''
  - name: IQE_NETLOG
    value: "1"
    description: Enable HTTP network logging in IQE (keep for API debugging)
  - name: NAME_STUB
    value: 'e2e-tests'
  - name: IQE_PARALLEL_ENABLED
    value: "false"
```

### Review by @rverdile - Commented on October 31, 2025 at 02:46 PM UTC

### Review by @jlsherrill - Commented on November 03, 2025 at 02:47 PM UTC

### Review by @swadeley - Commented on November 03, 2025 at 09:06 PM UTC

### Review by @swadeley - Commented on November 04, 2025 at 10:29 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1260*

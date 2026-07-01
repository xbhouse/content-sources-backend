---
type: pull_request
number: 125
title: "Fixes 233: update to go 1.18"
state: merged
author: jlsherrill
created: 2022-10-20T11:24:48Z
updated: 2022-11-01T12:50:51Z
closed: 2022-11-01T12:50:47Z
merged: 2022-11-01T12:50:47Z
base_branch: main
head_branch: 233
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/125
---

# Pull Request #125: Fixes 233: update to go 1.18

**Author**: @jlsherrill
**Created**: October 20, 2022 at 11:24 AM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `233`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on October 20, 2022 at 11:30 AM UTC

https://issues.redhat.com/browse/HMSCONTENT-233

### Comment by @Andrewgdewar on October 21, 2022 at 02:06 PM UTC

I vaguely remember there being some dependency that we were keeping this at 1.16 for?

### Comment by @jlsherrill on October 24, 2022 at 12:38 PM UTC

It was because the docker images didn't support a newer go, but the provisioning team found a good one that does :)

### Comment by @avisiedo on October 25, 2022 at 05:09 PM UTC

update: rebased from main

checking the image

### Comment by @jlsherrill on October 27, 2022 at 04:26 PM UTC

@avisiedo updated with your changes

### Comment by @avisiedo on October 31, 2022 at 09:15 AM UTC

/test

### Comment by @jlsherrill on October 31, 2022 at 08:38 PM UTC

actually saw this again:

```
 --- FAIL: TestRpmSuite (0.08s)
    --- FAIL: TestRpmSuite/TestRpmList (0.00s)
        suite.go:77: test panicked: runtime error: index out of range [0] with length 0
            goroutine 576 [running]:
            runtime/debug.Stack()
            	/opt/hostedtoolcache/go/1.18.7/x64/src/runtime/debug/stack.go:24 +0x65
            github.com/stretchr/testify/suite.failOnPanic(0xc000104ea0, {0x11b21a0, 0xc000165368})
            	/home/runner/go/pkg/mod/github.com/stretchr/testify@v1.8.0/suite/suite.go:77 +0x3b
            github.com/stretchr/testify/suite.Run.func1.1()
            	/home/runner/go/pkg/mod/github.com/stretchr/testify@v1.8.0/suite/suite.go:161 +0x252
            panic({0x11b21a0, 0xc000165368})
            	/opt/hostedtoolcache/go/1.18.7/x64/src/runtime/panic.go:838 +0x207
            github.com/content-services/content-sources-backend/pkg/dao.(*RpmSuite).TestRpmList(0xc00007f820)
            	/home/runner/work/content-sources-backend/content-sources-backend/pkg/dao/rpm_test.go:103 +0x9d8
            reflect.Value.call({0xc000586e40?, 0xc0004113d8?, 0xc000505d40?}, {0x11ff000, 0x4}, {0xc00032de70, 0x1, 0xc00032dd80?})
            	/opt/hostedtoolcache/go/1.18.7/x64/src/reflect/value.go:556 +0x845
            reflect.Value.Call({0xc000586e40?, 0xc0004113d8?, 0xc00007f820?}, {0xc00032de70, 0x1, 0x1})
            	/opt/hostedtoolcache/go/1.18.7/x64/src/reflect/value.go:339 +0xbf
            github.com/stretchr/testify/suite.Run.func1(0xc000104ea0)
            	/home/runner/go/pkg/mod/github.com/stretchr/testify@v1.8.0/suite/suite.go:175 +0x4b6
            testing.tRunner(0xc000104ea0, 0xc00048dd40)
            	/opt/hostedtoolcache/go/1.18.7/x64/src/testing/testing.go:1439 +0x102
            created by testing.(*T).Run
            	/opt/hostedtoolcache/go/1.18.7/x64/src/testing/testing.go:1486 +0x35f
```

---

## Reviews

### Review by @jlsherrill - Commented on October 20, 2022 at 11:25 AM UTC

### Review by @avisiedo - Changes Requested on October 26, 2022 at 07:58 PM UTC

Small changes detected for macos:

* at `mk/docker.mk` the rule `docker-build` need to specify the file by using `-f` option as `docker` use as default name `Dockerfile`; `podman` recognize both default files; `buildah` recognize both files too:

```makefile
docker-build:
    $(DOCKER) build $(DOCKER_BUILD_OPTS) -t "$(DOCKER_IMAGE)" $(DOCKER_CONTEXT_DIR) -f "$(DOCKER_DOCKERFILE)"
```

* It needs delete a line at `mk/variables.mk`, else it fails in macos because of the `override undefine ...`; deleting the line at `mk/variables.mk` fix it:

```makefile
override undefine LOAD_DB_CFG_WITH_YQ
```

As it is not used in any compare for `ifndef` or `ifdef` it does not impact any place, and that variable is not used in any other place; that line was added to clean-up the value.

---

In short:
- With the above changes, the image build in linux by using podman and in macos by using docker; after this it lgtm.
- It builds for x86_64 and arm64 architectures.
- The x86_64 version worked in the ephemeral environment.
- The base image is built from this [Dockerfile](https://github.com/quay/claircore/blob/main/etc/Dockerfile).

I tried to use the tool `dive`, which is very cool to browse the different layers from an image, but it does not let me go further than the current image definition: https://github.com/wagoodman/dive   :confused:

> I suggest to use the ubi go-toolset image when the ubi golang 1.18 is available for x86_64 and arm64 architectures.


### Review by @avisiedo - Approved on October 31, 2022 at 11:56 AM UTC

lgtm

I have filed a card for the test which is not related with this pr at: https://issues.redhat.com/browse/HMSCONTENT-252

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/125*

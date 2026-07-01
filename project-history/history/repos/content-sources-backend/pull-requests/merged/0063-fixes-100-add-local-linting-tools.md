---
type: pull_request
number: 63
title: "Fixes 100: add local linting tools"
state: merged
author: rverdile
created: 2022-07-28T17:55:11Z
updated: 2022-08-26T11:40:55Z
closed: 2022-08-26T11:40:55Z
merged: 2022-08-26T11:40:55Z
base_branch: main
head_branch: lint
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/63
---

# Pull Request #63: Fixes 100: add local linting tools

**Author**: @rverdile
**Created**: July 28, 2022 at 05:55 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `lint`

## Description

Adds two ways of linting, pre-commit or just golangci-lint, so everyone is happy :)

To install the pre-commit linter `make install-pre-commit`, and then it will run before committing.

To use the golangci-lint linter, `make install-golangci-lint` then `make lint`. the golangci-lint binary installs to ./release/golangci-lint

---

## Discussion

### Comment by @jlsherrill on July 28, 2022 at 06:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-100

### Comment by @jlsherrill on July 28, 2022 at 06:00 PM UTC

why is pre-commit needed?  can't you use https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks directly?  (keep in mind i've never actualy implemented a githook, so i may have  no idea what i'm talking about)

### Comment by @mshriver on July 28, 2022 at 07:15 PM UTC

@jlsherrill pre-commit provides some cool functionality, like creating its own environment to install all of the hook dependencies into instead of your development environment (or worse, system packages). 

It also provides a _ton_ of hooks to use right out of the gate.

https://pre-commit.com/hooks.html

If we want to take things a step further, pre-commit.ci exists too, and will auto-update the config for us.
https://pre-commit.ci/



### Comment by @jlsherrill on July 29, 2022 at 12:52 PM UTC

@mshriver  yeah i guess in my view having to install other dependencies in order to have the hooks run kinda defeats the point. 
 
If the goal of this is to help 'new' developers that may not be accustomed to working in go (or using a fancy ide), having the commit work fine unless they install a dependency seems to defeat the point (correct me if i'm wrong here, but i tried it and commited/pushed with no issue).  On the plus side it lets developers 'opt in', on the downside i think it defeats the purpose that you highlighted when you pushed a pr that failed linting.  Thoughts?

### Comment by @mshriver on July 29, 2022 at 02:11 PM UTC

I don't think its necessarily to help new developers, its to help meet PR check expectations before you push.  If you don't care to fail PR checks and have to amend/push additional commits, you don't install/use pre-commit and deal with the churn.  If you don't want that churn and to meet PR checks before pushing, you install pre-commit.

These dependencies aren't installed in production containers, aren't necessary for building, they're just for developers to meet coding standards. With Python we even use a tool like `Black` to completely enforce formatting, which will rewrite files completely if the person wrote in their own style. We're not going to that extent here with linting, but it at least lets you fix linting before letting GH fail it.

### Comment by @Andrewgdewar on July 29, 2022 at 03:11 PM UTC

Problem: I want to commit my code because my house is on fire.

Expected: I commit my code, and my house burns down. No problem, at least I got that commit in (_That's dedication right there_).

Actual: I commit my code, something prevents it.. I'm confused and attempt to debug the issues.. *burns alive*.

I don't think I'll be installing said dependencies.. 😆 


### Comment by @rverdile on July 29, 2022 at 03:45 PM UTC

I'll personally just run `make lint` before I'm ready to push to github. Pre-commit won't be useful for me, so if we implement it, I'm not opinionated about how we do it. With that said, I could see it blocking me more than helping me, and therefore always using the `-n` flag to get around it.

### Comment by @jlsherrill on July 29, 2022 at 03:52 PM UTC

[test]

### Comment by @mshriver on July 29, 2022 at 06:22 PM UTC

FWIW I've been using pre-commit on about a dozen repos for over 2 years and the above situations have happened maybe 3 times, where I've just re-run the commit with `-n` because I needed to move on from that working tree *this second*. (My homeowners insurance doesn't let me use candles anymore)

### Comment by @mshriver on July 29, 2022 at 09:38 PM UTC

[test]

### Comment by @jlsherrill on August 01, 2022 at 12:23 PM UTC

[test]

### Comment by @avisiedo on August 01, 2022 at 12:56 PM UTC

I think is good to have the tool available at least with `make lint` or `make lint-all`; the hooks can be configured for `commit` or `push` event (https://stackoverflow.com/questions/63820683/with-pre-commit-how-to-use-some-hooks-before-commit-and-others-before-push); if lint the files in every commit could be a barrier, maybe config it for push only could help more. If we want to push the changes into the remote repository, then I think do not install the pre-commit hook is the path (I have the opinion that I prefer to have the precommit hook installed anyway).

If dependencies is a problem, the python dependencies could be installed in a virtual environment adding a dependency to it into the rules; something like the below:

```makefile
ADD_PYTHON_ENV := source .venv/bin/activate && 

.PHONY: lint
lint: .venv
    $(ADD_PYTHON_ENV) pre-commit run

.PHONY: lint-all
lint-all: .venv
    $(ADD_PYTHON_ENV) pre-commit run --all

.venv:
    python -m venv .venv && $(ADD_PYTHON_ENV)  pip install -U pip && pip install pre-commit
```

In the above way, the python dependencies are installed into a local virtual environment (we keep our environment cleaner, and reduce the risk that some dependency could impact into our environment), so we don't have to install extra packages (well, python if it were not installed).

**upsides**:
- We align the changes with code standard before reach out the pr or even the remote branch (shift-left).
- It makes easy to add additional linters for different file types (golangci-lint for golang, hadolint for Dockerfiles, shellcheck for scripting files, markdown documents, yaml files, or even adding new hooks locally).
- Same tool can be used in pipeline, so that we can execute same configuration from pipeline and local cloned repository ([pre-commit.ci](https://pre-commit.ci/)). (In theory, I have to say that I have not used this in a pipeline yet).

**downsides**:
- when the pre-commit hook is installed, It could block to push or commit when we get linter hints; as it is said by @mshriver `--no-verify` or `-n` can be used and the hooks would not be executed.

### Comment by @mshriver on August 01, 2022 at 02:36 PM UTC

[test]


### Comment by @jlsherrill on August 01, 2022 at 07:03 PM UTC

ok to test

### Comment by @jlsherrill on August 01, 2022 at 07:03 PM UTC

/ok to test

### Comment by @jlsherrill on August 01, 2022 at 07:03 PM UTC

/ok-to-test

### Comment by @jlsherrill on August 01, 2022 at 07:32 PM UTC

[test]
ok to test
/ok-to-test

### Comment by @rverdile on August 08, 2022 at 12:17 PM UTC

updated

### Comment by @avisiedo on August 24, 2022 at 06:32 AM UTC

I can see pipeline is using version `v1.46.2` and 'install-golangci-lint' use `v1.47.2`; May I suggest to match `lint.mk` with the current version in the pipeline so that `install-golangci-lint` install `v1.46.2`? by the way, I can not replay the issue that is evoked in the pipeline from the local workstation; I should be missing something :pensive: ! for the rest, it LGTM! :+1: 

### Comment by @avisiedo on August 24, 2022 at 09:48 AM UTC

update: only the rebase to wipe the conflict

### Comment by @rverdile on August 25, 2022 at 03:34 PM UTC

changed the installed golangci-lint version to 1.46.2

### Comment by @avisiedo on August 26, 2022 at 04:51 AM UTC

lgtm

---

## Reviews

### Review by @mshriver - Commented on July 28, 2022 at 07:17 PM UTC

### Review by @mshriver - Commented on July 28, 2022 at 07:19 PM UTC

### Review by @rverdile - Commented on July 28, 2022 at 07:49 PM UTC

### Review by @rverdile - Commented on July 29, 2022 at 01:40 PM UTC

### Review by @avisiedo - Commented on July 29, 2022 at 04:56 PM UTC

### Review by @avisiedo - Commented on July 29, 2022 at 06:51 PM UTC

### Review by @avisiedo - Commented on July 30, 2022 at 07:20 AM UTC

### Review by @avisiedo - Commented on August 02, 2022 at 08:19 PM UTC

### Review by @avisiedo - Commented on August 02, 2022 at 08:30 PM UTC

### Review by @avisiedo - Commented on August 02, 2022 at 08:31 PM UTC

### Review by @avisiedo - Commented on August 02, 2022 at 08:32 PM UTC

### Review by @rverdile - Commented on August 04, 2022 at 03:05 PM UTC

### Review by @avisiedo - Commented on August 05, 2022 at 09:51 PM UTC

### Review by @avisiedo - Commented on August 05, 2022 at 09:52 PM UTC

### Review by @avisiedo - Commented on August 05, 2022 at 09:52 PM UTC

### Review by @avisiedo - Changes Requested on August 05, 2022 at 10:09 PM UTC

Just a couple of small comments!! thanks for this pr :+1: 

### Review by @avisiedo - Approved on August 26, 2022 at 04:50 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/63*

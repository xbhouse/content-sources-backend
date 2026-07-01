---
type: pull_request
number: 70
title: "Fixes 142: support cert via env variable"
state: merged
author: jlsherrill
created: 2022-08-03T19:57:23Z
updated: 2022-08-08T07:24:26Z
closed: 2022-08-08T07:24:26Z
merged: 2022-08-08T07:24:26Z
base_branch: main
head_branch: 142
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/70
---

# Pull Request #70: Fixes 142: support cert via env variable

**Author**: @jlsherrill
**Created**: August 03, 2022 at 07:57 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `142`

## Description

as well as via filename

---

## Discussion

### Comment by @jlsherrill on August 03, 2022 at 08:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-142

### Comment by @jlsherrill on August 03, 2022 at 09:11 PM UTC

did deployment changes as part of a 2nd pr:  https://github.com/content-services/content-sources-backend/pull/71



### Comment by @rverdile on August 04, 2022 at 03:52 PM UTC

I'm not sure how to test this, could you add a couple of testing steps?

### Comment by @jlsherrill on August 04, 2022 at 04:25 PM UTC

oh yes, that would be helpful!

First import the repos:

```
make repos-import
```

I'll send you a keypair to use.  Today you can set cert_path in your config.yaml to point to the cert:
```
certs:
  cert_path: "/home/jlsherri/cdncert/cert.pemz"
```

so when you run 'go run cmd/external-repos/main.go introspect-all`  it will use that cert to fetch red hat repo metadata.

With this pr, that should still work, but if you specify the cert contents via an ENV variable, it will use that instead:

```
export RH_CDN_CERT_PAIR=`cat /home/jlsherri/cdncert/cert.pem`
go run cmd/external-repos/main.go introspect-all
```
and will ignore whatever cert_path is set to.  If you don't set the env variable, it should still use cert_path

### Comment by @avisiedo on August 05, 2022 at 08:46 PM UTC

Launched from inside the pod by exporting the env variable by hand before executing the command:

```raw
sh-4.4$ time ./external-repos introspect-all
{"level":"error","error":"Config File \"config\" Not Found in \"[/configs]\"","time":"2022-08-05T20:04:38Z"}
{"level":"debug","time":"2022-08-05T20:04:38Z","message":"Introspecting http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/"}
{"level":"debug","time":"2022-08-05T20:09:06Z","message":"Introspecting http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/"}
{"level":"debug","time":"2022-08-05T20:10:36Z","message":"Introspecting http://mirror.centos.org/centos/8-stream/extras/x86_64/os/"}
{"level":"debug","time":"2022-08-05T20:10:36Z","message":"Introspecting http://mirror.stream.centos.org/9-stream/AppStream/x86_64/os/"}
{"level":"debug","time":"2022-08-05T20:12:13Z","message":"Introspecting http://mirror.stream.centos.org/9-stream/BaseOS/x86_64/os/"}
{"level":"debug","time":"2022-08-05T20:12:28Z","message":"Introspecting https://cdn.redhat.com/content/dist/layered/rhel8/x86_64/ansible/2/os"}
{"level":"debug","time":"2022-08-05T20:12:28Z","message":"Introspecting https://cdn.redhat.com/content/dist/rhel8/8.4/x86_64/appstream/os"}
{"level":"debug","time":"2022-08-05T20:16:33Z","message":"Introspecting https://cdn.redhat.com/content/dist/rhel8/8.4/x86_64/baseos/os"}
{"level":"debug","time":"2022-08-05T20:17:32Z","message":"Introspecting https://cdn.redhat.com/content/dist/rhel8/8.5/x86_64/appstream/os"}
{"level":"debug","time":"2022-08-05T20:22:49Z","message":"Introspecting https://cdn.redhat.com/content/dist/rhel8/8.5/x86_64/baseos/os"}
{"level":"debug","time":"2022-08-05T20:24:02Z","message":"Introspecting https://cdn.redhat.com/content/dist/rhel8/8.6/x86_64/appstream/os"}
{"level":"debug","time":"2022-08-05T20:30:53Z","message":"Introspecting https://cdn.redhat.com/content/dist/rhel8/8.6/x86_64/baseos/os"}
{"level":"debug","time":"2022-08-05T20:32:18Z","message":"Introspecting https://cdn.redhat.com/content/dist/rhel9/9.0/x86_64/appstream/os"}
{"level":"debug","time":"2022-08-05T20:32:36Z","message":"Introspecting https://cdn.redhat.com/content/dist/rhel9/9.0/x86_64/baseos/os"}
{"level":"debug","time":"2022-08-05T20:32:40Z","message":"Introspecting https://packages.cloud.google.com/yum/repos/cloud-sdk-el8-x86_64"}
{"level":"debug","time":"2022-08-05T20:32:48Z","message":"Introspecting https://packages.cloud.google.com/yum/repos/cloud-sdk-el9-x86_64"}
{"level":"debug","time":"2022-08-05T20:32:54Z","message":"Introspecting https://packages.cloud.google.com/yum/repos/google-compute-engine-el8-x86_64-stable"}
{"level":"debug","time":"2022-08-05T20:32:54Z","message":"Introspecting https://packages.cloud.google.com/yum/repos/google-compute-engine-el9-x86_64-stable"}
{"level":"debug","time":"2022-08-05T20:32:54Z","message":"Successfully Inserted 94614 packages"}

real	28m16.073s
user	7m47.465s
sys	0m33.434s

```

LGTM

---

## Reviews

### Review by @jlsherrill - Commented on August 03, 2022 at 07:57 PM UTC

### Review by @rverdile - Approved on August 04, 2022 at 07:26 PM UTC

looks good

was able to use a cert via env variable and introspect all the repos. got the right error without the env variable.

### Review by @swadeley - Approved on August 08, 2022 at 07:20 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/70*

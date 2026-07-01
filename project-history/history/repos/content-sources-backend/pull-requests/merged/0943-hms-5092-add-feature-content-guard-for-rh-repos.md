---
type: pull_request
number: 943
title: "HMS-5092: Add feature content guard for RH repos"
state: merged
author: jlsherrill
created: 2025-01-16T17:33:50Z
updated: 2025-01-27T14:08:39Z
closed: 2025-01-27T14:08:35Z
merged: 2025-01-27T14:08:35Z
base_branch: main
head_branch: 5092
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/943
---

# Pull Request #943: HMS-5092: Add feature content guard for RH repos

**Author**: @jlsherrill
**Created**: January 16, 2025 at 05:33 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `5092`

## Description

## Summary

* For RH repos, create a feature content guard if there is a feature and its not RHEL-OS-x86_64
* removes a snapshot attribute that was no longer needed "ContentGuardsAdded", this was used in the past but no longer. We can simplify it a bit now by removing it  (followup PR will delete the column)

## Testing steps

Due to the fact that accessing the subs-as-features api locally requires a custom ca and the content guard doesn't support that currently, this can only really be tested in ephemeral.

Use this deploy command:

```
bonfire deploy --frontends=true --set-image-tag quay.io/cloudservices/content-sources-backend=pr-943-latest --set-parameter content-sources-backend/OPTIONS_FEATURE_FILTER="OPENSHIFT-OCP-x86_64"  content-sources  --set-parameter content-sources-backend/FEATURES_ADMIN_TASKS_ENABLED=true    --ref-env=insights-stage  --no-remove-resources=pulp  --set-parameter content-sources-backend/CLIENTS_PULP_CUSTOM_REPO_CONTENT_GUARDS=true --set-parameter content-sources-backend/SUSPEND_CRON_JOB=true
```

Once that is up, login and verify that the OCP repos show up in your repos list under red hat repos

Login to the worker pod and have it snapshot one of the OCP repos:

```
$ oc get pods | grep task-worker
content-sources-backend-task-worker-796d75f76b-cfsxc            1/1     Running     0               156m

$ oc rsh content-sources-backend-task-worker-796d75f76b-cfsxc 

./external-repos snapshot https://cdn.redhat.com/content/dist/layered/rhel9/x86_64/rhocp/4.16/os/
```

NOTE: for some reason i get a db connection error about 9/10 times.   Keep running it until it runs without error.

Grab the URL of the latest snapshot  (Be sure to change the password and address to match that from `bonfire namespace describe`:

```
curl -X GET  -u jdoe:Wc6o0hs8YqrxWh3Z  --location "https://ee-7gzvgcdx.apps.crc-eph.r9lp.p1.openshiftapps.com/api/content-sources/v1.0/repositories/?origin=red_hat"        -H "Content-Type: application/json"  | jq '.data[] |  "\(.name) \(.last_snapshot.url)"'

```

look for:
```
"Red Hat OpenShift Container Platform 4.16 for RHEL 9 x86_64 (RPMs) http://pulp-content:8000/api/pulp-content/cs-redhat/b8691bb4-c409-4f65-b18b-3811422fa734/40e11bb1-594a-4d1a-8aab-86ef5cbbafe3/"
```

Now grab that URL, and exec into one of the pulp pods:

```
$ oc get pods | grep pulp | grep worker
pulp-worker-b668fb7c6-5vf8p                                     1/1     Running     2 (65m ago)     3h13m


$ oc rsh pulp-worker-b668fb7c6-5vf8p  
```

This should be denied, as there wasn't any auth provided:

```
curl  http://pulp-content:8000/api/pulp-content/cs-redhat/b8691bb4-c409-4f65-b18b-3811422fa734/40e11bb1-594a-4d1a-8aab-86ef5cbbafe3/"
```
```
403: Access denied
```

Since we're in the cluster, we can just specify the identity header  ./header.sh.  I've provided some examples here for stage:

```
curl  http://pulp-content:8000/api/pulp-content/cs-redhat/b8691bb4-c409-4f65-b18b-3811422fa734/40e11bb1-594a-4d1a-8aab-86ef5cbbafe3/  -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiIxNzY4NDYzOCIsICJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJmb28ifSwiYWNjb3VudF9udW1iZXIiOiIxNzY4NDYzOCIsImludGVybmFsIjp7Im9yZ19pZCI6IjE3Njg0NjM4In19fQo="
```

This is for org id 17684638 which has access.  

This uses a valid org id that doesn't have access:  18079510
```
curl  http://pulp-content:8000/api/pulp-content/cs-redhat/b8691bb4-c409-4f65-b18b-3811422fa734/40e11bb1-594a-4d1a-8aab-86ef5cbbafe3/ -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiIxODA3OTUxMCIsICJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiJmb28ifSwiYWNjb3VudF9udW1iZXIiOiIxODA3OTUxMCIsImludGVybmFsIjp7Im9yZ19pZCI6IjE4MDc5NTEwIn19fQo="
```
```
403: Access denied
```


And this uses an invalid org id (999):

```
curl  http://pulp-content:8000/api/pulp-content/cs-redhat/b8691bb4-c409-4f65-b18b-3811422fa734/40e11bb1-594a-4d1a-8aab-86ef5cbbafe3/ -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5OTkiLCAidHlwZSI6IlVzZXIiLCJ1c2VyIjp7InVzZXJuYW1lIjoiZm9vIn0sImFjY291bnRfbnVtYmVyIjoiOTk5IiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiOTk5In19fQo="
```
```
403: Access denied
```



Other things to test:

Snapshotting a normal RHEL repo, should not result in a guard being created:

```
./external-repos/main.go snapshot https://cdn.redhat.com/content/dist/layered/rhel8/x86_64/ansible/2/os/
```
 
Perform the same steps to grab the latest snapshot url.  It should not require any auth header to pull contetn



---

## Discussion

### Comment by @jlsherrill on January 16, 2025 at 06:00 PM UTC

https://issues.redhat.com/browse/Fixes HMS-5092

### Comment by @jlsherrill on January 16, 2025 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-5092

### Comment by @jlsherrill on January 20, 2025 at 05:57 PM UTC

/retest

---

## Reviews

### Review by @dominikvagner - Approved on January 24, 2025 at 01:48 PM UTC

looks good, no complaints from me, works as expected! 👏🏼🎉  
thanks for the great testing steps 🙇🏼‍♂️ 

approved! ✅

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/943*

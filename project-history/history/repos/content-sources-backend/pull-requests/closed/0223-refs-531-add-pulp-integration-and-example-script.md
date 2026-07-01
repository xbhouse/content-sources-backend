---
type: pull_request
number: 223
title: "Refs 531: Add pulp integration and example script"
state: closed
author: Andrewgdewar
created: 2023-02-28T21:44:58Z
updated: 2023-04-05T15:11:41Z
closed: 2023-04-05T15:11:41Z
base_branch: main
head_branch: CONTENT-531
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/223
---

# Pull Request #223: Refs 531: Add pulp integration and example script

**Author**: @Andrewgdewar
**Created**: February 28, 2023 at 09:44 PM UTC
**Status**: Closed
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `CONTENT-531`

## Description

## Summary

This adds a demonstration script for integrating with our new Pulp Binding. 
Documentation around this integration can be referenced [here](https://docs.google.com/document/d/1PBzYDDBwnNYRPPTNHKVnbKFOLZYjyLqXzPCLDk2-Zos/edit#).

## Testing steps

Run the following:

- make compose-clean   
- DEPLOY_PULP=true make compose-up
- make run 

Then use the API or the UI to add a few repositories to sync with pulp below. 

Then run the following command in your console: 
   
- go run cmd/external-repos/main.go pulp-create ||Enter URL of added repo||


---

## Discussion

### Comment by @jlsherrill on February 28, 2023 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-531

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/223*

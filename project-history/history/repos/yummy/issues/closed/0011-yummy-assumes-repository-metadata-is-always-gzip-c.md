---
type: issue
number: 11
title: "yummy assumes repository metadata is always gzip-compressed"
state: closed
author: dralley
created: 2023-02-13T04:19:56Z
updated: 2023-05-30T20:20:10Z
closed: 2023-05-30T20:20:10Z
labels: []
url: https://github.com/content-services/yummy/issues/11
---

# Issue #11: yummy assumes repository metadata is always gzip-compressed

**Author**: @dralley
**Created**: February 13, 2023 at 04:19 AM UTC
**Status**: Closed
**Labels**: None

## Description

https://github.com/content-services/yummy/blob/96f347aab2055ef2166bf0ea65bd6af831f225e9/pkg/yum/repository.go#L278

Though this is generally true this isn't a valid assumption, metadata can be compressed with LZMA (.xz), and soon ZSTD (.zst).

ZSTD compression is a significant enough improvement in every aspect (file size, compression speed, decompression speed) that it is more likely to see adoption in the future.

---

## Discussion

### Comment by @dralley on April 28, 2023 at 03:41 AM UTC

I would like to add, that zstd is likely to become a default option in the near future for createrepo_c, and therefore Fedora and CentOS.

https://lists.fedoraproject.org/archives/list/devel@lists.fedoraproject.org/thread/25J55FCQKTMCEOBS7EZCTP34RKLUOUQ3/

### Comment by @jlsherrill on April 28, 2023 at 12:44 PM UTC

Thanks for pinging this @dralley!  I've opened a jira to track our team on this:  https://issues.redhat.com/browse/HMS-1693

---

*Archived from: https://github.com/content-services/yummy/issues/11*

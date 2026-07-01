---
type: pull_request
number: 1
title: "Add Functionality and readme.md"
state: merged
author: Andrewgdewar
created: 2022-05-11T14:03:43Z
updated: 2022-05-17T15:19:08Z
closed: 2022-05-17T15:19:08Z
merged: 2022-05-17T15:19:08Z
base_branch: master
head_branch: ReviewBranch
labels: []
url: https://github.com/content-services/yummy/pull/1
---

# Pull Request #1: Add Functionality and readme.md

**Author**: @Andrewgdewar
**Created**: May 11, 2022 at 02:03 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `ReviewBranch`

## Description

To test this, you can use the following main.go: 

```
package main

import (
	"fmt"

	"github.com/content-services/utilities/Yum"
)

var rpmUrls = []string{
	"http://yum.theforeman.org/plugins/1.2/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.3/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.4/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.5/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.6/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.7/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.8/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.9/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.10/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.11/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.12/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.13/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.14/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.15/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.16/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.17/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.18/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.19/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.20/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.21/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.22/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.23/el7/x86_64/",
	"http://yum.theforeman.org/plugins/1.24/el7/x86_64/",
	"https://fedorapeople.org/groups/katello/fakerepos/zoo/",
	"https://download-i2.fedoraproject.org/pub/epel/7/x86_64/",
	"https://download-i2.fedoraproject.org/pub/fedora/linux/releases/35/Everything/x86_64/os/",
	"http://windhelm.usersys.redhat.com/pub/rhel7_os/",
}

func main() {
	var totalPackages = 0
	for i, url := range rpmUrls {
		data, err := Rpm.Extract(url)
		if err != nil {
			fmt.Errorf("GET error: %v\n", err)
		}
		amount := len(data)
		fmt.Printf("Repo %v contains %v rpm packages.\n", i+1, amount)
		totalPackages = totalPackages + amount
	}

	fmt.Printf("Number of packages parsed: %v\n", totalPackages)
}
```

The above is using the pkg.go imported version of this module. 

If you want to run the code directly you'll need to create a "utilities" directory to drop it into within your main.go directory.

And then add the following to your go.mod:

```
module main

go 1.16

require github.com/content-services/utilities v0.0.3

replace github.com/content-services/utilities => ./utilities
```

Then run a `go get -u` to update your go.sum appropriately.


---

## Discussion

### Comment by @jlsherrill on May 11, 2022 at 05:41 PM UTC

My initial thought for the Rpm part of this was that it could be a go 'yum metadata parser' library instead of a 'content-services' utilities' library.  The idea being that we may could get outside contributions to a 'yum metadata parser' for Go.    The Time part of it kinda complicates that, but i'm curious your thoughts

### Comment by @Andrewgdewar on May 11, 2022 at 05:56 PM UTC

> My initial thought for the Rpm part of this was that it could be a go 'yum metadata parser' library instead of a 'content-services' utilities' library. The idea being that we may could get outside contributions to a 'yum metadata parser' for Go. The Time part of it kinda complicates that, but i'm curious your thoughts

I have no thoughts! 😄  , this was just a proof of concept for me. However it may be used is okay with me. The RPM part likely needs better naming, and I can cut out the time bit if we want as well. Those methods aren't particularly complicated, and can be easily re-created elsewhere.

_I would add, knowing now that the parser is quite quick, that we likely would want to add quite a few other available package fields if this is ever expected to be used by others._

### Comment by @Andrewgdewar on May 13, 2022 at 05:33 PM UTC

> standard project layout

Not sure what you mean, as this is a module I have set it up as per a few tutorials surrounding those. Can you paste a ref here? 

### Comment by @rverdile on May 13, 2022 at 07:49 PM UTC

this: https://github.com/golang-standards/project-layout

Not all the directories, just the ones that seem to fit. But this is a small repo, so maybe it's unnecessary for now.

---

## Reviews

### Review by @rverdile - Commented on May 11, 2022 at 03:04 PM UTC

I ran it a few times and it seems to work, but I haven't thoroughly tested it. Most of my comments revolve around naming. 

Also, Is your intention to export everything? Because all the structs/functions with title case get exported. If some of these functions or structs are really meant to be internal, I think it would less confusing from a user perspective to not export them.

### Review by @Andrewgdewar - Commented on May 11, 2022 at 04:18 PM UTC

### Review by @jlsherrill - Commented on May 11, 2022 at 05:42 PM UTC

### Review by @rverdile - Commented on May 11, 2022 at 06:57 PM UTC

Names look better and I like the comments added to describe the functions. 

+1 to Justin's idea about this being a dedicated yum metadata parser. A quick google suggests that could be new territory, which is cool.

Also, is this repo going to use the "standard project layout"?

### Review by @jlsherrill - Approved on May 17, 2022 at 01:27 PM UTC

ACK!

---

*Archived from: https://github.com/content-services/yummy/pull/1*

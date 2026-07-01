---
type: pull_request
number: 16
title: "simpler loading of .env file in Makefile"
state: closed
author: jlsherrill
created: 2022-05-25T18:01:05Z
updated: 2022-05-26T21:08:57Z
closed: 2022-05-26T21:08:57Z
base_branch: main
head_branch: env_var
labels: []
url: https://github.com/content-services/content-sources-backend/pull/16
---

# Pull Request #16: simpler loading of .env file in Makefile

**Author**: @jlsherrill
**Created**: May 25, 2022 at 06:01 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `env_var`

## Description

This makes loading .env more generic and a bit simpler.  Curious your thoughts. 


---

## Discussion

### Comment by @jlsherrill on May 25, 2022 at 06:01 PM UTC

It also adds a sample .env file.  I'll update the README after https://github.com/content-services/content-sources-backend/pull/14 goes in 

### Comment by @rverdile on May 25, 2022 at 06:38 PM UTC

@jlsherrill  i'm not familiar with the sed command. what's that doing?

### Comment by @jlsherrill on May 25, 2022 at 06:41 PM UTC

so say you have an env file with: 
```
FOO=BAR
BAZ=FOZ
```

the make target 'make .env-export' takes your .env and creates a .env-export file:
```
export FOO=BAR
export BAZ=FOZ
```

Then in any command we need to use the configuration in, we can just use  '. .env-export && COMMAND' 
and it will basically export all the variables into the process.  I'm a bit surprised how obtuse this all is, but this seemed like the easiest way.

### Comment by @jlsherrill on May 25, 2022 at 06:44 PM UTC

I think when we switch to viper, this may all be a bit easier though.  

### Comment by @rverdile on May 25, 2022 at 07:04 PM UTC

I like this idea. I get why it feels obtuse. I used to see Makefiles used in conjunction with shell script to avoid making the Makefile too complicated.

### Comment by @jlsherrill on May 26, 2022 at 09:08 PM UTC

actually, i think i'm just gonna integrate with viper

---

## Reviews

### Review by @rverdile - Commented on May 25, 2022 at 07:00 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/16*

---
type: pull_request
number: 435
title: "Fixes 2702,2887: Add orgID to snapshot list endpoint"
state: merged
author: Andrewgdewar
created: 2023-10-19T16:32:30Z
updated: 2023-11-13T02:30:32Z
closed: 2023-11-13T02:17:09Z
merged: 2023-11-13T02:17:09Z
base_branch: main
head_branch: HMS-2702
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/435
---

# Pull Request #435: Fixes 2702,2887: Add orgID to snapshot list endpoint

**Author**: @Andrewgdewar
**Created**: October 19, 2023 at 04:32 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-2702`

## Description

## Summary

HMS-2702
- Adds orgID as a requirement to snapshot list

HMS 2887
- Allows the package list and snapshotList endpoint to return red hat repository (orgID of "-1") related data.

Additional changes: 
- Adds the new variable "any" to the "fetchRepoConfig" method.
- Updated dao methods using the above function as needed, create/delete/update are now prevented from taking effect on red_hat repositories.

## Testing steps

- Easiest testing method would be to checkout the accompanied front-end PR [here](https://github.com/content-services/content-sources-frontend/pull/156), and follow steps there-in.


---

## Discussion

### Comment by @jlsherrill on October 19, 2023 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-2702

### Comment by @jlsherrill on October 25, 2023 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-Fixes

### Comment by @jlsherrill on October 31, 2023 at 05:21 PM UTC

https://issues.redhat.com/browse/HMS-2887

### Comment by @rverdile on November 03, 2023 at 06:47 PM UTC

I'm not sure why, but it looks like most of the test didn't run. Do you see that too?

### Comment by @swadeley on November 06, 2023 at 08:15 AM UTC

> I'm not sure why, but it looks like most of the test didn't run. Do you see that too?

Hi, `test_validate_repo_parameters`  failed because the same repo was used in the previous filter test and not cleaned up. I have added another repo with invalid metadata to the filter test to avoid this conflict.

Not sure why test run was incomplete (I see `28 selected` and I counted 26 PASSED. Then later I see `= 1 failed, 26 passed, 1 skipped, 9 deselected`).

### Comment by @swadeley on November 06, 2023 at 08:16 AM UTC

/retest

### Comment by @jlsherrill on November 06, 2023 at 08:14 PM UTC

/retest

### Comment by @jlsherrill on November 07, 2023 at 07:50 PM UTC

@swadeley we've added a command to run when deploying to ephemeral that should snapshot a small rhel repo (Red Hat Ansible Engine 2 for RHEL 8 x86_64 (RPMs))   Hopefully there's enough memory for it to snapshot.  Note that it doesn't 'block' deployment on the snapshot, it happens in the background.  

I'll spin it up to verify as well once the image is built

### Comment by @jlsherrill on November 07, 2023 at 08:33 PM UTC

I've updated the init container to first wait until pulp is up before triggering the task.  I haven't been able to test this in ephemeral, as i can't seem to get the deployment.yaml changes to take effect when pointing to it with my local bonfire config.  I'm not 100% sure why.

### Comment by @jlsherrill on November 07, 2023 at 09:50 PM UTC

So i got it working, about half the time :/   the other half the time, there is some timing issue and the snapshot task is not scheduled.   Pushing up some debugging

### Comment by @swadeley on November 08, 2023 at 05:50 AM UTC

> @swadeley we've added a command to run when deploying to ephemeral that should snapshot a small rhel repo (Red Hat Ansible Engine 2 for RHEL 8 x86_64 (RPMs)) Hopefully there's enough memory for it to snapshot. Note that it doesn't 'block' deployment on the snapshot, it happens in the background.
> 
> I'll spin it up to verify as well once the image is built

Hi, I have had this PR and the frontend [1] deployed for over 24 hours. I saw the Ansible repo was snapshotted right away, but still no RHEL repos due to "Killed by signal 9." as mentioned before [2].

Let me know when I should redeploy to test a new image.

Thank you

[1] https://github.com/content-services/content-sources-frontend/pull/156
[2] https://github.com/content-services/content-sources-frontend/pull/156#issuecomment-1797120004

### Comment by @jlsherrill on November 08, 2023 at 03:59 PM UTC

@swadeley yes, we'll likely not ever see these large rhel repos be snapshotable in ephemeral.  

### Comment by @jlsherrill on November 08, 2023 at 03:59 PM UTC

@swadeley so this should be good to merge if you are okay with it.  I have removed my debugging output

### Comment by @swadeley on November 09, 2023 at 02:08 AM UTC

/retest

### Comment by @swadeley on November 10, 2023 at 01:44 AM UTC

/retest

### Comment by @Andrewgdewar on November 10, 2023 at 08:55 PM UTC

> It says in your summary that create/update/delete do not have an effect for red hat repositories. Right now, if my org ID is -1, I can delete/update but not create. Obviously, that's not a real use case, but I feel the behavior should probably be consistent unless there's a reason for it not to be.
> 
> Do we expect to want to delete/update red hat repositories from the API? Or should that be blocked the way create is blocked?

@rverdile 
If you can highlight somewhere that the dao layer would allow this, please do. 
I think I covered all cases of CrUD not being possible.

oh... 
I think I see what you mean in your comment now.
if the ordID is -1, then you could delete/update repositories.
So to prevent that we would need to specifically check that is not the orgID
Is that necessary? auth gives us the ordID (not the user) so I couldn't really see this happening.

---

## Reviews

### Review by @rverdile - Commented on October 23, 2023 at 08:28 PM UTC

### Review by @Andrewgdewar - Commented on October 23, 2023 at 09:18 PM UTC

### Review by @rverdile - Commented on October 24, 2023 at 02:18 PM UTC

### Review by @Andrewgdewar - Commented on October 24, 2023 at 02:35 PM UTC

### Review by @Andrewgdewar - Commented on October 25, 2023 at 09:39 PM UTC

### Review by @rverdile - Commented on October 30, 2023 at 05:03 PM UTC

### Review by @rverdile - Commented on October 30, 2023 at 05:04 PM UTC

### Review by @rverdile - Commented on October 30, 2023 at 05:04 PM UTC

### Review by @rverdile - Commented on October 30, 2023 at 05:11 PM UTC

### Review by @rverdile - Commented on October 30, 2023 at 05:18 PM UTC

### Review by @Andrewgdewar - Commented on October 30, 2023 at 08:32 PM UTC

### Review by @Andrewgdewar - Commented on October 30, 2023 at 08:43 PM UTC

### Review by @Andrewgdewar - Commented on October 31, 2023 at 07:55 PM UTC

### Review by @Andrewgdewar - Commented on October 31, 2023 at 08:09 PM UTC

### Review by @rverdile - Commented on November 07, 2023 at 06:28 PM UTC

It says in your summary that create/update/delete do not have an effect for red hat repositories. Right now, if my org ID is -1, I can delete/update but not create. Obviously, that's not a real use case, but I feel the behavior should probably be consistent unless there's a reason for it not to be.

Do we expect to want to delete/update red hat repositories from the API? Or should that be blocked the way create is blocked? 

### Review by @rverdile - Commented on November 07, 2023 at 06:36 PM UTC

### Review by @Andrewgdewar - Commented on November 10, 2023 at 08:49 PM UTC

### Review by @rverdile - Approved on November 10, 2023 at 09:26 PM UTC

nice!!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/435*

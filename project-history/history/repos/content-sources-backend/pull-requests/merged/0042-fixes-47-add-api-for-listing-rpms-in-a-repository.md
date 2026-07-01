---
type: pull_request
number: 42
title: "Fixes 47: add api for listing rpms in a repository"
state: merged
author: avisiedo
created: 2022-06-22T18:51:19Z
updated: 2022-08-26T15:38:10Z
closed: 2022-07-08T18:48:35Z
merged: 2022-07-08T18:48:35Z
base_branch: main
head_branch: hmscontent-47
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/42
---

# Pull Request #42: Fixes 47: add api for listing rpms in a repository

**Author**: @avisiedo
**Created**: June 22, 2022 at 06:51 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `hmscontent-47`

## Description

As a developer I want to retrieve the packages that belong to a repository that meet with:
- An api endpoint "/repositories/:uuid/rpms" where uuid is the id for the repository.
- A new table repositoriesand repository_configurations table relate to it.
- A new table repository_rpms where the rpm information will be stored with:
  - a uuid that identify the register.
  - the rpm information we want to store into the database.
  - a foreign key that relate to the repositories table that the rpm belong to.
- Add gorm models.
- The migration scripts for the above changes.
- Add the new table to 'dbmigrate seed' so that we can populate with random data.
- Tests.

---

## Discussion

### Comment by @jlsherrill on June 28, 2022 at 08:40 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-47

### Comment by @jlsherrill on June 29, 2022 at 05:51 PM UTC

https://issues.redhat.com/browse/HMSCONTENT--

### Comment by @avisiedo on July 01, 2022 at 02:27 PM UTC

I finish to debug a situation with TestRpmCreate in the dao layer and I push the new updates.

### Comment by @jlsherrill on July 05, 2022 at 03:26 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @jlsherrill on July 06, 2022 at 07:57 PM UTC

I think this pr "Fixes" and not "Refs" right?  I'll adjust, but feel free to change back if you disagree

### Comment by @jlsherrill on July 06, 2022 at 09:07 PM UTC

Code looks good!  Found a couple of issues in testing:

1) After creating a repository with the api, the URL is blank:
```
$ curl  -H "`./scripts/header.sh 1 1`"  -X GET http://localhost:8000/api/content_sources/v1.0/repositories/2a831a10-b667-41cb-9dab-12a1f4e6db8c
{"uuid":"2a831a10-b667-41cb-9dab-12a1f4e6db8c","name":"Foo","url":"","distribution_versions":[],"distribution_arch":"","account_id":"1","org_id":"1"}
```
its set properly in the db within the repositories table

2) when trying to fetch rpms, i get an error:

curl  -H "`./scripts/header.sh 1 1`"  -X GET http://localhost:8000/api/content_sources/v1.0/repositories/2a831a10-b667-41cb-9dab-12a1f4e6db8c/rpms
{"message":"repository_uuid = 2a831a10-b667-41cb-9dab-12a1f4e6db8c is not owned"}

I think https://github.com/content-services/content-sources-backend/pull/42/files#diff-6585c6f15d60726520ed4e0be761b0fe6bb1b269792bc972f7dc551aa268009cR26  should be checking uuid and not repository_uuid, but don't quote me on that :)



### Comment by @jlsherrill on July 06, 2022 at 09:12 PM UTC

3.  i changed isOwnedRepository to just return true (to bypass comment 2 above) , and i'm not getting any rpms back from the api:
{"data":[],"meta":{"limit":100,"offset":0,"count":0},"links":{"first":"/api/content_sources/v1.0/repositories/2a831a10-b667-41cb-9dab-12a1f4e6db8c/rpms?limit=100\u0026offset=0","last":"/api/content_sources/v1.0/repositories/2a831a10-b667-41cb-9dab-12a1f4e6db8c/rpms?limit=100\u0026offset=0"}}

I think the join here: https://github.com/content-services/content-sources-backend/pull/42/files#diff-6585c6f15d60726520ed4e0be761b0fe6bb1b269792bc972f7dc551aa268009cR61
isn't quite right, and is doing something similar where its using the wrong uuid.


EDIT, okay i changed some code there to translate between the repository_config uuid to the repository uuid and it worked great!  My guess is in your testing you were passing in the repository uuid and not the repository_configuration's uuid.  The user will only see the repository_configuration's uuid.  that's their view of the repository.

To give you an idea of what i did, i just changed your query to do this (Obviously no error checking being done):


```diff
@@ -55,16 +57,21 @@ func (r rpmDaoImpl) List(orgID string, uuidRepo string, limit int, offset int) (
                        fmt.Errorf("repository_uuid = %s is not owned", uuidRepo)
        }
 
+       repoConfig := models.RepositoryConfiguration{}
+       r.db.Where("UUID = ?", uuidRepo).Find(&repoConfig)
+
        // Select all the rpms from a repository
        if err := r.db.
                Model(&repoRpms).
                Joins(strings.Join([]string{"inner join", models.TableNameRpmsRepositories, "on uuid = rpm_uuid"}, " ")).
-               Where("repository_uuid = ?", uuidRepo).
+               Where("repository_uuid = ?", repoConfig.RepositoryUUID).
                Count(&totalRpms).
                Offset(offset).
                Limit(limit).
                Find(&repoRpms).
                Error; err != nil {
                return api.RepositoryRpmCollectionResponse{}, totalRpms, err
        }

```

### Comment by @avisiedo on July 07, 2022 at 08:30 AM UTC

thanks! I have fixed the `uuid` that is used to recover the rpm list.

I am having a look to the code that retrieve the Repository.url; the information is into the database.

### Comment by @avisiedo on July 07, 2022 at 08:43 AM UTC

thanks! I have fixed the empty `url`

### Comment by @jlsherrill on July 07, 2022 at 05:13 PM UTC

All 3 items are looking much better.  I do notice that url is still empty when fetching the repository (but not listing):

```
$ curl  -H "`./scripts/header.sh 1 1`"  -X GET http://localhost:8000/api/content_sources/v1.0/repositories/2a831a10-b667-41cb-9dab-12a1f4e6db8c   | python -m json.tool
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   150  100   150    0     0  55207      0 --:--:-- --:--:-- --:--:-- 75000
{
    "uuid": "2a831a10-b667-41cb-9dab-12a1f4e6db8c",
    "name": "Foo",
    "url": "",
    "distribution_versions": [],
    "distribution_arch": "",
    "account_id": "1",
    "org_id": "1"
}

```

hrmmm, which is strange because they are using the same code...
ahh, the dao Fetch also needs a preload of repository.

Would you mind adding an assertion to the dao layer list & fetch test to make sure that url is being returned?

### Comment by @jlsherrill on July 07, 2022 at 07:13 PM UTC

I think thats the last thing and then this is good to go

### Comment by @jlsherrill on July 19, 2022 at 08:20 PM UTC

https://issues.redhat.com/browse/14649800

---

## Reviews

### Review by @jlsherrill - Commented on June 22, 2022 at 08:48 PM UTC

### Review by @jlsherrill - Commented on June 23, 2022 at 02:29 AM UTC

### Review by @jlsherrill - Commented on June 23, 2022 at 02:30 AM UTC

### Review by @jlsherrill - Commented on June 23, 2022 at 02:31 AM UTC

### Review by @jlsherrill - Commented on June 23, 2022 at 02:32 AM UTC

### Review by @jlsherrill - Commented on June 23, 2022 at 02:36 AM UTC

### Review by @jlsherrill - Commented on June 23, 2022 at 02:38 AM UTC

### Review by @jlsherrill - Commented on June 23, 2022 at 07:16 PM UTC

### Review by @jlsherrill - Commented on June 24, 2022 at 04:47 PM UTC

### Review by @jlsherrill - Commented on June 24, 2022 at 04:58 PM UTC

### Review by @rverdile - Commented on June 24, 2022 at 05:03 PM UTC

### Review by @rverdile - Commented on June 24, 2022 at 05:12 PM UTC

### Review by @avisiedo - Commented on June 27, 2022 at 06:33 AM UTC

### Review by @avisiedo - Commented on June 27, 2022 at 06:48 AM UTC

### Review by @avisiedo - Commented on June 27, 2022 at 06:52 AM UTC

### Review by @avisiedo - Commented on June 27, 2022 at 10:20 AM UTC

### Review by @avisiedo - Commented on June 27, 2022 at 01:41 PM UTC

### Review by @rverdile - Commented on June 27, 2022 at 03:43 PM UTC

### Review by @rverdile - Commented on June 27, 2022 at 04:32 PM UTC

### Review by @swadeley - Commented on June 27, 2022 at 06:57 PM UTC

### Review by @swadeley - Commented on June 27, 2022 at 06:58 PM UTC

### Review by @swadeley - Commented on June 27, 2022 at 06:58 PM UTC

### Review by @jlsherrill - Commented on July 01, 2022 at 01:53 AM UTC

### Review by @jlsherrill - Commented on July 01, 2022 at 01:53 AM UTC

### Review by @jlsherrill - Commented on July 01, 2022 at 01:55 AM UTC

### Review by @jlsherrill - Commented on July 01, 2022 at 01:57 AM UTC

### Review by @jlsherrill - Commented on July 01, 2022 at 02:10 AM UTC

### Review by @jlsherrill - Commented on July 01, 2022 at 02:31 AM UTC

### Review by @jlsherrill - Commented on July 01, 2022 at 02:32 AM UTC

### Review by @jlsherrill - Commented on July 01, 2022 at 02:33 AM UTC

### Review by @jlsherrill - Commented on July 01, 2022 at 02:35 AM UTC

### Review by @jlsherrill - Commented on July 01, 2022 at 02:46 AM UTC

### Review by @jlsherrill - Commented on July 01, 2022 at 02:46 AM UTC

### Review by @jlsherrill - Commented on July 01, 2022 at 02:54 AM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 07:08 AM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 07:08 AM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 07:14 AM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 07:15 AM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 07:15 AM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 07:17 AM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 07:19 AM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 07:20 AM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 07:36 AM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 07:36 AM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 07:37 AM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 07:39 AM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 10:28 AM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 10:29 AM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 10:32 AM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 02:26 PM UTC

### Review by @avisiedo - Commented on July 04, 2022 at 02:15 PM UTC

### Review by @avisiedo - Commented on July 04, 2022 at 02:20 PM UTC

### Review by @avisiedo - Commented on July 04, 2022 at 02:21 PM UTC

### Review by @avisiedo - Commented on July 04, 2022 at 02:26 PM UTC

### Review by @avisiedo - Commented on July 04, 2022 at 05:06 PM UTC

### Review by @avisiedo - Commented on July 04, 2022 at 05:12 PM UTC

### Review by @avisiedo - Commented on July 05, 2022 at 05:54 AM UTC

### Review by @rverdile - Commented on July 05, 2022 at 04:34 PM UTC

### Review by @jlsherrill - Commented on July 05, 2022 at 08:01 PM UTC

### Review by @jlsherrill - Commented on July 05, 2022 at 08:01 PM UTC

### Review by @jlsherrill - Commented on July 05, 2022 at 08:11 PM UTC

### Review by @jlsherrill - Commented on July 05, 2022 at 08:12 PM UTC

### Review by @avisiedo - Commented on July 06, 2022 at 07:24 AM UTC

### Review by @avisiedo - Commented on July 06, 2022 at 07:44 AM UTC

### Review by @avisiedo - Commented on July 06, 2022 at 08:30 AM UTC

### Review by @jlsherrill - Commented on July 06, 2022 at 12:09 PM UTC

### Review by @jlsherrill - Commented on July 06, 2022 at 12:53 PM UTC

### Review by @jlsherrill - Commented on July 06, 2022 at 12:53 PM UTC

### Review by @jlsherrill - Commented on July 06, 2022 at 12:56 PM UTC

### Review by @avisiedo - Commented on July 06, 2022 at 02:24 PM UTC

### Review by @jlsherrill - Commented on July 06, 2022 at 02:40 PM UTC

### Review by @rverdile - Commented on July 06, 2022 at 03:32 PM UTC

### Review by @avisiedo - Commented on July 06, 2022 at 04:40 PM UTC

### Review by @avisiedo - Commented on July 06, 2022 at 05:16 PM UTC

### Review by @avisiedo - Commented on July 06, 2022 at 05:24 PM UTC

### Review by @rverdile - Commented on July 07, 2022 at 03:18 PM UTC

### Review by @rverdile - Commented on July 07, 2022 at 03:20 PM UTC

### Review by @rverdile - Commented on July 07, 2022 at 03:32 PM UTC

### Review by @jlsherrill - Approved on July 08, 2022 at 01:53 PM UTC

ACK! 🔥 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/42*

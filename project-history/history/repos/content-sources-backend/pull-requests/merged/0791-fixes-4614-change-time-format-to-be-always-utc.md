---
type: pull_request
number: 791
title: "Fixes 4614: change time format to be always UTC"
state: merged
author: dominikvagner
created: 2024-08-28T13:14:40Z
updated: 2024-09-16T10:00:25Z
closed: 2024-09-16T09:56:00Z
merged: 2024-09-16T09:56:00Z
base_branch: main
head_branch: 4614
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/791
---

# Pull Request #791: Fixes 4614: change time format to be always UTC

**Author**: @dominikvagner
**Created**: August 28, 2024 at 01:14 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4614`

## Description

## Summary
Changes all outputted dates/time info to be always in UTC and not in local system timezone (apart from the `Templates` part, that is done in #785).

## Testing steps
Go through endpoints that are displaying any time info and verify that all the time/date fields are in UTC timezone, i.e. aren't outputted with the offset.



---

## Discussion

### Comment by @jlsherrill on August 28, 2024 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-4614

### Comment by @swadeley on August 28, 2024 at 03:04 PM UTC

/restest

### Comment by @swadeley on August 28, 2024 at 03:16 PM UTC

/retest

### Comment by @dominikvagner on August 29, 2024 at 07:28 AM UTC

/retest

### Comment by @jlsherrill on August 30, 2024 at 03:30 PM UTC

Would it make sense to create a 'AfterFind' hook at the model level to convert these to UTC coming out of the database?     The benefit would be the dao layer doesn't have to care about it and we'd probably have to worry about it in fewer places.   

(haven't tried this and can't say for sure if it'd work, wanted to get your thoughts)

### Comment by @dominikvagner on September 10, 2024 at 06:00 AM UTC

> Would it make sense to create a 'AfterFind' hook at the model level to convert these to UTC coming out of the database? The benefit would be the dao layer doesn't have to care about it and we'd probably have to worry about it in fewer places.
> 
> (haven't tried this and can't say for sure if it'd work, wanted to get your thoughts)

(thought I replied already, the comment must have not gotten posted 🤔😄)

oh, didn't know `gorm` has a hook like that, good to know 😄 
I think it could work, but if any point the date is changed to a new one it would probably pick up the local timezone (don't know if reassigning it from model to API struct would be the case, but just changing it to a newly created date would unless that has `.UTC()` on it), this way of doing so in the DAO layer, right before outputting, we don't have to worry about that 😅

I can try/test it out if you think it would be better 👍🏼 

### Comment by @jlsherrill on September 10, 2024 at 07:27 PM UTC

I'm not quite sure what you mean, you could give an example?  

Just to make sure what i brought up actually worked, i tried it out, adding this to models/template.go:
```
func (t *Template) AfterFind(tx *gorm.DB) error {
	t.Date = t.Date.UTC()
	return nil
}
```

now, if i create or update a template and set the date to my timezone:

```
PATCH http://localhost:8000/api/content-sources/v1.0/templates/80555a80-7c36-4036-a79d-83a867df6bbc
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
Content-Type: application/json

{
  "date":"2024-09-19T15:09:43-04:00"

}
```
I get back UTC:

```

{
  "uuid": "80555a80-7c36-4036-a79d-83a867df6bbc",
  "name": "test6",
  "org_id": "9",
  "description": "",
  "arch": "x86_64",
  "version": "9",
  "date": "2024-09-19T19:09:43Z",
  "repository_uuids": [
    "f04efa17-00e3-46f8-a36f-9d838d7beaf4"
  ],
  "rhsm_environment_id": "80555a807c364036a79d83a867df6bbc",
  "created_by": "foo",
  "last_updated_by": "foo",
  "created_at": "2024-09-10T19:23:13.398768Z",
  "updated_at": "2024-09-10T19:23:48.202238Z",
  "use_latest": false
}
```

Note i only set 'date' here, not created_at, etc..   It seems to work well with create, update, and list.  adding similar things to models/task_info.go means that any object that includes a task (such repository_config), will automatically have that tasks dates converted.  

I could be completely misunderstanding your point though, so if its still valid, maybe an example would help me understand. 

### Comment by @dominikvagner on September 11, 2024 at 11:22 AM UTC

I looked into it a bit more and I think you're right. The hook would probably be a better solution, but to more explain what I meant:
If there was something like this in the code:
```go
// pseudo code, let's say this was a handler
func update(c echo.Ctx, uuid, date string) api.Template {
    template := daoReg.tDao.Fetch(uuid)
    template.Date = time.Parse(date, format)
    daoReg.tDao.Update(template)
    return c.JSON(200, daoReg.tDao.modelToApiResp(template))
} 
```
^ This would return a time _not in_ UTC, because it was changed after it was loaded from the DB, if this was handled in the `modelToApiResp()` it would be fine. 
This actually isn't an issue as the custom date struct in templates has a `UnmarshalJSON()` method, where it's converted to UTC anyway.
\
And it doesn't seem to be an issue with the `updated_at` field either, which I though it could be. That is auto updated by `Gorm` and when a `db.Update()` is called, we don't return the updated struct, but re-fetch it from the DB, where it would get changed to UTC by the hook.
\
So this shouldn't actually ever happen, but it could 😅. If it was converted in the DAO right before outputting, it wouldn't matter what you do with it before that, but it for is solution which requires to be in more places.

### Comment by @jlsherrill on September 11, 2024 at 12:45 PM UTC

ahh!   in that case, we could also implement  AfterSave()  and have it just call AfterFind ?

### Comment by @dominikvagner on September 12, 2024 at 11:31 AM UTC

> ahh! in that case, we could also implement AfterSave() and have it just call AfterFind ?

ahh!, tried it with the `AfterFind` and `AfterSave` together and it seems to work correctly, created some issues in tests, but fixed those and it should be ready 😄

### Comment by @dominikvagner on September 12, 2024 at 12:06 PM UTC

/retest

### Comment by @swadeley on September 16, 2024 at 08:22 AM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Commented on September 13, 2024 at 05:39 PM UTC

### Review by @jlsherrill - Approved on September 13, 2024 at 05:39 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/791*

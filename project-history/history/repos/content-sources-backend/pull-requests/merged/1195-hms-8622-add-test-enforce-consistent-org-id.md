---
type: pull_request
number: 1195
title: "HMS-8622: Add test enforce consistent org id"
state: merged
author: swadeley
created: 2025-09-09T13:01:40Z
updated: 2025-10-09T12:05:48Z
closed: 2025-10-09T12:05:48Z
merged: 2025-10-09T12:05:48Z
base_branch: main
head_branch: swadeley/HMS-8622
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1195
---

# Pull Request #1195: HMS-8622: Add test enforce consistent org id

**Author**: @swadeley
**Created**: September 09, 2025 at 01:01 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `swadeley/HMS-8622`

## Description

## Summary

[HMS-8622](https://issues.redhat.com/browse/HMS-8622) check responses for objects from other orgs

## Testing steps
Tests pass.


---

## Discussion

### Comment by @jlsherrill on September 09, 2025 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-8622

### Comment by @swadeley on September 12, 2025 at 01:45 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Commented on September 10, 2025 at 01:30 PM UTC

### Review by @rverdile - Commented on September 10, 2025 at 01:31 PM UTC

### Review by @rverdile - Commented on September 10, 2025 at 01:46 PM UTC

for the middleware to run, you'd also have to register it in here: https://github.com/content-services/content-sources-backend/blob/main/pkg/router/route.go#L56

### Review by @rverdile - Commented on September 10, 2025 at 02:26 PM UTC

### Review by @swadeley - Commented on September 11, 2025 at 10:23 AM UTC

### Review by @swadeley - Commented on September 11, 2025 at 10:24 AM UTC

### Review by @rverdile - Commented on September 16, 2025 at 08:13 PM UTC

### Review by @rverdile - Commented on September 16, 2025 at 08:13 PM UTC

### Review by @rverdile - Commented on September 16, 2025 at 08:14 PM UTC

### Review by @rverdile - Commented on September 16, 2025 at 08:15 PM UTC

### Review by @rverdile - Commented on September 16, 2025 at 08:26 PM UTC

### Review by @rverdile - Commented on September 16, 2025 at 08:27 PM UTC

### Review by @swadeley - Commented on September 19, 2025 at 04:28 PM UTC

### Review by @swadeley - Commented on September 19, 2025 at 04:28 PM UTC

### Review by @swadeley - Commented on September 19, 2025 at 04:28 PM UTC

### Review by @swadeley - Commented on September 19, 2025 at 04:28 PM UTC

### Review by @swadeley - Commented on September 19, 2025 at 04:29 PM UTC

### Review by @swadeley - Commented on September 19, 2025 at 04:29 PM UTC

### Review by @rverdile - Commented on September 30, 2025 at 03:48 PM UTC

It's looking closer! Here is an example of how the test should look. What we want is to register the actual handler, and register the middleware you've written, so the the test fully tests the handler passing the response to the middleware. Also, this lets us mock the database returning the wrong items

```
type EnforceSuite struct {
	suite.Suite
	reg    *dao.MockDaoRegistry
	tcMock *client.MockTaskClient
	pcMock *pulp_client.MockPulpGlobalClient
	fsMock *feature_service_client.MockFeatureServiceClient
}

func (suite *EnforceSuite) TestExample() {
	/* This block of code would be its own function so it can be reused by other test */
	/* router := setupTestHandler() */
	/* ============================================================================== */
	router := echo.New()
	router.Use(WrapMiddlewareWithSkipper(identity.EnforceIdentity, SkipMiddleware))
	router.Use(EnforceConsistentOrgId)
	router.HTTPErrorHandler = config.CustomHTTPErrorHandler
	pathPrefix := router.Group(api.FullRootPath())

	rh := handler.RepositoryHandler{
		DaoRegistry:          *suite.reg.ToDaoRegistry(),
		TaskClient:           suite.tcMock,
		FeatureServiceClient: suite.fsMock,
	}
	handler.RegisterRepositoryRoutes(pathPrefix, suite.reg.ToDaoRegistry(), &rh.TaskClient, &rh.FeatureServiceClient)

	/* =============================================================================== */

	collection := api.RepositoryCollectionResponse{
		Data: []api.RepositoryResponse{
			{
				UUID:  "test-uuid-1",
				Name:  "Test Repo",
				URL:   "http://example.com/repo1",
				OrgID: "valid-org-id",
			},
			// This one should cause an expected failure
			{
				UUID:  "test-uuid-2",
				Name:  "Test Repo 2",
				URL:   "http://example.com/repo2",
				OrgID: "invalid-org-id",
			},
		},
	}
	paginationData := api.PaginationData{Limit: 100, Offset: 0}
	// mocks the database returning the something with the wrong org
	suite.reg.RepositoryConfig.WithContextMock().On("List", test.MockCtx(), test_handler.MockOrgId, paginationData, api.FilterData{}).Return(collection, int64(1), nil)

	path := fmt.Sprintf("%s/repositories/", api.FullRootPath())
	req := httptest.NewRequest(http.MethodGet, path, nil)
	req.Header.Set(api.IdentityHeader, test_handler.EncodedIdentity(suite.T()))

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	response := rr.Result()
	defer response.Body.Close()

	assert.Equal(suite.T(), http.StatusInternalServerError, response.StatusCode)

}
```

### Review by @rverdile - Commented on October 07, 2025 at 05:10 PM UTC

### Review by @rverdile - Commented on October 07, 2025 at 05:10 PM UTC

### Review by @swadeley - Commented on October 08, 2025 at 10:00 AM UTC

### Review by @swadeley - Commented on October 08, 2025 at 10:00 AM UTC

### Review by @rverdile - Approved on October 08, 2025 at 04:15 PM UTC

looks good, nice job!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1195*

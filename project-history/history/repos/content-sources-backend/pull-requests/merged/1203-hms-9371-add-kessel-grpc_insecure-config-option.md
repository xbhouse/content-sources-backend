---
type: pull_request
number: 1203
title: "HMS-9371: add kessel grpc_insecure config option"
state: merged
author: rverdile
created: 2025-09-16T14:28:09Z
updated: 2025-09-16T18:11:19Z
closed: 2025-09-16T18:11:15Z
merged: 2025-09-16T18:11:15Z
base_branch: main
head_branch: kessel-auth
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1203
---

# Pull Request #1203: HMS-9371: add kessel grpc_insecure config option

**Author**: @rverdile
**Created**: September 16, 2025 at 02:28 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `kessel-auth`

## Description

## Summary
This adds a config option to skip auth and use an insecure connect for the GRPC client

## Testing steps



---

## Discussion

### Comment by @jlsherrill on September 16, 2025 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-9371

### Comment by @rverdile on September 16, 2025 at 02:36 PM UTC

@coderabbitai review

### Comment by @coderabbitai on September 16, 2025 at 02:36 PM UTC

<!-- This is an auto-generated reply by CodeRabbit -->
<details>
<summary>✅ Actions performed</summary>

Review triggered.

> Note: CodeRabbit is an incremental review system and does not re-review already reviewed commits. This command is applicable only when automatic reviews are paused.

</details>

### Comment by @coderabbitai on September 16, 2025 at 02:48 PM UTC

<!-- This is an auto-generated comment: summarize by coderabbit.ai -->
<!-- walkthrough_start -->

## Walkthrough
Adds a new grpc_insecure flag to the Kessel client configuration:
- configs/config.yaml.example: introduces clients.kessel.auth.grpc_insecure.
- deployments/deployment.yaml: adds parameter and env var CLIENTS_KESSEL_AUTH_GRPC_INSECURE (default 'false').
- pkg/config/config.go: extends KesselAuth struct with GrpcInsecure bool (mapstructure:"grpc_insecure") and sets default to false.
- pkg/clients/kessel_client/client.go: updates buildGRPCConnection to make the gRPC connection insecure when auth is disabled or GrpcInsecure is true; otherwise it uses OAuth2 credentials.

## Estimated code review effort
🎯 2 (Simple) | ⏱️ ~10 minutes

## Poem
> A tiny flag in config's fold,  
> grpc_insecure, brave and bold.  
> Env and code now sing in tune,  
> Default safe beneath the moon.  
> Swap the switch when needs unfold.

<!-- walkthrough_end -->


<!-- pre_merge_checks_walkthrough_start -->

## Pre-merge checks and finishing touches
<details>
<summary>❌ Failed checks (2 warnings)</summary>

|     Check name     | Status     | Explanation                                                                                                                                                                                                                                                                                                                                                                                                                  | Resolution                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| :----------------: | :--------- | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|  Description Check | ⚠️ Warning | The PR description includes the required "## Summary" heading and a short summary of the change, but the "## Testing steps" section from the repository template is present and left empty, so there are no verification instructions. Because a required template section is incomplete, the description does not fully meet the repository's template criteria. This prevents reviewers from easily validating the change. | Please populate "## Testing steps" with concrete reproduction and verification steps: show how to enable the flag (e.g., set clients.kessel.auth.grpc_insecure in configs/config.yaml.example or set the deployment parameter CLIENTS_KESSEL_AUTH_GRPC_INSECURE / env var CLIENTS_KESSEL_AUTH_GRPC_INSECURE), include commands to start the service or run relevant unit/integration tests, and state the expected observable outcomes or log messages that confirm the insecure gRPC path is used when enabled and OAuth2 when disabled. Adding a short sample config snippet and any test commands will make the description conform to the template and ready for review. |
| Docstring Coverage | ⚠️ Warning | Docstring coverage is 0.00% which is insufficient. The required threshold is 80.00%.                                                                                                                                                                                                                                                                                                                                         | You can run `@coderabbitai generate docstrings` to improve docstring coverage.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |

</details>
<details>
<summary>✅ Passed checks (1 passed)</summary>

|  Check name | Status   | Explanation                                                                                                                                                                                                                                                                                                                                                                                |
| :---------: | :------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Title Check | ✅ Passed | The title "HMS-9371: add kessel grpc_insecure config option" directly and concisely describes the primary change — adding a kessel gRPC insecure configuration flag and references the associated ticket. It is a single focused sentence that avoids file lists, emojis, or vague wording, making the PR intent clear to reviewers. Therefore it meets the repository's title guidelines. |

</details>

<!-- pre_merge_checks_walkthrough_end -->

<!-- finishing_touch_checkbox_start -->

<details>
<summary>✨ Finishing touches</summary>

- [ ] <!-- {"checkboxId": "7962f53c-55bc-4827-bfbf-6a18da830691"} --> 📝 Generate Docstrings
<details>
<summary>🧪 Generate unit tests</summary>

- [ ] <!-- {"checkboxId": "f47ac10b-58cc-4372-a567-0e02b2c3d479", "radioGroupId": "utg-output-choice-group-3299190064"} -->   Create PR with unit tests
- [ ] <!-- {"checkboxId": "07f1e7d6-8a8e-4e23-9900-8731c2c87f58", "radioGroupId": "utg-output-choice-group-3299190064"} -->   Post copyable unit tests in a comment

</details>

</details>

<!-- finishing_touch_checkbox_end -->

<!-- announcements_start -->

> [!TIP]
> <details>
> <summary>👮 Agentic pre-merge checks are now available in preview!</summary>
> 
> Pro plan users can now enable pre-merge checks in their settings to enforce checklists before merging PRs.
> 
> 	- Built-in checks – Quickly apply ready-made checks to enforce title conventions, require pull request descriptions that follow templates, validate linked issues for compliance, and more.
> 	- Custom agentic checks – Define your own rules using CodeRabbit’s advanced agentic capabilities to enforce organization-specific policies and workflows. For example, you can instruct CodeRabbit’s agent to verify that API documentation is updated whenever API schema files are modified in a PR. Note: Upto 5 custom checks are currently allowed during the preview period. Pricing for this feature will be announced in a few weeks.
> 
> Please see the [documentation](https://docs.coderabbit.ai/pr-reviews/pre-merge-checks) for more information.
> 
> Example:
> 
> ```yaml
> reviews:
>   pre_merge_checks:
>     custom_checks:
>       - name: "Undocumented Breaking Changes"
>         mode: "warning"
>         instructions: |
>           Pass/fail criteria: All breaking changes to public APIs, CLI flags, environment variables, configuration keys, database schemas, or HTTP/GraphQL endpoints must be documented in the "Breaking Change" section of the PR description and in CHANGELOG.md. Exclude purely internal or private changes (e.g., code not exported from package entry points or explicitly marked as internal).
> ```
> 
> Please share your feedback with us on this [Discord post](https://discord.com/channels/1134356397673414807/1414771631775158383).
> 
> </details>

<!-- announcements_end -->

<!-- tips_start -->

---

Thanks for using CodeRabbit! It's free for OSS, and your support helps us grow. If you like it, consider giving us a shout-out.

<details>
<summary>❤️ Share</summary>

- [X](https://twitter.com/intent/tweet?text=I%20just%20used%20%40coderabbitai%20for%20my%20code%20review%2C%20and%20it%27s%20fantastic%21%20It%27s%20free%20for%20OSS%20and%20offers%20a%20free%20trial%20for%20the%20proprietary%20code.%20Check%20it%20out%3A&url=https%3A//coderabbit.ai)
- [Mastodon](https://mastodon.social/share?text=I%20just%20used%20%40coderabbitai%20for%20my%20code%20review%2C%20and%20it%27s%20fantastic%21%20It%27s%20free%20for%20OSS%20and%20offers%20a%20free%20trial%20for%20the%20proprietary%20code.%20Check%20it%20out%3A%20https%3A%2F%2Fcoderabbit.ai)
- [Reddit](https://www.reddit.com/submit?title=Great%20tool%20for%20code%20review%20-%20CodeRabbit&text=I%20just%20used%20CodeRabbit%20for%20my%20code%20review%2C%20and%20it%27s%20fantastic%21%20It%27s%20free%20for%20OSS%20and%20offers%20a%20free%20trial%20for%20proprietary%20code.%20Check%20it%20out%3A%20https%3A//coderabbit.ai)
- [LinkedIn](https://www.linkedin.com/sharing/share-offsite/?url=https%3A%2F%2Fcoderabbit.ai&mini=true&title=Great%20tool%20for%20code%20review%20-%20CodeRabbit&summary=I%20just%20used%20CodeRabbit%20for%20my%20code%20review%2C%20and%20it%27s%20fantastic%21%20It%27s%20free%20for%20OSS%20and%20offers%20a%20free%20trial%20for%20proprietary%20code)

</details>

<sub>Comment `@coderabbitai help` to get the list of available commands and usage tips.</sub>

<!-- tips_end -->

<!-- internal state start -->


<!-- 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 -->

<!-- internal state end -->

### Comment by @rverdile on September 16, 2025 at 03:03 PM UTC

@coderabbitai review

### Comment by @coderabbitai on September 16, 2025 at 03:04 PM UTC

<!-- This is an auto-generated reply by CodeRabbit -->
<details>
<summary>✅ Actions performed</summary>

Review triggered.

> Note: CodeRabbit is an incremental review system and does not re-review already reviewed commits. This command is applicable only when automatic reviews are paused.

</details>

### Comment by @rverdile on September 16, 2025 at 05:45 PM UTC

related MR: https://gitlab.cee.redhat.com/service/app-interface/-/merge_requests/157682

---

## Reviews

### Review by @coderabbitai - Commented on September 16, 2025 at 02:48 PM UTC

**Actionable comments posted: 3**

> [!CAUTION]
> Some comments are outside the diff and can’t be posted inline due to platform limitations.
> 
> 
> 
> <details>
> <summary>⚠️ Outside diff range comments (2)</summary><blockquote>
> 
> <details>
> <summary>deployments/deployment.yaml (1)</summary><blockquote>
> 
> `1214-1216`: **Add a parameter description to discourage prod usage and clarify precedence**
> 
> Suggest adding a description:
> 
> ```diff
> -  - name: CLIENTS_KESSEL_AUTH_GRPC_INSECURE
> -    default: 'false'
> +  - name: CLIENTS_KESSEL_AUTH_GRPC_INSECURE
> +    default: 'false'
> +    description: "Use plaintext/insecure gRPC to Kessel (overrides auth). Do NOT enable in production."
> ```
> 
> </blockquote></details>
> <details>
> <summary>pkg/clients/kessel_client/client.go (1)</summary><blockquote>
> 
> `31-42`: **Don't require OAuth2 creds when grpc_insecure=true**
> 
> NewKesselClient currently errors if auth.enabled is true but creds are missing even when GrpcInsecure is set; GetRootWorkspaceID also always tries to fetch a token when auth.enabled is true. Gate both checks on GrpcInsecure.
> 
> - File: pkg/clients/kessel_client/client.go — NewKesselClient: change the validation to skip when GrpcInsecure is true.
> ```diff
> -	if kesselConfig.Auth.Enabled && (kesselConfig.Auth.ClientID == "" || kesselConfig.Auth.ClientSecret == "" || kesselConfig.Auth.OIDCIssuer == "") {
> +	if kesselConfig.Auth.Enabled && !kesselConfig.GrpcInsecure &&
> +		(kesselConfig.Auth.ClientID == "" || kesselConfig.Auth.ClientSecret == "" || kesselConfig.Auth.OIDCIssuer == "") {
>  		return nil, fmt.Errorf("kessel authentication is enabled but not configured")
>  	}
> ```
> 
> - File: pkg/clients/kessel_client/client.go — GetRootWorkspaceID: avoid calling getToken / adding the Authorization header when GrpcInsecure is true.
> ```diff
> -	if config.Get().Clients.Kessel.Auth.Enabled {
> +	kesselConfig := config.Get().Clients.Kessel
> +	if kesselConfig.Auth.Enabled && !kesselConfig.GrpcInsecure {
>  		log.Debug().Msgf("[Kessel] GetRootWorkspaceID: orgID %v", orgID)
>  		token, err := k.getToken(ctx)
>  		if err != nil {
>  			return "", statusCode, fmt.Errorf("error getting auth token: %w", err)
>  		}
>  		req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))
>  	}
> ```
> 
> </blockquote></details>
> 
> </blockquote></details>

<details>
<summary>📜 Review details</summary>

**Configuration used**: CodeRabbit UI

**Review profile**: ASSERTIVE

**Plan**: Pro

<details>
<summary>📥 Commits</summary>

Reviewing files that changed from the base of the PR and between e18c6916e932862324039e908d3fda7dd1daa5f1 and e6907ad4e3bfca3ad57fd6c4698ba9474b332421.

</details>

<details>
<summary>📒 Files selected for processing (4)</summary>

* `configs/config.yaml.example` (1 hunks)
* `deployments/deployment.yaml` (2 hunks)
* `pkg/clients/kessel_client/client.go` (1 hunks)
* `pkg/config/config.go` (2 hunks)

</details>

<details>
<summary>⏰ Context from checks skipped due to timeout of 90000ms. You can increase the timeout in your CodeRabbit configuration to a maximum of 15 minutes (900000ms). (3)</summary>

* GitHub Check: playwright-tests
* GitHub Check: Red Hat Konflux / content-sources-backend-bonfire-tekton / content-sources-backend
* GitHub Check: Test

</details>

<details>
<summary>🔇 Additional comments (2)</summary><blockquote>

<details>
<summary>pkg/config/config.go (1)</summary><blockquote>

`327-328`: **Default key location assumes grpc_insecure under auth (good if you move the field)**

No change needed if you adopt the struct fix above. If you instead keep the field at Kessel level, please change this default to clients.kessel.grpc_insecure and align other files accordingly.

</blockquote></details>
<details>
<summary>deployments/deployment.yaml (1)</summary><blockquote>

`372-373`: **Env var is wired correctly; please add a cautionary comment**

Looks good. Consider adding a brief comment noting this bypasses auth/TLS and should not be enabled in prod. This helps operators avoid foot‑guns.

</blockquote></details>

</blockquote></details>

</details>

<!-- This is an auto-generated comment by CodeRabbit for review status -->

### Review by @coderabbitai - Commented on September 16, 2025 at 03:09 PM UTC

**Actionable comments posted: 1**

> [!CAUTION]
> Some comments are outside the diff and can’t be posted inline due to platform limitations.
> 
> 
> 
> <details>
> <summary>⚠️ Outside diff range comments (2)</summary><blockquote>
> 
> <details>
> <summary>deployments/deployment.yaml (1)</summary><blockquote>
> 
> `1214-1216`: **Add a parameter description to flag the security risk.**
> 
> A brief description helps operators avoid enabling this outside dev.
> 
> 
> ```diff
> -  - name: CLIENTS_KESSEL_AUTH_GRPC_INSECURE
> -    default: 'false'
> +  - name: CLIENTS_KESSEL_AUTH_GRPC_INSECURE
> +    default: 'false'
> +    description: 'DEV ONLY. If true, Kessel gRPC uses an insecure (plaintext) connection and skips OAuth.'
> ```
> 
> </blockquote></details>
> <details>
> <summary>pkg/clients/kessel_client/client.go (1)</summary><blockquote>
> 
> `36-41`: **Don’t require OAuth creds when grpc_insecure is true.**
> 
> With grpc_insecure enabled, we should not fail initialization due to missing OAuth settings.
> 
> 
> ```diff
>    kesselConfig := config.Get().Clients.Kessel
> -  if kesselConfig.Auth.Enabled && (kesselConfig.Auth.ClientID == "" || kesselConfig.Auth.ClientSecret == "" || kesselConfig.Auth.OIDCIssuer == "") {
> +  if kesselConfig.Auth.Enabled && !kesselConfig.Auth.GrpcInsecure &&
> +     (kesselConfig.Auth.ClientID == "" || kesselConfig.Auth.ClientSecret == "" || kesselConfig.Auth.OIDCIssuer == "") {
>      return nil, fmt.Errorf("kessel authentication is enabled but not configured")
>    }
> ```
> 
> </blockquote></details>
> 
> </blockquote></details>

<details>
<summary>♻️ Duplicate comments (1)</summary><blockquote>

<details>
<summary>configs/config.yaml.example (1)</summary><blockquote>

`160-161`: **Set an explicit default and warn this is dev-only.**

Leaving grpc_insecure empty is ambiguous. Let’s make the safe default obvious and document the risk.



```diff
       oidc_issuer:
-      grpc_insecure:
+      # WARNING: dev/local only. Disables TLS/auth for the Kessel gRPC client.
+      grpc_insecure: false
```

</blockquote></details>

</blockquote></details>

<details>
<summary>📜 Review details</summary>

**Configuration used**: CodeRabbit UI

**Review profile**: ASSERTIVE

**Plan**: Pro

<details>
<summary>📥 Commits</summary>

Reviewing files that changed from the base of the PR and between e6907ad4e3bfca3ad57fd6c4698ba9474b332421 and 89b8b99be962afe62cc7d0fa8f150c7982079d84.

</details>

<details>
<summary>📒 Files selected for processing (4)</summary>

* `configs/config.yaml.example` (1 hunks)
* `deployments/deployment.yaml` (2 hunks)
* `pkg/clients/kessel_client/client.go` (1 hunks)
* `pkg/config/config.go` (2 hunks)

</details>

<details>
<summary>⏰ Context from checks skipped due to timeout of 90000ms. You can increase the timeout in your CodeRabbit configuration to a maximum of 15 minutes (900000ms). (3)</summary>

* GitHub Check: Red Hat Konflux / content-sources-backend-on-pull-request
* GitHub Check: playwright-tests
* GitHub Check: Test

</details>

<details>
<summary>🔇 Additional comments (4)</summary><blockquote>

<details>
<summary>deployments/deployment.yaml (1)</summary><blockquote>

`372-373`: **Ensure all pods that use Kessel gRPC receive the flag.**

You added CLIENTS_KESSEL_AUTH_GRPC_INSECURE to the service container. If task-worker (or any job) calls Kessel via gRPC, it also needs this env var.


Would you like me to add it to task-worker (and any jobs) in this PR?

</blockquote></details>
<details>
<summary>pkg/clients/kessel_client/client.go (1)</summary><blockquote>

`213-218`: **LGTM: insecure takes precedence over auth.**

The precedence and field path look correct now: Insecure when auth is disabled or grpc_insecure is true; otherwise OAuth2.

</blockquote></details>
<details>
<summary>pkg/config/config.go (2)</summary><blockquote>

`147-148`: **LGTM: field is in the correct struct with the right tag.**

GrpcInsecure under KesselAuth matches the YAML/env path clients.kessel.auth.grpc_insecure.

---

`327-328`: **LGTM: safe default is false.**

Defaulting grpc_insecure to false is the right posture.

</blockquote></details>

</blockquote></details>

</details>

<!-- This is an auto-generated comment by CodeRabbit for review status -->

### Review by @xbhouse - Approved on September 16, 2025 at 06:01 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1203*

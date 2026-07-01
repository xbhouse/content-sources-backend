---
type: pull_request
number: 131
title: "chore(config): update base configuration set up"
state: open
author: adonispuente
created: 2026-04-21T20:31:49Z
updated: 2026-06-25T20:57:15Z
base_branch: main
head_branch: config
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/131
---

# Pull Request #131: chore(config): update base configuration set up

**Author**: @adonispuente
**Created**: April 21, 2026 at 08:31 PM UTC
**Status**: Open
**Labels**: None
**Base**: `main` ← **Head**: `config`

## Description

In an effort to make onboarding easier for anyone trying to run things locally, I wanted to update the repo to include the custom routes and associated file to just simple be able to run `podman compose up proxy-iop  `

Summary

  Adds IOP (Insights on Premises) mode support for local development against Satellite/Foreman instances.
  
  Key Changes

  - IOP mode toggle: Set IOP=true to enable IOP-specific configuration
  - Location header rewriting: Keeps developers on the local proxy during redirects (prevents navigation to IOP
  instance URLs)
  - strip_prefix support: Enables path prefix stripping for IOP asset routing
  - Conditional Host header: Omits Host header in IOP mode (required for Foreman compatibility)
  - Caddyfile.iop: Dedicated config with tls_insecure_skip_verify for self-signed IOP certificates

  Testing Locally

  1. Comment out the test step in Dockerfile (lines 23-28)
  - Tests hang in local Docker builds due to Caddy startup when sourced
  - This is a pre-existing issue, not specific to IOP changes
  - Tests pass when run directly outside Docker (13/13)

  2. Build the proxy image:
  podman build -t ghcr.io/redhatinsights/frontend-development-proxy:latest .

  3. Test with vulnerability-ui:
  cd vulnerability-ui
  IOP_URL="https://ip-10-0-167-79.rhos-01.prod.psi.rdu2.redhat.com" npm run start:proxy:iop:local
  
  Access at: https://iop.foo.redhat.com:1337/insights/vulnerability

  Testing

  Tested against IOP instance running Foreman 3.16. Verified:
  - Location header rewrites keep users on local proxy URL
  - Custom route merging works correctly
  - Local code changes load through proxy successfully
  
  Related PRs

  - https://github.com/RedHatInsights/frontend-components/pull/2308
  - vulnerability-ui: foreman-3.16 branch IOP setup:  https://github.com/RedHatInsights/vulnerability-ui/pull/2695
  - insights-advisor-frontend: foreman-3.16 branch IOP setup: https://github.com/RedHatInsights/insights-advisor-frontend/pull/2036
  -proxy: https://github.com/RedHatInsights/frontend-development-proxy/pull/131

---

## Discussion

### Comment by @adonispuente on June 24, 2026 at 07:42 PM UTC

@charlesmulder I believe ive addressed everything now, thank you so much for the comprehensive review. a bunch of good points 

### Comment by @adonispuente on June 24, 2026 at 08:13 PM UTC

@charlesmulderThe issue is that removing tls_insecure_skip_verify breaks IOP connection because the IOP instances we use like
  (https://ip-10-0-167-79.rhos-01.prod.psi.rdu2.redhat.com) uses a certificate that's not trusted by mkcert.. That took me a minute but thats why it doesnt work

---

## Reviews

### Review by @charlesmulder - Commented on May 06, 2026 at 09:58 AM UTC

## Review Summary

Integration test validates implementation works end-to-end with frontend-components PR #2308. Core functionality solid, but need fixes before merge.

## Issues to Address

### 1. Hardcoded IP not portable (package.json:7)
```json
"IOP": "IOP=true HCC_ENV=iop HCC_ENV_URL=https://ip-10-0-167-196.rhos-01.prod.psi.rdu2.redhat.com ..."
```
Internal IP won't work for all developers. Use env var with fallback:
```json
"IOP": "IOP=true HCC_ENV=iop HCC_ENV_URL=${HCC_IOP_URL:-https://console.stage.redhat.com} ..."
```

### 2. Missing newline at EOF (entrypoint.sh:87)
POSIX requires newline at end of file. Add newline after final line.

### 3. Security warning too weak (README.md:76-79)
Current text understates risk of `tls_insecure_skip_verify`. Replace with:
```markdown
**⚠️ SECURITY: IOP MODE IS DEVELOPMENT ONLY**

In IOP mode specifically:
- The fallback proxy uses TLS transport with `tls_insecure_skip_verify` which **disables all certificate validation**
- **DO NOT use in production or with untrusted networks** - vulnerable to MITM attacks
- Only use on trusted local development networks
- Generated local route `reverse_proxy` blocks omit `header_up Host {http.reverse_proxy.upstream.hostport}`.
```

### 4. Unnecessary flag (entrypoint.sh:87)
```bash
/usr/bin/caddy run --config "$CADDY_CONFIG_PATH" --adapter caddyfile
```
Caddy auto-detects format from filename. Remove `--adapter caddyfile`.

## What Works
✅ Non-breaking - default behavior unchanged  
✅ Clear docs - README explains IOP mode  
✅ Simple toggle - `IOP=true` explicit  
✅ Docker Compose - easier local dev  
✅ Route separation - IOP routes isolated  
✅ Integration verified - works with frontend-components PR #2308

### Review by @charlesmulder - Commented on May 06, 2026 at 10:01 AM UTC

## Additional Issue: Unnecessary package.json

### 5. package.json in non-JavaScript project

Repository is Go/Caddy/Bash - no JavaScript code. package.json only used as task runner for docker compose commands:

```json
"scripts": {
  "dev-proxy": "docker compose up --build",
  "dev-proxy:down": "docker compose down",
  "IOP": "IOP=true HCC_ENV=iop ..."
}
```

**Problem:** Adds npm dependency + package.json/package-lock.json just to alias shell commands. Confusing for Go project.

**Better alternatives:**

**Option 1 - Makefile** (standard for Go projects):
```make
.PHONY: dev-proxy dev-proxy-down iop

dev-proxy:
	docker compose up --build

dev-proxy-down:
	docker compose down

iop:
	IOP=true HCC_ENV=iop HCC_ENV_URL=${HCC_IOP_URL:-https://console.stage.redhat.com} docker compose up --build
```

**Option 2 - Just document commands in README:**
```bash
# Normal mode
docker compose up --build

# IOP mode  
IOP=true HCC_ENV=iop HCC_ENV_URL=${HCC_IOP_URL:-https://console.stage.redhat.com} docker compose up --build

# Stop
docker compose down
```

Recommend removing package.json/package-lock.json and using Makefile or direct commands.

### Review by @charlesmulder - Commented on May 06, 2026 at 10:07 AM UTC

## Security Concern: TLS Verification Skip

### 6. tls_insecure_skip_verify in fallback handler

Caddyfile.iop disables TLS certificate validation for all fallback traffic (requests not matching custom routes):

```caddyfile
handle {
  reverse_proxy ${HCC_ENV_URL} {
    transport http {
      tls_insecure_skip_verify  # ← Disables all cert validation
    }
  }
}
```

**Risk:**
- All non-custom-route traffic (auth, APIs, etc) to `https://ip-10-0-167-196.rhos-01.prod.psi.rdu2.redhat.com` vulnerable to MITM
- No certificate validation = attacker can intercept/modify traffic
- Applies to sensitive data (auth tokens, user data)

**Question:** Why is TLS skip needed for IOP environment?

**Better alternatives:**

**If internal env has self-signed certs → Add CA to container:**
```dockerfile
# In Dockerfile
COPY internal-ca.crt /usr/local/share/ca-certificates/
RUN update-ca-certificates
```
Remove `tls_insecure_skip_verify` entirely.

**If internal env doesn't need TLS → Use HTTP:**
```bash
HCC_ENV_URL=http://ip-10-0-167-196.rhos-01.prod.psi.rdu2.redhat.com
```
Remove `transport http` block.

**If skip truly required → Document restriction:**
- Add to README: IOP mode MUST only run on isolated internal networks
- Document threat model (assumes network-level isolation)
- Consider env var toggle: `IOP_SKIP_TLS_VERIFY` (explicit opt-in)

Please clarify why TLS verification skip is necessary and consider alternatives.

### Review by @charlesmulder - Commented on May 06, 2026 at 10:09 AM UTC

## Recommended Solution for TLS Skip Issue

### mkcert approach (already used in insights-chrome)

The insights-chrome project already solves this problem using [mkcert](https://github.com/FiloSottile/mkcert) for locally-trusted certificates.

**Same approach works for proxy IOP mode:**

**1. User one-time setup:**
```sh
# Install mkcert
# See: https://github.com/FiloSottile/mkcert#installation

# Install local CA
mkcert -install
```

**2. Add CA to proxy container:**
```dockerfile
# In Dockerfile - add before final stage
COPY mkcert-ca.crt /usr/local/share/ca-certificates/
RUN update-ca-certificates
```

**3. Remove tls_insecure_skip_verify from Caddyfile.iop:**
```diff
handle {
  reverse_proxy ${HCC_ENV_URL} {
    header_up Accept-Encoding "gzip;q=0,deflate,sdch"
    header_up -Origin
-   transport http {
-     tls_insecure_skip_verify
-   }
  }
}
```

**4. Document in README:**
```markdown
### IOP Mode TLS Setup

IOP environment may use self-signed certificates. To enable proper TLS verification:

1. Install mkcert (one-time): https://github.com/FiloSottile/mkcert#installation
2. Install local CA: \`mkcert -install\`
3. Copy CA to repo: \`cp ~/.local/share/mkcert/rootCA.pem ./mkcert-ca.crt\`
4. Rebuild image: \`podman build -t frontend-development-proxy .\`

Container will trust mkcert CA and validate IOP certificates properly.
```

**Benefits:**
- Proper TLS validation (no security warnings)
- Consistent with insights-chrome approach
- Works with self-signed internal certs
- No MITM vulnerability

See insights-chrome README lines 98-127 for reference implementation.

### Review by @adonispuente - Commented on June 24, 2026 at 09:14 PM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/131*

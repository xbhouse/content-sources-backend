---
type: pull_request
number: 1091
title: "HMS-5998: properly get newest version for modules"
state: merged
author: jlsherrill
created: 2025-04-25T15:58:37Z
updated: 2025-04-29T16:06:15Z
closed: 2025-04-29T15:31:06Z
merged: 2025-04-29T15:31:06Z
base_branch: main
head_branch: 5998_2
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1091
---

# Pull Request #1091: HMS-5998: properly get newest version for modules

**Author**: @jlsherrill
**Created**: April 25, 2025 at 03:58 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `5998_2`

## Description

## Summary

The previous PR for HMS-5998 had a bug where if a module stream existed in a different repository with a higher version, that entire module stream would be ignored. 

## Testing steps

to reproduce the original issue, import and sync  all the appstream repos:

```
make repos-import
go run cmd/external-repos/main.go introspect  https://cdn.redhat.com/content/dist/rhel8/8/x86_64/appstream/os
go run cmd/external-repos/main.go introspect  https://cdn.redhat.com/content/dist/rhel9/9/x86_64/appstream/os
go run cmd/external-repos/main.go introspect https://cdn.redhat.com/content/dist/rhel9/9/aarch64/appstream/os
```

wait for them to finish.

Then trigger rpm search:

```
curl -X POST --location "http://localhost:8000/api/content-sources/v1.0/rpms/names" \
    -H "x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiIxIn0sImFjY291bnRfbnVtYmVyIjoiZm9vIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMSJ9fX0K" \
    -H "Content-Type: application/json" \
    -d '{"urls":["https://cdn.redhat.com/content/dist/rhel9/9/aarch64/appstream/os"],
          "include_package_sources": true,
          "search":"nginx"}'
```


You should see modules:

```
[
  {
    "package_name": "nginx",
    "summary": "A high performance web server and reverse proxy server",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  },
  {
    "package_name": "nginx-all-modules",
    "summary": "A meta package that installs all available Nginx modules",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  },
  {
    "package_name": "nginx-core",
    "summary": "nginx minimal core",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  },
  {
    "package_name": "nginx-filesystem",
    "summary": "The basic directory layout for the Nginx server",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  },
  {
    "package_name": "nginx-mod-devel",
    "summary": "Nginx module development files",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  },
  {
    "package_name": "nginx-mod-http-image-filter",
    "summary": "Nginx HTTP image filter module",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  },
  {
    "package_name": "nginx-mod-http-perl",
    "summary": "Nginx HTTP perl module",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  },
  {
    "package_name": "nginx-mod-http-xslt-filter",
    "summary": "Nginx XSLT module",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  },
  {
    "package_name": "nginx-mod-mail",
    "summary": "Nginx mail modules",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  },
  {
    "package_name": "nginx-mod-stream",
    "summary": "Nginx stream modules",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  }
]
```



---

## Discussion

### Comment by @jlsherrill on April 25, 2025 at 04:32 PM UTC

https://issues.redhat.com/browse/HMS-5998

### Comment by @swadeley on April 28, 2025 at 12:33 PM UTC

/retest

### Comment by @swadeley on April 29, 2025 at 03:30 PM UTC

Hi

I will  merge and test in stage as I need the appstream not just codeready repo.

### Comment by @swadeley on April 29, 2025 at 04:06 PM UTC

Hi

There are 20 modules.

lgtm

```
[
  {
    "package_name": "nginx",
    "summary": "A high performance web server and reverse proxy server",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  },
  {
    "package_name": "nginx-all-modules",
    "summary": "A meta package that installs all available Nginx modules",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  },
  {
    "package_name": "nginx-core",
    "summary": "nginx minimal core",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  },
  {
    "package_name": "nginx-filesystem",
    "summary": "The basic directory layout for the Nginx server",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  },
  {
    "package_name": "nginx-mod-devel",
    "summary": "Nginx module development files",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  },
  {
    "package_name": "nginx-mod-http-image-filter",
    "summary": "Nginx HTTP image filter module",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  },
  {
    "package_name": "nginx-mod-http-perl",
    "summary": "Nginx HTTP perl module",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  },
  {
    "package_name": "nginx-mod-http-xslt-filter",
    "summary": "Nginx XSLT module",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  },
  {
    "package_name": "nginx-mod-mail",
    "summary": "Nginx mail modules",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  },
  {
    "package_name": "nginx-mod-stream",
    "summary": "Nginx stream modules",
    "package_sources": [
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.22",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324053651",
        "description": "nginx 1.22 webserver module"
      },
      {
        "type": "module",
        "name": "nginx",
        "stream": "1.24",
        "context": "9",
        "arch": "aarch64",
        "version": "9050020250324055038",
        "description": "nginx 1.24 webserver module"
      }
    ]
  }
]
```

---

## Reviews

### Review by @rverdile - Approved on April 25, 2025 at 06:19 PM UTC

works!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1091*

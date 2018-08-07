> # Click! [![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Link%20Manager%20as%20a%20Service&url=https://kamilsk.github.io/click/&via=ikamilsk&hashtags=go,service,link-manager,link-storage,link-shortener,url-shortener)
> [![Analytics](https://ga-beacon.appspot.com/UA-109817251-20/click/readme?pixel)](https://kamilsk.github.io/click/)
> Link Manager as a Service &mdash; your personal link storage and URL shortener.

[![Patreon](https://img.shields.io/badge/patreon-donate-orange.svg)](https://www.patreon.com/octolab)
[![Build Status](https://travis-ci.org/kamilsk/click.svg?branch=master)](https://travis-ci.org/kamilsk/click)
[![Code Coverage](https://scrutinizer-ci.com/g/kamilsk/click/badges/coverage.png?b=master)](https://scrutinizer-ci.com/g/kamilsk/click/?branch=master)
[![Code Quality](https://scrutinizer-ci.com/g/kamilsk/click/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/kamilsk/click/?branch=master)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Roadmap

- [x] v1: [MVP](https://github.com/kamilsk/click/projects/1)
  - [**May 31, 2018**](https://github.com/kamilsk/click/milestone/1)
  - Main concepts and working prototype.
- [ ] v2: [URL shortener and CLI CRUD](https://github.com/kamilsk/click/projects/2)
  - [**August 31, 2018**](https://github.com/kamilsk/click/milestone/2)
  - Command line interface for create, read, update and delete operations above gRPC.
  - URL shortener functionality.
- [ ] v3: [DSL for rules and CSI](https://github.com/kamilsk/click/projects/3)
  - [**September 30, 2018**](https://github.com/kamilsk/click/milestone/3)
  - Client-side integration.
  - Domain-specific language to define target rules.
- [ ] v4: [GUI CRUD](https://github.com/kamilsk/click/projects/4)
  - [**October 31, 2018**](https://github.com/kamilsk/click/milestone/4)
  - Graphical user interface and admin panel to perform create, read, update and delete operations.
- [ ] Click!, SaaS
  - **December 31, 2018**
  - Ready to apply on Cloud.
  - Move to [OctoLab](https://github.com/octolab/) organization.

## Motivation

- We need better integration with static site generators like [Hugo](https://gohugo.io/).
- We want better products than [Bitly](https://bitly.com/) or [Ow.ly](http://ow.ly/).
- We have to full control over our users' data and protect it from third parties.

## Quick start

Requirements:

- Docker 17.09.0-ce or above
- Docker Compose 1.16.1 or above
- Go 1.9.2 or above
- GNU Make 3.81 or above

```bash
$ make up demo status

     Name                    Command               State                                  Ports
-------------------------------------------------------------------------------------------------------------------------------
click_db_1        docker-entrypoint.sh postgres    Up      0.0.0.0:5432->5432/tcp
click_server_1    /bin/sh -c envsubst '$SERV ...   Up      80/tcp, 0.0.0.0:80->8080/tcp
click_service_1   click run --with-profiling ...   Up      0.0.0.0:8080->80/tcp, 0.0.0.0:8090->8090/tcp, 0.0.0.0:8091->8091/tcp

$ curl http://localhost:8080/api/v1/a382922d-b615-4227-b598-6d3633c397aa
# {
#   "id": "a382922d-b615-4227-b598-6d3633c397aa",
#   "name": "Click! - Link Manager as a Service",
#   "status": "active",
#   "aliases": [
#     {
#       "id": 1,
#       "namespace": "global",
#       "urn": "github/click"
#     },
#     {
#       "id": 7,
#       "namespace": "global",
#       "urn": "github/click!"
#     }
#   ],
#   "targets": [
#     {
#       "id": 1,
#       "uri": "https://github.com/kamilsk/click",
#       "rule": {
#         "description": "Project location",
#         "tags": ["src"]
#       }
#     },
#     {
#       "id": 2,
#       "uri": "https://kamilsk.github.io/click/",
#       "rule": {
#         "description": "Promotion page",
#         "alias": 7,
#         "tags": ["promo"],
#         "match": 1
#       }
#     }
#   ]
# }
$ curl -v --cookie "token=41ca5e09-3ce2-4094-b108-3ecc257c6fa4" http://localhost:8080/github/click!
# > GET /github/click! HTTP/1.1
# > Host: localhost:8080
# > User-Agent: curl/7.54.0
# > Accept: */*
# > Cookie: token=41ca5e09-3ce2-4094-b108-3ecc257c6fa4
# >
# < HTTP/1.1 302 Found
# < Location: https://kamilsk.github.io/click/
# < Set-Cookie: token=41ca5e09-3ce2-4094-b108-3ecc257c6fa4; Path=/; HttpOnly; Secure
# < Date: Wed, 11 Apr 2018 17:37:48 GMT
# < Content-Length: 0
# <
```

## Specification

### API

You can find API specification [here](env/rest.http). Also, we recommend using [Insomnia](https://insomnia.rest)
HTTP client to work with the API - you can import data for it from the [file](env/insomnia.json).

### CLI

```bash
$ click --help
Click!

Usage:
  click [command]

Available Commands:
  completion  Print Bash or Zsh completion
  help        Help about any command
  migrate     Apply database migration
  run         Start HTTP server
  version     Show application version

Flags:
  -h, --help   help for click

Use "click [command] --help" for more information about a command.
```

#### Bash and Zsh completions

You can find completion files [here](https://github.com/kamilsk/shared/tree/dotfiles/bash_completion.d) or
build your own using these commands

```bash
$ click completion bash > /path/to/bash_completion.d/click.sh
$ click completion zsh  > /path/to/zsh-completions/_click.zsh
```

## Installation

### Brew

```bash
$ brew install kamilsk/tap/click
```

### Binary

```bash
$ export VER=1.0.0      # all available versions are on https://github.com/kamilsk/click/releases
$ export REQ_OS=Linux   # macOS and Windows are also available
$ export REQ_ARCH=64bit # 32bit is also available
$ wget -q -O click.tar.gz \
       https://github.com/kamilsk/click/releases/download/"${VER}/click_${VER}_${REQ_OS}-${REQ_ARCH}".tar.gz
$ tar xf click.tar.gz -C "${GOPATH}"/bin/ && rm click.tar.gz
```

### Docker Hub

```bash
$ docker pull kamilsk/click:1.x
```

### From source code

```bash
$ egg github.com/kamilsk/click@^1.0.0 -- make test install
```

#### Mirror

```bash
$ egg bitbucket.org/kamilsk/click@^1.0.0 -- make test install
```

> [egg](https://github.com/kamilsk/egg) is an `extended go get`.

## Update

This application is in a state of [MVP](https://en.wikipedia.org/wiki/Minimum_viable_product) and under active
development. [SemVer](https://semver.org/) is used for releases, and you can easily be updated within minor versions,
but major versions can be not [BC](https://en.wikipedia.org/wiki/Backward_compatibility)-safe.

## Notes

- [research](../../tree/research)
- tested on Go 1.9 and 1.10

---

[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/kamilsk/click)
[![@kamilsk](https://img.shields.io/badge/author-%40kamilsk-blue.svg)](https://twitter.com/ikamilsk)
[![@octolab](https://img.shields.io/badge/sponsor-%40octolab-blue.svg)](https://twitter.com/octolab_inc)

made with ❤️ by [OctoLab](https://www.octolab.org/)

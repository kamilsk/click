> # Click! [![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Link%20Manager%20as%20a%20Service&url=https://kamilsk.github.io/click/&via=ikamilsk&hashtags=go,service,link-manager,link-storage,link-shortener)
> [![Analytics](https://ga-beacon.appspot.com/UA-109817251-20/click/readme?pixel)](https://kamilsk.github.io/click/)
> Link Manager as a Service.

[![Patreon](https://img.shields.io/badge/patreon-donate-orange.svg)](https://www.patreon.com/octolab)
[![Build Status](https://travis-ci.org/kamilsk/click.svg?branch=master)](https://travis-ci.org/kamilsk/click)
[![Coverage Status](https://coveralls.io/repos/github/kamilsk/click/badge.svg)](https://coveralls.io/github/kamilsk/click)
[![GoDoc](https://godoc.org/github.com/kamilsk/click?status.svg)](https://godoc.org/github.com/kamilsk/click)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Quick start

```bash
$ make up demo status

    Name                   Command               State                                  Ports
-----------------------------------------------------------------------------------------------------------------------------
env_db_1        docker-entrypoint.sh postgres    Up      0.0.0.0:5432->5432/tcp
env_server_1    /bin/sh -c envsubst '$SERV ...   Up      80/tcp, 0.0.0.0:80->8080/tcp
env_service_1   click run --with-profile - ...   Up      0.0.0.0:8080->80/tcp, 0.0.0.0:8090->8090/tcp, 0.0.0.0:8091->8091/tcp

$ curl http://localhost:8080/api/v1/a382922d-b615-4227-b598-6d3633c397aa
$ curl -v http://localhost:8080/github/click!
> GET /github/click! HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 302 Found
< Location: https://github.com/kamilsk/click
< Date: Sat, 31 Mar 2018 20:40:42 GMT
< Content-Length: 0
<
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
   [command]

Available Commands:
  help        Help about any command
  migrate     Apply database migration
  run         Start HTTP server
  version     Show application version

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.
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
$ docker pull kamilsk/click:latest
```

### From source code

```bash
$ egg github.com/kamilsk/click@^1.0.0 -- make generate test install
```

#### Mirror

```bash
$ egg bitbucket.org/kamilsk/click@^1.0.0 -- make generate test install
```

> [egg](https://github.com/kamilsk/egg) is an `extended go get`.

#### Requirements

- Docker 17.09.0-ce or above
- Docker Compose 1.16.1 or above
- Go 1.9.2 or above
- GNU Make 3.81 or above

## Notes

- brief roadmap
  - [x] v1: MVP
  - [ ] v2: ...
  - [ ] v3: ...
  - [ ] v4: CRUD
  - [ ] v5: GUI
  - [ ] Click!, SaaS
- tested on Go 1.9 and 1.10

---

[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/kamilsk/click)
[![@kamilsk](https://img.shields.io/badge/author-%40kamilsk-blue.svg)](https://twitter.com/ikamilsk)
[![@octolab](https://img.shields.io/badge/sponsor-%40octolab-blue.svg)](https://twitter.com/octolab_inc)

made with ❤️ by [OctoLab](https://www.octolab.org/)

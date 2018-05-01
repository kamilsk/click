OPEN_BROWSER       =
SUPPORTED_VERSIONS = 1.9 1.10 latest


include makes/env.mk
include makes/docker.mk
include makes/local.mk
include env/cmd.mk
include env/docker.mk
include env/docker-compose.mk
include env/tools.mk


.PHONY: code-quality-check
code-quality-check: ARGS = \
	--exclude=".*_test\.go:.*error return value not checked.*\(errcheck\)$$" \
	--exclude="duplicate of.*_test.go.*\(dupl\)$$" \
	--exclude="static/bindata.go" \
	--exclude="mock_.*.go" \
	--vendor --deadline=5m ./... | sort
code-quality-check: docker-tool-gometalinter

.PHONY: code-quality-report
code-quality-report:
	time make code-quality-check | tail +7 | tee report.out


.PHONY: dev-up
dev-up: up demo stop-server stop-service clear status dev-server


.PHONY: pull-template
pull-template:
	rm -rf template
	git clone git@bitbucket.org:octotpl/materialkit.git template
	( \
	  cd template && \
	  git checkout 2.x && \
	  git describe --tags \
	)
	( \
	  cp -n template/Template/assets/css/material-kit.min.css docs/assets/css/ && \
	  cp -n template/Template/assets/js/bootstrap-material-design.min.js docs/assets/js/ \
	  cp -n template/Template/assets/js/material-kit.min.js docs/assets/js/ \
	)
	rm -rf template/.git

SHELL := /bin/sh
GO ?= go
KUBECTL ?= kubectl
GOCACHE ?= /tmp/urlog-go-cache

.PHONY: help
help:
	@printf '%s\n' 'Urlog make targets:'
	@printf '%s\n' '  make check          Run fmt, tests, YAML validation, and Kustomize render checks'
	@printf '%s\n' '  make fmt            Format Go code'
	@printf '%s\n' '  make test           Run Go unit tests'
	@printf '%s\n' '  make validate-yaml  Parse repository YAML files'
	@printf '%s\n' '  make kustomize      Render Urlog Kubernetes Kustomize overlays'

.PHONY: check
check: fmt test validate-yaml kustomize

.PHONY: fmt
fmt:
	GOCACHE=$(GOCACHE) $(GO) fmt ./...

.PHONY: test
test:
	GOCACHE=$(GOCACHE) $(GO) test ./...

.PHONY: validate-yaml
validate-yaml:
	ruby -e 'require "yaml"; ARGV.each { |f| YAML.load_stream(File.read(f)); puts "ok #{f}" }' $$(find bootstrap deploy examples modules -name '*.yaml' -o -name '*.yml')

.PHONY: kustomize
kustomize:
	$(KUBECTL) kustomize deploy/kubernetes/base >/tmp/urlog-kustomize-base.yaml
	$(KUBECTL) kustomize deploy/kubernetes/overlays/k3s >/tmp/urlog-kustomize-k3s.yaml
	$(KUBECTL) kustomize deploy/kubernetes/overlays/docker-desktop >/tmp/urlog-kustomize-docker-desktop.yaml
	$(KUBECTL) kustomize deploy/kubernetes/overlays/airgap >/tmp/urlog-kustomize-airgap.yaml

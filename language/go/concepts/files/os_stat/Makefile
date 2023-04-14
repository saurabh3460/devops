CURDIR=$(shell pwd)

ifeq (, $(shell which govvv))
$(shell go install github.com/ahmetb/govvv)	# This works for older go version
$(shell go install github.com/ahmetb/govvv@latest) # This works for new go version
endif


PKG := $(shell go list ./libs)
GIT_INFO := $(shell govvv -flags -pkg $(PKG))


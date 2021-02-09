GOMODHASH := $(shell sha256sum go.mod | cut -d" " -f1)
IMAGE_TAG ?= ge:$(GOMODHASH)
GINKGO_CMD ?= ginkgo -v --nodes=2 teamcity

all:
	docker build -t $(IMAGE_TAG) .
	docker run --rm -v $(PWD):/ginkgo-experiments -e GINKGO_REPORTERS="$(GINKGO_REPORTERS)" $(IMAGE_TAG) $(GINKGO_CMD)

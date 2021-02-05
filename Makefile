COMMIT := $(shell git rev-parse HEAD)
IMAGE_TAG ?= ge:$(COMMIT)
GINKGO_CMD ?= ginkgo -v --nodes=2 teamcity

all:
	docker build -t $(IMAGE_TAG) .
	docker run -v $(PWD):/ginkgo-experiments -e GINKGO_REPORTERS="$(GINKGO_REPORTERS)" $(IMAGE_TAG) $(GINKGO_CMD)

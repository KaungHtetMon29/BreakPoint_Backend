GO := /home/linuxbrew/.linuxbrew/bin/go
ENDPOINT_CFGS:= endpoint_cfgs/public/ping1/ping1.yaml endpoint_cfgs/public/ping/ping.yaml
OAPI_CFG:= ping-oapi.yaml

generate_sequential:=$(foreach cfg,$(ENDPOINT_CFGS), \
	$(GO) tool oapi-codegen -config $(cfg) $(OAPI_CFG); \
)

.PHONY: run-dev
run-dev:
	$(GO) run server.go

.PHONY: build
build:
ifneq ($(strip $(FILENAME)),)
	${GO} build -o $(FILENAME) server.go
else
	${GO} build server.go
endif

.PHONY: oapi-generate
oapi-generate:
	$(generate_sequential)

override ABSOLUTE_MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
override ABSOLUTE_PATH := $(patsubst %/,%,$(dir $(ABSOLUTE_MAKEFILE)))
override REL_PROJECT_PATH := $(subst $(PROJECT_ABS_PATH)/,,$(ABSOLUTE_PATH))

# saas services
SAAS_SERVICE_PROTO_FILES := $(shell find ./$(REL_PROJECT_PATH) -name "*.proto")
.PHONY: protoc-saas-conf-protobuf
# protoc :-->: generate saas-backend conf api protobuf
protoc-saas-conf-protobuf:
	@echo "# generate ping-service conf api protobuf"
	$(call protoc_protobuf,$(SAAS_SERVICE_PROTO_FILES))


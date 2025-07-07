override ABSOLUTE_MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
override ABSOLUTE_PATH := $(patsubst %/,%,$(dir $(ABSOLUTE_MAKEFILE)))
override REL_PROJECT_PATH := $(subst $(PROJECT_ABS_PATH)/,,$(ABSOLUTE_PATH))

SAAS_BACKEND_V1_API_PROTO := $(shell find ./$(REL_PROJECT_PATH) -name "*.proto")
SAAS_BACKEND_V1_INTERNAL_PROTO := "app/saas-backend/internal/conf/config.conf.proto"
SAAS_BACKEND_V1_PROTO_FILES := ""
ifneq ($(SAAS_BACKEND_V1_INTERNAL_PROTO), "")
	SAAS_BACKEND_V1_PROTO_FILES=$(SAAS_BACKEND_V1_API_PROTO) $(SAAS_BACKEND_V1_INTERNAL_PROTO)
else
	SAAS_BACKEND_V1_PROTO_FILES=$(SAAS_BACKEND_V1_API_PROTO)
endif
.PHONY: protoc-backend-v1-protobuf
# protoc :-->: generate saas backend protobuf
protoc-backend-v1-protobuf:
	@echo "# generate testdata service v1 protobuf"
	$(call protoc_protobuf,$(SAAS_BACKEND_V1_PROTO_FILES))

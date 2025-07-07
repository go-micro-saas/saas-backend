override ABSOLUTE_MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
override ABSOLUTE_PATH := $(patsubst %/,%,$(dir $(ABSOLUTE_MAKEFILE)))
override REL_PROJECT_PATH := $(subst $(PROJECT_ABS_PATH)/,,$(ABSOLUTE_PATH))

SAAS_BACKEND_API_PROTO := $(shell find ./$(REL_PROJECT_PATH) -name "*.proto")
SAAS_BACKEND_INTERNAL_PROTO := "app/saas-backend/internal/conf/config.conf.proto"
SAAS_BACKEND_PROTO_FILES := ""
ifneq ($(SAAS_BACKEND_INTERNAL_PROTO), "")
	SAAS_BACKEND_PROTO_FILES=$(SAAS_BACKEND_API_PROTO) $(SAAS_BACKEND_INTERNAL_PROTO)
else
	SAAS_BACKEND_PROTO_FILES=$(SAAS_BACKEND_API_PROTO)
endif
.PHONY: protoc-backend-protobuf
# protoc :-->: generate saas backend protobuf
protoc-backend-protobuf:
	@echo "# generate testdata service protobuf"
	$(call protoc_protobuf,$(SAAS_BACKEND_PROTO_FILES))

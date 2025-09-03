MAKEFILE_URL?=https://gitlab.com/quanata/projects/backend/be-makefile.git
ifeq '$(MAKEFILE_DIR)' ''
	MAKEFILE_BASE_DIR?=.makefile
	LATEST_EXISTING_TAG?=$(shell [ -d "$(MAKEFILE_BASE_DIR)" ] && cd $(MAKEFILE_BASE_DIR) && ls -d */ 2>/dev/null | sort -t "." -k1,1n -k2,2n -k3,3n | tail -n1 | sed 's/\/$$//')
	MAKEFILE_TAG?=$(LATEST_EXISTING_TAG)
	MAKEFILE_REF?=$(shell [ "$(MAKEFILE_TAG)" ] && echo "$(MAKEFILE_TAG)" || git ls-remote --tags "$(MAKEFILE_URL)" 2>/dev/null | cut -d/ -f3 | sort -rV | head -1 | sed -E 's/^([v0-9.]*)[\^\{\}]*$$/\1/')
	MAKEFILE_DIR?=$(MAKEFILE_BASE_DIR)/$(MAKEFILE_REF)
else
	MAKEFILE_REF?=$(shell cd $(MAKEFILE_DIR) && git describe --tags --always)
endif
include $(shell [ -d $(MAKEFILE_DIR) ] || git clone -b "$(MAKEFILE_REF)" --depth=1 --no-tags "$(MAKEFILE_URL)" $(MAKEFILE_DIR) 2>/dev/null; echo $(MAKEFILE_DIR)/base.mk)

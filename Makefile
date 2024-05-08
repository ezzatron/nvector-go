CHANGELOG_TAG_URL_PREFIX := https://github.com/ezzatron/nvector-go/releases/tag/

-include .makefiles/Makefile
-include .makefiles/pkg/go/v1/Makefile
-include .makefiles/pkg/changelog/v1/Makefile

.makefiles/%:
	@curl -sfL https://makefiles.dev/v1 | bash /dev/stdin "$@"

ifeq (migrations,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "good"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

migrations:
	goose -dir db/migrations create $(RUN_ARGS) sql
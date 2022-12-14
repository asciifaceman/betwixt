#########################
###      DEFS         ###
#########################

# Don't ask
SHELL=/bin/bash -e -o pipefail

#########################
###      OUTPUT       ###
#########################

COLOR = \
  use Term::ANSIColor; \
  printf("    %s %s\n", colored(["BOLD $$ARGV[0]"], "[$$ARGV[1]]"), join(" ", @ARGV[2..$$\#ARGV]));

COLOR_SECTION = \
  use Term::ANSIColor; \
  printf("\n  %s\n\n", colored(["BOLD GREEN"], join(" ", @ARGV)));

COLOR_INDENT = \
  use Term::ANSIColor; use Text::Wrap; \
  $$Text::Wrap::columns=128; $$Text::Wrap::separator="!!"; \
  $$INDENT = (length($$ARGV[1]) + 6 + 1); \
  @l = split(/!!/, wrap("", "", join(" ", @ARGV[2..$$\#ARGV]))); \
  printf("    %s %s\n", colored(["BOLD $$ARGV[0]"], "[$$ARGV[1]]"), shift(@l)); \
  for(@l) { printf("%s%s\n", " "x$$INDENT, $$_) };

HELP_FUNC = \
    %help; \
    while(<>) { \
        if(/^([a-z0-9_-]+):.*\#\#(?:@(\w+))?\s(.*)$$/) { \
            push(@{$$help{$$2}}, [$$1, $$3]); \
        } \
    }; \
    print "usage: VARS=VALUES make [target]\n"; \
    for ( sort keys %help ) { \
        print "$$_:\n"; \
        printf("  %-20s %s\n", $$_->[0], $$_->[1]) for @{$$help{$$_}}; \
        print "\n"; \
    }

RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[01;33m
NC=\033[0m

#########################
###      TARGETS      ###
#########################

.PHONY: help banner

.DEFAULT_GOAL := help

help: banner				##@miscellaneous Show this help
	@perl -e '$(HELP_FUNC)' $(MAKEFILE_LIST)

banner:						##@miscellaneous Display your swag banner
	@echo -e ""
	@echo -e "$(YELLOW)     ____  ___    _   ___   ____________  $(NC)"
	@echo -e "$(YELLOW)    / __ )/   |  / | / / | / / ____/ __ \ $(NC)"
	@echo -e "$(YELLOW)   / __  / /| | /  |/ /  |/ / __/ / /_/ / $(NC)"
	@echo -e "$(YELLOW)  / /_/ / ___ |/ /|  / /|  / /___/ _, _/  $(NC)"
	@echo -e "$(YELLOW) /_____/_/  |_/_/ |_/_/ |_/_____/_/ |_|   $(NC)"
	@echo -e ""


##### WORKFLOW #####

.PHONY: test clean build deploy-local cycle-local

clean:
	go clean
	rm -rf target/

build:
	go build -o target/

deploy-local:
	@echo -e "Installing to /usr/local/bin"
	sudo mv target/betwixt /usr/local/bin/betwixt

cycle-local: clean build deploy-local
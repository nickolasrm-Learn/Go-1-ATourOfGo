## ----------------------
## Commands for exercises
## ----------------------

help:  ## Show this message
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

install:  ## Install exercises dependencies
	@go mod download

list:  ## List all exercises
	@for file in $$(ls exercises); do echo $$(echo $$file | cut -d . -f 1); done

run:  ## Run a go file inside exercises
	@## Usage 'make run number=2'
	@if [ ! -z $(number) ]; then \
		file=$$(make list | grep -E "^$(number)-"); \
		if [ -z $$file ]; then \
			echo "Exercise not found" && exit 1; \
		fi; \
		go run exercises/$$file.go; \
	fi

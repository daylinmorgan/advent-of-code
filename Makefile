check-aoc-cookie:  ## ensures $AOC_SESSION_COOKIE env var is set
	@ test $${AOC_SESSION_COOKIE?env var not set}

skeleton: ## make skeleton main(_test).go files, optional: $DAY and $YEAR
	@ if [[ -n $$DAY && -n $$YEAR ]]; then \
		go run scripts/cmd/skeleton/main.go -day $(DAY) -year $(YEAR) ; \
	elif [[ -n $$DAY ]]; then \
		go run scripts/cmd/skeleton/main.go -day $(DAY); \
	else \
		go run scripts/cmd/skeleton/main.go; \
	fi

input: check-aoc-cookie ## get input, requires $AOC_SESSION_COOKIE, optional: $DAY and $YEAR
	@ if [[ -n $$DAY && -n $$YEAR ]]; then \
		go run scripts/cmd/input/main.go -day $(DAY) -year $(YEAR) -cookie $(AOC_SESSION_COOKIE); \
	elif [[ -n $$DAY ]]; then \
		go run scripts/cmd/input/main.go -day $(DAY) -cookie $(AOC_SESSION_COOKIE); \
	else \
		go run scripts/cmd/input/main.go -cookie $(AOC_SESSION_COOKIE); \
	fi

prompt: check-aoc-cookie ## get prompt, requires $AOC_SESSION_COOKIE, optional: $DAY and $YEAR
	@ if [[ -n $$DAY && -n $$YEAR ]]; then \
		go run scripts/cmd/prompt/main.go -day $(DAY) -year $(YEAR) -cookie $(AOC_SESSION_COOKIE); \
	elif [[ -n $$DAY ]]; then \
		go run scripts/cmd/prompt/main.go -day $(DAY) -cookie $(AOC_SESSION_COOKIE); \
	else \
		go run scripts/cmd/prompt/main.go -cookie $(AOC_SESSION_COOKIE); \
	fi

h help: ## Show this help
	@ printf '\033[35mUsage\033[0m: make <target>\n\nAvailable targets:\n'
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-17s\033[0m - %s\n", $$1, $$2}'

.PHONY: check-aoc-cookie skeleton input prompt h help

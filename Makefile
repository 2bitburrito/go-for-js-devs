# Makefile for running different examples in talk


img-go:
	@echo "__Go Implementation__"
	@echo " "
	@cd code/go/src/image-concurrency/ && go run .

img-js:
	@echo "__JS Implementation__"
	@echo " "
	@cd code/ts/src/image-processing/ && node resizing.ts

scraper:
	@echo "__Go Implementation__"
	@echo " "
	@cd code/go/src/web-scraper/ && go run .

	@echo " "
	@echo " "
	@echo "__JS Implementation__"
	@echo " "
	@cd code/ts/src/web-scraper/ && node webscraper-concurrent.ts

errors:
	@echo "__Go Error Checks__"
	@echo " "
	@cd code/go/src/errors/ && go run .

# .PHONY: img scraper errors

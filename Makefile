wifi:
	# the executable
	go build -o $@ -ldflags "-s -w" -tags osusergo,netgo

module.tar.gz: wifi
	# the bundled module
	rm -f $@
	tar czf $@ $^

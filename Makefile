wifi: *.go */*.go go.*
	# the executable
	go build -o $@ -ldflags "-s -w" -tags osusergo,netgo
	file $@

module.tar.gz: wifi first_run.sh meta.json
	# the bundled module
	rm -f $@
	tar czf $@ $^

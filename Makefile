PREFIX=hvkvp
DESCRIBE=$(git describe --tags)

TARGETS=$(addprefix $(PREFIX)-, centos7 fedora26)

build: $(TARGETS)

$(PREFIX)-%: build/Dockerfile.%
	mkdir -p out
	docker rmi -f $@ >/dev/null  2>&1 || true
	docker rm -f $@-extract > /dev/null 2>&1 || true
	echo "Building binaries for $@"
	docker build -t $@ -f $< .
	docker create --name $@-extract $@ sh
	docker cp $@-extract:/workspace/bin/hvkvp ./out/$@
	docker rm $@-extract || true
	#docker rmi $@ || true

clean:
	rm -f ./$(PREFIX)-*

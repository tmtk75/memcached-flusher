memcached-flusher: *.go
	go build

version=0.1.0

release: ./memcached-flusher
	rm -f pkg/*_amd64 pkg/*.exe
	ghr -u tmtk75 v$(version) pkg

compress:
	gzip -fk pkg/*_amd64
	zip pkg/memcached-flusher_windows_amd64.zip pkg/memcached-flusher_windows_amd64.exe

build:
	gox \
	  -os "linux darwin windows" \
	  -arch "amd64" \
	  -output "pkg/memcached-flusher_{{.OS}}_{{.Arch}}" \
	  .

clean:
	rm -f memcached-flusher

distclean: clean
	rm -rf pkg

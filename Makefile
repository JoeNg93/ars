install:
	pushd cmd/ars && go install; popd
build-mac:
	env GOOS=darwin go build -ldflags="-s -w" -o bin/osx/ars cmd/ars/main.go 
	pushd bin/osx && tar -czvf ars.tar.gz ars && rm -f ars; popd
build-linux:
	env GOOS=linux go build -ldflags="-s -w" -o bin/linux/ars cmd/ars/main.go
	pushd bin/linux && tar -czvf ars.tar.gz ars && rm -f ars; popd

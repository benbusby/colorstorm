debug:
	zig build
	rm -f ./colorstorm
	ln -s ./zig-out/bin/colorstorm .

test:
	@for f in $(shell ls -d -1 "${PWD}/src/"*.*); do \
		echo "Test: $${f}" && zig test $${f}; \
	done

release:
	zig build -Drelease-safe=true

release-all:
	mkdir -p out

	zig build -Drelease-safe=true -Dtarget=x86_64-windows
	mv ./zig-out/bin/colorstorm.exe ./out/
	zip -j ./out/colorstorm-windows-x86.zip ./out/colorstorm.exe

	zig build -Drelease-safe=true -Dtarget=x86_64-linux
	mv ./zig-out/bin/colorstorm ./out/
	zip -j ./out/colorstorm-linux-x86.zip ./out/colorstorm

	zig build -Drelease-safe=true -Dtarget=arm-linux
	mv ./zig-out/bin/colorstorm ./out/
	zip -j ./out/colorstorm-linux-arm.zip ./out/colorstorm

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

	# Windows x86
	zig build -Drelease-safe=true -Dtarget=x86_64-windows
	mv ./zig-out/bin/colorstorm.exe ./out/
	zip -j ./out/colorstorm-windows-x86.zip ./out/colorstorm.exe

	# Windows ARM
	zig build -Drelease-safe=true -Dtarget=aarch64-windows-gnu
	mv ./zig-out/bin/colorstorm.exe ./out/
	zip -j ./out/colorstorm-windows-aarch64.zip ./out/colorstorm.exe

	# Linux x86
	zig build -Drelease-safe=true -Dtarget=x86_64-linux
	mv ./zig-out/bin/colorstorm ./out/
	zip -j ./out/colorstorm-linux-x86.zip ./out/colorstorm

	# Linux ARM
	zig build -Drelease-safe=true -Dtarget=arm-linux
	mv ./zig-out/bin/colorstorm ./out/
	zip -j ./out/colorstorm-linux-arm.zip ./out/colorstorm

	# macOS x86
	zig build -Drelease-safe=true -Dtarget=x86_64-macos-gnu
	mv ./zig-out/bin/colorstorm ./out/
	zip -j ./out/colorstorm-macos-x86.zip ./out/colorstorm

	# macOS ARM
	zig build -Drelease-safe=true -Dtarget=aarch64-macos-gnu
	mv ./zig-out/bin/colorstorm ./out/
	zip -j ./out/colorstorm-macos-arm.zip ./out/colorstorm

	rm ./out/colorstorm
	rm ./out/colorstorm.exe

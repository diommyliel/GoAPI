# Need the ; to run https://stackoverflow.com/questions/73908737/permission-denied-when-running-go-from-makefile
build_api:
	echo "build"
	go build -o ./build/counter-player ; 
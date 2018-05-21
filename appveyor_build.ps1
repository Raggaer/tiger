If (Test-Path "buildOutput") {
	Remove-Item "buildOutput" -recurse
}

$date = (Get-Date).AddDays(-1).ToString('MM-dd-yyyy_HH:mm:ss')
$version = git rev-parse HEAD

echo "$date - $version"

$getDepCommand = 'go get github.com/golang/dep/cmd/dep'

echo "Versioning dep tool downloaded"

iex $getDepCommand

echo "Populating vendor directory"

$depCommand = 'dep ensure'

iex $depCommand

echo "Building for Windows amd64"
$env:GOOS = "windows"
$env:GOARCH = "amd64"

$winCommand = 'go build -o buildOutput\tiger_win_amd64.exe'

iex $winCommand

echo "Building for Linux amd64"
$env:GOOS = "linux"
$env:GOARCH = "amd64"

$linuxCommand = 'go build -o buildOutput\tiger_linux_amd64'

iex $linuxCommand

$linuxCommand = 'go build -o buildOutput\tiger_linux_arm64'

echo "Building for Linux arm64"
$env:GOOS = "linux"
$env:GOARCH = "arm64"

iex $linuxCommand
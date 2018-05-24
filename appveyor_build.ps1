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

$linuxCommand = 'go build -o buildOutput\tiger_linux_amd64.exe'

iex $linuxCommand

$linuxCommand = 'go build -o buildOutput\tiger_linux_arm64.exe'

echo "Building for Linux arm64"
$env:GOOS = "linux"
$env:GOARCH = "arm64"

iex $linuxCommand

echo "Creating template directories"

Copy-Item template buildOutput\data\template -recurse
Copy-Item config.toml.sample buildOutput\data\config.toml.sample

echo "Compressing data directories"

$files = Get-ChildItem -Path "buildOutput\data\*"

Compress-Archive -Path $files -CompressionLevel Optimal -DestinationPath buildOutput\release.zip

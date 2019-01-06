If (Test-Path "buildOutput") {
	Remove-Item "buildOutput" -recurse
}

$env:GO111MODULE = "on"
$date = (Get-Date).AddDays(-1).ToString('MM-dd-yyyy_HH:mm:ss')
$version = git rev-parse HEAD

echo "$date - $version"
echo "Building for Windows amd64"

$env:GOOS = "windows"
$env:GOARCH = "amd64"

$winCommand = 'go build -o buildOutput\tiger_win_amd64.exe -ldflags "-X github.com/raggaer/tiger/app/controllers.ApplicationVersion=$version -X github.com/raggaer/tiger/app/controllers.BuildDate=$date"'

iex $winCommand

echo "Building for Linux amd64"
$env:GOOS = "linux"
$env:GOARCH = "amd64"

$linuxCommand = 'go build -o buildOutput\tiger_linux_amd64 -ldflags "-X github.com/raggaer/tiger/app/controllers.ApplicationVersion=$version -X github.com/raggaer/tiger/app/controllers.BuildDate=$date"'

iex $linuxCommand

$linuxCommand = 'go build -o buildOutput\tiger_linux_arm64 -ldflags "-X github.com/raggaer/tiger/app/controllers.ApplicationVersion=$version -X github.com/raggaer/tiger/app/controllers.BuildDate=$date"'

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

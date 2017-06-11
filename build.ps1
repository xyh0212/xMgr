# chcp 65001 
function xPause()
{
    Write-Host "回车结束暂停"
    $s = Read-Host
}

$strCurPath = Split-Path -Parent $MyInvocation.MyCommand.Definition
$strCurDirName  = Split-Path -leaf $strCurPath

taskkill /f /im $strCurDirName".exe"
$bExist = Test-Path bin/$strCurDirName".exe"
if ($bExist -eq 1)
{
    Remove-Item bin/$strCurDirName".exe"
}  

set OLDGOPATH=$strCurPath
$env:GOPATH=$strCurPath
echo $env:GOPATH
gofmt -w src
#go build -ldflags "-X main.buildstamp `date '+%Y-%m-%d_%I:%M:%S'` -X main.githash `git rev-parse HEAD`" src/$strCurDirName/main.go
go install $strCurDirName
if ($LASTEXITCODE -ne 0)
{
    xPause
}
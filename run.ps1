# chcp 65001 

$strCurPath = Split-Path -Parent $MyInvocation.MyCommand.Definition
$strCurDirName  = Split-Path -leaf $strCurPath

&./build.ps1

cd bin
&./$strCurDirName

Write-Host "press any key"
$s = Read-Host

cd $strCurPath

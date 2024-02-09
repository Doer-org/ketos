$KetosUrl = "https://github.com/Doer-org/ketos/releases/latest/download/ketos_Windows_x86_64.zip"

$DownloadPath = "$env:USERPROFILE\Downloads\ketos_Windows_x86_64.zip"

Invoke-WebRequest -Uri $KetosUrl -OutFile $DownloadPath

$ExtractPath = "$env:USERPROFILE\Ketos"

If (-Not (Test-Path $ExtractPath)) {
    New-Item -ItemType Directory -Force -Path $ExtractPath
}

Expand-Archive -LiteralPath $DownloadPath -DestinationPath $ExtractPath

Write-Host "Ketos has been installed to $ExtractPath"

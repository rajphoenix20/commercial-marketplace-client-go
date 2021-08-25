$current_dir=(Get-Location).Path

$generated_code_dir="${current_dir}/../sdk/"
$meter_path="${generated_code_dir}/meter"
$saas_path="${generated_code_dir}/saas"
$packageName="commercial-marketplace-client-go"
$sdk_path="../sdk/"
$sdk_version='1.0.0'
$autorest_generator='@autorest/go@4.0.0-preview.26'
$saas = "saas"
$metering = "metering"

# Powershell up through 7.03 continues to have an issue with calling
# Remove-Item -Recurse
# on folders that live in a OneDrive folder. The following code
# is functionally equivalent to Remove-Item and does not suffer
# from the bug.
function deleteFolder($folderName){
    if ($false -eq (Test-Path $folderName -PathType Container)){
        return
    }

    $directories = Get-ChildItem $folderName -Directory

    Write-Host "Deleting folder $folderName"
    foreach($dir in $directories){
        deleteFolder($dir.FullName)
    }

    $files = Get-ChildItem $folderName -File

    foreach ($file in $files){
        Remove-Item $file.FullName
    }

    Remove-Item $folderName
}

function updatePackage($folder, $packageName){
    $packageNameValue = Join-Path -Path $current_dir -ChildPath $folder -Resolve
    $packageNameString = "package " + $packageName
    $finalPackageNameString = "package commercialmarketplace"
    $files = Get-ChildItem $packageNameValue -File

    foreach ($file in $files){
        $allText = Get-Content -Path $file.FullName -Raw
        $newText = $allText.Replace($packageNameValue, $packageName)
        #$newText = $newText.Replace("<version>", $version)
        $newText = $newText.Replace($packageNameString, $finalPackageNameString)
        Set-Content -Path $file.FullName -Value $newText
    }
}

function updateTelemetryInfoConstant(){
    $constantsFile = "./temp/saas/constants.go"
    $allText = Get-Content -Path $constantsFile -Raw
    $newText = $allText.Replace("azsdk-go-saas/", "azsdk-go-commercialmarketplace/")
    Set-Content -Path $constantsFile -Value $newText
}

function removeTelemetryInfoConstant(){
    $constantsFile = "./temp/metering/constants.go"
    $allText = Get-Content -Path $constantsFile -Raw
    $newText = $allText.Replace("const telemetryInfo", "const telemetryInfo_unused")
    Set-Content -Path $constantsFile -Value $newText
}

function removePopulateUnpopulateFunctions() {
    $modelsFile = "./temp/metering/models.go"
    $allText = Get-Content -Path $modelsFile -Raw
    $newText = $allText.Replace("func populate", "func populate_unused")
    $newText = $newText.Replace("func unpopulate", "func unpopulate_unused")
    Set-Content -Path $modelsFile -Value $newText
}

function updateScopes() {
    $connectionFile = "./temp/metering/connection.go"
    $allText = Get-Content -Path $connectionFile -Raw
    $newText = $allText.Replace("user_impersonation", "20e940b3-4c77-4b0b-9a53-9e16a1b010a7/.default")
    Set-Content -Path $connectionFile -Value $newText
}

if (Test-Path -Path $generated_code_dir -PathType Container)
{
    Write-Host "Cleaning out previously generated files"
    if ($true -eq (Test-Path -Path $meter_path -PathType Container))
    {
        Write-Host "Clearing out $meter_path"
        deleteFolder($meter_path)
    }
    if ($true -eq (Test-Path -Path $saas_path -PathType Container))
    {
        Write-Host "Clearing out $saas_path"
        deleteFolder($saas_path)
    }
}

if (-Not (Test-Path -Path $sdk_path -PathType Container))
{
    mkdir $sdk_path
}

deleteFolder(".\temp")

mkdir .\temp\saas
autorest --go `
    saas.md `
    --add-credentials `
    --clear-output-folder=true `
    --public-clients=true `
    --license-header=MICROSOFT_MIT_NO_VERSION `
    --namespace=microsoft.marketplace.saas `
    --package-version=$sdk_version `
    --generate-metadata=true `
    --output-folder=.\temp\saas `
    --credential-scope=20e940b3-4c77-4b0b-9a53-9e16a1b010a7/.default `
    --use=$autorest_generator `
    --module-version=$sdk_version `
    --openapi-type="data-plane" `
    --export-clients

updatePackage ".\temp\saas" "saas"
# The telemetryInfo constant is declared in the constants.go file for both SaaS and Metering.
# We want the value to be more "universal", so we change the value of the constant here.
updateTelemetryInfoConstant

# Rename files when the generated filename for both metering and saas is reused but the contents are unique.
mv ./temp/saas/models.go ./temp/saas/models_saas.go
mv ./temp/saas/response_types.go ./temp/saas/response_types_saas.go
mv ./temp/saas/constants.go ./temp/saas/constants_saas.go

# Move the generated files to the "right" location.
cp -r ./temp/saas/* ../sdk/commercialmarketplace -Force

mkdir .\temp\metering
autorest `
    --go metering.md `
    --add-credentials `
    --clear-output-folder=true `
    --public-clients=true `
    --license-header=MICROSOFT_MIT_NO_VERSION `
    --namespace=microsoft.marketplace.metering `
    --package-version=$sdk_version `
    --generate-metadata=true `
    --output-folder=.\temp\metering `
    --credential-scope=20e940b3-4c77-4b0b-9a53-9e16a1b010a7/.default `
    --use=$autorest_generator `
    --module-version=$sdk_version `
    --openapi-type="data-plane" `
    --export-clients

updatePackage ".\temp\metering" "metering"

# At this time, the scopes clause isn't being read properly from the credential-scope
# switch. So, force it.
updateScopes

# The telemetryInfo constant is declared in the constants.go file for both SaaS and Metering.
# We only want one version, so rename the one in Metering.
removeTelemetryInfoConstant
# The populate and unpopluate functions are declared in the models.go file for both SaaS and Metering.
# We only want one version, so rename the one in Metering.
removePopulateUnpopulateFunctions

# Rename files when the generated filename for both metering and saas is reused but the contents are unique.
mv ./temp/metering/models.go ./temp/metering/models_metering.go
mv ./temp/metering/response_types.go ./temp/metering/response_types_metering.go
mv ./temp/metering/constants.go ./temp/metering/constants_metering.go

# Move the generated files to the "right" location.
cp -r ./temp/metering/* ../sdk/commercialmarketplace -Force

# Get rid of the temp folder. Failure to do so may cause the golang engine to
# find and use these files, depending on how this source is called.
deleteFolder("./temp")

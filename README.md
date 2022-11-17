# Current status

This project is deprecated and not in active development.

# Project

This library makes use of some "in-flight updates" to the [Azure SDK for Go](https://github.com/Azure/azure-sdk-for-go). 
The library relies on a specific library. This dependence will change in the future. The following information 
will allow you to get things working against the [SaaS](https://docs.microsoft.com/azure/marketplace/partner-center-portal/pc-saas-fulfillment-api-v2) 
and [Metering](https://docs.microsoft.com/azure/marketplace/marketplace-metering-service-apis) APIs.

With golang installed, run the following commands for your project in order to make sure your go.mod file is correct:

``` 
go mod init main
go get -d github.com/Azure/azure-sdk-for-go/sdk/azcore@v0.16.2
go mod edit -require github.com/microsoft/commercial-marketplace-client-go/sdk/commercialmarketplace@latest
go mod tidy
```

Once that is done, you can run some code to see things in motion. For example, you can create a file called main.go
and put the following contents:

```go
package main

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/microsoft/commercial-marketplace-client-go/sdk/commercialmarketplace"
	"os"
	"time"
)

var (
	clientCredentials, clientCredentialsErr = azidentity.NewClientSecretCredential(
		os.Getenv("AAD_TENANT_ID"),
		os.Getenv("AAD_APP_CLIENT_ID"),
		os.Getenv("AAD_APP_CLIENT_SECRET"),
		nil)
)

func main() {
	fmt.Print("Hello, World!")

	var connection = commercialmarketplace.NewDefaultConnection(clientCredentials, nil)

	client := commercialmarketplace.NewFulfillmentOperationsClient(connection)
	var subscriptionsLister = client.ListSubscriptions(nil)
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Minute)
	defer cancel()

	var hasNextPage = subscriptionsLister.NextPage(ctx)
	if !hasNextPage {
		fmt.Print("No pages to retrieve. Check credentials and make sure this ID has access to 1+ existing SaaS offers")
	}
	var response = subscriptionsLister.PageResponse()
	var subscriptions = response.Subscriptions
	if len(subscriptions) == 0 {
		fmt.Print("Expected 1 or more subscriptions to be returned.")
	} else {
		fmt.Print(*(subscriptions[0].Name))
	}
}
```

Then, using an identity you have already setup and published in a SaaS application registration in 
[Partner Center Commercial Marketplace](https://partner.microsoft.com/en-us/dashboard/commercial-marketplace/overview),
set three environment variables (copying the Azure Active Directory tenant ID, Azure Active Directory application ID,
and the secret for that identity). The environment variables are:
* AAD_TENANT_ID: Azure Active Directory tenant ID
* AAD_APP_CLIENT_ID: Azure Active Directory application ID
* AAD_APP_CLIENT_SECRET: Secret associated with AAD_APP_CLIENT_ID

You can also use a certificate to authenticate with these APIs. That is not demonstrated here. 

## Contributing

This project welcomes contributions and suggestions.  Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.opensource.microsoft.com.

When you submit a pull request, a CLA bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., status check, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

## Trademarks

This project may contain trademarks or logos for projects, products, or services. Authorized use of Microsoft 
trademarks or logos is subject to and must follow 
[Microsoft's Trademark & Brand Guidelines](https://www.microsoft.com/en-us/legal/intellectualproperty/trademarks/usage/general).
Use of Microsoft trademarks or logos in modified versions of this project must not cause confusion or imply Microsoft sponsorship.
Any use of third-party trademarks or logos are subject to those third-party's policies.

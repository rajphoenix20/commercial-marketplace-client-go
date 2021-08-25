package commercialmarketplace

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"os"
	"testing"
	"time"
)

var (
	clientCredentials, clientCredentialsErr = azidentity.NewClientSecretCredential(
		os.Getenv("AAD_TENANT_ID"),
		os.Getenv("AAD_APP_CLIENT_ID"),
		os.Getenv("AAD_APP_CLIENT_SECRET"),
		nil)
)

func TestFulfillmentOperationsClient_ListSubscriptions(t *testing.T) {

	var connection = NewDefaultConnection(clientCredentials, nil)

	client := NewFulfillmentOperationsClient(connection)
	var subscriptionsLister = client.ListSubscriptions(nil)
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Minute)
	defer cancel()

	var hasNextPage = subscriptionsLister.NextPage(ctx)
	if !hasNextPage {
		t.Error("No pages to retrieve. Check credentials and make sure this ID has access to 1+ existing SaaS subscriptions")
	}
	var response = subscriptionsLister.PageResponse()
	var subscriptions = response.Subscriptions
	if len(subscriptions) == 0 {
		t.Error("Expected 1 or more subscriptions to be returned.")
	}
}

/*
func TestFulfillmentOperationsClient_ActivateSubscription(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		ctx            context.Context
		subscriptionId string
		body           SubscriberPlan
		options        *FulfillmentOperationsActivateSubscriptionOptions
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			got, err := client.ActivateSubscription(tt.args.ctx, tt.args.subscriptionId, tt.args.body, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("ActivateSubscription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ActivateSubscription() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentOperationsClient_DeleteSubscription(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		ctx            context.Context
		subscriptionId string
		options        *FulfillmentOperationsDeleteSubscriptionOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    FulfillmentOperationsDeleteSubscriptionResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			got, err := client.DeleteSubscription(tt.args.ctx, tt.args.subscriptionId, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteSubscription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteSubscription() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentOperationsClient_GetSubscription(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		ctx            context.Context
		subscriptionId string
		options        *FulfillmentOperationsGetSubscriptionOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    SubscriptionResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			got, err := client.GetSubscription(tt.args.ctx, tt.args.subscriptionId, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSubscription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubscription() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentOperationsClient_ListAvailablePlans(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		ctx            context.Context
		subscriptionId string
		options        *FulfillmentOperationsListAvailablePlansOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    SubscriptionPlansResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			got, err := client.ListAvailablePlans(tt.args.ctx, tt.args.subscriptionId, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAvailablePlans() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAvailablePlans() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentOperationsClient_Resolve(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		ctx                 context.Context
		xMSMarketplaceToken string
		options             *FulfillmentOperationsResolveOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    ResolvedSubscriptionResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			got, err := client.Resolve(tt.args.ctx, tt.args.xMSMarketplaceToken, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Resolve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Resolve() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentOperationsClient_UpdateSubscription(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		ctx            context.Context
		subscriptionId string
		body           SubscriberPlan
		options        *FulfillmentOperationsUpdateSubscriptionOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    FulfillmentOperationsUpdateSubscriptionResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			got, err := client.UpdateSubscription(tt.args.ctx, tt.args.subscriptionId, tt.args.body, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateSubscription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateSubscription() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentOperationsClient_activateSubscriptionCreateRequest(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		ctx            context.Context
		subscriptionId string
		body           SubscriberPlan
		options        *FulfillmentOperationsActivateSubscriptionOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *azcore.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			got, err := client.activateSubscriptionCreateRequest(tt.args.ctx, tt.args.subscriptionId, tt.args.body, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("activateSubscriptionCreateRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("activateSubscriptionCreateRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentOperationsClient_activateSubscriptionHandleError(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		resp *azcore.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			if err := client.activateSubscriptionHandleError(tt.args.resp); (err != nil) != tt.wantErr {
				t.Errorf("activateSubscriptionHandleError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFulfillmentOperationsClient_deleteSubscriptionCreateRequest(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		ctx            context.Context
		subscriptionId string
		options        *FulfillmentOperationsDeleteSubscriptionOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *azcore.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			got, err := client.deleteSubscriptionCreateRequest(tt.args.ctx, tt.args.subscriptionId, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("deleteSubscriptionCreateRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteSubscriptionCreateRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentOperationsClient_deleteSubscriptionHandleError(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		resp *azcore.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			if err := client.deleteSubscriptionHandleError(tt.args.resp); (err != nil) != tt.wantErr {
				t.Errorf("deleteSubscriptionHandleError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFulfillmentOperationsClient_deleteSubscriptionHandleResponse(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		resp *azcore.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    FulfillmentOperationsDeleteSubscriptionResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			got, err := client.deleteSubscriptionHandleResponse(tt.args.resp)
			if (err != nil) != tt.wantErr {
				t.Errorf("deleteSubscriptionHandleResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteSubscriptionHandleResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentOperationsClient_getSubscriptionCreateRequest(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		ctx            context.Context
		subscriptionId string
		options        *FulfillmentOperationsGetSubscriptionOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *azcore.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			got, err := client.getSubscriptionCreateRequest(tt.args.ctx, tt.args.subscriptionId, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSubscriptionCreateRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSubscriptionCreateRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentOperationsClient_getSubscriptionHandleError(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		resp *azcore.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			if err := client.getSubscriptionHandleError(tt.args.resp); (err != nil) != tt.wantErr {
				t.Errorf("getSubscriptionHandleError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFulfillmentOperationsClient_getSubscriptionHandleResponse(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		resp *azcore.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    SubscriptionResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			got, err := client.getSubscriptionHandleResponse(tt.args.resp)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSubscriptionHandleResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSubscriptionHandleResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentOperationsClient_listAvailablePlansCreateRequest(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		ctx            context.Context
		subscriptionId string
		options        *FulfillmentOperationsListAvailablePlansOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *azcore.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			got, err := client.listAvailablePlansCreateRequest(tt.args.ctx, tt.args.subscriptionId, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("listAvailablePlansCreateRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listAvailablePlansCreateRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentOperationsClient_listAvailablePlansHandleError(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		resp *azcore.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			if err := client.listAvailablePlansHandleError(tt.args.resp); (err != nil) != tt.wantErr {
				t.Errorf("listAvailablePlansHandleError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFulfillmentOperationsClient_listAvailablePlansHandleResponse(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		resp *azcore.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    SubscriptionPlansResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			got, err := client.listAvailablePlansHandleResponse(tt.args.resp)
			if (err != nil) != tt.wantErr {
				t.Errorf("listAvailablePlansHandleResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listAvailablePlansHandleResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentOperationsClient_listSubscriptionsCreateRequest(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		ctx     context.Context
		options *FulfillmentOperationsListSubscriptionsOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *azcore.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			got, err := client.listSubscriptionsCreateRequest(tt.args.ctx, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("listSubscriptionsCreateRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listSubscriptionsCreateRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentOperationsClient_listSubscriptionsHandleError(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		resp *azcore.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			if err := client.listSubscriptionsHandleError(tt.args.resp); (err != nil) != tt.wantErr {
				t.Errorf("listSubscriptionsHandleError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFulfillmentOperationsClient_listSubscriptionsHandleResponse(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		resp *azcore.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    SubscriptionsResponseResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			got, err := client.listSubscriptionsHandleResponse(tt.args.resp)
			if (err != nil) != tt.wantErr {
				t.Errorf("listSubscriptionsHandleResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listSubscriptionsHandleResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentOperationsClient_resolveCreateRequest(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		ctx                 context.Context
		xMSMarketplaceToken string
		options             *FulfillmentOperationsResolveOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *azcore.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			got, err := client.resolveCreateRequest(tt.args.ctx, tt.args.xMSMarketplaceToken, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("resolveCreateRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resolveCreateRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentOperationsClient_resolveHandleError(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		resp *azcore.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			if err := client.resolveHandleError(tt.args.resp); (err != nil) != tt.wantErr {
				t.Errorf("resolveHandleError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFulfillmentOperationsClient_resolveHandleResponse(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		resp *azcore.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    ResolvedSubscriptionResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			got, err := client.resolveHandleResponse(tt.args.resp)
			if (err != nil) != tt.wantErr {
				t.Errorf("resolveHandleResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resolveHandleResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentOperationsClient_updateSubscriptionCreateRequest(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		ctx            context.Context
		subscriptionId string
		body           SubscriberPlan
		options        *FulfillmentOperationsUpdateSubscriptionOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *azcore.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			got, err := client.updateSubscriptionCreateRequest(tt.args.ctx, tt.args.subscriptionId, tt.args.body, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("updateSubscriptionCreateRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateSubscriptionCreateRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFulfillmentOperationsClient_updateSubscriptionHandleError(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		resp *azcore.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			if err := client.updateSubscriptionHandleError(tt.args.resp); (err != nil) != tt.wantErr {
				t.Errorf("updateSubscriptionHandleError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFulfillmentOperationsClient_updateSubscriptionHandleResponse(t *testing.T) {
	type fields struct {
		con *Connection
	}
	type args struct {
		resp *azcore.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    FulfillmentOperationsUpdateSubscriptionResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &FulfillmentOperationsClient{
				con: tt.fields.con,
			}
			got, err := client.updateSubscriptionHandleResponse(tt.args.resp)
			if (err != nil) != tt.wantErr {
				t.Errorf("updateSubscriptionHandleResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateSubscriptionHandleResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFulfillmentOperationsClient(t *testing.T) {
	type args struct {
		con *Connection
	}
	tests := []struct {
		name string
		args args
		want *FulfillmentOperationsClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFulfillmentOperationsClient(tt.args.con); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFulfillmentOperationsClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/

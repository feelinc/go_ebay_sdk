
=================
unofficial eBay Go SDK
=================
Created in 2018. Need a lot of improvements and adding more resources support.

### INSTALLATION
```go get github.com/feelinc/go_ebay_sdk```

# Features!

  - Shopping
  -- GetMultipleItems
  -- GetSingleItem
  - Trading
  -- AddMemberMessageAAQToPartner
  -- AddMemberMessageRTQ
  -- CompleteSale
  -- GetFeedback
  -- GetItem
  -- GetOrderTransactions
  -- GetUser

### Example

    import "github.com/feelinc/go_ebay_sdk/trading"
    
    api := trading.NewConnection(trading.SetDevID("the-dev-id"),
	    trading.SetAppID("the-app-id"),
	    trading.SetCertID("the-cert-id"),
	    trading.SetToken("the-token"),
	    trading.SetSiteID(ebaysdk.SiteIds["US"]))
	    
	request := trading.NewGetUser("the-item-id", "the-user-id", false)
	
	response, err := api.Execute(request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)


### Todos
 - Unit and Integration test
 - Improve existing codes
 - Implement remaining Telesign resources

### FOUND BUGS
Please open a issue (please check if similar issue exist reported here, just comment). We will consider to fix or close without fixing it.

### IMPROVING
Thank you for your help improving it. Please fork and create push request.

License
----
[MIT license](http://opensource.org/licenses/MIT).


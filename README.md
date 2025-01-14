# Go Commercelayer SDK

This projects provides an API client to connect with the Commercelayer apis. It is automatically generated from the 
[openapi definition](https://data.commercelayer.app/schemas/openapi-no-ref.json) so it should be feature complete.

# Documentation
-  [api](./api/README.md)

# Usage
```
clientId := <client-id>
clientSecret := <client-secret>
apiEndpoint := <url>/api
authEndpoint := <ur>/oauth/token

ctx := context.Background()

credentials := clientcredentials.Config{
    ClientID:     clientId,
    ClientSecret: clientSecret,
    TokenURL:     authEndpoint,
    Scopes:       []string{},
}

httpClient := credentials.Client(ctx)

commercelayerClient := api.NewAPIClient(&api.Configuration{
    HTTPClient: httpClient,
    Debug:      true,
    Servers: []api.ServerConfiguration{
        {URL: apiEndpoint},
    },
})

data, resp, err := commercelayerClient.AddressesApi.GETAddresses(ctx).Execute()
```


# Development
- Make sure you have [nodejs and npm](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm) installed 
- Run `make` to regenerate the whole SDK. This will also install the dependencies required to generate the SDK through npm.

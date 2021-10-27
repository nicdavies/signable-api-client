# Signable GO API Client
A Go API Client for working with the [Signable API](https://developer.signable.co.uk/).

## Install
```
go get github.com/nicdavies/signable-api-client
```

## Features
// todo

## Documentation

### Usage
```go
// create a new api client
client := client.NewClient("apikey")

// get all contacts
res, err := client.GetContacts(context.TODO(), &types.ListContactOptions{
    Offset: 0,
    Limit: 10,
})

// get a contact
res, err := client.GetContact(context.TODO(), &types.ListContactOptions{
    Id: 1000,
})

// create a contact
res, err := client.CreateContact(context.TODO(), &types.CreateContactOptions{
	Name: "Hello World",
	Email: "hi@example.org",
})

// update a contact
res, err := client.CreateContact(context.TODO(), 1000, &types.UpdateContactOptions{
    Name: "Hello World",
    Email: "hi@example.org",
})

// delete a contact
res, err := client.CreateContact(context.TODO(), &types.DeleteContactOptions{
    Id: 1000,
})
```
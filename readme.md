# Vetable Market appication

## Introduction
This is a simple distributed application for vegetable sales system written using Go language and Remote Procedure Calls (RPC). 

This application has two components
 1. server --> (can be found in server/main.go)
 2. client --> (can be found in client/client.go)


The Server has following capabilities

1. Query the file and output names & details of all available vegetables.
2. Output the price per kg of a given vegetable.
3. Output the available amount of kg for a given vegetable.
4. Add new vegetable to the file with price per kg and amount of kg.
5. Update the price or available amount of a given vegetable.

Accordingly, clients can use server functions to do the following tasks. Client is CLI application which run these tasks dynamically.
1. Receive a list of all available vegetables and display.
2. Receive a name list of all available vegetables and display.
3. Get all the details of a given vegetable and display.
4. Get the price per kg of a given vegetable and display.
5. Get the available amount of kg of a given vegetable and display.
6. Send a new vegetable name to the server to be added to the server file.
7. Send new price & available amount for a given vegetable to be updated in the server file.
8. Send new price for a given vegetable to be updated in the server file.
9. Send new available amount for a given vegetable to be updated in the server file.



## Run

### Run the server

Execute following command

```
cd server
go run main.go
```

### Run the client 

First go to the client folder by executing following command
```
cd client
```

1. Send a new vegetable name to the server to be added to the server file.
```
go run client.go add {veg Name} {veg price per kg} {veg amount in kg}
```
    eg: go run client.go add carrot 10 100

2. Receive a list of all available vegetables and display.
```
go run client.go get
```

3. Receive a name list of all available vegetables and display.
```
go run client.go get names
```
4. Get all the details of a given vegetable and display.
```
go run client.go get {vegtableName}
```
    eg: go run client.go get carrot

5. Get the price per kg of a given vegetable and display.

```
go run client.go get price {vegtableName}
```
    eg: go run client.go get price carrot

6. Get the available amount of kg of a given vegetable and display.

```
go run client.go get quantity {vegtableName}
```
    eg: go run client.go get quantity carrot

7. Send new price & available amount for a given vegetable to be updated in the server file.
```
go run client.go update {vegtableName} {veg price per kg} {veg amount in kg}
```
    eg: go run client.go update carrot 12 110

8. Send new price for a given vegetable to be updated in the server file.
```
go run client.go update {vegtableName} {veg price per kg}
```
    eg: go run client.go update price carrot 15 

9. Send available amount for a given vegetable to be updated in the server file.
```
go run client.go update {vegtableName} veg amount in kg}
```
    eg: go run client.go update quantity carrot 125 


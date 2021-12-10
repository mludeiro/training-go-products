# Microservices training Go

The goal of this project is to create 3 microservices, one for clients, one for invoices and one for products in Golang. Each microservice will be in it's own folder with the corresponding name.

## Products microservice

This endpoint will manage products and will interact with other microservices. It will expose the endpoint in port **5002**
The endpoints to create for this service are:

### Get all products

**GET**

**/products**

*Response:*
```
[
	{
		"Id": 1,
		"Name": "Computer"
	},
	...
]
```

### Get product by id

**GET**

**/products/{id}**

*Response:*
```
{
	"Id": 1,
	"Name": "Computer"
}
```

### Create product

**POST**

**/products**

*Request data*
```
{
	"Name": "Computer"
}
```

*Response:*
```
{
	"Id": 1,
	"Name": "Computer"
}
```

*Special considerations*

The id of the product will be autoincremental determined by the endpoint or storage

### Update product

**PUT**

**/products/{id}**

*Request data*
```
{
	"Name": "Computer"
}
```

*Response:*
```
{
	"Id": 1,
	"Name": "Computer"
}
```

### Delete product

**DELETE**

**/products/{id}**

*Special considerations*

This endpoint will use the invoice endpoint (GET /invoices/product/{id}) to check if the product has been used in any invoice and can't be deleted if it has been used
		
## Tips and special considerations
- Use GORM [example usage](https://github.com/mludeiro/go-micro)
- Use in memory database [example usage](https://github.com/mludeiro/go-micro)

## Next steps
- Unit tests
- Docker usage
- Test push
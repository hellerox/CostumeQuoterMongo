# Costume Quoter API (using MongoDB)

API for my test application for Artist City

## Install requirements

dep ensure

## Run

go run main.go

## API

### **getMaterials**

Returns all Materials from collection Materials

- **URL**

  /material

- **Method:**

  `GET`

### **createMaterial**

Create new material based on JSON request

- **URL**

  /material

- **Method:**

  `POST`

### **getMaterialByObjectId**

Returns one Material by ObjectId

- **URL**

  /material/<id>

- **Method:**

  `GET`


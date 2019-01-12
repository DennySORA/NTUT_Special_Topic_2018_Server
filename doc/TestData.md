# Test Data

## Query

### Login

    query Login{
    LogIn(ID:"abcde@gmail.com",Password:"123456789"){
        Status{
        StatusCode
        Description
        }
        GetTimes
        AccountID
        AccountToken
    }
    }

### CheckAccount

    query CheckAccount{
    CheckAccountHas(ID:"abcde@gmail.com"){
        Status{
        StatusCode
        Description
        }
        Has
    }
    }

### GetCarID

    query GetCarID {
    GetCarID(
    ID:"abcde@gmail.com",
    Token:""
    ){
        Status{
        StatusCode
        Description
        }
        CarID
        CarName
        RefreshTime
        CreateTime
    }
    }

### GetTemporarilyToken

    query GetTemporarilyToken {
    GetTemporarilyToken(
        ID:"abcde@gmail.com",
    Token:""
    ){
        Status{
        StatusCode
        Description
        }
        Token
        GetTimes
    }
    
    }

### GetUser

    query GetUser {
    GetUser(
        ID:"abcde@gmail.com",
        Token:""
    ){
        Status{
        StatusCode
        Description
        }
        Car{
        CarID
        CarName
            CreateTime
        RefreshTime
        }
        Profile{
        Name
        Gender
        Phone{
            Country
            Number
        }
        }
        Accesse{
        Certification
        PermitTime
        Level
        Permit_log{
            Level
            Times
            Authority
        }
        }
        SiginHistory{
        Times
        Types
        Device
        UseToken
        }
        LogoutHistory{
        Times
        Types
        Device
        UseToken
        }
    }
    }


---

## Mutation

### AddCarID

    mutation AddCarID {
    AddCarID(
        Token:{Token:""},
        InputCarNews:{
        ID:"abcde@gmail.com",
        CarID:"41d5sa4f5sd",
        CarName:"我愛車",
        TemporarilyToken:"694100"
        }
    ){
    Status{
        StatusCode
        Description
        }
        ID
        CarID
        Token
    }
    }

### Create

    mutation Create{
    CreateAccount (
        AccountIDPW: {
        Account:"abcde@gmail.com",
        Password:"123456789"
        },
        User: {
        Name: "abcde",
        Gender: 1,
        Country: "+886",
        Number:"123456789"
        }){
        Status{
        StatusCode
        Description
        }
        ID
    }
    }
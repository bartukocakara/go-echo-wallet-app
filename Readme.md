# Wallet App Backend Application
| No. | Topic                                                                   |
| --- | ----------------------------------------------------------------------- |
| 1   | [**Purpose**](#Purpose) |
| 2   | [**Packages**](#Packages) |
| 3   | [**Server**](#Server) |
| 4   | [**Docker**](#Docker) |
| 5   | [**Endpoints**](#Endpoints) |


## <a href="https://github.com/bartukocakara/roof-stacks-case/blob/master/Technical.md"> --> For DIAGRAM</a>
### Semantic Versioning 2.0.0 https://semver.org/lang/tr/

### Purpose
- Users can create wallet, make transaction(deposit, withdraw) and report the transactions 

### Packages
- POSTGRES - Database https://github.com/jmoiron/sqlx (https://github.com/lib/pq)
- ECHO - Framework https://echo.labstack.com/guide/
- CACHE - Caching https://github.com/patrickmn/go-cache

### Server
- Application will run on internal server
### Docker
```
docker compose up -d --build
docker compose up
```
### !!!!!!! After Create table please run currencies sql !!!!!!!!!!!!!!!!!!!

### Endpoints
- BASE URL = http://127.0.0.1:8080
- PREFIX = api
- VERSION = v1
#### FULL URL = http://127.0.0.1:8080/api/v1/
| Uri  | Description |  Methods | Params | Header |
| :------:|  :-----------:| :-----------:| :-----------:| :-----------:|
| /auth/register   | Register User  | POST | email, password | - |
| /auth/login   | Login User  | POST | email, password | - |
| /wallets   | Wallet List  | GET | - | Bearer {TOKEN} |
| /wallets   | Wallet Create  | POST | title,currency_id,balance,limit,amount | Bearer {TOKEN} |
| /transactions/wallets/{{walletId}}   | Transaction Withdraw  | POST | currency_id,action_type('WITHDRAW,DEPOSIT'),amount | Bearer {TOKEN} |
| /transactions   | Transaction List Query  | GET | QP ? from,to,limit | Bearer {TOKEN} |
| /transactions   | Transaction Report | GET | QP ? from,to,limit,reportable=1 | Bearer {TOKEN} |

Response Codes
| CODE  | Message |
| :------:|  :-----------:|
| 200 | OK |
| 201 | Created |
| 400 | Bad Request |
| 401 | Unauthorized |
| 404 | Not Found |
| 500 | Internal Server Error |

Response Example
```
{
    "status": true,
    "message": "Created",
    "errors": null,
     "data": [
        {
            "id": 5,
            "status": true,
            "type": "WITHDRAW",
            "currency_type": 1,
            "to_wallet_id": 1,
            "list_price": 100,
            "created_at": "2022-09-13T01:59:33.90382+03:00",
            "updated_at": "2022-09-13T01:59:33.90382+03:00"
        },
        {
            "id": 6,
            "status": true,
            "type": "WITHDRAW",
            "currency_type": 1,
            "to_wallet_id": 1,
            "list_price": 100,
            "created_at": "2022-09-13T01:59:34.632979+03:00",
            "updated_at": "2022-09-13T01:59:34.632979+03:00"
        }
    ]
  }
}
```

### Author
| Role  | Personal Info | Contact |
| :------:| :-----------:| :-----------:|
| Developer | Bartu Kocakara | kocakarabartu@gmail.com |

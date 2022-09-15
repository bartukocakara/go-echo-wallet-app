# Multiplatform Backend Application
| No. | Topic                                                                   |
| --- | ----------------------------------------------------------------------- |
| 1   | [**Purpose**](#Purpose) |
| 2   | [**Packages**](#Packages) |
| 3   | [**Server**](#Server) |
| 4   | [**Docker**](#Docker) |
| 5   | [**Endpoints**](#Endpoints) |

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
```
### Endpoints
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
  "code": 200,
  "message": "OK",
  "data" : {
    "list" : [
      "id" : 1
    ]
  }
}
```

- BASE URL = http://127.0.0.1:8080
- PREFIX = api
- VERSION = v1
#### FULL URL = http://127.0.0.1:8080/api/v1/
| Endpoints  | Description |  Methods | Params | Header |
| :------:|  :-----------:| :-----------:| :-----------:| :-----------:|
| /auth/register   | Register User  | POST | email, password, name | - |
| /auth/login   | Login User  | POST | email, password | - |

### Author
| Role  | Personal Info | Contact |
| :------:| :-----------:| :-----------:|
| Developer | Bartu Kocakara | kocakarabartu@gmail.com |

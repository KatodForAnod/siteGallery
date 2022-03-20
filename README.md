# siteGallery
## Website for storing images with a recommendation system
Main tasks:
1. authorization/registration system - done
2. loading images - done
4. caching html pages - not done
5. tagging images - not done
6. unit-tests - not done
7. indexing db tables for fast searching tagged images - not done
8. module recently watched - not done
9. module recommendation by users history - not done

## How to run?
### Install docker and run it
### Run command
```
go get github.com/KatodForAnod/siteGallery
```
### Change directory to project siteGallery
### Create file .env
```
POSTGRES_USER=?
POSTGRES_PASSWORD=?
POSTGRES_DB=?
ACCESS_SECRET=?
```
### Create file conf.config
```
{
    "db_config":{
        "user":"",
        "host":"",
        "port": ,
        "password":"",
        "dbname":"",
        "sslmode":""
    },
    "sv_config":{
        "host":"",
        "port":""
    }
}
```
### Run command
```
docker compose up
```

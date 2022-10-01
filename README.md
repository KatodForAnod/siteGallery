# siteGallery
## Website for storing images with a recommendation system
Main tasks:
1. authorization/registration system - done
2. loading images - done
4. caching html pages - not done
5. tagging images - not done
6. unit-tests - done
7. indexing db tables for fast searching tagged images - not done
8. module recently watched - not done
9. module recommendation by users history - not done

## How to run?
1. ### Install docker and run it
2. ### Run command
```
git clone https://github.com/KatodForAnod/siteGallery.git
```
3. ### Change directory to project siteGallery
4. ### Create file .env
```
POSTGRES_USER=?
POSTGRES_PASSWORD=?
POSTGRES_DB=?
ACCESS_SECRET=?
```
5. ### Create file conf.config
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
6. ### Run command
```
docker compose up
```

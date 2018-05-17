# lssue

## Running Locally
```
// update ui
git submodule update --init

// install dependencies
dep ensure

// build app
docker-compose build app

// run database & app
docker-compose up -d db
docker-compose up app
```
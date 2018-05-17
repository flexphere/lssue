# lssue 
![CircleCI](https://circleci.com/gh/flexphere/lssue.svg?style=svg&circle-token=e7d51209c755f338dc7b22ba7132e96ea4eab12f)

## Getting Started
1. create a new OAuth App in github from [here.](https://github.com/settings/applications/new)
   the `Authorization callback URL` should be set as `http://127.0.0.1:8080/oauth2/callback` for default use
   
2. create .env file and set github app credentials
    ```
    $ cp .env.sample .env
    $ vim .env
    ~
    GITHUB_CLIENT_ID=your_github_client_id
    GITHUB_CLIENT_SECRET=your_github_client_secret
    ~
    ```
3. install dependencies and run app
    ```
    docker-compose build
    docker-compose up -d db
    docker-compose up app
    ```
4. open up [http://127.0.0.1:8080/](http://127.0.0.1:8080/) in your browser
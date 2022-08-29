# Commit Smart CBS

Commit Smart CBS is a simple core banking system. It is a part of backend assignment for Commit Smart.

Link for API documentation can be found   [here](https://documenter.getpostman.com/view/7431834/VUxKTUqx).


### Key Features
- Customer Registration
- Customer Login
- Deposit Fund
- Withdraw Fund
- Transfer Fund between Customer Accounts
- List Transactions
- List Accounts

### Pre-requisites
- Go
- PostgreSQL
- Migrate CLI [here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).

### Getting Started

- Clone the repository 
    ```
    git clone git@github.com:mandeepkhatry/commit-smart-core-banking-system.git
    ```
- Run following command to setup database: 
    ```
    sudo -u postgres psql
    create database commit_smart_cbs;
    create user postgres with encrypted password 'admin';
    grant all privileges on database commit_smart_cbs to postgres;
    ```
- Run Migration
    ```
    make migrateup
    ```
- Update dependencies
    ```
    make tidy
    ```
- Run Server
    ```
    make run_server
    ```




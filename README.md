# dante-api

**API user crud management based on [Dante Socks5 (1.4.3)](https://www.inet.no/dante)**

## Project Setup

```shell
git clone https://github.com/trumanwong/dante-api.git
cd dante-api/scripts && docker-compose up -d
```

### Create/Update User

```shell
curl --location --request POST 'your_dante_api_container_ip:2022/api/user' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "name": "dante",
    "password": "dante"
}'
```

### Query User

```shell
curl --location --request GET 'your_dante_api_container_ip:2022/api/user'
```

### Delete User

```shell
curl --location --request DELETE 'your_dante_api_container_ip:2022/api/user' --form 'name="dante"'
```

### Verify

```go
curl https://ifconfig.co --socks5 your_dante_api_container_ip:2020 --proxy-user sockd:sockd
```


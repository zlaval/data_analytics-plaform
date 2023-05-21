docker compose -f platform-compose.yml up

docker compose -f airbyte-compose.yml -f platform-compose.yml  --env-file .env up -d

localhost:8080

airbyte/password
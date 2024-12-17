docker run --name postgres-dev \
  -e POSTGRES_USER=oldaim \
  -e POSTGRES_PASSWORD=oldaim \
  -e POSTGRES_DB=mydb \
  -p 5432:5432 \
  -d postgres:latest
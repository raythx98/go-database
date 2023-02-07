docker stop database-go
docker rm database-go
docker rmi database-go
docker build -t database-go .
docker run \
  --name database-go \
  --rm \
  -p 9282:9282 \
  database-go
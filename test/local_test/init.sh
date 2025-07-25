cd ../../

# First compose down if it exists
docker compose down
# kill any processes that might be using psql port 5432
sudo lsof -ti :5432 | xargs -r sudo kill -9

docker compose build
docker compose up -d

# ... (TODO) add tests and commands to run them

# terminate
docker compose down
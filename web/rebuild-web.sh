if [ "$1" = "build" ]
then
  echo "Rebuilding the web_builder"
  docker build -f web/Dockerfile-build -t web_builder .
fi
rm web/web
docker run -it -v "$(pwd)/web":/app web_builder
docker exec web pkill -f web

trap "pkill -f web" SIGHUP
trap "exit" SIGTERM

while [[ 1 ]]; do
  ./web;
done

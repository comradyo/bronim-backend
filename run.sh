docker_info_str=$(docker ps -a | grep "bronim");
docker_arr=($(echo $docker_info_str | tr " " "\n"));
container_id=${docker_arr[0]};
docker stop $container_id;
docker rm $container_id;
git pull;
docker build -t bronim .;
docker run -p 5000:5000 bronim;

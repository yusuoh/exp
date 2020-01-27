docker build . -f ./Dockerfile.ubuntu_dev -t ubuntu_dev
docker container create -i -t -P -h ubuntu_dev --name ubuntu_dev ubuntu_dev
docker container start ubuntu_dev

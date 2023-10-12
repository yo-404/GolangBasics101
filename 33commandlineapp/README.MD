docker build --target dev . -t go

docker run -it -v ${PWD}:/work go sh

<!--not to future self for my local env -->
docker run -it videoscmd get --help

docker run -it videoscmd get --help
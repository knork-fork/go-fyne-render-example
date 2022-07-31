
# Sample multi-platform Go docker project

Simple Go render example using Fyne library with multi-platform compiler inside docker image.

Thanks to https://github.com/fyne-io/fyne/issues/3052 for supplied example code.

## Running the container

`docker-compose` commands need to be run from inside the docker folder, so enter it first: 

```
cd docker
```

Start container:

```
docker-compose up --build -d
```

Or if already built:

```
docker-compose up -d
```

Stop container:

```
docker-compose down
```


Rebuild container:

```
docker-compose build
```

## Install vendor folder

```
docker/go-mod-vendor
```

## Building

Build for Linux and Windows:

```
docker/build_for_linux
```

```
docker/build_for_windows
```

## Extras

Enter the shell:

```
docker/shell
```

Go fmt & vet

```
docker/go-fmt
```

```
docker/go-vet
```
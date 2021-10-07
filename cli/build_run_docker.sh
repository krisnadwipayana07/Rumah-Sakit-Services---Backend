docker rm
docker build -t rumah_sakit:1.0.0 ../.
docker run -d -it -p 8080:8080 --name=rumah_sakit_container rumah_sakit:1.0.0
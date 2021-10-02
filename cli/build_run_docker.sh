docker rm
docker build -t rumah_sakit:1.0.0 .
docker run -d -it -p 8001:8001 --name=rumah_sakit rumah_sakit
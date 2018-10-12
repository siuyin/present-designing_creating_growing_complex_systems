# Designing, creating and growing complex software systems 
Drawing inspiration from human communities and business
organisation in designing complex software systems.

## If running in a docker container
docker run -it --name present -v godata:/home/siuyin/go -p 3999:3999 siuyin/go:dev
setup .bashrc to have ~/go/bin in PATH or export PATH=~/go/bin:$PATH
present -http 0.0.0.0:3999 -orighost 192.168.99.100 or 127.0.0.1

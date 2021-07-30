FROM ubuntu
# Create app directory
WORKDIR ~/go-cli
# Install app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
# where available (npm@5+)
COPY . .
ENV GOROOT=/usr/local/go
ENV GOPATH=$HOME/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH
RUN apt-get update && \
    apt-get install -y software-properties-common && \
    rm -rf /var/lib/apt/lists/* && \
    add-apt-repository -y ppa:ethereum/ethereum && \
    add-apt-repository -y ppa:ethereum/ethereum-dev && \
    apt-get update && \
    apt-get -y install curl && \
    apt-get install -y  wget && \
    apt-get install -y build-essential && \
    apt-get -y install solc && \
    apt install -y protobuf-compiler && \ 
    curl -sL https://deb.nodesource.com/setup_12.x -o nodesource_setup.sh && \
    chmod +x nodesource_setup.sh && \
    ./nodesource_setup.sh && \
    apt-get install -y nodejs && \
    npm install --silent && \
    mkdir tmp/ && cd tmp/ && wget https://golang.org/dl/go1.16.6.linux-amd64.tar.gz && \
    tar -xvf go1.16.6.linux-amd64.tar.gz && \
    mv -f go ../../../usr/local && pwd && \
    cd ../ && rm -rf tmp && \
    go version && go get -u github.com/ethereum/go-ethereum@v1.10.6 && \
    cd ../../go/pkg/mod/github.com/ethereum/go-ethereum@v1.10.6 && chmod +777 build &&  make && make devtools
RUN go mod tidy && npm run dockerize-build
CMD ["/bin/bash"]

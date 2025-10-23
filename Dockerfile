FROM docker.io/golang:1.25.1-trixie

# install dependencies
RUN apt-get update && \
  apt-get install -y \
  build-essential \
  zsh \
  && rm -rf /var/lib/apt/lists/*

# install just
RUN curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to /usr/local/bin

# setup zsh and oh-my-zsh
RUN git clone --single-branch --depth 1 https://github.com/robbyrussell/oh-my-zsh.git ~/.oh-my-zsh
RUN cp ~/.oh-my-zsh/templates/zshrc.zsh-template ~/.zshrc
RUN chsh -s /bin/zsh

CMD [ "/bin/zsh" ]

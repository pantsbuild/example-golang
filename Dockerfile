FROM docker.io/golang:1.24.9-trixie
# install dependencies
RUN apt-get update && \
  apt-get install -y \
  build-essential \
  zsh \
  && rm -rf /var/lib/apt/lists/*
# install just
# install Pants
COPY get-pants.sh /tmp/get-pants.sh
RUN bash /tmp/get-pants.sh --bin-dir /usr/local/bin && \
  export PATH="/usr/local/bin:$PATH" && \
  git clone --single-branch --depth 1 https://github.com/robbyrussell/oh-my-zsh.git ~/.oh-my-zsh && \
  cp ~/.oh-my-zsh/templates/zshrc.zsh-template ~/.zshrc && \
  chsh -s /bin/zsh
CMD [ "/bin/zsh" ]
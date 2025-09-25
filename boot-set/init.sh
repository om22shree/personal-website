#!/bin/bash

# Personal Setup Script
# Installs: Git, Docker, Go, Python tools (in venv), Minikube, kubectl, Neovim

set -e  # Exit on error

# Update package lists
sudo apt update

# Install Git
sudo apt install -y git
git config --global user.name "om22shree"
git config --global user.email "om22shree@gmail.com"

# Install Docker
sudo apt remove -y docker docker-engine docker.io containerd runc 2>/dev/null || true
sudo apt install -y \
    ca-certificates \
    curl \
    gnupg \
    lsb-release
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt update
sudo apt install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin
sudo usermod -aG docker $USER

# Install Neovim
sudo apt install -y neovim
# Configure Neovim with line numbers
mkdir -p ~/.config/nvim
cat > ~/.config/nvim/init.vim << 'EOF'
set number
set relativenumber
set cursorline
set numberwidth=4
set signcolumn=yes
EOF

# Install Go
GO_VERSION=$(curl -s https://go.dev/VERSION?m=text | head -n1)
wget "https://go.dev/dl/$GO_VERSION.linux-amd64.tar.gz"
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf "$GO_VERSION.linux-amd64.tar.gz"
rm "$GO_VERSION.linux-amd64.tar.gz"
export PATH=$PATH:/usr/local/go/bin
go install golang.org/x/tools/gopls@latest
go install github.com/cweill/gotests/gotests@latest
go install github.com/fatih/gomodifytags@latest
go install github.com/josharian/impl@latest
go install github.com/haya14busa/goplay/cmd/goplay@latest
go install github.com/go-delve/delve/cmd/dlv@latest
go install honnef.co/go/tools/cmd/staticcheck@latest

# Install Python tools in virtual environment
sudo apt install -y python3-venv python3-dev
python3 -m venv ~/.local/share/python-dev-tools-venv
~/.local/share/python-dev-tools-venv/bin/pip install --upgrade pip
~/.local/share/python-dev-tools-venv/bin/pip install \
    python-lsp-server[all] \
    ruff \
    black \
    isort \
    flake8 \
    mypy \
    debugpy \
    pyright

# Add Python venv to PATH
if ! grep -q ".local/share/python-dev-tools-venv/bin" ~/.bashrc; then
    echo 'export PATH=$PATH:$HOME/.local/share/python-dev-tools-venv/bin' >> ~/.bashrc
fi

# Install Minikube and kubectl
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
rm minikube-linux-amd64 kubectl

# Add Go to PATH permanently
if ! grep -q "/usr/local/go/bin" ~/.bashrc; then
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
fi

# Add vi=nvim alias
if ! grep -q "alias vi=nvim" ~/.bashrc; then
    echo 'alias vi=nvim' >> ~/.bashrc
fi

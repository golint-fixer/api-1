# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.box = "dekstroza/kube-overlay-xfs"
  config.vm.network "public_network"

  config.vm.provider "virtualbox" do |vb|
    vb.cpus = 4
    vb.gui = false
    vb.memory = 10240
    vb.name = "k8s"
  end

  config.vm.provision "shell", inline: <<-SHELL
    sudo apt-get update
    sudo apt-get install -y apache2
  SHELL
end

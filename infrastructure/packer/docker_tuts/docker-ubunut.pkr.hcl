/*
The packer {} block contains Packer settings, including specifying a required Packer version
*/
packer {

  /*
    required_plugins block in the Packer block, which specifies all 
    the plugin required by the template to build your image. 
  */
  required_plugins {
    /*
      Each plugin block contains a version and source attribute. 
      Packer will use these attributes to download the appropriate plugin(s).
    */
    docker = {
      version = ">= 0.0.7"
      source = "github.com/hashicorp/docker"
    }
  }
}

/*
The source block configures a specific builder plugin, which is then invoked by a build block.

Source blocks use builders and communicators to define what kind of virtualization to use, 
how to launch the image you want to provision, and how to connect to it. 
*/

source "docker" "ubuntu" {
  image  = "ubuntu:xenial"
  commit = true
}

build {
  name    = "learn-packer"
  sources = [
    "source.docker.ubuntu"
  ]
}

# https://learn.hashicorp.com/tutorials/packer/docker-get-started-build-image?in=packer/docker-get-started
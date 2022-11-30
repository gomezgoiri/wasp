# The name of your project. A project typically maps 1:1 to a VCS repository.
# This name must be unique for your Waypoint server. If you're running in
# local mode, this must be unique to your machine.
project = "isc"

# Labels can be specified for organizational purposes.
labels = { "team" = "isc" }

variable "adminWhitelist" {
    type = list(string)
    // type = string
    // default = dynamic("vault", {
    //     path = "secret/data/nomad/jobs/wasp"
    //     key  = "/data/adminWhitelist"
    // })
}

variable "ghcr" {
    type = object({
        username = string
        password = string
        server_address = string
    })
    // type = string
    // default = dynamic("vault", {
    //     path = "secret/data/nomad/jobs/wasp"
    //     key  = "/data/ghcr"
    // })
}

# An application to deploy.
app "wasp-evm" {
    runner {
        enabled = true
        profile = "nomad"

        data_source "git" {
            url  = "https://github.com/iotaledger/wasp.git"
            ref = "v0.3.8"
        }
    }

    # Build specifies how an application should be deployed. In this case,
    # we'll build using a Dockerfile and keeping it in a local registry.
    build {
        use "docker" {
            disable_entrypoint = true
            buildkit   = true
            dockerfile = "./Dockerfile.noncached"
            build_args = {
                BUILD_TAGS = "rocksdb"
                BUILD_LD_FLAGS = "-X github.com/iotaledger/wasp/packages/wasp.VersionHash=${gitrefhash()}"
                BUILD_TARGET = "..."
            }
        }

        registry {
            use "docker" {
                image = "ghcr.io/luke-thorne/wasp"
                tag = gitrefhash()
                auth = var.ghcr
                // username = var.ghcr.username
                // password = var.ghcr.password
                // encoded_auth = base64encode(jsonencode(var.ghcr))
            }
        }
    }

    # Deploy to Nomad
    deploy {
        use "nomad-jobspec" {
            // Templated to perhaps bring in the artifact from a previous
            // build/registry, entrypoint env vars, etc.
            jobspec = templatefile("${path.app}/wasp.nomad.tpl", { 
                artifact = artifact
                adminWhitelist = jsonencode(var.adminWhitelist)
                auth = var.ghcr
                workspace = workspace.name
            })
        }
    }
}

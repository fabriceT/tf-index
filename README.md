# tf-index

A ridiculous tools that generates an index for a version of opentofu/terraform provider.

Why ?

At least to files are required when mirroring provider:

- index.json
- <version>.json

The <version>.json contains a [h1 hash](https://pkg.go.dev/golang.org/x/mod/sumdb/dirhash#Hash1) of the zip file of the provider. I couldn't find a CLI tool to calculate it, so I wrote a simple wrapper to do that.

---
## Note to myself

The ~/.terraformrc must contains
```hcl
  network_mirror {
    url = "https://localhost/providers"
    include = [ "my/provider" ] # change with the providers you are mirroring.
  }
```

Adjust with your setup.

If you're using

```hcl
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "3.0.2"
    }
  }

```

Terraform will query

```txt
"GET /registry.terraform.io/kreuzwerker/docker/index.json HTTP/2.0" 200 39 "-" "Terraform/1.11.0"
"GET /registry.terraform.io/kreuzwerker/docker/3.0.2.json HTTP/2.0" 200 198 "-" "Terraform/1.11.0"
"HEAD /registry.terraform.io/kreuzwerker/docker/terraform-provider-docker_3.0.2_linux_amd64.zip HTTP/2.0" 200 0 "-" "Terraform/1.11.0"
```

The `index.json`

```json
{
  "versions": {
    "3.0.2": {}
  }
}
```

The `3.0.2.json`

```json
{
  "archives": {
    "linux_amd64": {
      "hashes": [
        "h1:cT2ccWOtlfKYBUE60/v2/4Q6Stk1KYTNnhxSck+VPlU="
      ],
      "url": "terraform-provider-docker_3.0.2_linux_amd64.zip"
    }
  }
}
```

This tools creates the content of the `<version>.json` in a very basic way.

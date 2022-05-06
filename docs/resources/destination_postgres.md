---
page_title: "rudderstack_destination_postgres Resource - terraform-provider-rudderstack"
subcategory: ""
description: |-
  
---

# rudderstack_destination_postgres (Resource)

This resource represents a PostgreSQL warehouse destination. For more information check 
https://www.rudderstack.com/docs/data-warehouse-integrations/postgresql

## Example Usage

```terraform
resource "rudderstack_destination_postgres" "example" {
  name = "my-postgres"

  config {
    host                 = "localhost"
    user                 = "postgres"
    password             = "postgres"
    port                 = "5432"
    namespace            = "example"
    database             = "example"
    ssl_mode             = "disable"
    use_rudder_storage   = false

    # verify_ca {
    #   client_key  = "-----BEGIN RSA PRIVATE KEY-----...-----END CERTIFICATE-----"
    #   client_cert = "-----BEGIN RSA PRIVATE KEY-----...-----END CERTIFICATE-----"
    #   server_ca   = "-----BEGIN RSA PRIVATE KEY-----...-----END CERTIFICATE-----"
    # }

    # s3 {
    #   bucket_name   = ""
    #   access_key_id = ""
    #   access_key    = ""
    # }

    # gcs {
    #   bucket_name = ""
    #   credentials = ""
    # }

    # azure_blob {
    #   container_name = ""
    #   account_name   = ""
    #   account_key    = ""
    # }

    # minio {
    #   bucket_name       = ""
    #   endpoint          = ""
    #   access_key_id     = ""
    #   secret_access_key = ""
    #   use_ssl           = ""
    # }
  }

  #   sync {
  #     frequency = "30"
  #     sync_start_at    = "???"
  #     exclude_start_at = "???"
  #     exclude_end_at   = "???"
  #   }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `config` (Block List, Min: 1, Max: 1) Destination specific configuration. Check the nested block documenation for more information. (see [below for nested schema](#nestedblock--config))
- `name` (String) Human readable name of the destination. The value has to be unique across all destinations.

### Optional

- `enabled` (Boolean) An enabled destination allows data to be sent to it.

### Read-Only

- `created_at` (String) Time when the resource was created, in ISO 8601 format.
- `id` (String) The ID of this resource.
- `updated_at` (String) Time when the resource was last updated, in ISO 8601 format.

<a id="nestedblock--config"></a>
### Nested Schema for `config`

Required:

- `database` (String)
- `host` (String)
- `namespace` (String)
- `password` (String, Sensitive)
- `port` (String)
- `ssl_mode` (String)
- `user` (String)

Optional:

- `azure_blob` (Block List, Max: 1) (see [below for nested schema](#nestedblock--config--azure_blob))
- `gcs` (Block List, Max: 1) (see [below for nested schema](#nestedblock--config--gcs))
- `minio` (Block List, Max: 1) (see [below for nested schema](#nestedblock--config--minio))
- `s3` (Block List, Max: 1) (see [below for nested schema](#nestedblock--config--s3))
- `sync` (Block List, Max: 1) (see [below for nested schema](#nestedblock--config--sync))
- `use_rudder_storage` (Boolean)
- `verify_ca` (Block List, Max: 1) (see [below for nested schema](#nestedblock--config--verify_ca))

<a id="nestedblock--config--azure_blob"></a>
### Nested Schema for `config.azure_blob`

Required:

- `account_key` (String, Sensitive)
- `account_name` (String)
- `container_name` (String)


<a id="nestedblock--config--gcs"></a>
### Nested Schema for `config.gcs`

Required:

- `bucket_name` (String)
- `credentials` (String, Sensitive)


<a id="nestedblock--config--minio"></a>
### Nested Schema for `config.minio`

Required:

- `access_key_id` (String)
- `bucket_name` (String)
- `endpoint` (String)
- `secret_access_key` (String, Sensitive)

Optional:

- `use_ssl` (Boolean)


<a id="nestedblock--config--s3"></a>
### Nested Schema for `config.s3`

Required:

- `access_key` (String, Sensitive)
- `access_key_id` (String)
- `bucket_name` (String)


<a id="nestedblock--config--sync"></a>
### Nested Schema for `config.sync`

Required:

- `frequency` (String)

Optional:

- `exclude_window_end_time` (String)
- `exclude_window_start_time` (String)
- `start_at` (String)


<a id="nestedblock--config--verify_ca"></a>
### Nested Schema for `config.verify_ca`

Required:

- `client_cert` (String, Sensitive)
- `client_key` (String, Sensitive)
- `server_ca` (String, Sensitive)
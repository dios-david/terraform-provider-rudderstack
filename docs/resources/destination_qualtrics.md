---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "rudderstack_destination_qualtrics Resource - terraform-provider-rudderstack"
subcategory: ""
description: |-
  
---

# rudderstack_destination_qualtrics (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `config` (Block List, Min: 1, Max: 1) Destination specific configuration. Check the nested block documenation for more information. (see [below for nested schema](#nestedblock--config))
- `name` (String) Human readable name of the destination. The value has to be unique across all the destinations.

### Optional

- `enabled` (Boolean) An enabled destination allows data to be sent to it.

### Read-Only

- `created_at` (String) Time when the resource was created, in ISO 8601 format.
- `id` (String) The ID of this resource.
- `updated_at` (String) Time when the resource was last updated, in ISO 8601 format.

<a id="nestedblock--config"></a>
### Nested Schema for `config`

Required:

- `brand_id` (String) Enter your Brand ID.
- `project_id` (String, Sensitive) Enter your Project ID.

Optional:

- `enable_generic_page_title` (Boolean) This setting enables Generic Page Title.
- `event_filtering` (Block List, Max: 1) RudderStack lets you determine which events should be allowed to flow through or blocked. (see [below for nested schema](#nestedblock--config--event_filtering))
- `onetrust_cookie_categories` (List of String) Specify the OneTrust category name for mapping the OneTrust consent settings to RudderStack's consent purposes.
- `use_native_sdk` (Block List, Max: 1) Enable this setting to send the events through SDKs. (see [below for nested schema](#nestedblock--config--use_native_sdk))

<a id="nestedblock--config--event_filtering"></a>
### Nested Schema for `config.event_filtering`

Optional:

- `blacklist` (List of String) Enter the event names to be blacklisted.
- `whitelist` (List of String) Enter the event names to be whitelisted.


<a id="nestedblock--config--use_native_sdk"></a>
### Nested Schema for `config.use_native_sdk`

Optional:

- `android` (Boolean)
- `ios` (Boolean)
- `web` (Boolean)


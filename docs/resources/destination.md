# Resource rudderstack_destination
Manages a Rudderstack CDP destination.

<a id="example"></a>
## Example Usage
```
resource "rudderstack_destination" "dst1" {
  name = "tfdestination"
  type = "GA"
  config = {
    object = {
      "trackingID": { str = "UA-908213012-193" },
      "doubleClick": { bool = true },
      "enhancedLinkAttribution": { bool = true },
      "includeSearch": { bool = true },
      "enableServerSideIdentify": { bool = true },
      "serverSideIdentifyEventCategory": { str = "mnd,msdnf" },
      "serverSideIdentifyEventAction": { str = ",mn,m" },
      "disableMd5": { bool = true },
      "anonymizeIp": { bool = true },
      "enhancedEcommerce": { bool = true },
      "nonInteraction": { bool = true },
      "sendUserId": { bool = true },
      "dimensions": {
        objects_list = [
          {
             object = {
               "from": { str = "mas." },
               "to": { str = "3" },
             }
          }
        ]
      },
      "metrics": {
        objects_list = [
          {
             object = {
               "from": { str = "kksad1222" },
               "to": { str = "2" },
             }
          }
        ]
      },
      "contentGroupings": {
        objects_list = [
          {
             object = {
               "from": { str = "lkjdlkjsdf" },
               "to": { str = "lkjlkjsdf" },
             }
          }
        ]
      },
    },
  }
}
```

## Argument Reference

### Required

- **name** (String, Required) Specifies name of the resource.
- **type** (String, Required) Selects category of the CDP destination to be created. Examples include GA(Google Analytics), Webhook, Kissmetrics, etc.  Consult Rudderstack documentation for list of supported source types.
- **config** (Attributes, Required) (Check [this](../guides/config.md) for schema)

### Read-Only

- **id** (String) The ID of this resource.
# Updating the YAML file

* Update **console-api-XXX-swagger.json** file
* Manually modify updated file as described bellow
* Run **make generate**


::: tip
Some manual changes are required due to limitations/errors in the client generation tool.
:::

## ResourceDTO

Modify type, value should be:

```json
"value" : {
"type" : "string"
}
```

## IdentityProviderDTO

Modify type, add identity lookups property:

```json
"identityLookups" : {
            "uniqueItems" : true,
            "type" : "array",
            "items" : {
              "$ref" : "#/components/schemas/IdentityLookupDTO"
            }
          },
```
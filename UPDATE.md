# Updating the YAML file

Some manual changes are required due to limitations/errors in the client generation tool

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
# json-schema-description-filler

Primarily intended to fill in missing `description` fields in BigQuery json schema files. Uses matching `name` fields to search if `description` is filled in another file for the same column.
## Building:
```
go build -o cmd ./  # building binary
```
```
docker build .  # for docker
```
## Usage:
```
./cmd comma,delimited,paths  # Recursively searches all provided paths for json files
```
### As a pre-commit-hook (requires docker running) - add following to .pre-commit-config.yaml:

```
    - repo: https://github.com/defilercompiler/json-schema-description-filler
      rev: 'fill-in-commit-sha'
      hooks:
        - id: json-schema-description-filler-docker | json-schema-description-filler-golang
          args: ['path/to/your/jsons/,another/path/']

```
It is recommended to have a json formatter in place as this tool disregards formatting.
Eg.
### table1.json
```
[
  {
    "description": "",
    "mode": "REQUIRED",
    "name": "patient_id",
    "type": "INTEGER"
  }
]
```
### table2.json
```
[
  {
    "description": "The unique identifier of a patient",
    "mode": "REQUIRED",
    "name": "patient_id",
    "type": "INTEGER"
  }
]
```

### Updated table1.json
```
[
  {
    "description": "The unique identifier of a patient",
    "mode": "REQUIRED",
    "name": "patient_id",
    "type": "INTEGER"
  }
]
```
### table2.json - no changes

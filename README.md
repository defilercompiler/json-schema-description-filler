# json-schema-description-filler

Primarily intended to fill in missing `description` fields in BigQuery json schema files. Uses matching `name` fields to search if `description` is filled in another file for the same column.
## Building:
```
go build -o cmd ./
```
## Usage:
```
./cmd comma,delimited,paths  # Recursively searches all provided paths for json files
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

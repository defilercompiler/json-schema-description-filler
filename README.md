# json-schema-description-filler

Primarily intended to fill in missing `description` fields in BigQuery json schema files. Uses matching `name` fields to search if `description` is filled in another json schema files (another table).
Eg.
```
table1.json
[
  {
    "name": "patient_id",
    "mode": "REQUIRED",
    "type": "INTEGER",
    "description": ""
  },
  {
  more columns ...
  }
]
table2.json
[
  {
    "name": "patient_id",
    "mode": "REQUIRED",
    "type": "INTEGER",
    "description": "The unique identifier of a patient"
  },
  {
  more columns ...
  }
]
```
becomes
```
table1.json
[
  {
    "name": "patient_id",
    "mode": "REQUIRED",
    "type": "INTEGER",
    "description": "The unique identifier of a patient"
  },
  {
  more columns ...
  }
]
table2.json
[
  {
    "name": "patient_id",
    "mode": "REQUIRED",
    "type": "INTEGER",
    "description": "The unique identifier of a patient"
  },
  {
  more columns ...
  }
]
```

# Cloud Storage Api

## spman

```
ssh root@178.128.44.201
```

Run this command after login to ssh

```
spman
```

**Output:**

```
create: Create a new space
list: List all spaces
info: Show details by ID
edit: Edit space name by ID
delete: Delete a space by ID
```



### Create new space

```shell
spman create testspace
```

**Output:**

```
New space entry created.
Space ID: 4
Name: testspace
Access Token: b0c8fbcf-d8ae-428e-8b5f-6e25a7d96695
```

### List All spaces

```
spman list
```

**Output**

```
1. testspace (ID:4)
```

### View Space details (with space id)

```
spman info 4
```

**Output**

```
Space ID: 4
Name: testspace
Access Token: b0c8fbcf-d8ae-428e-8b5f-6e25a7d96695
```

### Edit space name  (with space id)

```
spamn edit 4
```

**Output:**

```
Space ID: 4
Current Name: testspace
New Name: test
Successfully change named to test
```

### Delete space (with space id)

```
spman delete 4
```

**Output**

```
Delete space: 4
Space ID: 4
Current Name: test
Please enter "test" to confirm delete: test
```



# REST API

**Request**

```curl
curl --location --request GET '178.128.44.201:8000'
```

**Response:**

```
Cloud storage api
```



## Upload a file

All necessary folder will be automatically created from given path

**Request**

```shell
curl --location --request POST '178.128.44.201:8000/files' \
--header 'Space-ID: 4' \
--header 'Access-Token: b0c8fbcf-d8ae-428e-8b5f-6e25a7d96695' \
--form 'name=Screencast' \
--form 'file=@/home/princebillygk/2020-10-28 16-25-06.webm' \
--form 'path=videos/prince'
```

**Response**

```json
{
    "data": {
        "name": "Screencast",
        "path": "videos/prince",
        "sizeInBytes": 9397802,
        "filePath": "videos/prince/Screencast.webm",
        "downloadURL": "178.128.44.201:8000/downloads/4/videos/prince/Screencast.webm"
    },
    "status": "success",
    "message": "File Uploaded Successfully"
}
```



## Delete a file

**Request**

```shell
curl --location --request DELETE '178.128.44.201:8000/files?path=videos/prince/Screencast.webm' \
--header 'Space-ID: 4' \
--header 'Access-Token: b0c8fbcf-d8ae-428e-8b5f-6e25a7d96695'
```

**Response**

```json
{
    "data": {
        "path": "videos/prince/Screencast.webm"
    },
    "status": "success",
    "message": "Deleted successfully"
}
```



## Delete a Folder 

**Request**

```shell
curl --location --request DELETE '178.128.44.201:8000/files?path=videos' \
--header 'Space-ID: 4' \
--header 'Access-Token: b0c8fbcf-d8ae-428e-8b5f-6e25a7d96695'
```

```json
{
    "data": {
        "path": "videos"
    },
    "status": "success",
    "message": "Deleted successfully"
}
```



## Error Codes
| Error code | Description                 |
| ---------- | --------------------------- |
| 991        | Request with invalid passed |
| 992        | Invalid or no file passed   |
| 993        | Invalid space id            |




# Assembling API
The API for PaaS Assembling.

Table of Contents

1. [create a project.](#projects)

<a name="projects"></a>

## projects

| Specification | Value |
|-----|-----|
| Resource Path | /projects |
| API Version | 1.0.0 |
| BasePath for the API | http://10.67.147.211:9090/api/v1 |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /projects | [POST](#Create) | create a project. |
| /projects | [GET](#List) | list projects by namespace. |
| /projects/\{project_id\} | [GET](#Get) | get project by id. |
| /projects/\{project_id\} | [DELETE](#Delete) | delete a project. |
| /projects/\{project_id\}/webhook | [POST](#Webhook) | Transport a push event. |
| /projects/\{project_id\}/builds | [GET](#List) | list builds for a project. |
| /projects/\{project_id\}/builds | [POST](#Build) | start a build. |
| /projects/\{project_id\}/builds/\{build_id\} | [GET](#Get) | get a build. |



<a name="Create"></a>

#### API: /projects (POST)


create a project.



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| namespace | form | string | namespace | Yes |
| name | form | string | the name of app project | Yes |
| active | form | string | on / off | Yes |
| type | form | string | software / docker_image | Yes |
| repo_url | form | string | http://code.huawei.com/xxx/yyy | Yes |
| branch_filter | form | string | seprated with comma |  |
| language | form | string | Go1, used if type=software |  |
| environments | form | string | one per line (ex: FOO=bar), used if type=software |  |
| commands | form | string | one per line, used if type=software |  |
| artifacts | form | string | one per line, used if type=software |  |
| image_name | form | string | Docker image name, used if type=docker_image |  |
| image_tag | form | string | Docker image tag, used if type=docker_image |  |
| dockerfile_dir | form | string | the directory which Dockerfile is contained, used if type=docker_image |  |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Project](#code.huawei.com.containerops.assembling.models.Project) | success |
| 400 | object | [Error](#code.huawei.com.containerops.assembling.controllers.Error) | failure |


<a name="List"></a>

#### API: /projects (GET)


list projects by namespace.



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| namespace | query | string | namespace | Yes |
| type | query | string | software / docker_image | Yes |
| page | query | int | page > 0, default 1 |  |
| per_page | query | int | 0 < per_page <= 200, default 20 |  |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | array | [ProjectListResult](#code.huawei.com.containerops.assembling.controllers.ProjectListResult) | success |
| 400 | object | [Error](#code.huawei.com.containerops.assembling.controllers.Error) | failure |


<a name="Get"></a>

#### API: /projects/\{project_id\} (GET)


get project by id.



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| project_id | path | string | project ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Project](#code.huawei.com.containerops.assembling.models.Project) | success |
| 404 | object | [Error](#code.huawei.com.containerops.assembling.controllers.Error) | failure |


<a name="Delete"></a>

#### API: /projects/\{project_id\} (DELETE)


delete a project.



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| project_id | path | string | project ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | string | success |
| 404 | object | [Error](#code.huawei.com.containerops.assembling.controllers.Error) | failure |


<a name="Webhook"></a>

#### API: /projects/\{project_id\}/webhook (POST)


Transport a push event.



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| project_id | path | string | project ID | Yes |
| Ak-Sha | header | string | ak + sha(sk) | Yes |
| other | body | string | event body emitted by github/gitlab | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | string | success |
| 400 | object | string | failure |


<a name="List"></a>

#### API: /projects/\{project_id\}/builds (GET)


list builds for a project.



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| project_id | path | string | project ID | Yes |
| page | query | int | page > 0, default 1 |  |
| per_page | query | int | 0 < per_page <= 200, default 20 |  |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | array | [BuildsListResult](#code.huawei.com.containerops.assembling.controllers.BuildsListResult) | success |
| 400 | object | [Error](#code.huawei.com.containerops.assembling.controllers.Error) | failure |


<a name="Build"></a>

#### API: /projects/\{project_id\}/builds (POST)


start a build.



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| project_id | path | string | project ID | Yes |
| image_name | form | string | Docker image name, use project setting as default |  |
| image_tag | form | string | Docker image tag, use project setting as default |  |
| dockerfile | form | string | The content of Dockferfile |  |
| repo_url | form | string | ex. http://code.huawei.com/xxx/yyy, used if dockerfile not provided |  |
| ref | form | string | git reference, used if dockerfile not provided |  |
| token | form | string | git token, used if dockerfile not provided |  |
| dockerfile_dir | form | string | the directory which Dockerfile is contained, used if dockerfile not provided |  |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Build](#code.huawei.com.containerops.assembling.models.Build) | success |
| 404 | object | [Error](#code.huawei.com.containerops.assembling.controllers.Error) | failure |


<a name="Get"></a>

#### API: /projects/\{project_id\}/builds/\{build_id\} (GET)


get a build.



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| project_id | path | string | project ID | Yes |
| build_id | path | string | build ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Build](#code.huawei.com.containerops.assembling.models.Build) | success |
| 404 | object | [Error](#code.huawei.com.containerops.assembling.controllers.Error) | failure |




### Models

<a name="code.huawei.com.containerops.assembling.controllers.BuildsListResult"></a>

#### BuildsListResult

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| branch | string | Git branch of this build |
| commit | string | Git commit ID |
| commit_message | string | Git commit message |
| commit_timestamp | Time | Git commit timestamp |
| committer | string | Git committer |
| id | string | Build ID |
| state | string | Build State is one of Building/Success/Failed |

<a name="code.huawei.com.containerops.assembling.controllers.Error"></a>

#### Error

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| message | string |  |

<a name="code.huawei.com.containerops.assembling.controllers.ProjectListResult"></a>

#### ProjectListResult

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| id | string | Project ID |
| name | string | Project Name |
| repo_url | string | The binding git repo URL |
| state | string | Project State is one of None/Building/Success/Failed |

<a name="code.huawei.com.containerops.assembling.models.Build"></a>

#### Build

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| branch | string | Git branch of this build |
| commit | string | Git commit ID |
| commit_message | string | Git commit message |
| commit_timestamp | Time | Git commit timestamp |
| committer | string | Git committer |
| finished | Time | Timestamp of build finish |
| id | uint64 | Build ID |
| log | string | Log details of build |
| project_id | string | Project ID of build |
| started | Time | Timestamp of build start |
| state | string | State is one of Building/Success/Failed |

<a name="code.huawei.com.containerops.assembling.models.Project"></a>

#### Project

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| active | bool | Active |
| artifacts | string | Indicate what files are build expected |
| branch_filter | string | Branch filter for push event |
| commands | string | Commands when build |
| created_at | Time | Timestamp of project creation |
| environments | string | Environments setting before bulid |
| id | string | Project ID |
| language | string | Project language |
| name | string | Project Name |
| namespace | string | Namespace |
| repo_url | string | The binding git repo URL |
| type | string | type is one of software/docker_image |



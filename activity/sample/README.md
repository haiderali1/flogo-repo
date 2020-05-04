# sample Utility
Sample activity to understand creation of activity in flogo

## Installation

Navigate to the Flogo app directory [Essential !] and issue the following command

```bash
flogo install github.com/haiderali1/flogo-repo/activity/sample
```

## Schema
Inputs and Outputs:

```json
{
"inputs":[
      {
        "name": "string1",
        "required": true,
        "type": "string"
      },
      {
        "name": "string2",
        "required": true,
        "type": "string"
      },
      {
        "name": "string3",
        "type": "string"
      }
    ],
    "outputs": [
      {
        "name": "output_string",
        "type": "string"
      }  }
```

## Settings
| Setting   | Description      |
|:----------|:-----------------|
| String1   | First String     |
| String2   | Second String    |
| String3   | Third String     |

## Outputs
| Output            | Description                                 |
|:------------------|:--------------------------------------------|
| output_string     | Concatenation of above three strings        |


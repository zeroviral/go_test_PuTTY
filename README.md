# go_test_PuTTY
A sampler project for being introduced to GO, and Crypto Tech Stacks.

To run locally:

You should be within the following folder:
```log
<path_to_project_root/go_test_PuTTY>
```

Run the following from within the root:

```shell script
go run .
```

After this, the server (with default settings should then be running on)
```html
http://localhost:3000
```

## Usage 

`GET / transactions`

```json
{
  "EID": "<your_ethereum_id>"
}
```

**Arguments**

- `"EID":string"` A string value which should be your ethereum ID.
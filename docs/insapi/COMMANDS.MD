1. **feature vrrp**
  1. Message Format: json
  2. Command Type: cli_conf
  3. Request Body:
  ```
  {
    "ins_api": {
      "version": "1.2",
      "type": "cli_conf",
      "chunk": "0",
      "sid": "1",
      "input": "feature vrrp",
      "output_format": "json"
    }
  }
  ```
  
  4. Response Body:
  ```
  {
    "ins_api": {
      "type": "cli_conf",
      "version": "1.2",
      "sid": "eoc",
      "outputs": {
        "output": {
          "msg": "Success",
          "code": "200",
          "body": "For proper functioning of VRRP protocol , LAN_BASE_SERVICES_PKG  license is required\n"
        }
      }
    }
  }
```


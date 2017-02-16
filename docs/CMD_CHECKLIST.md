# Checklist of Nexus Commands for dev purposes.

## API Type:
ins_api

### Tested Hardware:
Cisco Nexus 5548 Chassis ("O2 32X10GE/Modular Universal Platform Supervisor")

### System Version:
7.3(1)N1(1)

### Commands:
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

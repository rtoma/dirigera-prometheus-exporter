# light controller

## Websocket

### deviceConfigurationChanged

```
{
  "id": "00000015-0000-0000-0000-100000000015",
  "time": "2024-07-14T16:10:15.000Z",
  "specversion": "3.154.0",
  "source": "urn:com:ikea:homesmart:iotc:zigbee",
  "type": "deviceConfigurationChanged",
  "data": {
    "id": "f6a7b8c9-d0e1-2345-fabc-456789012345_1",
    "type": "controller",
    "deviceType": "lightController",
    "createdAt": "2024-07-14T18:10:11.000Z",
    "isReachable": true,
    "lastSeen": "2024-07-14T18:10:15.000Z",
    "attributes": {
      "customName": "Controller 1",
      "firmwareVersion": "2.4.11",
      "hardwareVersion": "1",
      "manufacturer": "IKEA of Sweden",
      "model": "Remote Control N2",
      "productCode": "E2001",
      "serialNumber": "AABBCCFFFEEE0044",
      "batteryPercentage": 78,
      "isOn": false,
      "lightLevel": 1,
      "permittingJoin": false,
      "otaPolicy": "autoUpdate",
      "otaProgress": 0,
      "otaScheduleEnd": "00:00",
      "otaScheduleStart": "00:00",
      "otaState": "readyToCheck",
      "otaStatus": "upToDate"
    },
    "capabilities": {
      "canSend": [
        "isOn",
        "lightLevel"
      ],
      "canReceive": [
        "customName"
      ]
    },
    "deviceSet": [],
    "remoteLinks": []
  }
}
```


## Device info

```
  {
    "id": "a7b8c9d0-e1f2-3456-abcd-567890123456_1",
    "type": "controller",
    "deviceType": "lightController",
    "createdAt": "2024-07-14T17:17:41.000Z",
    "isReachable": true,
    "lastSeen": "2024-07-14T15:19:59.000Z",
    "attributes": {
      "customName": "Bedroom",
      "firmwareVersion": "2.4.11",
      "hardwareVersion": "1",
      "manufacturer": "IKEA of Sweden",
      "model": "Remote Control N2",
      "productCode": "E2001",
      "serialNumber": "AABBCCFFFEEE0055",
      "batteryPercentage": 82,
      "isOn": false,
      "lightLevel": 1,
      "permittingJoin": false,
      "otaPolicy": "autoUpdate",
      "otaProgress": 0,
      "otaScheduleEnd": "00:00",
      "otaScheduleStart": "00:00",
      "otaState": "readyToCheck",
      "otaStatus": "upToDate",
      "circadianPresets": []
    },
    "capabilities": {
      "canSend": [
        "isOn",
        "lightLevel"
      ],
      "canReceive": [
        "customName"
      ]
    },
    "room": {
      "id": "c9d0e1f2-a3b4-5678-cdef-789012345678",
      "name": "Bedroom 1",
      "color": "ikea_lilac_no_3",
      "icon": "rooms_bed"
    },
    "deviceSet": [],
    "remoteLinks": [],
    "isHidden": false
  },
  ```
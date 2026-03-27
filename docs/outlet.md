# Outlet

## Websocket

### Turn on/off

```
{"id":"00000012-0000-0000-0000-100000000012","time":"2024-06-27T07:53:11.000Z","specversion":"3.154.0","source":"urn:com:ikea:homesmart:iotc:zigbee","type":"deviceStateChanged","data":{"id":"d4e5f6a7-b8c9-0123-defa-234567890123_1","type":"outlet","deviceType":"outlet","createdAt":"2024-06-23T17:36:29.000Z","isReachable":true,"lastSeen":"2024-06-27T09:53:11.000Z","attributes":{"isOn":true},"remoteLinks":[]}}
```

```
{"id":"00000013-0000-0000-0000-100000000013","time":"2024-06-27T07:53:15.000Z","specversion":"3.154.0","source":"urn:com:ikea:homesmart:iotc:zigbee","type":"deviceStateChanged","data":{"id":"d4e5f6a7-b8c9-0123-defa-234567890123_1","type":"outlet","deviceType":"outlet","createdAt":"2024-06-23T17:36:29.000Z","isReachable":true,"lastSeen":"2024-06-27T09:53:15.000Z","attributes":{"isOn":false},"remoteLinks":[]}}
```

### identity

```
{"id":"00000014-0000-0000-0000-100000000014","time":"2024-06-27T13:27:08.000Z","specversion":"3.154.0","source":"urn:com:ikea:homesmart:iotc:zigbee","type":"deviceStateChanged","data":{"id":"d4e5f6a7-b8c9-0123-defa-234567890123_1","type":"outlet","deviceType":"outlet","createdAt":"2024-06-23T17:36:29.000Z","isReachable":true,"lastSeen":"2024-06-27T15:27:08.000Z","attributes":{"identifyPeriod":10,"identifyStarted":"2024-06-27T15:27:08.000Z"},"remoteLinks":[]}}
```

## Device info

```
{
  "id": "d4e5f6a7-b8c9-0123-defa-234567890123_1",
  "type": "outlet",
  "deviceType": "outlet",
  "createdAt": "2024-06-23T17:36:29.000Z",
  "isReachable": true,
  "lastSeen": "2024-06-27T09:53:15.000Z",
  "attributes": {
    "customName": "Lamp",
    "model": "TRETAKT Smart plug",
    "manufacturer": "IKEA of Sweden",
    "firmwareVersion": "2.4.4",
    "hardwareVersion": "1",
    "serialNumber": "AABBCCFFFEEE0033",
    "productCode": "E2204",
    "isOn": false,
    "startupOnOff": "startPrevious",
    "lightLevel": 100,
    "startUpCurrentLevel": -1,
    "childLock": false,
    "statusLight": true,
    "identifyStarted": "2000-01-01T00:00:00.000Z",
    "identifyPeriod": 0,
    "permittingJoin": false,
    "otaStatus": "upToDate",
    "otaState": "readyToCheck",
    "otaProgress": 0,
    "otaPolicy": "autoUpdate",
    "otaScheduleStart": "00:00",
    "otaScheduleEnd": "00:00"
  },
  "capabilities": {
    "canSend": [],
    "canReceive": [
      "customName",
      "isOn",
      "lightLevel",
      "childLock",
      "statusLight"
    ]
  },
  "room": {
    "id": "b8c9d0e1-f2a3-4567-bcde-678901234567",
    "name": "Living Room",
    "color": "ikea_green_no_65",
    "icon": "rooms_sofa"
  },
  "deviceSet": [],
  "remoteLinks": [],
  "isHidden": false
}
```
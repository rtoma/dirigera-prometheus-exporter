# Air purifier


## Websocket

### filterElapsedTime

```
{"id":"00000009-0000-0000-0000-100000000009","time":"2024-06-27T12:35:45.000Z","specversion":"3.154.0","source":"urn:com:ikea:homesmart:iotc:zigbee","type":"deviceStateChanged","data":{"id":"b2c3d4e5-f6a7-8901-bcde-f12345678901_1","type":"airPurifier","deviceType":"airPurifier","createdAt":"2024-06-23T16:16:23.000Z","isReachable":true,"lastSeen":"2024-06-27T14:35:45.000Z","attributes":{"filterElapsedTime":5219},"remoteLinks":[]}}
```

### identity

```
{"id":"0000000a-0000-0000-0000-10000000000a","time":"2024-06-27T13:21:38.000Z","specversion":"3.154.0","source":"urn:com:ikea:homesmart:iotc:zigbee","type":"deviceStateChanged","data":{"id":"b2c3d4e5-f6a7-8901-bcde-f12345678901_1","type":"airPurifier","deviceType":"airPurifier","createdAt":"2024-06-23T16:16:23.000Z","isReachable":true,"lastSeen":"2024-06-27T15:21:38.000Z","attributes":{"identifyPeriod":10,"identifyStarted":"2024-06-27T15:21:38.000Z"},"remoteLinks":[]}}
```

### motorState

In 'auto' mode:
```
{"id":"0000000b-0000-0000-0000-10000000000b","time":"2024-06-27T07:55:46.000Z","specversion":"3.154.0","source":"urn:com:ikea:homesmart:iotc:zigbee","type":"deviceStateChanged","data":{"id":"b2c3d4e5-f6a7-8901-bcde-f12345678901_1","type":"airPurifier","deviceType":"airPurifier","createdAt":"2024-06-23T16:16:23.000Z","isReachable":true,"lastSeen":"2024-06-27T09:55:46.000Z","attributes":{"fanMode":"auto","fanModeSequence":"lowMediumHighAuto","motorState":10},"remoteLinks":[]}}
```

In 'low' mode:
```
{"id":"0000000c-0000-0000-0000-10000000000c","time":"2024-06-27T07:53:34.000Z","specversion":"3.154.0","source":"urn:com:ikea:homesmart:iotc:zigbee","type":"deviceStateChanged","data":{"id":"b2c3d4e5-f6a7-8901-bcde-f12345678901_1","type":"airPurifier","deviceType":"airPurifier","createdAt":"2024-06-23T16:16:23.000Z","isReachable":true,"lastSeen":"2024-06-27T09:53:34.000Z","attributes":{"fanMode":"low","fanModeSequence":"lowMediumHighAuto","motorState":25},"remoteLinks":[]}}
```

### start of auto mode sequence

```
# enable 'auto' mode
2024/06/27 09:53:46 Received: {"id":"0000000d-0000-0000-0000-10000000000d","time":"2024-06-27T07:53:46.000Z","specversion":"3.154.0","source":"urn:com:ikea:homesmart:iotc:zigbee","type":"deviceStateChanged","data":{"id":"b2c3d4e5-f6a7-8901-bcde-f12345678901_1","type":"airPurifier","deviceType":"airPurifier","createdAt":"2024-06-23T16:16:23.000Z","isReachable":true,"lastSeen":"2024-06-27T09:53:46.000Z","attributes":{"fanMode":"auto","fanModeSequence":"lowMediumHighAuto"},"remoteLinks":[]}}

# start motor at 30
2024/06/27 09:53:46 Received: {"id":"0000000e-0000-0000-0000-10000000000e","time":"2024-06-27T07:53:46.000Z","specversion":"3.154.0","source":"urn:com:ikea:homesmart:iotc:zigbee","type":"deviceStateChanged","data":{"id":"b2c3d4e5-f6a7-8901-bcde-f12345678901_1","type":"airPurifier","deviceType":"airPurifier","createdAt":"2024-06-23T16:16:23.000Z","isReachable":true,"lastSeen":"2024-06-27T09:53:46.000Z","attributes":{"fanMode":"auto","fanModeSequence":"lowMediumHighAuto","motorState":30},"remoteLinks":[]}}

# 30s later, reduce motor to 25
2024/06/27 09:54:16 Received: {"id":"0000000f-0000-0000-0000-10000000000f","time":"2024-06-27T07:54:16.000Z","specversion":"3.154.0","source":"urn:com:ikea:homesmart:iotc:zigbee","type":"deviceStateChanged","data":{"id":"b2c3d4e5-f6a7-8901-bcde-f12345678901_1","type":"airPurifier","deviceType":"airPurifier","createdAt":"2024-06-23T16:16:23.000Z","isReachable":true,"lastSeen":"2024-06-27T09:54:16.000Z","attributes":{"fanMode":"auto","fanModeSequence":"lowMediumHighAuto","motorState":25},"remoteLinks":[]}}

# 30s later, reduce motor to 20
2024/06/27 09:54:46 Received: {"id":"00000010-0000-0000-0000-100000000010","time":"2024-06-27T07:54:46.000Z","specversion":"3.154.0","source":"urn:com:ikea:homesmart:iotc:zigbee","type":"deviceStateChanged","data":{"id":"b2c3d4e5-f6a7-8901-bcde-f12345678901_1","type":"airPurifier","deviceType":"airPurifier","createdAt":"2024-06-23T16:16:23.000Z","isReachable":true,"lastSeen":"2024-06-27T09:54:46.000Z","attributes":{"fanMode":"auto","fanModeSequence":"lowMediumHighAuto","motorState":20},"remoteLinks":[]}}

# 30s later, reduce motor to 15
2024/06/27 09:55:16 Received: {"id":"00000011-0000-0000-0000-100000000011","time":"2024-06-27T07:55:16.000Z","specversion":"3.154.0","source":"urn:com:ikea:homesmart:iotc:zigbee","type":"deviceStateChanged","data":{"id":"b2c3d4e5-f6a7-8901-bcde-f12345678901_1","type":"airPurifier","deviceType":"airPurifier","createdAt":"2024-06-23T16:16:23.000Z","isReachable":true,"lastSeen":"2024-06-27T09:55:16.000Z","attributes":{"fanMode":"auto","fanModeSequence":"lowMediumHighAuto","motorState":15},"remoteLinks":[]}}

# 30s later, reduce motor to 10
2024/06/27 09:55:46 Received: {"id":"0000000b-0000-0000-0000-10000000000b","time":"2024-06-27T07:55:46.000Z","specversion":"3.154.0","source":"urn:com:ikea:homesmart:iotc:zigbee","type":"deviceStateChanged","data":{"id":"b2c3d4e5-f6a7-8901-bcde-f12345678901_1","type":"airPurifier","deviceType":"airPurifier","createdAt":"2024-06-23T16:16:23.000Z","isReachable":true,"lastSeen":"2024-06-27T09:55:46.000Z","attributes":{"fanMode":"auto","fanModeSequence":"lowMediumHighAuto","motorState":10},"remoteLinks":[]}}
```




## Device info

```
{
  "id": "b2c3d4e5-f6a7-8901-bcde-f12345678901_1",
  "type": "airPurifier",
  "deviceType": "airPurifier",
  "createdAt": "2024-06-23T16:16:23.000Z",
  "isReachable": true,
  "lastSeen": "2024-06-27T14:35:45.000Z",
  "attributes": {
    "customName": "Air Purifier 1",
    "model": "STARKVIND Air purifier",
    "manufacturer": "IKEA of Sweden",
    "firmwareVersion": "1.1.001",
    "hardwareVersion": "1",
    "serialNumber": "AABBCCFFFEEE0011",
    "productCode": "E2007",
    "fanMode": "auto",
    "fanModeSequence": "lowMediumHighAuto",
    "motorState": 10,
    "motorRuntime": 4904,
    "filterElapsedTime": 5219,
    "filterAlarmStatus": false,
    "filterLifetime": 259200,
    "childLock": false,
    "statusLight": true,
    "currentPM25": 3,
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
      "fanMode",
      "fanModeSequence",
      "motorState",
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
# Gateway

## Websocket

### OTA

OTA check when visiting the Dirigera "Hub/app version page" in the IKEA app.

#### otaState=checkInProgress

```
2024/06/27 10:10:35 Received: {"id":"00000001-0000-0000-0000-100000000001","time":"2024-06-27T08:10:35.000Z","specversion":"3.154.0","source":"urn:com:ikea:homesmart:iotc:ota","type":"deviceStateChanged","data":{"id":"a1b2c3d4-e5f6-7890-abcd-ef1234567890_1","type":"gateway","deviceType":"gateway","lastSeen":"2024-06-27T08:10:35.000Z","attributes":{"otaState":"checkInProgress"},"remoteLinks":[]}}
```

#### otaState=readyToCheck

```
2024/06/27 10:10:37 Received: {"id":"00000002-0000-0000-0000-100000000002","time":"2024-06-27T08:10:36.000Z","specversion":"3.154.0","source":"urn:com:ikea:homesmart:iotc:ota","type":"deviceStateChanged","data":{"id":"a1b2c3d4-e5f6-7890-abcd-ef1234567890_1","type":"gateway","deviceType":"gateway","lastSeen":"2024-06-27T08:10:36.000Z","attributes":{"otaState":"readyToCheck"},"remoteLinks":[]}}
```

#### otaState=upToDate

```
2024/06/27 10:10:37 Received: {"id":"00000003-0000-0000-0000-100000000003","time":"2024-06-27T08:10:36.000Z","specversion":"3.154.0","source":"urn:com:ikea:homesmart:iotc:ota","type":"deviceStateChanged","data":{"id":"a1b2c3d4-e5f6-7890-abcd-ef1234567890_1","type":"gateway","deviceType":"gateway","lastSeen":"2024-06-27T08:10:36.000Z","attributes":{"otaStatus":"upToDate"},"remoteLinks":[]}}
```

### Reset

```
{"id":"00000004-0000-0000-0000-100000000004","time":"2024-06-28T19:38:20.000Z","specversion":"3.154.0","source":"urn:com:ikea:homesmart:iotc:gatewayd","type":"deviceStateChanged","data":{"id":"a1b2c3d4-e5f6-7890-abcd-ef1234567890_1","type":"gateway","deviceType":"gateway","lastSeen":"2024-06-28T19:38:20.000Z","remoteLinks":[]}}

{"id":"00000005-0000-0000-0000-100000000005","time":"2024-06-28T19:38:51.273Z","specversion":"3.64.0","source":"urn:com:ikea:homesmart:iotc:cloudconnector","type":"deviceStateChanged","data":{"id":"a1b2c3d4-e5f6-7890-abcd-ef1234567890_1","type":"gateway","deviceType":"gateway","lastSeen":"2024-06-28T19:38:51.273Z","attributes":{"backendConnected":false},"remoteLinks":[]}} 

{"id":"00000006-0000-0000-0000-100000000006","time":"2024-06-28T19:39:10.000Z","specversion":"3.154.0","source":"urn:com:ikea:homesmart:iotc:gatewayd","type":"deviceStateChanged","data":{"id":"a1b2c3d4-e5f6-7890-abcd-ef1234567890_1","type":"gateway","deviceType":"gateway","lastSeen":"2024-06-28T19:39:10.000Z","remoteLinks":[]}}

{"id":"00000007-0000-0000-0000-100000000007","time":"2024-06-28T19:39:31.702Z","specversion":"3.64.0","source":"urn:com:ikea:homesmart:iotc:cloudconnector","type":"deviceStateChanged","data":{"id":"a1b2c3d4-e5f6-7890-abcd-ef1234567890_1","type":"gateway","deviceType":"gateway","lastSeen":"2024-06-28T19:39:31.702Z","attributes":{"backendConnected":true},"remoteLinks":[]}} 
```

## Device info

```
{
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890_1",
  "relationId": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "type": "gateway",
  "deviceType": "gateway",
  "createdAt": "2024-03-08T06:47:18.766Z",
  "isReachable": true,
  "lastSeen": "2024-06-27T12:37:00.096Z",
  "attributes": {
    "customName": "My Home",
    "model": "DIRIGERA Hub for smart products",
    "manufacturer": "IKEA of Sweden",
    "firmwareVersion": "2.588.0",
    "hardwareVersion": "P2.5",
    "serialNumber": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "identifyStarted": "2000-01-01T00:00:00.000Z",
    "identifyPeriod": 0,
    "otaStatus": "upToDate",
    "otaState": "readyToCheck",
    "otaProgress": 0,
    "otaPolicy": "autoDownload",
    "otaScheduleStart": "00:00",
    "otaScheduleEnd": "00:00",
    "permittingJoin": false,
    "backendConnected": true,
    "backendConnectionPersistent": true,
    "backendOnboardingComplete": true,
    "backendRegion": "eu-example-1",
    "backendCountryCode": "XX",
    "userConsents": [
      {
        "name": "analytics",
        "value": "disabled"
      },
      {
        "name": "diagnostics",
        "value": "enabled"
      }
    ],
    "logLevel": 3,
    "coredump": false,
    "timezone": "Europe/London",
    "nextSunSet": "2024-06-27T20:04:00.000Z",
    "nextSunRise": "2024-06-28T03:20:00.000Z",
    "homestate": "home",
    "countryCode": "XZ",
    "coordinates": {
      "latitude": 0.0,
      "longitude": 0.0,
      "accuracy": -1
    },
    "isOn": false
  },
  "capabilities": {
    "canSend": [],
    "canReceive": [
      "customName",
      "permittingJoin",
      "userConsents",
      "logLevel",
      "time",
      "timezone",
      "countryCode",
      "coordinates"
    ]
  },
  "deviceSet": [],
  "remoteLinks": []
}
```

### Timeservice

```
{
  "id": "00000008-0000-0000-0000-100000000008",
  "time": "2024-06-28T03:21:00.386Z",
  "specversion": "3.154.0",
  "source": "urn:com:ikea:homesmart:iotc:timeservice",
  "type": "deviceStateChanged",
  "data": {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890_1",
    "type": "gateway",
    "deviceType": "gateway",
    "lastSeen": "2024-06-28T03:21:00.386Z",
    "attributes": {
      "nextSunSet": "2024-06-28T20:04:00.000Z",
      "nextSunRise": "2024-06-29T03:21:00.000Z"
    },
    "remoteLinks": []
  }
}
```
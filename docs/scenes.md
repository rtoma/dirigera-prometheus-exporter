# Scenes

## Websocket

### sceneUpdated

```
{
  "id": "00000016-0000-0000-0000-100000000016",
  "time": "2024-06-27T21:07:00.258Z",
  "specversion": "3.154.0",
  "source": "urn:com:ikea:homesmart:iotc:rulesengine",
  "type": "sceneUpdated",
  "data": {
    "id": "e1f2a3b4-c5d6-7890-efab-901234567890",
    "info": {
      "name": "Lights Off",
      "icon": "scenes_clean_sparkles"
    },
    "type": "userScene",
    "triggers": [
      {
        "id": "00000017-0000-0000-0000-100000000017",
        "type": "app",
        "disabled": false
      },
      {
        "id": "00000018-0000-0000-0000-100000000018",
        "type": "time",
        "triggeredAt": "2024-06-27T21:07:00.257Z",
        "disabled": false,
        "trigger": {
          "days": [
            "Sun",
            "Fri",
            "Mon",
            "Thu",
            "Tue",
            "Wed",
            "Sat"
          ],
          "time": "23:07"
        },
        "nextTriggerAt": "2024-06-27T21:07:00.000Z"
      }
    ],
    "actions": [
      {
        "id": "d4e5f6a7-b8c9-0123-defa-234567890123_1",
        "type": "device",
        "deviceId": "d4e5f6a7-b8c9-0123-defa-234567890123_1",
        "enabled": true,
        "attributes": {
          "isOn": false
        }
      }
    ],
    "commands": [],
    "createdAt": "2024-06-23T21:05:54.979Z",
    "lastCompleted": "2024-06-26T21:07:00.940Z",
    "lastTriggered": "2024-06-26T21:07:00.940Z",
    "undoAllowedDuration": 30
  }
}
```


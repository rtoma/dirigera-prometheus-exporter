# dirigera-prometheus-exporter

A Prometheus exporter for the [IKEA Dirigera](https://www.ikea.com/us/en/p/dirigera-hub-for-smart-products-white-smart-20503252/) smart home hub.

Connects to the Dirigera hub via its local REST API and WebSocket event stream, and exposes device state as Prometheus metrics.

## Getting started

### 1. Obtain an access token

Use the included `cmd/access-token` tool to obtain an access token from your Dirigera hub:

```bash
go run ./cmd/access-token -host <your-hub-ip-or-hostname>
```

Follow the prompt: when instructed, press the action button on the bottom of the Dirigera hub, then press Enter. The access token will be printed to stdout.

### 2. Run the exporter

#### Docker

```bash
docker run \
  -e DIRIGERA_URL=https://<hub-ip>:8443 \
  -e DIRIGERA_ACCESS_TOKEN=<your-token> \
  -p 9090:9090 \
  ghcr.io/rtoma/dirigera-prometheus-exporter:latest
```

#### From source

```bash
DIRIGERA_URL=https://<hub-ip>:8443 \
DIRIGERA_ACCESS_TOKEN=<your-token> \
go run ./cmd/exporter
```

## Configuration

| Environment variable | Required | Default | Description |
|---|---|---|---|
| `DIRIGERA_URL` | yes | — | URL of the Dirigera hub, e.g. `https://192.168.1.x:8443` |
| `DIRIGERA_ACCESS_TOKEN` | yes | — | Access token obtained via `cmd/access-token` |
| `DIRIGERA_INTERVAL` | no | `30` | Polling interval in seconds (5–120) |
| `DIRIGERA_BIND_ADDRESS` | no | `:9090` | Address the HTTP server listens on |

## Metrics

All device metrics carry `device_id`, `device_type`, and `device_name` labels unless noted otherwise.

### All devices

| Metric | Description |
|---|---|
| `ikea_dirigera_device_reachable` | 1 if the device is reachable, 0 otherwise |
| `ikea_dirigera_device_last_seen` | Last seen unix timestamp (seconds) |
| `ikea_dirigera_device_child_lock_status` | Child lock enabled (1) or disabled (0) |
| `ikea_dirigera_device_status_light` | Status light enabled (1) or disabled (0) |

### Air purifier

| Metric | Description |
|---|---|
| `ikea_dirigera_airpurifier_pm25` | Current PM2.5 level |
| `ikea_dirigera_airpurifier_motor_state` | Motor state / fan speed |
| `ikea_dirigera_airpurifier_motor_runtime` | Total motor runtime |
| `ikea_dirigera_airpurifier_filter_elapsed` | Filter elapsed time (minutes) |
| `ikea_dirigera_airpurifier_filter_lifetime` | Filter lifetime (minutes) |
| `ikea_dirigera_airpurifier_filter_alarm_status` | Filter alarm active (1) or not (0) |

### Environment sensor

| Metric | Description |
|---|---|
| `ikea_dirigera_environment_sensor_temperature` | Temperature |
| `ikea_dirigera_environment_sensor_humidity` | Humidity |
| `ikea_dirigera_environment_sensor_pm25` | PM2.5 level |
| `ikea_dirigera_environment_sensor_max_pm25` | Max measured PM2.5 |
| `ikea_dirigera_environment_sensor_min_pm25` | Min measured PM2.5 |
| `ikea_dirigera_environment_sensor_voc_index` | VOC index |

### Light

| Metric | Description |
|---|---|
| `ikea_dirigera_light_is_on` | Light is on (1) or off (0) |
| `ikea_dirigera_light_level` | Brightness level |
| `ikea_dirigera_light_color_temperature` | Color temperature |
| `ikea_dirigera_light_color_temperature_min` | Minimum color temperature |
| `ikea_dirigera_light_color_temperature_max` | Maximum color temperature |

### Light controller

| Metric | Description |
|---|---|
| `ikea_dirigera_light_controller_is_on` | Is on (1) or off (0) |
| `ikea_dirigera_light_controller_light_level` | Light level |
| `ikea_dirigera_light_controller_battery_percentage` | Battery percentage |

### Outlet

| Metric | Description |
|---|---|
| `ikea_dirigera_outlet_is_on` | Outlet is on (1) or off (0) |
| `ikea_dirigera_outlet_light_level` | LED indicator level |

### Water sensor

| Metric | Description |
|---|---|
| `ikea_dirigera_water_sensor_water_leak_detected` | Water leak detected (1) or not (0) |
| `ikea_dirigera_water_sensor_battery_percentage` | Battery percentage |

### WebSocket

| Metric | Labels | Description |
|---|---|---|
| `ikea_dirigera_websocket_received_events` | `source`, `type` | Events received from the Dirigera WebSocket |
| `ikea_dirigera_websocket_errors` | `error` | WebSocket errors |

## API reference

The Dirigera hub API is documented here: https://codeberg.org/argrento/dirigera/src/branch/main/info/api.md


CREATE TABLE IF NOT EXISTS estore.events_by_device (
      device_uuid UUID
    , hour TIMESTAMP
    , kind INT
    , begin_time TIMESTAMP
    , uuid UUID
    , source INT
    , end_time TIMESTAMP
    , properties TEXT
    , PRIMARY KEY ((device_uuid, hour), kind, begin_time, uuid))
      WITH CLUSTERING ORDER BY (kind ASC, begin_time DESC, uuid ASC);
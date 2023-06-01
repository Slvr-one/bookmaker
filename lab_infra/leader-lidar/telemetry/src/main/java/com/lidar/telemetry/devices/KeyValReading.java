package com.lidar.telemetry.devices;

import java.util.HashMap;
import java.util.Map;

public class KeyValReading implements Reading{
    private final Map<String,String> vals;
    private final String name;

    public KeyValReading(String name) {
        this.vals = new HashMap<>();
        this.name = name;
    }

    public String getType() {
        return "m";
    }

    public String getName() {
        return name;
    }

    public String getData() {
        StringBuilder sb = new StringBuilder();
        sb.append(name).append(" d(");
        for (Map.Entry<String,String> pair : vals.entrySet()) {
            sb.append(pair.getKey()).append('=').append(pair.getValue()).append(':');
        }
        sb.append(')');
        return sb.toString();
    }

    public void set(String k,String v) {
        vals.put(k,v);
    }
}

package com.lidar.telemetry.devices;

public class StringReading implements Reading{
    private final String val;
    private final String name;

    public StringReading(String name, String val) {
        this.val = val;
        this.name = name;
    }

    public String getType() {
        return "s";
    }

    public String getName() {
        return name;
    }

    public String getData() {
        return String.format("%s s('%s')",name,val);
    }
}

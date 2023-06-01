package com.lidar.telemetry.devices;

public class NumericReading implements Reading {
    private final float val;
    private final String name;

    public NumericReading(String name, float val) {
        this.val = val;
        this.name = name;
    }

    public String getType() {
        return "f";
    }

    public String getName() {
        return name;
    }

    public String getData() {
        return String.format("%s f(%f)",name,val);
    }
}

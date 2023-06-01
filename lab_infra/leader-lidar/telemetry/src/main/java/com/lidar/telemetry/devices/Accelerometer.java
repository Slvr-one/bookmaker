package com.lidar.telemetry.devices;

public class Accelerometer implements Device{

    private float curVal=0;

    @Override
    public Reading read() {
        curVal+=Math.random()-.5;
        return new NumericReading("ACL",curVal);
    }
}

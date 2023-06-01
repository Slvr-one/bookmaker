package com.lidar.telemetry.devices;

public class Speedometer implements Device{

    private float curVal=0;

    @Override
    public Reading read() {
        curVal+=10*(Math.random()-.5);
        if (curVal<0) curVal = 0;
        if(curVal>160) curVal-=1;
        return new NumericReading("VEL",curVal);
    }
}

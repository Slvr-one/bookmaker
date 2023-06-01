package com.lidar.telemetry;

import com.lidar.telemetry.devices.Accelerometer;
import com.lidar.telemetry.devices.Device;
import com.lidar.telemetry.devices.Speedometer;

import java.util.ArrayList;

public class TelemetrySuite {

    private ArrayList<Device> telemetryDevices = new ArrayList<>();

    public TelemetrySuite() {
        telemetryDevices.add(new Accelerometer());
        telemetryDevices.add(new Speedometer());
    }

    public String read() {
        StringBuilder log = new StringBuilder();
        for (Device d : telemetryDevices) {
            log.append(d.read().getData()).append(';');
        }

        return log.toString();

    }

}

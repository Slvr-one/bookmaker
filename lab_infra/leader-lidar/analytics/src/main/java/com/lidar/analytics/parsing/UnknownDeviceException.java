package com.lidar.analytics.parsing;

public class UnknownDeviceException extends RuntimeException {
    public UnknownDeviceException(String deviceId) {
        super("Unknown device with code "+deviceId);
    }
}

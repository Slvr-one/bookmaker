package com.lidar.analytics.devices;

public interface ReadingParser {

    Reading parse(String name, String reading);
}

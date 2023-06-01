package com.lidar.analytics.parsing;

import com.lidar.analytics.devices.Reading;
import com.lidar.analytics.devices.ReadingParser;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Parser {
    Map<String, ReadingParser> deviceById = new HashMap<>();

    public void register(String id, ReadingParser d) { deviceById.put(id,d); }

    public List<Reading> parse(String log) {
        ArrayList<Reading> ret = new ArrayList<>();

        String[] readings = log.split(";");
        for (String r : readings) {
            String deviceId  = getDeviceId(r);
            if (deviceId==null) throw new ParseException(r);
            ReadingParser d = deviceById.get(deviceId);
            if (d==null) throw new UnknownDeviceException(deviceId);
            ret.add(d.parse(deviceId,r));
        }
        return ret;
    }

    private String getDeviceId(String reading) {
        String id=reading.trim().substring(0,3);
        return id;
    }


}

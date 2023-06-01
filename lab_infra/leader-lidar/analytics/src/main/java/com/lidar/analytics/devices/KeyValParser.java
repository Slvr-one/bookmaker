package com.lidar.analytics.devices;

import com.lidar.analytics.parsing.ParseException;

public class KeyValParser implements ReadingParser {
    @Override
    public Reading parse(String name, String reading) {
        int parLeft = reading.indexOf('(');
        int parRight = reading.indexOf(')');
        char type=reading.charAt(parLeft-1);
        if (type!='d') throw new ParseException("Expected dictionary as input and got unknown type "+type);
        String[] vals = reading.substring(parLeft + 1, parRight).split(":");
        KeyValReading ret = new KeyValReading(name);
        for (String a : vals) {
            String[] kv=a.split("=");
            ret.set(kv[0],kv[1]);
        }
        return ret;

    }
}

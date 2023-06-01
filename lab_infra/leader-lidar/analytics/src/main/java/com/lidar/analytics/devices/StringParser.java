package com.lidar.analytics.devices;

import com.lidar.analytics.parsing.ParseException;

public class StringParser implements ReadingParser {
    @Override
    public Reading parse(String name, String reading) {
        int parLeft = reading.indexOf('(');
        int parRight = reading.indexOf(')');
        char type=reading.charAt(parLeft-1);
        if (type!='s') throw new ParseException("Expected string input and got unknown type "+type);
        String val = reading.substring(parLeft + 2, parRight-1);
        return new StringReading(name, val);
    }
}

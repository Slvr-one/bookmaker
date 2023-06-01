package com.lidar.analytics.devices;

import com.lidar.analytics.parsing.ParseException;

public class NumericParser implements ReadingParser {
    @Override
    public Reading parse(String name, String reading) {
        int parLeft = reading.indexOf('(');
        int parRight = reading.indexOf(')');
        char type=reading.charAt(parLeft-1);
        if (type!='f') throw new ParseException("Expected numeric input and got unknown type "+type);
        String val = reading.substring(parLeft+1,parRight);
        return new NumericReading(name,Float.parseFloat(val));
    }
}

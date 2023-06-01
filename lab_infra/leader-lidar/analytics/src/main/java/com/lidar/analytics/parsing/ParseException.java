package com.lidar.analytics.parsing;

public class ParseException extends RuntimeException {
    public ParseException(String line) {
        super(String.format("Cannot parse '%s' - protocol version mismatch!?",line));
    }
}

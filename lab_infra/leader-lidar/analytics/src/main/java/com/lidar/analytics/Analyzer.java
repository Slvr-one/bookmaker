package com.lidar.analytics;

import com.lidar.analytics.devices.NumericParser;
import com.lidar.analytics.devices.Reading;
import com.lidar.analytics.parsing.Parser;

import java.util.ArrayList;
import java.util.List;

public class Analyzer {
    private static final double C = 1000000;
    private List<WarnRange> warnings = new ArrayList<>();
    private final Parser parser;

    private int noWarningBoundary = 0;

    public Analyzer() {
        registerWarning("PERFORM AUTOBREAK",3);
        registerWarning("WARN AHEAD",15);
        registerWarning("WARN LEFT",15);
        registerWarning("WARN RIGHT",5);

        this.parser = new Parser();
        parser.register("ACL", new NumericParser());
        parser.register("VEL", new NumericParser());
    }

    /**
     * registers warnings
     * @param w
     * @param probability - in reference to C, i.e. 1 is 1/C
     */
    private void registerWarning(String w, int probability) {
        noWarningBoundary +=probability;
        warnings.add(new WarnRange(w, noWarningBoundary));
    }

    public String analyze(String log) {
        List<Reading> readings = parser.parse(log);

        // Here we should use the readings to analyze car situation and issue warnings.
        // Naturally - this part does not exist. It is, after all, just a dev-ops excercise
        int event = (int) (Math.random()*C);

        if (event>=noWarningBoundary) return null;

        for (WarnRange range : warnings) {
            if (event<range.getTopRange()) return range.getWarning();
        }
        return "WTF "+ event;
    }

}

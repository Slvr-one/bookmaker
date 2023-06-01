package com.lidar.analytics;

public class WarnRange {
    private final String warning;
    private final int bottomRange;

    public WarnRange(String warning, int bottomRange) {
        this.warning = warning;
        this.bottomRange = bottomRange;
    }

    public String getWarning() {
        return warning;
    }

    public int getTopRange() {
        return bottomRange;
    }
}

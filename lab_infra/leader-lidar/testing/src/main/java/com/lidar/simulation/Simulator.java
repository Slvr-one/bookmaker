package com.lidar.simulation;

import com.lidar.analytics.Analyzer;
import com.lidar.telemetry.TelemetrySuite;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;

public class Simulator {

    private static final int SIMULATION_LENGTH = 200000;

    private TelemetrySuite car = new TelemetrySuite();
    private Analyzer analyzer = new Analyzer();

    public static void main(String[] args) {

        try {

            File testFile = new File("tests.txt");
            if (!testFile.exists()) {
                System.err.println("Could not find tests.txt");
                System.exit(1);
            }

            BufferedReader b = new BufferedReader(new FileReader(testFile));

            String testDescription = "";
            int testNumber=0;

            while ((testDescription = b.readLine()) != null) {
                testNumber++;
                System.out.println(String.format("============ Executing test %d:%s ================",testNumber,testDescription));
                int sleepTime = (int)(Math.random()*30);
                int length = (int)(Math.random()*SIMULATION_LENGTH);
                new Simulator().go(sleepTime,length);
            }

        } catch (IOException e) {
            e.printStackTrace();
        }
    }


    private void go(int sleepTime, int length) {
        for (int t=0;t<SIMULATION_LENGTH;t++) {
            String action = analyzer.analyze(car.read());
            if (action!=null) System.out.println(String.format("%07d:%s",t,action));
        }
    }

}

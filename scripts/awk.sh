#!/bin/bash

# Counting the occurrences of each word in a text file:

awk '{ for(i=1; i<=NF; i++) count[$i]++ } END { for(word in count) print word, count[word] }' file.txt
# using an associative array count to keep track of the word frequencies. 
# It iterates over each word in the input file, 
# increments the count for that word,
# prints the word and its count at the end.
# Extracting specific fields from a CSV file:

awk -F',' '{ print $1, $3 }' file.csv
# the -F option sets the field separator to a comma. 
# The script then prints the first and third fields of each line in the CSV file. 
# You can modify the field numbers or add additional fields as needed.

# Summing values in a column and calculating the average:

awk '{ sum += $1 } END { avg = sum / NR; print "Sum:", sum; print "Average:", avg }' file.txt
# calculates the sum of values in the first column of a file 
# then calculates the average by dividing the sum by the total number of records (NR).
# and prints result at the end.

# Filtering lines based on a condition:

awk '$3 > 50 { print $1, $2 }' file.txt

# filters the lines of a file based on a condition. 
# prints the first and second fields of lines where the third field is greater than 50. 
# modify the condition to suit your needs.

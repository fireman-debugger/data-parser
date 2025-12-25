"""
data-parser: A command-line tool for parsing and processing data files.

Usage:
    data-parser <input_file> [options]

Options:
    -h, --help         Show this help message and exit
    -o, --output       Specify the output file [default: stdout]
    -t, --type         Specify the type of data to parse [default: csv]
    -d, --delimiter    Specify the delimiter used in the input file [default:,]
"""

import argparse
import csv
import sys

def parse_csv(input_file, output_file, delimiter):
    with open(input_file, 'r') as f:
        reader = csv.reader(f, delimiter=delimiter)
        data = list(reader)

    with open(output_file, 'w') as f:
        writer = csv.writer(f, delimiter=delimiter)
        writer.writerows(data)

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('input_file', help='Input file to parse')
    parser.add_argument('-o', '--output', default='stdout', help='Output file to write to')
    parser.add_argument('-t', '--type', default='csv', help='Type of data to parse')
    parser.add_argument('-d', '--delimiter', default=',', help='Delimiter used in the input file')
    args = parser.parse_args()

    if args.output == 'stdout':
        output_file = sys.stdout
    else:
        output_file = open(args.output, 'w')

    if args.type == 'csv':
        parse_csv(args.input_file, output_file, args.delimiter)
    else:
        print('Error: Unsupported data type', file=sys.stderr)

if __name__ == '__main__':
    main()
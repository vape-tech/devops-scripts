#!/usr/bin/env python3

import os
import sys
import argparse
import subprocess

# Define the version of the project
__version__ = '0.1.0'

# Define the description of the project
__description__ = 'A collection of DevOps scripts'

# Define the parser for the script
parser = argparse.ArgumentParser(description=__description__)

# Add a version argument to the parser
parser.add_argument('-v', '--version', action='version', version=__version__)

# Add a help argument to the parser
parser.add_argument('-h', '--help', action='help', help='Show this help message and exit')

# Parse the arguments
args = parser.parse_args()

# Check if the arguments are valid
if args.version:
    print(f'devops-scripts version {__version__}')
    sys.exit(0)

if args.help:
    parser.print_help()
    sys.exit(0)

# Check if the repository is clean
if subprocess.call(['git', 'diff', '--quiet']) == 0:
    print('Repository is clean')
else:
    print('Repository is not clean')
    sys.exit(1)

# Check if the project is installed with pip
if subprocess.call(['pip', 'show', 'devops-scripts']) == 0:
    print('Project is installed with pip')
else:
    print('Project is not installed with pip')
    sys.exit(1)

# Check if the project is installed as a Git submodule
if os.path.exists('.gitmodules'):
    print('Project is installed as a Git submodule')
else:
    print('Project is not installed as a Git submodule')
    sys.exit(1)
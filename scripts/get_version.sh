#!/bin/bash
echo -n $(git describe --tags --abbrev=0) > version.txt
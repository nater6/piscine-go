#! /bin/bash
 find -name "*.sh" |rev| cut -d "/" -f1 | rev | sed 's/.sh//g' | sort -n -r 
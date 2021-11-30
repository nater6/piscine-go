#! /bin/bash
 find -name "*.sh" 
 find -name "*.sh" |rev
 find -name "*.sh" |rev| cut -d "/" -f1 
 find -name "*.sh" |rev| cut -d "/" -f1 | rev 
 find -name "*.sh" |rev| cut -d "/" -f1 | rev | sed 's/.sh//g' 
 find -name "*.sh" |rev| cut -d "/" -f1 | rev | sed 's/.sh//g' | sort -n -r 
$Regex_Pattern = '^\d\d(?|(-)|(---)|(:)|(\.))(\d\d\1){2}\d\d$';

$Test_String = <STDIN> ;
if($Test_String =~ /$Regex_Pattern/){
    print "true";
} else {
    print "false";
}

# $Regex_Pattern = '^\2(((tac)|\1tic\1)+)$';
$Regex_Pattern = '^(\2tic|(tac))+$';

$Test_String = <STDIN> ;
if($Test_String =~ /$Regex_Pattern/){
    print "true";
} else {
    print "false";
}

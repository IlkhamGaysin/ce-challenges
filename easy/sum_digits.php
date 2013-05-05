<?php
$fh = fopen($argv[1], "r");
$c = 0;
while (true) {
	$test = fgetc($fh);

	if ($test == "") {
		break;
	}
	elseif ($test == "\n") {
		echo $c, "\n";
		$c = 0;
	}
	else {
		$c += intval($test);
	}
}
?>

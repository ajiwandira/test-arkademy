<?php

function triangle(int $jumlah) {
	 
	// Outer loop for changing rows
	for ($i = $jumlah; $i >= 1; $i--) {
	 
		if ($i == $jumlah){
			for($j=$i;$j>=1;$j--){
				echo " "."*"." ";
				
    		}
			echo "</br>";
		}
		else {
			//Loop added for outer spacing
			 for ($j = $jumlah; $j > $i; $j--) {
				echo "  ";
			 }
			 echo "*";
			 
			 
			 //Loop added for internal spacing
			 for ($j = 1; $j < ($i - 1) * 2; $j++) {
			 echo "  ";
			 }
			 if ($i == 1)
			 echo "</br>";
			 else
			 echo "*</br>";
		}
		
	 
	}
	
}

triangle(10);

?>

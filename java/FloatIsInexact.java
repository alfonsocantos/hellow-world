/**
*   From http://www.odi.ch/prog/design/newbies.php#46
*/

class FloatIsInexact {
	public static void main (String args[]){
		float total = 0.0f;
		float addition = 0.30f;
		for (int i = 0; i<100; i++){
			total += addition;
		}
		System.out.println("Expecting 0.30 x 100 = 30.0 but... " 
				+ total);
	}
}

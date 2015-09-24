// http://www.odi.ch/prog/design/newbies.php
/**
*	Don't use Float for work with money calcularions,
*	just for measurements, not when precision is a req.
*
*   From http://www.odi.ch/prog/design/newbies.php#46
*/
import java.math.BigDecimal;
import java.math.RoundingMode;


class FloatIsInexact {
	public static void main (String args[]){
		
		// The wrong way.
		float total = 0.0f;
		float addition = 0.30f;
		for (int i = 0; i<100; i++){
			total += addition;
		}
		System.out.println("Expecting 0.30 x 100 = 30.0 but... " 
				+ total);

		// The correct way.
		BigDecimal totalBD = BigDecimal.ZERO;
		BigDecimal additionBD = new BigDecimal(0.30f);
		for (int i = 0; i<100; i++){
                        totalBD = new BigDecimal(totalBD.add(additionBD).toString());
                }
		System.out.println (totalBD.setScale(2, RoundingMode.HALF_UP));

	}
}

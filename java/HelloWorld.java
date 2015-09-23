public class HelloWorld {
    public static void main (String[] args){
	
	String _string = "";
	
	if (args == null || args.length <= 0){
	    _string = "World";
	} else {
	    _string = args[0];
	}

	System.out.println("Hello " + _string + "!");
    }
}

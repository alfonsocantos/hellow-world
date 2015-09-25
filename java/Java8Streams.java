/*
*   Requires java 8
*   Based on http://eyalgo.com/2015/01/06/java-8-stream-and-lambda-expressions-parsing-file-example/ by eyalgo
*
*/

import java.io.FileWriter;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Stream;

class Java8Streams {

	public static void main (String[] args){
	
		if (args.length == 2) {
			System.out.println("Procesing file... " + args[0] + " filter: " + args[1]);
			
			try {
				Stream<String> lines = Files.lines(Paths.get(args[0]));
				parse (lines, args[1]);
			} catch (Exception e) {
				System.out.println(e.getMessage());
			}
		} else {
			System.out.println (new StringBuilder().append("Usage:\n\tJava8Streams file filter"));
		}
	
	}

	private static void parse (Stream<String> lines, String filter) throws IOException {
		lines
			.filter(line -> line.contains(filter))
			.forEach(System.out::println);
		lines.close();
	}

}

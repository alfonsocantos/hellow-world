// http://www.odi.ch/prog/design/newbies.php
/** reading a file using or not BufferedReaders.
*  (the same for writting)
*/

import java.io.BufferedInputStream;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.io.InputStream;

class BufferedReaders {

        public static void main (String[] args){
                try {
                        readFromFile (new File("/dev/zero"));
			readFromFileBuffered (new File("/dev/zero"));
                } catch (Exception e) {
                        e.printStackTrace();
                }

        }

        private static void readFromFile (File file) throws FileNotFoundException {
                long started = System.currentTimeMillis();
                InputStream in = new FileInputStream(file);
                int b;
                int max = 1024 * 1024;
                for (int i = 0; i <= max; i++){
                        try {
                                b = in.read();
                        } catch (IOException e) {
                                e.printStackTrace();
                        }
                }
                System.out.println("Reading not buffered takes " + (System.currentTimeMillis() - started + " ms"));
        }

	 private static void readFromFileBuffered (File file) throws FileNotFoundException {
                long started = System.currentTimeMillis();
                InputStream in = new BufferedInputStream(new FileInputStream(file));
                int b;
                int max = 1024 * 1024;
                for (int i = 0; i <= max; i++){
                        try {
                                b = in.read();
                        } catch (IOException e) {
                                e.printStackTrace();
                        }
                }
                System.out.println("Reading buffered takes " + (System.currentTimeMillis() - started + " ms"));
        }
}


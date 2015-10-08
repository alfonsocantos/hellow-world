# include <stdio.h>
# include <stdlib.h>
# include <string.h>


int main ( int argc, char *argv[] ){
  if ( argc != 3) {
    printf("Usage: \n\t %s file_substrings whole_file \n", argv[0]);
    return -1;
  }
  FILE *finput, *fwhole;
  printf("Cheking if all lines in %s are in the file %s...\n",
    argv[1], argv[2]);

  finput = fopen(argv[1], "r");
  if ( finput == 0 ) cant(argv[1]);

  fwhole = fopen(argv[2], "r");
  if ( fwhole == 0 ) cant(argv[1]);

  // for the first file
  char *lineinput = NULL;
  size_t leninput = 0;
  ssize_t readinput;
  // for the second file
  char *linewhole = NULL;
  size_t lenwhole = 0;
  ssize_t readwhole;


  while ((readinput = getline(&lineinput, &leninput, finput)) != -1) {
      while ((readwhole = getline(&linewhole, &lenwhole, fwhole) != -1 ) ) {
// TODO me falta construir la condición de salida del bucle interno
        if (match (lineinput, linewhole) )
      }
  }


  fclose(finput);
  fclose(fwhole);
  return 0;
}

/*** Display if the line doesn't match the line ********/
int match (word, sent)
char *word;
char *sent;
{
  char *str = strstr(sent, word);
  if (str)
    return 1;
  return 0;
}

/*** Display usage summary *****************************/
usage(s)
char    *s;
{
   fprintf(stderr, "?GREP-E-%s\n", s);
   fprintf(stderr,
      "Usage: checkStrings file_patterns file\n");
   exit(1);
}


/*** Report unopenable file ****************************/
cant(s)
char *s;
{
   fprintf(stderr, "%s: cannot open\n", s);
   exit(1);
}

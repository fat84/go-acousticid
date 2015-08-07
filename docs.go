package acousticid

/*
 Program flow:
 * Read input directory -- DONE
  * For each file do the following -- DONE
  	* Check if the file has id3 tag //todo
  	* Is the id3 tag valid: Artist, Song Title, Album Name //todo
  	* If not valid: //todo
  	 * get the acoustic fingerprint?  -- DONE
  	 * Make the acoustic ID REST call  -- DONE
  	 * If valid result found: //todo needs refinement
  	  * Make ID3 tag struct  -- DONE
  	  * Write into the mp3 file  -- DONE
  	 * If no valid result found: //DONE
  	  * Check for details in file name //DONE
  	  * Create the id3 tag struct from file name //DONE
  	  * Write into the mp3 file //DONE
  	* If valid:
  	 * continue //todo
 * parallel execution //DONE
 Files:1
 acoustic-id.go

 id3.go

 id3-utils.go

 */

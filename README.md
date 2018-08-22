# file_obfuscator
Transform a file in another type by prepending a different header


This simple program allows you to wrap any file in a more innocuous file like wav, jpg, png or ico (the only 4 implemented)


There are two action defined:


	'c' means camoufage which will prepend the desired file's headers
 	'r' means reveal, it strips the fake header and recreate the original file

To hide a file: 

	./obf -a c -f cat -e png

(it will take the utility cat and make it look like a png image)

To re-create the original file:

	./obf -a r -f cat.png
  
(it will take the 'image' cat.png and return a executable)  


Few examples:

	$ file cat
	cat: Mach-O 64-bit executable x86_64
	$ file cat.png
	cat.png: PNG image data, 50331776 x 33554432, 16-bit grayscale, non-interlaced
	$ file cat.png_orig 
	cat.png_orig: Mach-O 64-bit executable x86_64
	$ file cat.ico
	cat.ico: MS Windows icon resource - 2 icons, 16x16, 32 bits/pixel, 32x32, 207 colors, 65261 planes, 7 bits/pixel


Tested on RHEL7 and MacOS 10, fairly sure will work everywhere

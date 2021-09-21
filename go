set -e

python chuckie_orig.py -a >chuckie_orig.asm
../tools/acme -o chuckie_orig.bin chuckie_orig.asm
xxd -o 4352 CH_EGG_1100_29AB_ORIG.bin >temp1.hex
xxd -o 4352 chuckie_orig.bin >temp2.hex
diff temp1.hex temp2.hex
rm temp1.hex
rm temp2.hex
